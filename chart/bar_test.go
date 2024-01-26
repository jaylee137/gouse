package chart

import (
	"os"
	"testing"
)

func TestCreateBarChart(t *testing.T) {
	// Prepare test data
	options := &BarChartOpts{
		Output:   "bar_test.html",
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		XAxis:    []string{"A", "B", "C"},
		Items: []BarChartItem{
			{
				Name:   "Series1",
				Values: []float64{1.0, 2.0, 3.0},
			},
			{
				Name:   "Series2",
				Values: []float64{4.0, 5.0, 6.0},
			},
		},
	}

	// Run the function
	CreateBarChart(options)

	// Verify the output file
	if _, err := os.Stat("bar_test.html"); os.IsNotExist(err) {
		t.Errorf("Expected output file 'bar_test.html' not found")
	}

	// Clean up the test file
	_ = os.Remove("bar_test.html")
}

func TestGenerateBarItems(t *testing.T) {
	// Test with sample values
	values := []float64{1.0, 2.0, 3.0}
	result := generateBarItems(values)

	// Verify the result
	if len(result) != len(values) {
		t.Errorf("Expected %d items, but got %d", len(values), len(result))
	}

	// Verify each item's value
	for i, v := range values {
		if result[i].Value != v {
			t.Errorf("Expected value %.2f at index %d, but got %.2f", v, i, result[i].Value)
		}
	}
}
