package shared

import (
	"testing"
	"reflect"
)

func TestConsole(t *testing.T) {
	colors := []string{DEFAULT, RED, GREEN, YELLOW, PURPLE, PINK, CYAN, WHITE, ORANGE, BLUE, MAGENTA}

	for _, color := range colors {
		if reflect.TypeOf(color).Kind() != reflect.String {
			t.Errorf("%s is not string", color)
		} else if len(color) == 0 {
			t.Errorf("%s is empty", color)
		} else if color[0] != '\033' {
			t.Errorf("%s is not a color", color)
		} else if color[len(color)-1] != 'm' {
			t.Errorf("%s is not a color", color)
		}
	}
}