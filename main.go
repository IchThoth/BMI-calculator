package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI CALCULATOR")
	fmt.Println("-------------------")
	//output information
	fmt.Print("please enter your weight (kg): ")
	weightinput, _ := reader.ReadString('\n')

	fmt.Print("please enter your height(m): ")
	heightinput, _ := reader.ReadString('\n')

	//save user input into variables

	weightinput = strings.Replace(weightinput, "\n", "", -1)
	heightinput = strings.Replace(heightinput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightinput, 64)
	height, _ := strconv.ParseFloat(heightinput, 64)

	//calculate bmi index
	bmi := weight / (height * height)

	//output bmi
	fmt.Printf("your bmi : %.2f", bmi)
}
