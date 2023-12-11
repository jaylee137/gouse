package strings

import (
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

func TestRandom(t *testing.T) {
	expected := 10
	got := Random(expected)

	if len(got) != expected {
		t.Errorf("Random() = %v, want %v", len(got), expected)
	}
}
