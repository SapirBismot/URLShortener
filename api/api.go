package api

import (
	"crypto/sha256"
	"math/big"
	"net/http"

	"github.com/alextanhongpin/base62"
	"github.com/gin-gonic/gin"
)

var baseUrl = "http://localhost:9808/"
var urlsMap = make(map[string]string)

type UrlCreateRequestBody struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "URL Shortener API!"})
}

func CreateUrl(c *gin.Context) {
	var createRequestBody UrlCreateRequestBody
	if err := c.ShouldBindJSON(&createRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	longUrl := createRequestBody.LongUrl
	shortUrl := generateUrl(longUrl)
	urlsMap[shortUrl] = longUrl
	c.JSON(http.StatusOK, gin.H{
		"short_url": baseUrl + shortUrl,
	})
}

func Redirect(c *gin.Context) {
	shortUrl := c.Param("short_url")
	longUrl := urlsMap[shortUrl]
	c.Redirect(http.StatusFound, "http://"+longUrl)
}

func generateUrl(longUrl string) string {
	if longUrl == "" {
		return ""
	}
	hashUrl := sha256.Sum256([]byte(longUrl))
	bytesNumber := new(big.Int).SetBytes(hashUrl[:]).Uint64()
	encodedUrl := base62.Encode(bytesNumber)
	return encodedUrl
}
