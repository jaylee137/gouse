package strings

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Count 1",
			args: args{
				s:      "Hello World",
				substr: "",
			},
			want: 11,
		},
		{
			name: "Count 2",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 3,
		},
		{
			name: "Count 3",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		var got int

		if tt.args.substr == "" {
			got = Count(tt.args.s)
		} else {
			got = Count(tt.args.s, tt.args.substr)
		}

		if got != tt.want {
			t.Errorf("Count() = %v, want %v", got, tt.want)
		}
	}
}

func TestLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Lines 1",
			args: args{
				s: "Hello World",
			},
			want: 1,
		},
		{
			name: "Lines 2",
			args: args{
				s: "Hello\nWorld",
			},
			want: 2,
		},
		{
			name: "Lines 3",
			args: args{
				s: "Hello\nWorld\n",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := Lines(tt.args.s); got != tt.want {
			t.Errorf("Lines() = %v, want %v", got, tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name      string
		args      args
		wantFirst int
		wantLast  int
	}{
		{
			name: "Index 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			wantFirst: 2,
			wantLast:  9,
		},
		{
			name: "Index 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			wantFirst: 2,
			wantLast:  2,
		},
		{
			name: "Index 3",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			wantFirst: -1,
			wantLast:  -1,
		},
		{
			name: "Index 4",
			args: args{
				s:      "Hello World",
				substr: " ",
			},
			wantFirst: 5,
			wantLast:  5,
		},
		{
			name: "Index 5",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			wantFirst: 6,
			wantLast:  6,
		},
	}

	for _, tt := range tests {
		gotFirst, gotLast := Index(tt.args.s, tt.args.substr)

		if gotFirst != tt.wantFirst {
			t.Errorf("Index() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
		}

		if gotLast != tt.wantLast {
			t.Errorf("Index() gotLast = %v, want %v", gotLast, tt.wantLast)
		}
	}
}

func TestFIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests1 := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FIndex 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 2,
		},
		{
			name: "FIndex 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 2,
		},
		{
			name: "FIndex 4",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			want: -1,
		},
		{
			name: "FIndex 5",
			args: args{
				s:      "Hello World",
				substr: " ",
			},
			want: 5,
		},
		{
			name: "FIndex 9",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			want: 6,
		},
	}

	for _, tt := range tests1 {
		if got := FIndex(tt.args.s, tt.args.substr); got != tt.want {
			t.Errorf("FIndex() = %v, want %v", got, tt.want)
		}
	}
}

func TestLIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests2 := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LIndex 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 9,
		},
		{
			name: "LIndex 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 2,
		},
		{
			name: "LIndex 4",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			want: -1,
		},
		{
			name: "LIndex 9",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			want: 6,
		},
	}

	for _, tt := range tests2 {
		got := LIndex(tt.args.s, tt.args.substr)

		if got != tt.want {
			t.Errorf("LIndex() = %v, want %v", got, tt.want)
		}
	}
}

func TestRandomStr(t *testing.T) {
	expectedLength := 10
	got := RandomStr(expectedLength)

	if len(got) != expectedLength {
		t.Errorf("RandomStr() = %v, want %v", len(got), expectedLength)
	}

	if reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("RandomStr() = %v, want %v", reflect.TypeOf(got).Kind(), reflect.String)
	}
}

func TestRandomNum(t *testing.T) {
	expectedLength := 6
	got := RandomNum(expectedLength)

	if len(got) != expectedLength {
		t.Errorf("RandomNum() = %v, want %v", len(got), expectedLength)
	}

	if reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("RandomNum() = %v, want %v", reflect.TypeOf(got).Kind(), reflect.String)
	}
}

func TestAt(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "At 1",
			args: args{
				s: "Hello World",
				i: 0,
			},
			want: "H",
		},
		{
			name: "At 2",
			args: args{
				s: "Hello World",
				i: 1,
			},
			want: "e",
		},
		{
			name: "At 3",
			args: args{
				s: "Hello World",
				i: 2,
			},
			want: "l",
		},
		{
			name: "At 4",
			args: args{
				s: "Hello World",
				i: 3,
			},
			want: "l",
		},
		{
			name: "At 5",
			args: args{
				s: "Hello World",
				i: 4,
			},
			want: "o",
		},
		{
			name: "At 6",
			args: args{
				s: "Hello World",
				i: 5,
			},
			want: " ",
		},
		{
			name: "At 7",
			args: args{
				s: "Hello World",
				i: 6,
			},
			want: "W",
		},
		{
			name: "At 8",
			args: args{
				s: "Hello World",
				i: 7,
			},
			want: "o",
		},
		{
			name: "At 9",
			args: args{
				s: "Hello World",
				i: 8,
			},
			want: "r",
		},
		{
			name: "At 10",
			args: args{
				s: "Hello World",
				i: 9,
			},
			want: "l",
		},
		{
			name: "At 11",
			args: args{
				s: "Hello World",
				i: 10,
			},
			want: "d",
		},
	}

	for _, tt := range tests {
		if got := At(tt.args.s, tt.args.i); got != tt.want {
			t.Errorf("At() = %v, want %v", got, tt.want)
		}
	}
}

func TestCodePointAt(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CodePointAt 1",
			args: args{
				s: "Hello World",
				i: 0,
			},
			want: 72,
		},
		{
			name: "CodePointAt 2",
			args: args{
				s: "Hello World",
				i: 1,
			},
			want: 101,
		},
		{
			name: "CodePointAt 3",
			args: args{
				s: "Hello World",
				i: 2,
			},
			want: 108,
		},
		{
			name: "CodePointAt 4",
			args: args{
				s: "Hello World",
				i: 3,
			},
			want: 108,
		},
		{
			name: "CodePointAt 5",
			args: args{
				s: "Hello World",
				i: 4,
			},
			want: 111,
		},
		{
			name: "CodePointAt 6",
			args: args{
				s: "Hello World",
				i: 5,
			},
			want: 32,
		},
		{
			name: "CodePointAt 7",
			args: args{
				s: "Hello World",
				i: 6,
			},
			want: 87,
		},
		{
			name: "CodePointAt 8",
			args: args{
				s: "Hello World",
				i: 7,
			},
			want: 111,
		},
		{
			name: "CodePointAt 9",
			args: args{
				s: "Hello World",
				i: 8,
			},
			want: 114,
		},
		{
			name: "CodePointAt 10",
			args: args{
				s: "Hello World",
				i: 9,
			},
			want: 108,
		},
		{
			name: "CodePointAt 11",
			args: args{
				s: "Hello World",
				i: 10,
			},
			want: 100,
		},
	}

	for _, tt := range tests {
		if got := CodePointAt(tt.args.s, tt.args.i); got != tt.want {
			t.Errorf("CodePointAt() = %v, want %v", got, tt.want)
		}
	}
}

func TestCodePoint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CodePoint 1",
			args: args{
				s: "Hello World",
			},
			want: []int{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
		},
	}

	for _, tt := range tests {
		got := CodePoint(tt.args.s)
		if len(got) != len(tt.want) {
			t.Errorf("CodePoint() = %v, want %v", len(got), len(tt.want))
		} else if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CodePoint() = %v, want %v", got, tt.want)
		}
	}
}

func TestFromCodePointAt(t *testing.T) {
	type args struct {
		codePoint int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FromCodePointAt 1",
			args: args{
				codePoint: 72,
			},
			want: "H",
		},
		{
			name: "FromCodePointAt 2",
			args: args{
				codePoint: 101,
			},
			want: "e",
		},
		{
			name: "FromCodePointAt 3",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 4",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 5",
			args: args{
				codePoint: 111,
			},
			want: "o",
		},
		{
			name: "FromCodePointAt 6",
			args: args{
				codePoint: 32,
			},
			want: " ",
		},
		{
			name: "FromCodePointAt 7",
			args: args{
				codePoint: 87,
			},
			want: "W",
		},
		{
			name: "FromCodePointAt 8",
			args: args{
				codePoint: 111,
			},
			want: "o",
		},
		{
			name: "FromCodePointAt 9",
			args: args{
				codePoint: 114,
			},
			want: "r",
		},
		{
			name: "FromCodePointAt 10",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 11",
			args: args{
				codePoint: 100,
			},
			want: "d",
		},
	}

	for _, tt := range tests {
		if got := FromCodePointAt(tt.args.codePoint); got != tt.want {
			t.Errorf("FromCodePointAt() = %v, want %v", got, tt.want)
		}
	}
}

func TestFromCodePoint(t *testing.T) {
	type args struct {
		codePoint []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FromCodePoint 1",
			args: args{
				codePoint: []int{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
			},
			want: "Hello World",
		},
	}

	for _, tt := range tests {
		if got := FromCodePoint(tt.args.codePoint...); got != tt.want {
			t.Errorf("FromCodePoint() = %v, want %v", got, tt.want)
		}
	}
}
