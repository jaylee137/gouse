package types

import (
	"reflect"
	"testing"

	"github.com/thuongtruong1009/gouse/shared"
)

func TestIsUUID(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "123e4567-e89b-12d3-a456-426614174000",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "123e4567-e89b-12d3-a456-42661417400",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "123e4567-e89b-12d3-a456-4266141740001",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 5",
			Input:   "123e4567-e89b-12d3-a456-42661417400-",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsUUID(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsUUID(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsUUID(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsGmail(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@gmail.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsGmail(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsGmail(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsGmail(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsYahoo(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@yahoo.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsYahoo(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsYahoo(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsYahoo(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsOutlook(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example@outlook.com",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsOutlook(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsOutlook(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsOutlook(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestUsername(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "example",
			Want:    true,
			WantErr: false,
		},
		{
			Name:    "Test 2",
			Input:   "example@",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 3",
			Input:   "example@.",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 4",
			Input:   "example@.com",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsUsername(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsUsername(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsUsername(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsPassword(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "12345678",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 2",
			Input:   "1234567",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsPassword(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsPassword(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsPassword(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}

func TestIsPhone(t *testing.T) {
	tests := []shared.ITest{
		{
			Name:    "Test 1",
			Input:   "12345678",
			Want:    false,
			WantErr: true,
		},
		{
			Name:    "Test 2",
			Input:   "1234567",
			Want:    false,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		isValid, err := IsPhone(tt.Input)
		if (err != nil) != tt.WantErr {
			t.Errorf("IsPhone(): %s error = %v, wantErr %v", tt.Name, err, tt.WantErr)
		}

		if (isValid != tt.Want) || !reflect.DeepEqual(isValid, tt.Want) {
			t.Errorf("IsPhone(): %s got = %v, want %v", tt.Name, isValid, tt.Want)
		}
	}
}
