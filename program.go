// go/src/http-client/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type users struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://gorest.co.in/public/v2/users")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)

	users := []users{}
	err2 := json.Unmarshal([]byte(body), &users)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("\nUsers:")
	fmt.Println(users)

	for _, user := range users {
		if user.Status == "inactive" {
			targetUrl := fmt.Sprint("https://gorest.co.in/public/v2/users/", user.Id)
			req, err := http.NewRequest("DELETE", targetUrl, nil)
			if err != nil {
				fmt.Printf("error %s", err)
				continue
			}
			resp, err := c.Do(req)
			if err != nil {
				fmt.Printf("error %s", err)
				continue
			}
			//defer resp.Body.Close()
			//body, err := ioutil.ReadAll(resp.Body)
			//if err != nil {
			//	fmt.Printf("error %s", err)
			//	continue
			//}
			fmt.Printf("\n%s Status: %s", targetUrl, resp.Status)
		}
	}
}
