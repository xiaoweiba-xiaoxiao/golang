package yaml

import (
	"testing"
)

func TestYaml(t *testing.T) {
	file := "test.yaml"
	yaml1 := NewYaml()
	yaml1.ReadFile(file)
	yaml1.WriteFile(file)
}
