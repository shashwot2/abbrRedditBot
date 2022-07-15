package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Quote string `json:"quote"`
}

func getenv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	redditsecret := getenv("rsecret")
	fmt.Println(redditsecret)
	response, err := http.Get("https://api.kanye.rest/")
	if err != nil {
		log.Fatalf("Error getting response from kanye.rest")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body")
	}
	var responseobject Response
	json.Unmarshal(responseData, &responseobject)
	fmt.Println(responseobject.Quote)
}
