package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"crypto/sha256"
	"github.com/alextanhongpin/base62"
	"math/big"
)

var port = 9808
var baseUrl = "http://localhost:"+ port + "/"
var urlsMap = make(map[string]string)

type UrlCreateRequestBody struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"message": "URL Shortener API!",})})
	r.POST("/create-url", createUrl)
	r.GET("/:short-url", redirect)

	err := r.Run(":" + port)
	if err != nil {
		log.fatal(err)
	}
}

func createUrl(c *gin.Context) {
	var createRequestBody UrlCreateRequestBody
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	longUrl := createRequestBody.LongUrl
	shortUrl := generateUrl(longUrl)
	urlMap[shortUrl] = longUrl
	c.JSON(200, gin.H{
		"short_url": baseUrl + shortUrl,
	})
}

func redirect(c *gin.Context) {
	shortUrl := c.Param("short-url")
	longUrl := urlMap[shortUrl]
	c.Redirect(302, longUrl)
}

func generateUrl(longUrl string) string {
	hashUrl := sha256.Sum256([]byte(longUrl)])
	bytesNumber := new(big.Int).SetBytes(hashUrl[:]).Uint64()
	encodedUrl := base62.Encode(bytesNumber)
	return encodedUrl
}
