package main

import (
	"errors"
	"fmt"
)

type IRatingService interface {
	CalculateRating(rating float32)
}

type Client struct {
	Id     int
	Name   string
	Rating int
	Driver Driver `json:"driver,nil"`
}

func (client *Client) AssignDrive(driver *Driver) error {
	b := client.Driver == (Driver{})
	if b {
		client.Driver = *driver
		return nil
	}

	return errors.New("driver is already assigned")
}

func (c *Client) GetRating(service IRatingService) {
	service.CalculateRating(5)
}

func (c *Client) ConfirmRide() {
	c.Driver.ConfirmRide = true
}

type Driver struct {
	Id          int
	Name        string
	Rating      float32
	TotalRide   int
	ConfirmRide bool `json:"confirmRide,omitempty"`
}

func (d *Driver) CalculateRating(rating float32) {
	status := d.ConfirmRide
	if status == true {
		currentPoint := d.Rating * float32(d.TotalRide)
		totalRide := d.TotalRide + 1
		currentRating := (currentPoint + rating) / float32(totalRide*1.0)

		d.Rating = currentRating
		d.TotalRide = d.TotalRide + 1
	}
}

func (d *Driver) GetDriver() *Driver {
	return d
}

func main() {
	driver := &Driver{
		Id:        1,
		Name:      "ravi",
		Rating:    4.6,
		TotalRide: 9,
	}
	client := &Client{
		Id:     1,
		Name:   "bob",
		Rating: 4,
	}
	fmt.Println("before driver assigning")
	fmt.Println(client)
	fmt.Println()

	fmt.Println("after driver assigning")
	err := client.AssignDrive(driver)
	fmt.Println(client)
	fmt.Println()
	client.ConfirmRide()
	client.GetRating(&client.Driver)
	if err != nil {
		return
	}
	fmt.Println(client.Driver)
}
