package main

import (
	"bytes"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func scan(c *gin.Context) {

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

func main() {
	r := gin.Default()

	r.GET("/scan.png", scan)

	for _, asset := range AssetNames() {
		log.Printf("Loading asset: %s\n", asset)
		r.GET(strings.TrimSuffix(strings.TrimPrefix(asset, "web"), "index.html"), serveAsset)
	}

	r.Run()
}
