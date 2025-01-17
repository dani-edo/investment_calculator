package main // declare about this file will be entry file of the project

// go standard library https://pkg.go.dev/std
import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount := getUserInput("Investment Amount: ")
	years := getUserInput("Years: ")
	expectedReturnRate := getUserInput("Expected Return Rate: ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)
	message := fmt.Sprintf(`Future Value: %.5f, Future Real Value: %.5f`, futureValue, futureRealValue)
	fmt.Printf(message)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue = futureValue / math.Pow((1+inflationRate/100), years)
	return futureValue, futureRealValue
}

func getUserInput(text string) (result float64) {
	fmt.Print(text)
	fmt.Scan(&result)
	return
}
