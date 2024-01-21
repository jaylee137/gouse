package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type BarChartItem struct {
	Name  string
	Value float64
}

type BarChartOpts struct {
	Title    string
	Subtitle string
	XAxis    []string
	Nums     int
	Items    []BarChartItem
}

func (opts *BarChartOpts) GetNums() int {
	return opts.Nums
}

func generateBarItems(nums int, val float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < nums; i++ {
		items = append(items, opts.BarData{Value: val})
	}
	return items
}

func CreateBarChart(options *BarChartOpts) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    options.Title,
		Subtitle: options.Subtitle,
	}))

	// bar.SetXAxis(options.XAxis).
	// 	AddSeries("Category A", generateBarItems()).
	// 	AddSeries("Category B", generateBarItems())

	for _, item := range options.Items {
		bar.SetXAxis(options.XAxis).
			AddSeries(item.Name, generateBarItems(options.Nums, item.Value))
	}

	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}
