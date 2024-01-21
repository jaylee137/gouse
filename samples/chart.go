package samples

import (
	"github.com/thuongtruong1009/gouse/chart"
)

func chartBar() {
	newChart := &chart.BarChartOpts{
		Title:    "Bar chart in Go",
		Subtitle: "This is fun to use!",
		XAxis:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		Nums:     6,
		Items: []chart.BarChartItem{
			{Name: "Category A", Value: 200},
			{Name: "Category B", Value: 300},
			{Name: "Category C", Value: 400},
		},
	}

	chart.CreateBarChart(newChart)
}
