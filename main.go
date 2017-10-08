package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	compressorRechargeTime = 4 * time.Minute
	blowTime               = 30 * time.Second
	zones                  = 5
)

var sprinlerIP = "192.168.1.173"

func main() {
	for {
		for i := 1; i <= zones; i++ {
			zone := i
			log.Printf("starting zone %d for %s", zone, blowTime)
			if err := startZone(zone); err != nil {
				log.Fatalf("failed to start zone %d", zone)
			}
			time.Sleep(blowTime)
			log.Printf("stopping all sprinkers")
			if err := StopAll(); err != nil {
				log.Fatalf("failed to stop all zones")
			}
			log.Printf("waiting %s for compressor to recharge", compressorRechargeTime)
			time.Sleep(compressorRechargeTime)
		}
	}
}

func StopAll() error {
	u := fmt.Sprintf("http://%s/stopSprinklers.htm", sprinlerIP)
	data := url.Values{}
	data.Add("stop", "active")
	resp, err := http.PostForm(u, data)
	if err != nil {
		return fmt.Errorf("failed to stop all zones: %v", err)
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("failed to stop all zones. Got code: %s", resp.Status)
	}

	return nil
}

func startZone(zone int) error {
	u := fmt.Sprintf("http://%s/program.htm", sprinlerIP)
	data := url.Values{}
	data.Add("doProgram", "1")
	data.Add(fmt.Sprintf("z%ddurMin", zone), "1")
	data.Add("runNow", "1")
	data.Add("pgmNum", "4")
	resp, err := http.PostForm(u, data)
	if err != nil {
		return fmt.Errorf("failed to start zone: %v", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to start zone. got: %s", resp.Status)
	}

	return nil
}
