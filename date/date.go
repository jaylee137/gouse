package date

import "time"

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
		return format(date[0], "2006-01-02T15:04:05.999Z")
	}
}

func Short(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "02/01/2006")
	} else {
		return format(date[0].(time.Time), "02/01/2006")
	}
}

func ShortNormal(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/02/2006")
	} else {
		return format(date[0].(time.Time), "01/02/2006")
	}
}

func ShortReverse(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006/01/02")
	} else {
		return format(date[0].(time.Time), "2006/01/02")
	}
}

func ShortDash(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02")
	} else {
		return format(date[0].(time.Time), "2006-01-02")
	}
}

func ShortDot(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006.01.02")
	} else {
		return format(date[0].(time.Time), "2006.01.02")
	}
}

func ShortUnderscore(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006_01_02")
	} else {
		return format(date[0].(time.Time), "2006_01_02")
	}
}

func ShortSpace(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006 01 02")
	} else {
		return format(date[0].(time.Time), "2006 01 02")
	}
}

func ShortMonth(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/2006")
	} else {
		return format(date[0].(time.Time), "01/2006")
	}
}

func Long(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Monday, January 2, 2006")
	} else {
		return format(date[0].(time.Time), "Monday, January 2, 2006")
	}
}

func UTC(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Jan 2, 2006 at 3:04pm (MST)")
	} else {
		return format(date[0].(time.Time), "Jan 2, 2006 at 3:04pm (MST)")
	}
}
