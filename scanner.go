package main

import (
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

var dev []sane.Device

func devices(c *gin.Context) {
	c.JSON(200, dev)
}

func options(c *gin.Context) {
	connection, err := sane.Open(c.Query("scanner"))
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Scanner Connection Error: %s", err)
		return
	}

	options := connection.Options()
	c.JSON(200, options)
}

func scan(c *gin.Context) {
	connection, err := sane.Open(c.Query("scanner"))
	defer connection.Close()
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Scanner Connection Error: %s", err)
		return
	}

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