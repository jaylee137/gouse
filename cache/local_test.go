package cache

import (
	"testing"
)

func TestGetLocal(t *testing.T) {
	c := NewLocal()
	c.SetLocal("test", "test")
	if result, err := c.GetLocal("test"); result != "test" {
		t.Error("GetLocal error", err)
	}
}

func TestSetLocal(t *testing.T) {
	c := NewLocal()
	c.SetLocal("test", "test")
	if result, err := c.GetLocal("test"); result != "test" {
		t.Error("SetLocal error", err)
	}
}

func TestDelLocal(t *testing.T) {
	c := NewLocal()
	c.SetLocal("test", "test")
	c.DelLocal("test")
	if result, err := c.GetLocal("test"); result != "" {
		t.Error("DelLocal error", err)
	}
}

func TestFlushLocal(t *testing.T) {
	c := NewLocal()
	c.SetLocal("test", "test")
	c.FlushLocal()
	if result, err := c.GetLocal("test"); result != "" {
		t.Error("FlushLocal error", err)
	}
}

func TestAllLocal(t *testing.T) {
	c := NewLocal()
	c.SetLocal("test", "test")
	if c.AllLocal()["test"] != "test" {
		t.Error("AllLocal error")
	}
}
