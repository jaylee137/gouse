package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineChartItem struct {
	Name   string
	Values []float64
}

type LineChartOpts struct {
	Output   string
	Title    string
	Subtitle string
	XAxis    []string
	Items    []LineChartItem
}

func generateLineItems(values []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.LineData{Value: values[i]})
	}
	return items
}

func CreateLineChart(options *LineChartOpts) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    options.Title,
			Subtitle: options.Subtitle,
		}),
	)

	// line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
	// 	AddSeries("Category A", generateLineItems()).
	// 	AddSeries("Category B", generateLineItems()).
	// 	SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	for _, item := range options.Items {
		line.SetXAxis(options.XAxis).
			AddSeries(item.Name, generateLineItems(item.Values)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	}

	f, _ := os.Create(options.Output)
	_ = line.Render(f)
}
