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
	pl := parkinglot.CreateParkingLot(0)
	count := 0
	// if there is no input given on command line
	if len(os.Args) == 2 {
		words := strings.Split(os.Args[1], ".")
		if words[1] == "txt" {
			ProcessFile(os.Args[1])
		} else {
			fmt.Println("Invalid file format")
			return
		}
	} else {
		//Creating a Scanner that will read the input from the console

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				break
			}
			count++
			pl = ProcessCommand(pl, scanner.Text(), count)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

//ProcessCommand ... Process each command written on command line
func ProcessCommand(pl parkinglot.ParkingLot, command string, count int) parkinglot.ParkingLot {
	if count == 1 {
		commandArr := strings.Split(command, " ")
		// first, create parking lot
		i1, _ := strconv.Atoi(commandArr[1])
		pl = parkinglot.CreateParkingLot(i1)
		return pl
	}
	// performs all operations from adding car to leaving car to status
	err := HandleInput(&pl, command)
	if err != nil {
		fmt.Println(err)
		return pl
	}
	return pl
}

//ProcessFile ... Read the incoming input file
func ProcessFile(filename string) {
	file, err := os.Open(filename)
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

	pl := parkinglot.CreateParkingLot(0)
	words := strings.Fields(inputArr[0])
	if words[0] == "create_parking_lot" { // reading `Create parking lot`
		if len(words) != 2 {
			fmt.Println("input file format is incorrect")
			return
		}
		//convert parking_lot number into integer
		i1, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Println("Can't convert string to integer")
			return
		}

		pl = parkinglot.CreateParkingLot(i1)
	} else {
		fmt.Println("Create a parking lot first")
	}

	for i := 1; i < len(inputArr); i++ {
		err = HandleInput(&pl, inputArr[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

//HandleInput ... Handles each line of input file
func HandleInput(pl *parkinglot.ParkingLot, inputArr string) error {
	words := strings.Fields(inputArr)
	if words[0] == "park" { //reading 'Add new car to parking lot'
		if len(words) != 3 {
			return errors.New("input file format is incorrect")
		}
		vehicle := parkinglot.Vehicle{
			RegNo:  words[1],
			Color:  words[2],
			Status: "park",
		}
		err := pl.AddVehicle(vehicle)
		if err != nil {
			fmt.Println(err)
		}

	} else if words[0] == "leave" { //reading 'Car leaves the parking lot'
		if len(words) != 2 {
			return errors.New("input file format is incorrect")
		}
		//convert alloted number into integer
		i1, err := strconv.Atoi(words[1])
		if err != nil {
			return errors.New("Can't convert string to integer")
		}

		pl.RemoveVehicle(i1)

	} else if words[0] == "status" { //list all the cars
		if len(words) != 1 {
			return errors.New("input file format is incorrect")
		}
		pl.ListAllVehicles()
	} else { //list cars with specific requirement
		if len(words) != 2 {
			return errors.New("input file format is incorrect")
		}
		pl.ListWithQuery(words[0], words[1])
	}
	return nil
}
