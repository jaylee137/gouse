package types

import "testing"

func TestMatchRegex(t *testing.T) {
	arr := []map[string]string{
		{
			"regex": "^[a-zA-Z0-9]{3,20}$",
			"word":  "zoomer",
		},
		{
			"regex": "^[a-zA-Z0-9]{3,20}$",
			"word":  "zoomer123",
		},
	}

	for _, item := range arr {
		if !MatchRegex(item["regex"], item["word"]) {
			t.Errorf("MatchRegex() = %v, want %v", false, true)
		}
	}

}
