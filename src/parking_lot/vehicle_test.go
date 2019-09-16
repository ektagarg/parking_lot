package parkinglot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = CreateParkingLot(3)

//TestCreateParkingLot ... Check if parking lot has the given capacity
func TestCreateParkingLot(t *testing.T) {
	//p = CreateParkingLot(3)
	assert.Equal(t, 3, cap(p.Vehicles), "Capacity of parkinglot in invalid")
}

//TestAddVehicle ... Add a vehicle to lot
func TestAddVehicle(t *testing.T) {
	v1 := Vehicle{
		RegNo:  "RJ327E384",
		Color:  "White",
		Status: "park",
	}

	lenVehicles := len(p.Vehicles)
	p.AddVehicle(v1)
	assert.Equal(t, len(p.Vehicles), lenVehicles+1, "Length of vehicles is not increased")
}

//TestAddVehicles ... Add multiple vehicles to lot
func TestAddVehicles(t *testing.T) {
	v2 := Vehicle{
		RegNo:  "RJ3227E384",
		Color:  "Black",
		Status: "park",
	}
	v3 := Vehicle{
		RegNo:  "RJ32deeE384",
		Color:  "Pink",
		Status: "park",
	}
	v4 := Vehicle{
		RegNo:  "RJ32d6eE384",
		Color:  "White",
		Status: "park",
	}
	p.AddVehicle(v2)
	p.AddVehicle(v3)
	p.AddVehicle(v4)
	p.AddVehicle(v4)
	p.AddVehicle(v4)
	assert.Equal(t, len(p.Vehicles), p.MaxCap, "Length of vehicles exceeded in parking lot")
}

//TestRemoveVehicle ... Remove vehicle from the lot
func TestRemoveVehicle(t *testing.T) {
	p.RemoveVehicle(2)
	assert.Equal(t, p.Vehicles[1].Status, "leave", "Vehicle status didn't get removed")
}

//TestListVehicles .... List all the vehicles
func TestListVehicles(t *testing.T) {
	p.ListAllVehicles()
}

//TestListWithQuery ... List vehicles as per the query
func TestListWithQuery(t *testing.T) {
	p.ListWithQuery("registration_numbers_for_cars_with_colour", "White")
	var vehicleNum string
	for _, v := range p.Vehicles {
		if v.Color == "White" {
			vehicleNum = v.RegNo
		}
	}
	assert.Equal(t, vehicleNum, "RJ327E384", "Vehicle regNo. is incorrect")
}
