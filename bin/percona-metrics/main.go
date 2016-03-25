/*
	Copyright (c) 2016, Percona LLC and/or its affiliates. All rights reserved.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"
)

type Exporter struct {
	Name         string   `yaml:"name"`
	Port         string   `yaml:"port"`
	Args         []string `yaml:"args"`
	InstanceUUID string   `yaml:"instance_uuid"`
	err          error
	errCount     uint
}

var exporters []*Exporter
var mux *sync.Mutex = &sync.Mutex{}

var (
	ErrDupePort = errors.New("duplicate port")
)

const (
	BIN             = "percona-metrics"
	MAX_ERRORS      = 3
	DEFAULT_BASEDIR = "/usr/local/percona/pmm-client"
	DEFAULT_LISTEN  = "127.0.0.1:9004"
)

var (
	flagBasedir string
	flagListen  string
)

func init() {
	flag.StringVar(&flagBasedir, "basedir", DEFAULT_BASEDIR, "pmm-client basedir")
	flag.StringVar(&flagListen, "listen", DEFAULT_LISTEN, "IP:port to listen on")
	flag.Parse()
}

var stopChan chan struct{}
var eStopChan chan *Exporter

func main() {
	if err := os.Chdir(flagBasedir); err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(filepath.Join(flagBasedir, "exporters.yml"))
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(bytes, &exporters); err != nil {
		log.Fatal(err)
	}

	log.Println(BIN + " running")

	// Start all the exporters saved in exporters.yml.
	stopChan = make(chan struct{})
	eStopChan = make(chan *Exporter, 10)

	for _, e := range exporters {
		go run(e)
	}

	// Start API after running exporters ^ to avoid race conditions.
	log.Printf("Server address: %s\n", serverAddr())
	api := NewAPI(flagListen)
	go api.Run()

	// //////////////////////////////////////////////////////////////////////
	// Main loop
	// //////////////////////////////////////////////////////////////////////
	sigTermChan := make(chan os.Signal, 1)
	signal.Notify(sigTermChan, os.Interrupt, syscall.SIGTERM)

RESPAWN_LOOP:
	for {
		select {
		case e := <-eStopChan: // exporter e stopped
			respawn(e)
		case <-sigTermChan:
			log.Println("Caught signal, terminating...")
			break RESPAWN_LOOP
		}
	}

	// Got signal, shut down...
	close(stopChan)

	// Give exporters 1s to terminate.
	timeout := time.After(1 * time.Second)
TIMEOUT_LOOP:
	for {
		select {
		case e := <-eStopChan:
			// Don't respawn because we're shutting down.
			log.Printf(e.Name+" stopped: %s", e.err)
		case <-timeout:
			break TIMEOUT_LOOP
		case <-sigTermChan:
			log.Println("Caught signal, terminating...")
			break TIMEOUT_LOOP
		}
	}

	log.Println(BIN + " stopped")
}

func add(e *Exporter) error {
	mux.Lock()
	defer mux.Unlock()

	for _, ee := range exporters {
		if e.Port == ee.Port {
			return ErrDupePort
		}
	}

	go run(e)

	exporters = append(exporters, e)
	return save()
}

func run(e *Exporter) {
	defer func() {
		eStopChan <- e // received in main loop
	}()

	logFileName, _ := filepath.Abs(filepath.Join(flagBasedir, e.Name+"_"+e.Port+".log"))
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	bin, _ := filepath.Abs(filepath.Join(flagBasedir, e.Name))
	cmd := exec.Command(bin, e.Args...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	if e.InstanceUUID != "" {
		in, err := GetInstance(e.InstanceUUID)
		if err != nil {
			log.Fatal(err)
		}
		if in.DSN != "" {
			env := os.Environ()
			env = append(env, "DATA_SOURCE_NAME="+in.DSN)
			cmd.Env = env
		}
	}

	runErrChan := make(chan error, 1)
	go func() {
		runErrChan <- cmd.Run() // received below
	}()

	log.Printf("Started %s:%s %s", e.Name, e.Port, strings.Join(e.Args, " "))

	select {
	case err := <-runErrChan:
		e.err = err
	case <-stopChan:
		e.err = cmd.Process.Kill()
	}
}

func respawn(e *Exporter) {
	e.errCount++
	log.Printf(e.Name+" stopped (%d/%d), restarting: %s", e.errCount, MAX_ERRORS, e.err)
	if e.errCount < MAX_ERRORS {
		time.Sleep(1 * time.Second)
		go run(e)
	} else {
		log.Printf("Not restarting %s because it has failed too mamy times", e.Name)
	}
}

func serverAddr() string {
	out, err := exec.Command("pmm-admin", "server").Output()
	if err != nil {
		log.Println(err)
	}
	return strings.TrimSpace(string(out))
}

func save() error {
	// CALLER MUST LOCK mux!
	filename := filepath.Join(flagBasedir, "exporters.yml")
	bytes, _ := yaml.Marshal(exporters)
	return ioutil.WriteFile(filename, bytes, 0644)
}
