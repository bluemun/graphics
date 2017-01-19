// Copyright 2017 The bluemun Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package engine Defines methods used to run the engine.
package engine

import (
	"runtime"
	"time"

	"github.com/op/go-logging"
)

func init() {
	runtime.LockOSThread()
}

// Loop used by opengl to do its calls, needs to be called by the main thread.
func Loop() {
	for f := range mainfunc {
		f()
	}
}

var mainfunc = make(chan func())

// Do runs a given function on the main thread when there is time.
func Do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		if Logger.IsEnabledFor(logging.DEBUG) {
			timer := time.NewTicker(time.Second * 10)
			defer timer.Stop()
			go func() {
				<-timer.C
				Logger.Critical("Main thread took more then 10 seconds to run a single function.")
			}()
		}

		f()
		done <- true
	}
	<-done
}
