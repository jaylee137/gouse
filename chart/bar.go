package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type BarChartItem struct {
	Name   string
	Values []float64
}

type BarChartOpts struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []BarChartItem
}

func generateBarItems(values []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}

func CreateBarChart(options *BarChartOpts) {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    options.Title,
		Subtitle: options.Subtitle,
	}))

	for _, item := range options.Items {
		bar.SetXAxis(options.XAxis).
			AddSeries(item.Name, generateBarItems(item.Values))
	}

	f, _ := os.Create(options.Output)
	_ = bar.Render(f)
}
