package v2

import (
	"fmt"
	chart "generics-poc/charts/v2"
)

func PrintChart(chart chart.Chart, process bool) bool {
	fmt.Println("~chart 2~")
	fmt.Println(process)
	fmt.Println(chart.Version)
	return true
}
