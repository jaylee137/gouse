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

func Ok() {
	date := time.Now()
  
	formattedDate, err := FormatDateOrString(date, "2006-01-02")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Formatted Date:", formattedDate)

	stringDate := "2023-11-22T15:04:05Z"
	formattedStringDate, err := FormatDateOrString(stringDate, "Jan 2, 2006 at 3:04pm (MST)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Formatted String Date:", formattedStringDate)
}