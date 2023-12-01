package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
)

func main() {
	content, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Error has occurred")
		return
	}
	fmt.Println(content)

	f, err := SqrtErr(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.3f", f)

}

// SqrtErr Making custom error
func SqrtErr(number int) (float64, error) {
	if number < 0 {
		return 0, errors.New("the number is not valid")
	}
	fmt.Println("without error")
	return math.Sqrt(float64(number)), nil
}
