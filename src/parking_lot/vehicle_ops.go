package parkinglot

import (
	"errors"
	"fmt"
)

//CreateParkingLot ... Create a parking lot of given capacity
func CreateParkingLot(len int, cap int) []Vehicle {
	parkingLot := make([]Vehicle, len, cap)
	fmt.Printf("Created a parking lot with %d slots", cap)
	fmt.Println()
	return parkingLot
}

//AddToParkingLot ... Add a Vehicle to parking lot
func (Vehicle *Vehicle) AddToParkingLot(parkingLot []Vehicle) ([]Vehicle, error) {
	// Insert if there is any "Vehicle left" space in parking lot
	for i := 0; i < len(parkingLot); i++ {
		if parkingLot[i].Status == "leave" {
			emptySpace := &parkingLot[i]
			emptySpace.RegNo = Vehicle.RegNo
			emptySpace.Color = Vehicle.Color
			emptySpace.Status = "park"
			fmt.Printf("Allocated slot number:%d", i+1)
			fmt.Println()
			return parkingLot, nil
		}
	}

	// Insert if parking lot capacity has space for a Vehicle
	if cap(parkingLot) == len(parkingLot) {
		fmt.Println("Sorry, parking lot is full")
		return parkingLot, nil
	}

	parkingLot = append(parkingLot, *Vehicle)
	fmt.Printf("Allocated slot number: %d", len(parkingLot))
	fmt.Println()
	return parkingLot, nil
}

//RemoveFromParkingLot ... 	Remove a Vehicle from parking lot
func RemoveFromParkingLot(toremove int, parkingLot []Vehicle) ([]Vehicle, error) {
	//Check if allocated no. is valid

	if toremove < cap(parkingLot) {
		//Check if Vehicle is present on allocated no.
		if parkingLot[toremove-1].Status == "park" {
			allocatedSpace := &parkingLot[toremove-1]
			allocatedSpace.Status = "leave"
			fmt.Printf("Slot number %d is free", toremove+1)
			fmt.Println()
			return parkingLot, nil
		}
		return parkingLot, errors.New("Allocated no. in parking lot has no Vehicle")

	}
	return parkingLot, errors.New("Invalid allocated number")

}

//ListAllVehicles ... list all the Vehicles available in the parking lot
func ListAllVehicles(parkingLot []Vehicle) {
	fmt.Println("Slot No.   Registration No.   Color")
	for k, v := range parkingLot {
		fmt.Printf("%d          %s      %s", k+1, v.RegNo, v.Color)
		fmt.Println()
	}
}

//ListWithQuery ... list cars with specific requirement
func ListWithQuery(parkingLot []Vehicle, query string, property string) {
	if query == "registration_numbers_for_cars_with_colour" {
		for _, v := range parkingLot {
			if v.Color == property {
				fmt.Printf("%s ", v.RegNo)
			}
		}
	} else if query == "slot_numbers_for_cars_with_colour" {
		for k, v := range parkingLot {
			if v.Color == property {
				fmt.Printf("%d ", k+1)
			}
		}
	} else if query == "slot_number_for_registration_number" {
		var status int
		for k, v := range parkingLot {
			if v.RegNo == property {
				fmt.Printf("%d ", k+1)
				status = 1
			}
		}
		if status != 1 {
			fmt.Printf("Not found")
		}
	} else {
		fmt.Printf("Not found")
	}
	fmt.Println()
}
