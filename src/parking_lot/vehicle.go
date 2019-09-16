package parkinglot

import (
	"errors"
	"fmt"
)

//Vehicle ... details of a vehicle
type Vehicle struct {
	RegNo  string
	Color  string
	Status string
}

//ParkingLot ... list of all vehicles and capacity of lot
type ParkingLot struct {
	Vehicles []Vehicle
	MaxCap   int
}

//CreateParkingLot ... Create a parking lot of specific size
func CreateParkingLot(cap int) ParkingLot {
	p := ParkingLot{
		Vehicles: make([]Vehicle, 0, cap),
		MaxCap:   cap,
	}
	if cap > 0 {
		fmt.Printf("Created a parking lot with %d slots", cap)
		fmt.Println()
	}
	return p
}

//AddVehicle ... Case when a vehicle comes to parking lot
func (p *ParkingLot) AddVehicle(v Vehicle) error {
	// Insert if there is any "Vehicle left" space in parking lot
	for i := 0; i < len(p.Vehicles); i++ {
		if p.Vehicles[i].Status == "leave" {
			emptySpace := &p.Vehicles[i]
			emptySpace.RegNo = v.RegNo
			emptySpace.Color = v.Color
			emptySpace.Status = "park"
			fmt.Printf("Allocated slot number:%d", i+1)
			fmt.Println()
			return nil
		}
	}

	// Insert if parking lot capacity has space for a Vehicle
	if len(p.Vehicles) == p.MaxCap {
		return errors.New("Sorry, parking lot is full")
	}

	p.Vehicles = append(p.Vehicles, v)
	//parkingLot = append(parkingLot, *Vehicle)
	fmt.Printf("Allocated slot number: %d", len(p.Vehicles))
	fmt.Println()
	return nil
}

//RemoveVehicle ... Case when a vehicle leaves parking lot
func (p *ParkingLot) RemoveVehicle(toremove int) {
	//Check if allocated no. is valid
	if toremove < p.MaxCap {
		//Check if Vehicle is present on allocated no.
		if p.Vehicles[toremove-1].Status == "park" {
			allocatedSpace := &p.Vehicles[toremove-1]
			allocatedSpace.Status = "leave"
			fmt.Printf("Slot number %d is free", toremove)
			fmt.Println()
		} else {
			fmt.Println("Allocated no. in parking lot has no parked Vehicle")
			return
		}
	} else {
		fmt.Println("Invalid allocated number")
	}
}

//ListAllVehicles ... list all the Vehicles available in the parking lot
func (p *ParkingLot) ListAllVehicles() {
	fmt.Println("Slot No.   Registration No.   Color")
	for k, v := range p.Vehicles {
		if v.Status == "park" {
			fmt.Printf("%d          %s      %s", k+1, v.RegNo, v.Color)
			fmt.Println()
		}
	}
}

//ListWithQuery ... list cars with specific requirement
func (p *ParkingLot) ListWithQuery(query string, property string) {
	if query == "registration_numbers_for_cars_with_colour" {
		for _, v := range p.Vehicles {
			if v.Color == property {
				fmt.Printf("%s ", v.RegNo)
			}
		}
	} else if query == "slot_numbers_for_cars_with_colour" {
		for k, v := range p.Vehicles {
			if v.Color == property {
				fmt.Printf("%d ", k+1)
			}
		}
	} else if query == "slot_number_for_registration_number" {
		var status int
		for k, v := range p.Vehicles {
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
