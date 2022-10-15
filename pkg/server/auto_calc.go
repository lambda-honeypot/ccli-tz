package server

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func shouldCalc() bool {
	return false
}

func doCalc() error {
	return nil
}

func calcLoop(tickerChannel <-chan time.Time) {
	for {
		select {
		case <-tickerChannel:
			runLoop()
		}
	}
}

func runLoop() {
	if shouldCalc() {
		err := doCalc()
		if err != nil {
			log.Warnf("error autocalculating: %v", err)
		}
	}
}
