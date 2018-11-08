package main

import (
	"bytes"
	"log"
	"mime"
	"net/http"
	"os"
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

	dev, err = sane.Devices()
	if err != nil {
		panic(err)
	}
	for _, device := range dev {
		// TODO: blacklist
		connection, err := sane.Open(device.Name)
		if err != nil {
			// TODO: scanners can be offline, this should be a warning?
			panic(err)
		}
		con[device.Name] = connection
		opt[device.Name] = connection.Options()
		// TODO: default options
		defer connection.Close()
	}
	log.Printf("Initialized devices: %d found. (2/4)\n", len(dev))

	buttons()
	log.Printf("Initialized buttons. (3/4)\n")

	if r, _ := os.LookupEnv("GIN_RELEASE"); r != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/devices.json", devices)
	r.GET("/options.json", options)
	r.GET("/scan.png", scan)

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
