package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func testMap1() {
	var a map[string]int
	//a["stu01"] = 100 不能直接赋值，map没有初始化
	if a == nil {
		a = make(map[string]int, 16)
		a["stu001"] = 100
	}
	fmt.Printf("%#v", a)
}

/* 声明式初始化 */
func testMap2() {
	var a map[string]int = map[string]int{
		"stu001": 100,
		"stu002": 200,
		"stu003": 300,
	}
	fmt.Printf("a=%#v", a)
}

/*map 插入操作 */
func testMap3() {
	var a map[string]int = make(map[string]int, 16)
	a["stu001"] = 10
	a["stu002"] = 20
	a["stu003"] = 30
	fmt.Printf("a=%#v", a)
}

/*通过key访问map*/
func testMap4() {
	a := make(map[string]int, 16)
	a["100"] = 1
	a["123"] = 2
	b := "123"
	fmt.Println("key", b, "value", a[b])
}

/*判断key是否存在*/
func testMap5() {
	a := make(map[string]int, 16)
	a["100"] = 1
	a["123"] = 2
	key := "124"
	value, ok := a[key]
	if !ok {
		fmt.Printf("key %s is not exsit\n", key)
		fmt.Printf("value of a[%s] is %d", key, value)
	} else {
		fmt.Println(value)
	}

}

/* map 遍历 */
func testMap6() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)
	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = rand.Intn(1000)
	}
	for key, value := range a {
		fmt.Printf("value of a[%s] id %d\n", key, value)
	}
}

func testMap7() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)
	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = rand.Intn(1000)
	}
	for key, _ := range a {
		fmt.Printf("a[%s]=%d\n", key, a[key])
	}
}

/* map 删除元素 */
func testMap8() {
	var a map[string]int = make(map[string]int, 16)
	a["stu001"] = 10
	a["stu002"] = 20
	a["stu003"] = 30
	fmt.Printf("a=%#v\n", a)
	key := "stu003"
	delete(a, key)
	fmt.Printf("a=%#v\n", a)
}

/* map 的长度 */
func testMap9() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)
	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = rand.Intn(1000)
	}
	fmt.Printf("length of map is %d\n", len(a))
	delete(a, "key0")
	fmt.Printf("length of map is %d\n", len(a))
}

/* 默认情况下map的key是无序排列的，对map按key进行有序排列 */
func testMap10() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)
	// keys := make([]string, 0)
	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("key%d", i)
		// keys = append(keys, key)
		a[key] = rand.Intn(1000)
	}
	var keys []string
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("a[%s] is %d\n", key, a[key])
	}
}

/* map是引用数据类型 */
func testMap11() {
	var a map[string]int = map[string]int{
		"001": 10,
		"002": 20,
		"003": 30,
	}
	b := a
	fmt.Println(a)
	b["001"] = 20
	fmt.Println(a)
}

/* map 的切片 */
func testMap12() {
	var s []map[string]int
	s = make([]map[string]int, 10)
	s[0] = make(map[string]int, 16)
	s[0]["100"] = 100
	fmt.Println(s)
}

func main() {
	//testMap1()
	// testMap2()
	// testMap3()
	// testMap4()
	// testMap5()
	//testMap6()
	// testMap8()
	// testMap9()
	// testMap10()
	// testMap11()
	testMap12()
}
