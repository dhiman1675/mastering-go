package main

import (
	"fmt"
	"math"
	"os"
)

const inflationRate = 3.5

func main() {
	// var investmentAmount = 10_000
	// var expectedReturnRate = 5.5
	// var years = 10

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// fmt.Println(futureValue)

	// The repeated use of float64 is increasing the verbosity of the code.
	// to fix this we can add explicit type annotation

	// Different ways of declaring variables
	// var investmentAmount, years float64 = 10_000, 10
	// investmentAmount, years := 10_000.0, 10.0

	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Please enter your investment amount: ")
	if _, err := fmt.Scanln(&investmentAmount); err != nil {
		fmt.Println("❌ Invalid input for investment amount. Please enter a number.")
		os.Exit(1)
	}

	fmt.Print("Enter expected return rate (%): ")
	if _, err := fmt.Scanln(&expectedReturnRate); err != nil {
		fmt.Println("❌ Invalid input for return rate. Please enter a number.")
		os.Exit(1)
	}

	fmt.Print("Enter number of years: ")
	if _, err := fmt.Scanln(&years); err != nil {
		fmt.Println("❌ Invalid input for years. Please enter a number.")
		os.Exit(1)
	}

	futureValue, futureRealValue := calcFutureValue(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Println(`Future Value (A/C Inflation):

	`, futureRealValue)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	return futureValue, futureValue / math.Pow(1+inflationRate/100, years)
}
