## golang 学习笔记

### 一. go语言基础

#### 1.变量，常量， 基本数据类型

##### 1.1 变量

```go
变量的申明 不能以数字开头
var a int = 1 //可用用于申明全局变量 
a := 1 //不用用于申明全局变量

变量申明不能以关键命名
go 中的关键字
break default func interface select
case defer go map struct
chan else goto package switch
const fallthough if range type
continue for import return var

多变量申明
var (
    a int
    b bool
    c string
    d int = 8
    f float64 = 3.1415    
)
```

##### 1.2 常量

```go
常量声明
const a int = 16go
批量申明
const (
    a = 1 << iota
    b
    c
)
```

##### 1.3 数据类型:

###### 1.3.1 数据类型介绍

```go
整数类型：
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
var a int
a = 16

bool类型
var b bool
b = true

浮点数类型
float32 float64
var f float64 = 3.1415

字符串类型
var str string = "hello"

整数型的初始化为 0, 浮点数初始化为 0.0, 字符串为 "", bool为 false
其他数据类型
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型

```

###### 1.3.2 操作符

```go
bool 类型
var b bool = false
!b 取反
|| 或
&& 与
== 等等于
%t 格式化shc

int float数值类型
+ - * /
==
%d

string 类型
"" 可以包含控制字符
`` 所有字原样输出

常用操作
len(str) 获取长度
+,fmt.Sprintf 拼接字段
strings.Split(str,"") 切割
strings.Contains 包含
strings.HasPrefix, strings.HasSuffix 判断前缀和后缀
strings.Index(), strings.LastIndex() ⼦串出现的位置
strings.Join(a[]string, sep string) join操作

操作符
逻辑操作符
==,!=, >, <, >=, <=,
运算算
+ - * / %
```

##### 1.4 类型推导

```go
go 可以根据赋值类型进行变量推导
示例1
b := 1 由1可以推导 b 为int类型
c := false 由值推导出c是bool类型
```



##### 1.5 go 程序的基本结构

```go
1. 任何⼀个代码⽂件⾪属于⼀个包

2. import 关键字，引⽤其他包：

   import(“fmt”)

   import(“os”)

   通常习惯写成：

   import (

    “fmt”

    “os”

   )

   开发可执⾏程序，package **main**，

    并且有且只有⼀个main⼊⼝函数

   4. 包中函数调⽤：

   a. 同⼀个包中函数，直接⽤函数名调⽤

   b. 不同包中函数，通过包名+点+

   函数名进⾏调⽤

   5. 包访问控制规则：

   a. ⼤写意味着这个函数/变量是可导出的

   b. ⼩写意味着这个函数/变量是私有的，包外部不能访问
```



#### 2. 函数

函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。

##### 2.1 函数的形式

```go
1.函数的关键字 func functionname(args)(int) {}

package main

import (
    "fmt"
)

func main(){
    fmt.Println("hello world")
}
2.带参数的函数
func testFunc(a,b int){
    fmt.Printf("a + b =%d",a+b)
}
3.带不定长参数的函数
func testFunc(a ...interface{}){
    for i := 0;i < len(a);i++{
        fmt.Printf("args[%d] is %v",i,a[i])
    }
}
4.函数的返回值
func testFunc(a,b int)(int){
    return a * b
}
5.多返回值
func testFunc(a,b int)(sum int,mod int){
    sum = a+b
    mod = a%b
    return //当返回值被定义时，直接返回皆可
}
```

##### 2.2  函数的调用

```go
1. 函数调用示例
package main

import (
	"fmt"
)

func Add(a,b int)(int){
    return a + b
}

func main(){
    a,b := 1,3
    d := Add(a,b)//调用函数Add
    fmt.Printf(d)
}

```

##### 2.3 匿名函数和闭包

```go
1.没有函数名的函数被称为匿名函数

示例
package main

import (
	"fmt"
)

func main(){
    result := func (a int,b int) (int){
        retrun a + b
    }(1,2) //匿名函数以及传参形式
    fmt.Println(result)
}

2.函数也是值可以像其他的值一样传递,也可以将函数作为参数

2.1 函数作为值传递
package main

import (
    "fmt"
)

func main(){
 	result := func (a int,b int) (int){
        retrun a + b
    }//将函数赋值给result
    fmt.Println(result(1,2))   
}
2.2 函数作为参数
package main

import (
    "fmt"
)

func testFunc(fn func(int,int) int,a,b int) int{ //函数作为参数
    return fn(a,b)
}

func main(){
    result := func (a int,b int) int{
        return a + b
    }
    fmt.Println(testFunc(result,1,3))    
}
2.3 闭包
闭包就是返回一个函数
func add() func(int) int{
    sum := 0
    return funx(x int) int{ //返回一个函数
        sum += x
        return sum
    }
}

func main(){
    result := add()
    for i :=0;i < 10;i++ {
        fmt.Println(result(i))
    }
}
2.3.1 闭包实例，写一个Fibonacci
package main

import (
    "fmt"
)

func Fibonacci() func() int{
    x,y,res := 0,1,0
    return func () int{
        res = x
        x,y = y,x+y
        return x
    }
}

func main(){
    fibo := Fibonacci()
    for i := 0;i< 20;i++{
        fmt.Println(fibo())
    }
}
2.3.2 利用channal 实现(提前预习)
package main

import (
    "fmt"
)

func fibonacci(fobi chan int,count int){
    x,y,res := 0,1,0
    for i :=0;i<count;i++{
        res,x,y = x,y,x+y
        fibo <- res
    }
    close(fibo)
}

func main() {
    fobi := make(chan int,10)
    go fibonacci(fobi,10)
    for val := range fobi{
        fmt.Printf(" val")
    }
}


```

#### 3.  if , switch,for

##### 3.1 if 

if 是go语言的条件表达式

```go
3.1.1 if 基本用法
if 条件 { 条件成立执行
    ...
}else{ //否则执行
    ...
} 
3.1.2 多条件判断
if {
    ...
}else if{
    ...
}else if{
    ...
}else if{
    ...
}else{
    ...
}
3.1.3 if 可以有一个简单的表达式，然后再进行判断
func c(a int) int{
    return a * (-1)
}
func main(){
    if a := c(a); a < 0{ //可以添加简单表达式a := c(a)
        fmt.Println("a is uint")
    }
}

```

##### 3.2 go语言没有while go 语言的循环是使用for实现

```go
3.2.1 for
package main

import (
    "fmt"
)

func main(){
    for i := 0;i < 10;i++{
        fmt.Println(i)
    }
}
3.2.2 for range 
for := range 需要两个接受值，第一值为下标，第二个值为下标所对应的值。 
示例：
package main

import (
	"fmt"
)

func main(){
    str := "hello chunyan"
    for index,val := range str{ 
        fmt.Println("str[%d] is letter %s",index,val)
    }
}

3.2.3 嵌套循环
嵌套循环实现99乘法表
for i := 1; i < 10;i ++ {
    for j := 1;j < i+1;i++ {
        fmt.Printf("%dx%d=%d\t",j,i,i*j)
    }
    fmt.Println()
}
3.2.3 死循环
死循环 格式for {} 
for {
    
}
3.2.4 continue break
continue 跳过之后的代码进行下一次循环
break 退出循环，常常和死循环搭配
示例:
利用死循环，从字符串中找到第一次出现的位置并打印。
package main

import (
    "fmt"
)

str = "hello world"
childstr = " "
index := 0
for {
    if str[index] == childstr {
        fmt.Printf("%s is the %d letter of %s",childstr,index,str)
        break
    }
}
3.2.5 特殊 for
for (退出条件) {
}


```

##### 3.3 switch

```go
switch 用于多分支判断可以优化if的代码，提供代码的可读性
swith {
case
}
示例：
判断某个变量的类型

package main

import (
"fmt"
)
func main(){
	var a int
    switch a.(type) {
    case int:
        fmt.Println("a is int")
    case string:
        fmt.Println("a is string")
    default:
        fmt.Println("can not reconized a's type")
    }    
}

```

#### 4. 数组，切片，map

##### 4.1 数组

数组是存储一组相同类型的数据，数组的长度不可变

```go
4.1.1 数组声明
var arrInt [8]int 声明一个长度为8,元素类型为int的数组
4.1.2 数组初始化
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
4.1.3 数组赋值
var arrInt [8]int = [8]int{}
arrInt[0]=1
fmt.Println(arrInt)
4.1.4 数组遍历
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
for index,val = range arrInt{
    fmt.Printf("arrInt[%d] is %d",index,val)
}
注：数组不能改变长度
```



##### 5.面向对象

##### 6. io

##### 7.接口

##### 8.反射

