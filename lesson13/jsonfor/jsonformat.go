package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string
	Sex  string
	Age  int
}

type Woman struct {
	Sanwei  string
	Peoples []*People
}

func testJson() {
	var buty *Woman = &Woman{
		Sanwei: "28,35,68",
	}
	for i := 0; i < 10; i++ {
		pe := &People{
			Name: fmt.Sprintf("meinv%d", i),
			Sex:  "women",
			Age:  18,
		}
		buty.Peoples = append(buty.Peoples, pe)
	}

	data, err := json.Marshal(buty)
	if err == nil {
		fmt.Printf("%s", string(data))
	}
}

var Tostring string = `{"Sanwei":"28,35,68","Peoples":[{"Name":"meinv0","Sex":"women","Age":18},{"Name":"meinv1","Sex":"women","Age":18},{"Name":"meinv2","Sex":"women","Age":18},{"Name":"meinv3","Sex":"women","Age":18},{"Name":"meinv4","Sex":"women","Age":18},{"Name":"meinv5","Sex":"women","Age":18},{"Name":"meinv6","Sex":"women","Age":18},{"Name":"meinv7","Sex":"women","Age":18},{"Name":"meinv8","Sex":"women","Age":18},{"Name":"meinv9","Sex":"women","Age":18}]}`

func jsontostring() {
	var c1 *Woman = &Woman{}
	err := json.Unmarshal([]byte(Tostring), c1)
	if err != nil {
		return
	}
	fmt.Printf("%#v", c1)
}

func main() {
	// testJson()
	jsontostring()
}
