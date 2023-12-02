package constants

import (
	"testing"
	"reflect"
)

func TestConst(t *testing.T) {
	if len(KeyStr) != 94 {
		t.Errorf("KeyStr length is not equal to %d", 94)
	}

	if reflect.TypeOf(KeyStr).Kind() != reflect.String {
		t.Errorf("KeyStr is not string")
	}

	if KeyStr != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]|:<>?/.,';][=-`~" {
		t.Errorf("KeyStr is not correct")
	}

	if KeyStr[0] != 'a' {
		t.Errorf("KeyStr[0] is not correct")
	}

	if KeyStr[len(KeyStr)-1] != '~' {
		t.Errorf("KeyStr[%d] is not correct", len(KeyStr)-1)
	}
}