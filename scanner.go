package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/url"
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
	err := doScan(c.Query("device"), c.Writer, url.Values{}, enc.Encode)
	errh(err, c)
}

func scanJpg(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")
	err := doScan(c.Query("device"), c.Writer, url.Values{}, func(w io.Writer, m image.Image) error {
		return jpeg.Encode(w, m, &jpeg.Options{
			Quality: 80,
		})
	})
	errh(err, c)
}

func scanPdf(c *gin.Context) {
	f, err := os.OpenFile("/tmp/sane-webscan-convert.jpg", os.O_CREATE|os.O_RDWR, 0600)
	if errh(err, c) {
		return
	}

	err = doScan(c.Query("device"), f, url.Values{}, func(w io.Writer, m image.Image) error {
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

func doScan(device string, w io.Writer, config url.Values, enc func(io.Writer, image.Image) error) error {
	connection, ok := con[device]
	if !ok {
		return errors.New("Device not found")
	}

	var defaultConfig = map[string]interface{}{
		"source":        "Flatbed",
		"preview":       false,
		"depth":         16,
		"custom-gamma":  false,
		"swdskew":       true,
		"swcrop":        true,
		"swdespeck":     true,
		"swderotate":    true,
		"lamp-off-time": 2,
		"lamp-off-scan": false,

		"mode":       "Color",
		"resolution": 600,

		"brightness": 0,
		"contrast":   0,

		"tl-x": 0,
		"tl-y": 0,
		"br-x": 216,
		"br-y": 300,
	}
	for option, value := range defaultConfig {
		info, err := connection.SetOption(option, value)
		log.Printf("Setting default option %s to '%s': %+v%s", option, value, info, err)
	}

	for option, value := range config {
		info, err := connection.SetOption(option, value)
		log.Printf("Setting option %s to '%s': %+v%s", option, value, info, err)
	}

	image, err := connection.ReadImage()
	if err != nil {
		return errors.New("Scanning Error: " + err.Error())
	}

	err = enc(w, image)
	if err != nil {
		return errors.New("Image Encoding Error: " + err.Error())
	}

	return nil
}
