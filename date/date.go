package date

import (
	"fmt"
	"time"
)

func FormatDateOrString(input interface{}, format string) (string, error) {
	var t time.Time

	switch v := input.(type) {
	case time.Time:
		t = v
	case string:
		var err error
		t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unsupported type: %T", input)
	}

	return t.Format(format), nil
}

func format(input interface{}, format string) string {
	var t time.Time

	switch v := input.(type) {
	case time.Time:
		t = v
	case string:
		var err error
		t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return ""
		}
	default:
		return ""
	}

	return t.Format(format)
}

func ISO(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02T15:04:05.999Z")
	} else {
		return format(date[0].(time.Time), "2006-01-02T15:04:05.999Z")
	}
}

func Short(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "02/01/2006")
	} else {
		return format(date[0].(time.Time), "02/01/2006")
	}
}

func ShortInverse(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/02/2006")
	} else {
		return format(date[0].(time.Time), "01/02/2006")
	}
}

func Long(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Monday, January 2, 2006")
	} else {
		return format(date[0].(time.Time), "Monday, January 2, 2006")
	}
}

// func Custom(date interface{}, format string) string {
// 	return format(date, format)
// }

func Ok() {
	date := time.Now()

	formattedDate, err := FormatDateOrString(date, "2006-01-02")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Formatted Date:", formattedDate)

	formattedStringDate1, err := FormatDateOrString(date, "Jan 2, 2006 at 3:04pm (MST)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Formatted String Date:", formattedStringDate1)

	stringDate := "2023-11-22T15:04:05Z"
	formattedStringDate2, err := FormatDateOrString(stringDate, "Jan 2, 2006 at 3:04pm (MST)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Formatted String Date:", formattedStringDate2)
}