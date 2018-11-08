package main

import (
	"bytes"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tjgq/sane"
)

func main() {
	r := gin.Default()

	if err := sane.Init(); err != nil {
		panic(err)
	}
	defer sane.Exit()

	buttons()

	r.GET("/devices.json", devices)
	r.GET("/:scanner/options.json", options)
	r.GET("/:scanner/scan.png", scan)

	for _, asset := range AssetNames() {
		log.Printf("Loading asset: %s\n", asset)
		r.GET(strings.TrimSuffix(strings.TrimPrefix(asset, "web"), "index.html"), serveAsset)
	}

	r.Run()
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
