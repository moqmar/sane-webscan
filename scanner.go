package main

import (
	"image/png"

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
	connection, ok := con[c.Query("device")]
	if !ok {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Device not found")
	}

	/*c.Request.ParseForm()
	for option, value := range c.Request.PostForm {
		info, err := connection.SetOption(option, value[0])
		log.Printf("Setting option %s to '%s': %+v%s", option, value[0], info, err)
	}*/
	connection.SetOption("brightness", 0)
	connection.SetOption("contrast", 0)
	connection.SetOption("lamp-off-time", 2)

	connection.SetOption("mode", 0)
	connection.SetOption("source", 0)
	connection.SetOption("preview", 1)
	connection.SetOption("depth", 1)
	connection.SetOption("resolution", 3)

	connection.SetOption("tl-x", 0)
	connection.SetOption("tl-y", 0)
	connection.SetOption("br-x", 216)
	connection.SetOption("br-y", 300)

	image, err := connection.ReadImage()
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Scanning Error: %s", err)
		return
	}

	c.Header("Content-Type", "image/png")
	c.Status(200)
	err = png.Encode(c.Writer, image)
	if err != nil {
		c.String(500, "Image Encoding Error: %s", err)
	}
}
