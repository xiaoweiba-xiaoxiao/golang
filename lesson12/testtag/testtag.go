package testtag

type Testtag struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

func New(name, sex string, age int) *Testtag {
	var testtag *Testtag = &Testtag{
		Name: name,
		Sex:  sex,
		Age:  age,
	}
	return testtag
}
