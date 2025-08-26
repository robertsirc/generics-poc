package cmd

import (
	"fmt"
	actionsV2 "generics-poc/actions/v2"
	actionsV3 "generics-poc/actions/v3"
	v2 "generics-poc/charts/v2"
	v3 "generics-poc/charts/v3"
)

func ProcessChart(chart any, process bool) {
	fmt.Println("~Process Chart~")
	switch c := (chart).(type) {
	case v2.Chart:
		actionsV2.PrintChart(c, process)
	case v3.Chart:
		actionsV3.PrintChart(c, process)
	default:
		fmt.Println("~no chart found~")
	}
}
