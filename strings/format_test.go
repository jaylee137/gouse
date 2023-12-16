package strings

import (
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"test", "e", []string{"t", "st"}},
		{"test", "t", []string{"", "es", ""}},
		{"test", "s", []string{"te", "t"}},
		{"test", "x", []string{"test"}},
		{"test", "", []string{"t", "e", "s", "t"}},
	}

	for _, test := range tests {
		actual := Split(test.input, test.sep)
		if len(actual) != len(test.expected) {
			t.Errorf("Split(%q, %q): length expected %d, length actual %d", test.input, test.sep, len(test.expected), len(actual))
		} else {
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("Split(%q, %q): expected %q, actual %q", test.input, test.sep, test.expected, actual)
				}
			}
		}
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"test", []string{"t", "e", "s", "t"}},
		{"test test", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test\ntest", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test\ttest", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test test test", []string{"t", "e", "s", "t", "t", "e", "s", "t", "t", "e", "s", "t"}},
	}

	for _, test := range tests {
		actual := Words(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("Words(%q): length expected %d, length actual %d", test.input, len(test.expected), len(actual))
		} else {
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("Words(%q): expected %q, actual %q", test.input, test.expected, actual)
				}
			}
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "tset"},
		{"test test", "tset tset"},
		{"test\ntest", "tset\ntset"},
		{"test\ttest", "tset\ttset"},
		{"test test test", "tset tset tset"},
	}

	for _, test := range tests {
		actual := Reverse(test.input)
		if actual != test.expected {
			t.Errorf("Reverse(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLowers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"TEST", "test"},
		{"test TEST", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\tTEST", "test\ttest"},
		{"test TEST test", "test test test"},
	}

	for _, test := range tests {
		actual := Lowers(test.input)
		if actual != test.expected {
			t.Errorf("Lowers(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'t', 't'},
		{'T', 't'},
		{'\n', '\n'},
		{'\t', '\t'},
		{' ', ' '},
	}

	for _, test := range tests {
		actual := Lower(test.input)
		if actual != test.expected {
			t.Errorf("Lower(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"TEST", "tEST"},
		{"test TEST", "test TEST"},
		{"test\ntest", "test\ntest"},
		{"test\tTEST", "test\tTEST"},
		{"test TEST test", "test TEST test"},
	}

	for _, test := range tests {
		actual := LowerFirst(test.input)
		if actual != test.expected {
			t.Errorf("LowerFirst(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'t', 'T'},
		{'T', 'T'},
		{'\n', '\n'},
		{'\t', '\t'},
		{' ', ' '},
	}

	for _, test := range tests {
		actual := Upper(test.input)
		if actual != test.expected {
			t.Errorf("Upper(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUppers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "TEST"},
		{"TEST", "TEST"},
		{"test TEST", "TEST TEST"},
		{"test\ntest", "TEST\nTEST"},
		{"test\tTEST", "TEST\tTEST"},
		{"test TEST test", "TEST TEST TEST"},
	}

	for _, test := range tests {
		actual := Uppers(test.input)
		if actual != test.expected {
			t.Errorf("Uppers(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"TEST", "TEST"},
		{"test TEST", "Test TEST"},
		{"test\ntest", "Test\ntest"},
		{"test\tTEST", "Test\tTEST"},
		{"test TEST test", "Test TEST test"},
	}

	for _, test := range tests {
		actual := UpperFirst(test.input)
		if actual != test.expected {
			t.Errorf("UpperFirst(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected string
	}{
		{"test", 1, "test"},
		{"test", 2, "testtest"},
		{"test", 3, "testtesttest"},
		{"test", 4, "testtesttesttest"},
		{"test", 5, "testtesttesttesttest"},
	}

	for _, test := range tests {
		actual := Repeat(test.input, test.count)
		if actual != test.expected {
			t.Errorf("Repeat(%q, %q): expected %q, actual %q", test.input, test.count, test.expected, actual)
		}
	}
}

func TestTruncate(t *testing.T) {
	testsDefault := []struct {
		input    string
		length   int
		expected string
	}{
		{"test", 1, "t..."},
		{"test", 2, "te..."},
		{"test", 3, "tes..."},
		{"test", 4, "test"},
		{"test", 5, "test"},
	}

	for _, test := range testsDefault {
		actualDefault := Truncate(test.input, test.length)
		if actualDefault != test.expected {
			t.Errorf("Truncate(%q, %q): expected %q, actual %q", test.input, test.length, test.expected, actualDefault)
		}
	}

	testsOmission := []struct {
		input    string
		length   int
		omission string
		expected string
	}{
		{"test", 1, "...", "t..."},
		{"test", 2, "@@@", "te@@@"},
		{"test", 3, "$$$", "tes$$$"},
		{"test", 4, "!!!", "test"},
		{"test", 5, "---", "test"},
	}

	for _, test := range testsOmission {
		actualOmission := Truncate(test.input, test.length, test.omission)
		if actualOmission != test.expected {
			t.Errorf("Truncate(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.omission, test.expected, actualOmission)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		input    string
		old      string
		new      string
		expected string
	}{
		{"test", "e", "a", "tast"},
		{"test", "t", "a", "aesa"},
		{"test", "s", "a", "teat"},
		{"test", "x", "a", "test"},
		{"test", "", "a", "test"},
	}

	for _, test := range tests {
		actual := Replace(test.input, test.old, test.new)
		if actual != test.expected {
			t.Errorf("Replace(%q, %q, %q): expected %q, actual %q", test.input, test.old, test.new, test.expected, actual)
		} else {
			if len(actual) != len(test.input) {
				t.Errorf("Replace(%q, %q, %q): length expected %d, length actual %d", test.input, test.old, test.new, len(test.input), len(actual))
			}
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test"},
		{" test ", "test"},
		{" test test ", "test test"},
	}

	for _, test := range tests {
		actual := Trim(test.input)
		if actual != test.expected {
			t.Errorf("Trim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test "},
		{" test ", "test "},
		{" test test ", "test test "},
	}

	for _, test := range tests {
		actual := LTrim(test.input)
		if actual != test.expected {
			t.Errorf("LTrim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestRTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", " test"},
		{"test ", "test"},
		{" test ", " test"},
		{" test test ", " test test"},
	}

	for _, test := range tests {
		actual := RTrim(test.input)
		if actual != test.expected {
			t.Errorf("RTrim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestTrimBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test"},
		{" test ", "test"},
		{" test test ", "test test"},
		{"\ntest", "test"},
		{"test\n", "test"},
		{"\ntest\n", "test"},
		{"\ntest\ntest\n", "test\ntest"},
		{"\ttest", "test"},
		{"test\t", "test"},
		{"\ttest\t", "test"},
		{"\ttest\ttest\t", "test\ttest"},
		{"\ttest\ntest\t", "test\ntest"},
		{"\ntest\ttest\n", "test\ttest"},
	}

	for _, test := range tests {
		actual := TrimBlank(test.input)
		if actual != test.expected {
			t.Errorf("TrimBlank(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestTrimPrefix(t *testing.T) {
	tests := []struct {
		input    string
		prefix   string
		expected string
	}{
		{"test", "te", "st"},
		{"test", "st", "test"},
		{"test", "x", "test"},
		{"test", "", "test"},
	}

	for _, test := range tests {
		actual := TrimPrefix(test.input, test.prefix)
		if actual != test.expected {
			t.Errorf("TrimPrefix(%q, %q): expected %q, actual %q", test.input, test.prefix, test.expected, actual)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		input    string
		suffix   string
		expected string
	}{
		{"test", "st", "te"},
		{"test", "te", "test"},
		{"test", "x", "test"},
		{"test", "", "test"},
	}

	for _, test := range tests {
		actual := TrimSuffix(test.input, test.suffix)
		if actual != test.expected {
			t.Errorf("TrimSuffix(%q, %q): expected %q, actual %q", test.input, test.suffix, test.expected, actual)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input    []string
		sep      string
		expected string
	}{
		{[]string{"test", "test"}, " ", "test test"},
		{[]string{"test", "test"}, "\n", "test\ntest"},
		{[]string{"test", "test"}, "\t", "test\ttest"},
		{[]string{"test", "test"}, "x", "testxtest"},
		{[]string{"test", "test"}, "", "testtest"},
	}

	for _, test := range tests {
		actual := Join(test.input, test.sep)
		if actual != test.expected {
			t.Errorf("Join(%q, %q): expected %q, actual %q", test.input, test.sep, test.expected, actual)
		}
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"test", "test"}, "testtest"},
		{[]string{"test", "test", "test"}, "testtesttest"},
		{[]string{"test", "test", "test", "test"}, "testtesttesttest"},
		{[]string{"test", "test", "test", "test", "test"}, "testtesttesttesttest"},
	}

	for _, test := range tests {
		actual := Concat(test.input...)
		if actual != test.expected {
			t.Errorf("Concat(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestSubStr(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		end      int
		expected string
	}{
		{"test", 0, 1, "t"},
		{"test", 0, 2, "te"},
		{"test", 0, 3, "tes"},
		{"test", 0, 4, "test"},
		{"test", 0, 5, "test"},
		{"test", 1, 1, ""},
		{"test", 1, 2, "e"},
		{"test", 1, 3, "es"},
		{"test", 1, 4, "est"},
		{"test", 1, 5, "est"},
		{"test", 2, 2, ""},
		{"test", 2, 3, "s"},
		{"test", 2, 4, "st"},
		{"test", 2, 5, "st"},
		{"test", 3, 3, ""},
		{"test", 3, 4, "t"},
		{"test", 3, 5, "t"},
		{"test", 4, 4, ""},
		{"test", 4, 5, ""},
		{"test", 5, 5, ""},
	}

	for _, test := range tests {
		actual := SubStr(test.input, test.start, test.end)
		if actual != test.expected {
			t.Errorf("SubStr(%q, %q, %q): expected %q, actual %q", test.input, test.start, test.end, test.expected, actual)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		end      int
		expected string
	}{
		{"test", 0, 1, "t"},
		{"test", 0, 2, "te"},
		{"test", 0, 3, "tes"},
		{"test", 0, 4, "test"},
		{"test", 0, 5, "test"},
		{"test", 1, 1, ""},
		{"test", 1, 2, "e"},
		{"test", 1, 3, "es"},
		{"test", 1, 4, "est"},
		{"test", 1, 5, "est"},
		{"test", 2, 2, ""},
		{"test", 2, 3, "s"},
		{"test", 2, 4, "st"},
		{"test", 2, 5, "st"},
		{"test", 3, 3, ""},
		{"test", 3, 4, "t"},
		{"test", 3, 5, "t"},
		{"test", 4, 4, ""},
		{"test", 4, 5, ""},
		{"test", 5, 5, ""},
	}

	for _, test := range tests {
		actual := Slice(test.input, test.start, test.end)
		if actual != test.expected {
			t.Errorf("Slice(%q, %q, %q): expected %q, actual %q", test.input, test.start, test.end, test.expected, actual)
		}
	}
}

func TestSplice(t *testing.T) {
	tests := []struct {
		input       string
		start       int
		deleteCount int
		items       []string
		expected    string
	}{
		{"test", 0, 1, []string{"a"}, "aest"},
		{"test", 0, 2, []string{"a"}, "atest"},
		{"test", 1, 1, []string{"a"}, "tast"},
		// {"test", 1, 2, []string{"a"}, "tast"},
		{"test", 2, 1, []string{"a", "b"}, "teabt"},
		{"test", 3, 1, []string{"a"}, "tesa"},
		{"test", 4, 1, []string{"a"}, "testa"},
		{"test", 4, 2, []string{"a"}, "testa"},
		{"test", 4, 3, []string{"a"}, "testa"},
		{"test", 4, 4, []string{"a"}, "testa"},
	}

	for _, test := range tests {
		actual := Splice(test.input, test.start, test.deleteCount, test.items...)
		if actual != test.expected {
			t.Errorf("Splice(%q, %q, %q, %q): expected %q, actual %q", test.input, test.start, test.deleteCount, test.items, test.expected, actual)
		}
	}
}

func TestEscape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test test", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\ttest", "test\ttest"},
		{"test test test", "test test test"},
	}

	for _, test := range tests {
		actual := Escape(test.input)
		if actual != test.expected {
			t.Errorf("Escape(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUnescape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test test", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\ttest", "test\ttest"},
		{"test test test", "test test test"},
	}

	for _, test := range tests {
		actual := Unescape(test.input)
		if actual != test.expected {
			t.Errorf("Unescape(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestPadStart(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		pad      byte
		expected string
	}{
		{"test", 1, 'a', "test"},
		{"test", 2, 'a', "test"},
		{"test", 3, 'a', "test"},
		{"test", 4, 'a', "test"},
		{"test", 5, 'a', "atest"},
		{"test", 6, 'a', "aatest"},
	}

	for _, test := range tests {
		actual := PadStart(test.input, test.length, test.pad)
		if actual != test.expected {
			t.Errorf("PadStart(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.pad, test.expected, actual)
		}
	}
}

func TestPadEnd(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		pad      byte
		expected string
	}{
		{"test", 1, 'a', "test"},
		{"test", 2, 'a', "test"},
		{"test", 3, 'a', "test"},
		{"test", 4, 'a', "test"},
		{"test", 5, 'a', "testa"},
		{"test", 6, 'a', "testaa"},
	}

	for _, test := range tests {
		actual := PadEnd(test.input, test.length, test.pad)
		if actual != test.expected {
			t.Errorf("PadEnd(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.pad, test.expected, actual)
		}
	}
}
