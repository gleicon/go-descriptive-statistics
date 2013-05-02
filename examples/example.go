package main

import (
	"fmt"
	"github.com/gleicon/go-descriptive-statistics/descriptive-statistics"
)

var e = descriptive_statistics.Enum{2, 6, 9, 3, 5, 1, 8, 3, 6, 9, 2}

func main() {
	fmt.Println("List size: ", e.Len())
	fmt.Println("List Mean value: ", e.Mean())
	fmt.Println("List Median value: ", e.Median())
	fmt.Println("List Variance: ", e.Variance())
	fmt.Println("List Standard Deviation: ", e.StandardDeviation())
	fmt.Println("List 90% Percentile: ", e.Percentile(90.0))
}
