package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      string `json:"age"`
	Social   struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	} `json:"social"`
	Like struct {
		Food  string `json:"food"`
		Brand string `json:"brand"`
	} `json:"like"`
}

func main() {
	filepath := "./data.json"
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read file err", err.Error())
	}

	var users []User
	err = json.Unmarshal(file, &users)
	if err != nil {
		fmt.Println("unmarshal json error", err.Error())
	}

	fmt.Printf("users=%v \n", users)

}
