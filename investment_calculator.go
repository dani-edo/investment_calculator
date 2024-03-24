package main // declare about this file will be entry file of the project

// go standard library https://pkg.go.dev/std
import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println(futureValue)
}
