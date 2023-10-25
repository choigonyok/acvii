package config

import (
	"testing"

	"github.com/choigonyok/goopt/pkg/env"
)

func TestInit(t *testing.T) {
	if env.Get("ACVCTL").Name != "ACVCTL" {
		t.Error()
	}
}

func TestInitConfigFromViper(t *testing.T) {
	if InitConfigFromViper() != nil {
		t.Error()
	}
}
