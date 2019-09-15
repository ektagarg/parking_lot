package main

import (
	"bufio"
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
		words := strings.Fields(inputArr[i])

		//parkingLot := make([]parkinglot.Vehicle, 0)

		if words[0] == "create_parking_lot" { // reading `Create parking lot`
			if len(words) != 2 {
				fmt.Println("Input file format is incorrect!")
			} else {
				//convert parking_lot number into integer
				i1, err := strconv.Atoi(words[1])
				if err != nil {
					fmt.Println("Can't convert string to integer")
					return
				}

				parkingLot = parkinglot.CreateParkingLot(0, i1)

			}
		} else if words[0] == "park" { //reading 'Add new car to parking lot'
			//fmt.Println(cap(parkingLot))
			if len(words) != 3 {
				fmt.Println("Input file format is incorrect!")
			} else {
				vehicle := parkinglot.Vehicle{
					RegNo:  words[1],
					Color:  words[2],
					Status: "park",
				}
				parkingLot, err = vehicle.AddToParkingLot(parkingLot)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else if words[0] == "leave" { //reading 'Car leaves the parking lot'
			if len(words) != 2 {
				fmt.Println("Input file format is incorrect!")
				return
			}
			//convert alloted number into integer
			i1, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println("Can't convert string to integer")
				return
			}

			parkingLot, err = parkinglot.RemoveFromParkingLot(i1, parkingLot)
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if words[0] == "status" { //list all the cars
			if len(words) != 1 {
				fmt.Println("Input file format is incorrect!")
				return
			}
			parkinglot.ListAllVehicles(parkingLot)
		} else { //list cars with specific requirement
			if len(words) != 2 {
				fmt.Println("Input file format is incorrect!")
				return
			}

			parkinglot.ListWithQuery(parkingLot, words[0], words[1])

		}
	}
}

func HandleInput(inputLine string) {

}
