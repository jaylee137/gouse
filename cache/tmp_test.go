package cache

import (
	"testing"
	"time"
)

func TestGetTmp(t *testing.T) {
	tmp := NewTmp(1 * time.Hour)
	tmp.SetTmp("test", "test", 0)
	if tmp.GetTmp("test") != "test" {
		t.Error("TestGetTmp failed")
	}
}

func TestSetTmp(t *testing.T) {
	tmp := NewTmp(1 * time.Hour)
	tmp.SetTmp("test", "test", 0)
	if tmp.GetTmp("test") != "test" {
		t.Error("TestSetTmp failed")
	}
}

func TestDelTmp(t *testing.T) {
	tmp := NewTmp(1 * time.Hour)
	tmp.SetTmp("test", "test", 0)
	tmp.DelTmp("test")
	if tmp.GetTmp("test") != nil {
		t.Error("TestDelTmp failed")
	}
}

func TestFlushTmp(t *testing.T) {
	tmp := NewTmp(1 * time.Hour)
	tmp.SetTmp("test", "test", 0)
	tmp.FlushTmp()
	if tmp.GetTmp("test") != nil {
		t.Error("TestFlushTmp failed")
	}
}

func TestAllTmp(t *testing.T) {
	tmp := NewTmp(1 * time.Hour)
	tmp.SetTmp("test", "test", 0)
	if tmp.AllTmp()["test"] != "test" {
		t.Error("TestAllTmp failed")
	}
}
