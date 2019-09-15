package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	parkinglot "./src/parking_lot"
)

func main() {
	ReadFile()
}

//ReadFile ... Read the incoming input file
func ReadFile() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : " + os.Args[0] + " file name")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		scanned := scanner.Text()
		inputArr = append(inputArr, scanned)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var parkingLot []parkinglot.Vehicle
	for i := 0; i < len(inputArr); i++ {
		parkingLot, err = HandleInput(parkingLot, inputArr[i], i)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//HandleInput ... Handles each line of input file
func HandleInput(parkingLot []parkinglot.Vehicle, inputArr string, i int) ([]parkinglot.Vehicle, error) {
	words := strings.Fields(inputArr)

	if words[0] == "create_parking_lot" { // reading `Create parking lot`
		if len(words) != 2 {
			return parkingLot, errors.New("input file format is incorrect")
		}
		//convert parking_lot number into integer
		i1, err := strconv.Atoi(words[1])
		if err != nil {
			return parkingLot, errors.New("Can't convert string to integer")
		}

		parkingLot = parkinglot.CreateParkingLot(0, i1)
		return parkingLot, nil

	} else if words[0] == "park" { //reading 'Add new car to parking lot'
		if len(words) != 3 {
			return parkingLot, errors.New("input file format is incorrect")
		}
		vehicle := parkinglot.Vehicle{
			RegNo:  words[1],
			Color:  words[2],
			Status: "park",
		}
		parkingLot, err := vehicle.AddToParkingLot(parkingLot)
		if err != nil {
			return parkingLot, err
		}
		return parkingLot, nil

	} else if words[0] == "leave" { //reading 'Car leaves the parking lot'
		if len(words) != 2 {
			return parkingLot, errors.New("input file format is incorrect")
		}
		//convert alloted number into integer
		i1, err := strconv.Atoi(words[1])
		if err != nil {
			return parkingLot, errors.New("Can't convert string to integer")
		}

		parkingLot, err = parkinglot.RemoveFromParkingLot(i1, parkingLot)
		if err != nil {
			fmt.Println(err)
			return parkingLot, nil
		}
		return parkingLot, nil

	} else if words[0] == "status" { //list all the cars
		if len(words) != 1 {
			return parkingLot, errors.New("input file format is incorrect")
		}
		parkinglot.ListAllVehicles(parkingLot)
		return parkingLot, nil
	} else { //list cars with specific requirement
		if len(words) != 2 {
			return parkingLot, errors.New("input file format is incorrect")
		}
		parkinglot.ListWithQuery(parkingLot, words[0], words[1])
		return parkingLot, nil
	}
}
