package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	ID       uint   `json:"id"`
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

var db map[uint]User

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

	fmt.Printf("get all users=%v \n", users)

	fmt.Println("get all into DB")
	db = map[uint]User{}
	initDB(db, users)

	fmt.Printf("GetAll = %v \n", GetAll(db))
	u, _ := Get(db, 4107)
	fmt.Printf("Get id = %v \n", u)
}

func initDB(db map[uint]User, users []User) {
	for _, user := range users {
		db[user.ID] = user
	}
}

func GetAll(db map[uint]User) []User {
	res := []User{}
	for _, val := range db {
		res = append(res, val)
	}
	return res
}

func Get(db map[uint]User, id uint) (User, error) {
	if val, ok := db[id]; ok {
		return val, nil
	}
	return User{}, fmt.Errorf("no such record")
}
