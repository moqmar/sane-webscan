package main

import (
	"log"
	"os/exec"
	"time"
)

func buttons() {
	for {
		for device, connection := range con {
			v, err := connection.GetOption("copy")
			if err == nil && v.(bool) == true {

				// Scan
				scan <- Options{
					Config: map[string]interface{}{
						"resolution": 300,
					},
					Device: device,
					Format: TIFF,
				}
				err := <-wait
				if err != nil {
					log.Printf("Copy error: %s (scan)\n", err)
				}

				// Print
				cmd := exec.Command("lp", "-o", "portrait", "-o", "fit-to-page", "-o", "media=iso_a4_210x297mm", "/tmp/sane-webscan.tif")
				err = cmd.Run()
				if err != nil {
					log.Printf("Copy error: %s (print)\n", err)
				}

			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
