package main

import (
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func buttons() {
	for {
		for device, connection := range con {
			v, err := connection.GetOption("copy")
			if err == nil && v.(bool) == true {
				f, err := os.OpenFile("/tmp/sane-webscan-copy.jpg", os.O_CREATE|os.O_RDWR, 0600)
				if err != nil {
					log.Printf("Copy error: %s (OpenFile)\n", err)
				}
				err = doScan(device, f, map[string]interface{}{}, func(w io.Writer, m image.Image) error {
					return jpeg.Encode(w, m, &jpeg.Options{
						Quality: 80,
					})
				})
				if err != nil {
					log.Printf("Copy error: %s (Scan)\n", err)
				}
				cmd := exec.Command("lp", "-o", "portrait", "-o", "fit-to-page", "-o", "media=A4", "/tmp/sane-webscan-copy.jpg")
				err = cmd.Run()
				if err != nil {
					log.Printf("Copy error: %s (lp)\n", err)
				}
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
