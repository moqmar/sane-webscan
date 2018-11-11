package main

import (
	"errors"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"os/exec"

	"golang.org/x/image/tiff"

	"github.com/tjgq/sane"
)

var dev []sane.Device
var opt = map[string][]sane.Option{}
var con = map[string]*sane.Conn{}

type Format int

const (
	JPEG Format = 0
	PNG  Format = 1
	PDF  Format = 2
	TIFF Format = 3
)

type Options struct {
	Config map[string]interface{}
	Format Format
	Device string
}

var scan = make(chan Options)
var wait = make(chan error)

func scanLoop() {
	for o := <-scan; ; {
		if o.Device == "" {
			o.Device = dev[0].Name
		}

		log.Printf("Connecting to %s\n", o.Device)
		connection, ok := con[o.Device]
		if !ok {
			wait <- errors.New("Device not found")
			continue
		}

		var defaultConfig = map[string]interface{}{
			"preview":      false,
			"depth":        16,
			"custom-gamma": false,
			//"swdskew":       true,
			//"swcrop":        true,
			//"swdespeck":     true,
			//"swderotate":    true,

			"mode":       "Color",
			"resolution": 600,

			"brightness": 10,
			"contrast":   0,

			"tl-x": 0.0,
			"tl-y": 0.0,
			"br-x": 210.0,
			"br-y": 297.0,
		}
		log.Printf("Applying default configuration...\n")
		for option, value := range defaultConfig {
			info, err := connection.SetOption(option, value)
			log.Printf("Setting default option %s to '%s': %+v%s", option, value, info, err)
		}

		log.Printf("Applying additional configuration...\n")
		for option, value := range o.Config {
			info, err := connection.SetOption(option, value)
			log.Printf("Setting option %s to '%s': %+v%s", option, value, info, err)
		}

		log.Printf("Reading image...")
		image, err := connection.ReadImage()
		if err != nil {
			wait <- errors.New("Scanning Error: " + err.Error())
			continue
		}

		log.Printf("Encoding image...")
		switch o.Format {
		case JPEG:
			var file *os.File
			file, err = os.OpenFile("/tmp/sane-webscan.jpg", os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			err := jpeg.Encode(file, image, &jpeg.Options{
				Quality: 80,
			})
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			file.Close()
		case PNG:
			var file *os.File
			file, err = os.OpenFile("/tmp/sane-webscan.png", os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			enc := png.Encoder{CompressionLevel: png.BestCompression}
			err = enc.Encode(file, image)
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			file.Close()
		case TIFF:
			var file *os.File
			file, err = os.OpenFile("/tmp/sane-webscan.tif", os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			err = tiff.Encode(file, image, &tiff.Options{
				Compression: tiff.Deflate,
				Predictor:   true,
			})
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
		case PDF:
			var file *os.File
			file, err = os.OpenFile("/tmp/sane-webscan.jpg", os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			err := jpeg.Encode(file, image, &jpeg.Options{
				Quality: 80,
			})
			if err != nil {
				wait <- err
				file.Close()
				continue
			}
			file.Close()

			err = exec.Command("convert", "/tmp/sane-webscan.jpg", "/tmp/sane-webscan.pdf").Run()
			if err != nil {
				wait <- err
				continue
			}

			err = exec.Command("pdfsandwich", "-tesso", "-l deu", "/tmp/sane-webscan.pdf").Run()
			if err != nil {
				wait <- err
				continue
			}
		}
		log.Printf("Scan complete.")

		wait <- nil
	}
}
