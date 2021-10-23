package main

import (
	"BmiCalculator/info"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(info.WelcomeMessage)
	fmt.Println(info.Separator)

	fmt.Println(info.HeightPrompt)
	var heightTextInput, _ = reader.ReadString('\n')
	var heightText = strings.TrimSuffix(heightTextInput, "\n")

	var height, heightConvError = strconv.ParseFloat(heightText, 32)
	if heightConvError != nil {
		panic(heightConvError)
	}

	fmt.Println(info.WeightPrompt)
	var weightTextInput, _ = reader.ReadString('\n')
	var weightText = strings.TrimSuffix(weightTextInput, "\n")

	var weight, weightConvError = strconv.ParseFloat(weightText, 32)
	if weightConvError != nil {
		panic(weightConvError)
	}

	var bmi = weight / (height * height)

	fmt.Printf("your bmi is: %v \n", bmi)

	fmt.Println(info.EndProgramMessage)
}
