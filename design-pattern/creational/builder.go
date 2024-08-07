package creational

import "fmt"

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// A builder of type car
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// A builder of type car
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// using
func runBuilder() {
	manufacoringDirector := &ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacoringDirector.SetBuilder(carBuilder)
	manufacoringDirector.Construct()

	car := carBuilder.GetVehicle()

	fmt.Printf("%v\n", car)

	bikeBuilder := &BikeBuilder{}
	manufacoringDirector.Construct()
	manufacoringDirector.SetBuilder(bikeBuilder)

	bike := bikeBuilder.GetVehicle()

	fmt.Printf("%v\n", bike)
}
