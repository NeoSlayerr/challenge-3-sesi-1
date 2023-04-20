package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestBody struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		apiUrl := "https://jsonplaceholder.typicode.com/posts"

		data := RequestBody{
			Water: rand.Intn(15) + 1,
			Wind:  rand.Intn(15) + 1,
		}

		bs, err := json.Marshal(data)

		if err != nil {
			log.Panicf("error while converting struct to json => %s \n", err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bs))

		if err != nil {
			log.Panicf("error while defining the request instance => %s \n", err.Error())
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		response, err := client.Do(request)

		if err != nil {
			log.Panicf("error while sending the api request => %s \n", err.Error())
		}

		defer response.Body.Close()

		responseBody, err := ioutil.ReadAll(response.Body)

		var requestBody RequestBody

		err = json.Unmarshal(responseBody, &requestBody)

		data = RequestBody{
			Water: requestBody.Water,
			Wind:  requestBody.Wind,
		}

		bs, err = json.MarshalIndent(data, "", " ")

		if err != nil {
			log.Panicf("error while converting struct to json => %s \n", err.Error())
		}

		fmt.Println(string(bs))

		if data.Water > 8 {
			fmt.Println("status water : bahaya")
		} else if data.Water >= 6 && data.Water <= 8 {
			fmt.Println("status water : siaga")
		} else if data.Water < 5 {
			fmt.Println("status water : aman")
		}

		if data.Wind > 15 {
			fmt.Println("status wind : bahaya")
		} else if data.Wind >= 7 && data.Wind <= 15 {
			fmt.Println("status wind : siaga")
		} else if data.Wind < 6 {
			fmt.Println("status wind : aman")
		}
		fmt.Println("")
		time.Sleep(time.Second * 15)
	}
}
