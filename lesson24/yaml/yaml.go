package yaml

import (
	"fmt"
	"strings"
)

type Yaml interface {
	ReadFile(file string) ([]byte, error)
	WriteFile(file string) error
}

type yaml struct {
}

func insyaml(file string) (err error) {
	ok1 := strings.Contains(file, ".yaml")
	ok2 := strings.Contains(file, ".yml")
	if ok1 || ok2 {
		return
	}
	return fmt.Errorf("%s is yaml file", file)
}

func (y *yaml) ReadFile(file string) (data []byte, err error) {
	return
}

func (y *yaml) WriteFile(file string) error {
	fmt.Printf("write file %s", file)
	return nil
}

func NewYaml() Yaml {
	return &yaml{}
}
