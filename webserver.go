package main

import (
	"github.com/gin-gonic/gin"
)

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

func scanPng(c *gin.Context) {
	scan <- Options{
		Format: PNG,
		Device: c.Query("device"),
	}
	err := <-wait
	if err != nil {
		c.String(500, err.Error())
	}
	c.File("/tmp/sane-webscan.png")
}

func scanJpg(c *gin.Context) {
	scan <- Options{
		Format: JPEG,
		Device: c.Query("device"),
	}
	err := <-wait
	if err != nil {
		c.String(500, err.Error())
	}
	c.File("/tmp/sane-webscan.jpg")
}

func scanPdf(c *gin.Context) {
	scan <- Options{
		Format: PDF,
		Device: c.Query("device"),
		Config: map[string]interface{}{
			"resolution": 300,
		},
	}
	err := <-wait
	if err != nil {
		c.String(500, err.Error())
	}
	c.File("/tmp/sane-webscan_ocr.pdf")
}
