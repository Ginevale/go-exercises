package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Id          string `json:"id"`
	ResultType  string `json:"resultType"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SearchResponse struct {
	SearchType string     `json:"searchType"`
	Expression string     `json:"expression"`
	Responses  []Response `json:"results"`
}

func main() {

	var searchType string
	fmt.Println("Choose what type of search")
	fmt.Scanln(&searchType)
	searchType = string(searchType)

	var title string
	fmt.Println("Choose a title")
	fmt.Scanln(&title)
	title = string(title)

	url := "https://imdb-api.com/en/API/" + searchType + "/k_jfnccyjp/" + title
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBody := SearchResponse{}
	err = json.Unmarshal([]byte(body), &jsonBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jsonBody.Responses[0])
}
