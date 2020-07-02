package lesson1

import "testing"

func TesttestString(t *testing.T) {
	if testString() != "olleh" {
		t.Error("err")
	}
}
