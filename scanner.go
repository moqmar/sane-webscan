package main

import (
	"errors"
	"image/png"
	"io"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

var dev []sane.Device
var opt = map[string][]sane.Option{}
var con = map[string]*sane.Conn{}

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
	c.Request.ParseForm()
	err := doScan(c.Query("device"), c.Writer, c.Request.PostForm)
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, err.Error())
		return
	}
}

func doScan(device string, w io.Writer, config url.Values) error {
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

	err = png.Encode(w, image)
	if err != nil {
		return errors.New("Image Encoding Error: " + err.Error())
	}

	return nil
}
