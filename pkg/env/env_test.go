package env

import (
	"testing"
)

func TestNewStringVar(t *testing.T) {
	if NewStringVar("test1", "desc1", "default1") != nil {
		t.Error()
	}
	v := Get("test1")
	if v.Name != "test1" || v.Description != "desc1" || v.DefaultValue != "default1" || v.Type != STRING {
		t.Error()
	}
}

func TestNewIntVar(t *testing.T) {
	if NewIntVar("test2", "desc2", "default2") != nil {
		t.Error()
	}
	v := Get("test2")
	if v.Name != "test2" || v.Description != "desc2" || v.DefaultValue != "default2" || v.Type != INT {
		t.Error()
	}
}

func TestNewBoolVar(t *testing.T) {
	if NewBoolVar("test3", "desc3", "default3") != nil {
		t.Error()
	}
	v := Get("test3")
	if v.Name != "test3" || v.Description != "desc3" || v.DefaultValue != "default3" || v.Type != BOOL {
		t.Error()
	}
}
