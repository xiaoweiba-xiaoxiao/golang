package main

import "fmt"

//fmt.Scanf()以空格作为分割符，占位符和格式化输出⼀致
func testScanf() {
	var a int
	var b string
	var c float64
	fmt.Scanf("%d %s %f", &a, &b, &c)
	fmt.Printf("a=%d b=%s d=%f\n", a, b, c)
	/* 只能读取第一个的输入，例如中断输入10，相当于输入10\r\n
	Scanf,不会把\r\n给处理掉，所以输出结果是a=10 b= c=0.000000
	fmt.Scanf("%d",a)
	fmt.Scanf("%s",b)
	fmt.Scanf("%f",c)
	*/
	//Scanf加上\n表示值读到换行符为止，就能把\r\n给处理掉。
	fmt.Scanf("%d\n", a)
	fmt.Scanf("%s\n", b)
	fmt.Scanf("%f\n", c)
	fmt.Printf("a=%d b=%s d=%f\n", a, b, c)
}

//Scan()以空格和换行符做为分割符
func testScan() {
	var a int
	var b string
	var c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d b=%s d=%f\n", a, b, c)
}

//Scanln()以空格作为分割符，以换行符结束。
func testScanln() {
	var a int
	var b string
	var c float64
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("a=%d b=%s d=%f\n", a, b, c)
}

func main() {
	// testScanf()
	// testScan()
	testScanln()
}
