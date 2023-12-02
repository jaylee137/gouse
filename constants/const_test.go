package constants

import (
	"testing"
	"reflect"
)

func TestConst(t *testing.T) {
	if(len(KeyStr) != 62) {
		t.Errorf("KeyStr length is not 62")
	}

	if(reflect.TypeOf(KeyStr).Kind() != reflect.String) {
		t.Errorf("KeyStr is not string")
	}

	if(KeyStr != "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") {
		t.Errorf("KeyStr is not correct")
	}

	if(KeyStr[0] != 'a') {
		t.Errorf("KeyStr[0] is not correct")
	}

	if(KeyStr[61] != '9') {
		t.Errorf("KeyStr[61] is not correct")
	}
}