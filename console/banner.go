package console

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/shared"
	"github.com/thuongtruong1009/gouse/strings"
)

func Banner(font shared.FontBannerType, s string) {
	split := strings.Split(strings.Uppers(s), "")
	for i := 0; i < 3; i++ {
		for _, v := range split {
			fmt.Print(font[v][i])
		}
		fmt.Print("\n")
	}
}
