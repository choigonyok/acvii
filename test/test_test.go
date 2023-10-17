package test_test

import (
	"testing"

	"github.com/choigonyok/goopt/internal/test"
)

func TestTest(t *testing.T) {
	if test.Test() != "TEST" {
		t.Error("TEST FAILED")
	}
}
