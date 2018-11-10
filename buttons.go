package main

import (
	"image"
	"image/jpeg"
	"io"
	"log"
	"os/exec"
	"time"

	"golang.org/x/image/tiff"
)

func buttons() {
	for {
		for device, connection := range con {
			v, err := connection.GetOption("copy")
			if err == nil && v.(bool) == true {
				cmd := exec.Command("lp", "-o", "portrait", "-o", "fit-to-page", "-o", "media=iso_a4_210x297mm")
				r, w := io.Pipe()
				cmd.Stdin = r
				go doScan(device, w, map[string]interface{}{}, func(w io.Writer, m image.Image) error {
					return tiff.Encode(w, m, &jpeg.Options{
						Quality: 80,
					})
				})
				err = cmd.Run()
				if err != nil {
					log.Printf("Copy error: %s (lp)\n", err)
				}
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
