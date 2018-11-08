package main

import (
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

func devices(c *gin.Context) {
	devices, err := sane.Devices()
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Device List Error: %s", err)
		return
	}

	c.String(200, "%+v\n", devices)
}

func options(c *gin.Context) {
	connection, err := sane.Open(c.Param("scanner"))
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.String(500, "Scanner Connection Error: %s", err)
		return
	}

	options := connection.Options()
	c.String(200, "%+v\n", options)
}

func scan(c *gin.Context) {
	connection, err := sane.Open(c.Param("scanner"))
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
