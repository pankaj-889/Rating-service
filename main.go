package main

import (
	. "RatingService/model"
	"fmt"
)

func main() {
	dr := Driver{
		Id:        1,
		Name:      "ravi",
		Rating:    4.6,
		TotalRide: 9,
	}
	driver := Newdriver(dr)
	cl := Client{
		Id:     1,
		Name:   "bob",
		Rating: 4,
	}
	client := NewClient(cl)
	fmt.Println("before driver assigning")
	fmt.Println(*client)
	fmt.Println()

	fmt.Println("after driver assigning")
	err := client.AssignDrive(*driver)
	fmt.Println(*client)
	fmt.Println()

	client.CalculateRating(5)
	if err != nil {
		return
	}
	fmt.Println(client.Driver)
}
