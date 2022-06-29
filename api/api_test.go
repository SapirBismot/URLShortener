package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type test_struct struct {
	ShortUrl string `json:"short_url"`
}

var longUrl = "http://www.google.com/search?q=shortener+url&oq=shortener+url&aqs=chrome..69i57j69i59l2j0i512l2j69i60l3.4152j0j7&sourceid=chrome&ie=UTF-8"

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}
func TestDefaultRequest(t *testing.T) {
	mockResponse := `{"message":"URL Shortener API!"}`
	r := SetUpRouter()
	r.GET("/", Hello)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestCreateUrlRequest(t *testing.T) {
	r := SetUpRouter()
	r.POST("/create_url", CreateUrl)
	reqBody := map[string]string{"long_url": longUrl}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/create_url", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRedirectRequest(t *testing.T) {
	r := SetUpRouter()
	r.POST("/create_url", CreateUrl)
	r.GET("/:short_url", Redirect)

	reqBody := map[string]string{"long_url": longUrl}
	jsonBody, _ := json.Marshal(reqBody)
	postReq, _ := http.NewRequest("POST", "/create_url", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, postReq)
	createUrlRes, _ := ioutil.ReadAll(w.Body)
	var testStruct test_struct
	json.Unmarshal(createUrlRes, &testStruct)
	assert.Equal(t, http.StatusOK, w.Code)

	getReqShortUrl, _ := http.NewRequest("GET", testStruct.ShortUrl, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, getReqShortUrl)
	resUrl, _ := w.Result().Location()
	assert.Equal(t, resUrl.String(), "http://"+longUrl)
}
