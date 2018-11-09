package main

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"time"
)

func buttons() {
	for {
		for device, connection := range con {
			v, err := connection.GetOption("copy")
			if err == nil && v.(bool) == true {
				f, err := os.OpenFile("/tmp/sane-webscan-copy.png", os.O_CREATE|os.O_RDWR, 0600)
				if err != nil {
					log.Printf("Copy error: %s (OpenFile)\n", err)
				}
				doScan(device, f, url.Values{})
				cmd := exec.Command("lp", "/tmp/sane-webscan-copy.png")
				err = cmd.Run()
				if err != nil {
					log.Printf("Copy error: %s (lp)\n", err)
				}
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
