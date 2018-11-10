package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

var dev []sane.Device
var opt = map[string][]sane.Option{}
var con = map[string]*sane.Conn{}

func errh(err error, c *gin.Context) bool {
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, err.Error())
		return true
	}
	return false
}

func devices(c *gin.Context) {
	c.JSON(200, dev)
}

func options(c *gin.Context) {
	options, ok := opt[c.Query("device")]
	if !ok {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Device not found")
	}
	c.JSON(200, options)
}

func scan(c *gin.Context) {
	c.Header("Content-Type", "image/png")
	enc := png.Encoder{CompressionLevel: png.BestCompression}
	device := c.Query("device")
	if device == "" {
		device = dev[0].Name
	}
	f, err := os.OpenFile("/tmp/sane-webscan.png", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Printf("Copy error: %s (OpenFile)\n", err)
	}
	err = doScan(device, f, map[string]interface{}{}, enc.Encode)
	if errh(err, c) {
		return
	}
	c.File("/tmp/sane-webscan.png")
}

func scanJpg(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")
	device := c.Query("device")
	if device == "" {
		device = dev[0].Name
	}
	f, err := os.OpenFile("/tmp/sane-webscan.jpg", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Printf("Copy error: %s (OpenFile)\n", err)
	}
	doScan(device, f, map[string]interface{}{}, func(w io.Writer, m image.Image) error {
		return jpeg.Encode(w, m, &jpeg.Options{
			Quality: 80,
		})
	})
	if errh(err, c) {
		return
	}
	c.File("/tmp/sane-webscan.jpg")
}

func scanPdf(c *gin.Context) {
	f, err := os.OpenFile("/tmp/sane-webscan-convert.jpg", os.O_CREATE|os.O_RDWR, 0600)
	if errh(err, c) {
		return
	}

	device := c.Query("device")
	if device == "" {
		device = dev[0].Name
	}
	err = doScan(device, f, map[string]interface{}{
		"resolution": 300,
	}, func(w io.Writer, m image.Image) error {
		return jpeg.Encode(w, m, &jpeg.Options{
			Quality: 80,
		})
	})
	if errh(err, c) {
		return
	}

	err = exec.Command("convert", "/tmp/sane-webscan-convert.jpg", "/tmp/sane-webscan-convert.pdf").Run()
	if errh(err, c) {
		return
	}

	err = exec.Command("pdfsandwich", "-tesso", "-l deu", "/tmp/sane-webscan-convert.pdf").Run()
	if errh(err, c) {
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.File("/tmp/sane-webscan-convert_ocr.pdf")
}

func doScan(device string, w io.Writer, config map[string]interface{}, enc func(io.Writer, image.Image) error) error {
	log.Printf("Connecting to %s\n", device)
	connection, ok := con[device]
	if !ok {
		return errors.New("Device not found")
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
	for option, value := range config {
		info, err := connection.SetOption(option, value)
		log.Printf("Setting option %s to '%s': %+v%s", option, value, info, err)
	}

	log.Printf("Reading image...")
	image, err := connection.ReadImage()
	if err != nil {
		log.Printf("Scanning error: %s\n", err)
		return errors.New("Scanning Error: " + err.Error())
	}

	log.Printf("Encoding image...")
	err = enc(w, image)
	if err != nil {
		log.Printf("Image Encoding error: %s\n", err)
		return errors.New("Image Encoding Error: " + err.Error())
	}
	log.Printf("Scan complete.")

	return nil
}
