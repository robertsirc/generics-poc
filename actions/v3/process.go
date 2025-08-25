package v3

import (
	"fmt"
	chart "generics-poc/charts/v3"
)

func PrintChart(chart chart.Chart, process bool) bool {
	fmt.Println("~chart 3~")
	fmt.Println(process)
	fmt.Println(chart.Version)
	return true
}
