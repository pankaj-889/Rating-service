package model

import "errors"

type Client struct {
	Id          int
	Name        string
	Rating      int
	Driver      interface{} `json:"driver,nil"`
	ConfirmRide bool        `json:"confirmRide,omitempty"`
}

func NewClient(client Client) *Client {
	return &Client{
		Id:          client.Id,
		Name:        client.Name,
		Rating:      client.Rating,
		Driver:      nil,
		ConfirmRide: false,
	}
}

func (client *Client) AssignDrive(driver Driver) error {
	if client.Driver == nil {
		client.Driver = driver
		client.ConfirmRide = true
		return nil
	}

	return errors.New("driver is already assigned")
}

func (c *Client) CalculateRating(rating float32) {
	status := c.ConfirmRide
	if status == true {
		driver := c.Driver.(Driver)
		currentPoint := driver.Rating * float32(driver.TotalRide)
		totalRide := driver.TotalRide + 1
		currentRating := (currentPoint + rating) / float32(totalRide*1.0)

		driver.Rating = currentRating
		driver.TotalRide = driver.TotalRide + 1
		c.Driver = driver
	}
}
