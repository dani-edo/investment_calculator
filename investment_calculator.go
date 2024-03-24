package main // declare about this file will be entry file of the project

// go standard library https://pkg.go.dev/std
import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)
}
