package model

type Driver struct {
	Id        int
	Name      string
	Rating    float32
	TotalRide int
}

func Newdriver(driver Driver) *Driver {
	return &Driver{
		Id:        driver.Id,
		Name:      driver.Name,
		Rating:    driver.Rating,
		TotalRide: driver.TotalRide,
	}
}

func (c *Driver) CalculateRating() {

}

func (c *Driver) GetDriver() *Driver {
	return c
}
