package main // declare about this file will be entry file of the project

// go standard library https://pkg.go.dev/std
import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)
}
