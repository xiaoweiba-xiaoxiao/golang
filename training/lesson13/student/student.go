package student

type Student struct {
	Name  string  `json:"name"`
	Sex   string  `json:"name"`
	Age   int     `josn:"age"`
	Grade string  `json:"grade"`
	Score float32 `json:"score"`
}

func New(name, sex string, age int, grade string, score float32) (stu *Student) {
	stu = &Student{
		Name:  name,
		Sex:   sex,
		Age:   age,
		Grade: grade,
		Score: score,
	}
	return
}
