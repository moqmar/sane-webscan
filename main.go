package main

import (
	"bytes"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

func main() {

	err := sane.Init()
	if err != nil {
		panic(err)
	}
	defer sane.Exit()
	log.Printf("Initialized sane. (1/4)\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	go func() {
		for range c {
			for _, c := range con {
				c.Cancel()
				c.Close()
			}
			os.Exit(0)
		}
	}()

	dev, err = sane.Devices()
	if err != nil {
		panic(err)
	}
	for i, device := range dev {
		// TODO: blacklist
		connection, err := sane.Open(device.Name)
		if err != nil {
			log.Printf("Warning: device is offline (%s): %s", err, device.Name)
			dev = append(dev[:i], dev[i+1:]...)
			continue
		}
		con[device.Name] = connection
		opt[device.Name] = connection.Options()
		// TODO: default options
	}
	log.Printf("Initialized devices: %d found. (2/4)\n", len(dev))

	go buttons()
	log.Printf("Initialized buttons. (3/4)\n")

	if r, _ := os.LookupEnv("GIN_RELEASE"); r != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/devices.json", devices)
	r.GET("/options.json", options)
	r.GET("/scan.png", scan)
	r.GET("/scan.jpg", scanJpg)
	r.GET("/scan.pdf", scanPdf)
	//r.POST("/convert/:format", convert)

	for _, asset := range AssetNames() {
		log.Printf("Loading asset: %s\n", asset)
		r.GET(strings.TrimSuffix(strings.TrimPrefix(asset, "web"), "index.html"), serveAsset)
	}

	log.Printf("Initialized webserver (default port: 8080). (4/4)\n")

	err = r.Run()
	if err != nil {
		panic(err)
	}
}

func serveAsset(c *gin.Context) {
	p := c.Request.URL.Path
	if strings.HasSuffix(p, "/") {
		p = p + "index.html"
	}
	a := MustAsset("web" + p)
	var contentType = mime.TypeByExtension(filepath.Ext(p))

	c.DataFromReader(http.StatusOK, int64(len(a)), contentType, bytes.NewReader(a), map[string]string{})
}
