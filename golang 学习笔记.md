





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
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
故上面结果abc为(1,2,3)
```

##### 1.3 数据类型:

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

###### 1.3.1数字类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **uint8** 无符号 8 位整型 (0 到 255)                         |
| 2    | **uint16** 无符号 16 位整型 (0 到 65535)                     |
| 3    | **uint32** 无符号 32 位整型 (0 到 4294967295)                |
| 4    | **uint64** 无符号 64 位整型 (0 到 18446744073709551615)      |
| 5    | **int8** 有符号 8 位整型 (-128 到 127)                       |
| 6    | **int16** 有符号 16 位整型 (-32768 到 32767)                 |
| 7    | **int32** 有符号 32 位整型 (-2147483648 到 2147483647)       |
| 8    | **int64** 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

###### 1.3.2浮点型

| 序号 | 类型和描述                        |
| :--- | :-------------------------------- |
| 1    | **float32** IEEE-754 32位浮点型数 |
| 2    | **float64** IEEE-754 64位浮点型数 |
| 3    | **complex64** 32 位实数和虚数     |
| 4    | **complex128** 64 位实数和虚数    |

------

###### 1.3.3其他数字类型

以下列出了其他更多的数字类型：

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **byte** 无符号 8 位整型 (0 到 255)                          |
| 2    | **rune** **int32** 有符号 32 位整型 (-2147483648 到 2147483647) |
| 3    | **uint** 32 或 64 位                                         |
| 4    | **int** 与 uint 一样大小                                     |
| 5    | **uintptr** 无符号整型，用于存放一个指针                     |

###### 1.3.3.4 数据类型介绍

```go
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

###### 1.3.3.5操作符

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
==,!=, >, <, >=, <=,&&，||,!
运算算
+ - * / % >> << ^ & |
A =   0011 1100
B =   0000 1101
A^B = 0011 0001
A&B = 0000 1100
A|B = 0011 1101
var a int =2 
a<<1 //4
a>>1//2
```

##### 1.4 类型推导

```go
go 可以根据赋值类型进行变量推导
示例1
b := 1 由1可以推导 b 为int类型
c := false 由值推导出c是bool类型
可以通过
fmt.Printf(”%T/n",c)//可以输出类型
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

##### 1.6golang关键字

| break    | default    | func   | interface | select |
| -------- | ---------- | ------ | --------- | ------ |
| case     | defer      | go     | map       | struct |
| chan     | else       | goto   | package   | switch |
| const    | fallthough | if     | range     | type   |
| continue | for        | import | return    | var    |

##### 1.7注释

注释可以帮我们很好的完成文档的工作，写得好的注释可以方便我们以后的维护。

/**/ 的块注释和 // 的单行注释两种注释风格， 

包注释、结构体（接口）注释、函数（方法）注释、代码逻辑注释以及注释规范方面进行说明。

###### 1.7.1包注释

- 每个包都应该有一个包注释，一个位于package子句之前行注释
- 包注释应该包含下面基本信息

```go
// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  请填写自己的真是姓名（需要改）  ${DATE} ${TIME}
// @Update  请填写自己的真是姓名（需要改）  ${DATE} ${TIME}
package ${GO_PACKAGE_NAME}
```

###### 1.7.2结构（接口）注释

每个自定义的结构体或者接口都应该有注释说明，该注释对结构进行简要介绍，放在结构体定义的前一行，格式为： 结构体名， 结构体说明。同时结构体内的每个成员变量都要有说明，该说明放在成员变量的后面（注意对齐），实例如下：

```go
// User   用户对象，定义了用户的基础信息
type User struct{
    Username  string // 用户名
    Email     string // 邮箱
}
```

###### 1.7.3函数（方法）注释

- 每个函数，或者方法（结构体或者接口下的函数称为方法）都应该有注释说明
- 函数的注释应该包括三个方面

```go
// @title    函数名称
// @description   函数的详细描述
// @auth      作者             时间（2019/6/18   10:57 ）
// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"
```

###### 1.7.4代码逻辑注释

- 每个代码块都要添加单行注释
- 注视使用TODO 开始  详细如下

```go
// TODO  代码块的执行解释
if   userAge < 18 {
     
}
```

###### 1.7.5注释要求

- 统一使用中文注释，对于中英文字符之间严格使用空格分隔， 这个不仅仅是中文和英文之间，英文和中文标点之间也都要使用空格分隔
- 全部使用单行注释，禁止使用多行注释
- 和代码的规范一样，单行注释不要过长，禁止超过 120 字符。

###### 1.7.6缩进和折行

- 折行方面，一行最长不超过120个字符，超过的请使用换行展示，尽量保持格式优雅

###### 1.7.7括号和空格

括号和空格方面，也可以直接使用 gofmt 工具格式化（go 会强制左大括号不换行，换行会报语法错误），所有的运算符和操作数之间要留空格。

###### 1.7.8import 规范

```go
// 单行引入
import  "fmt"
```



```go
// 多包引入，每包独占一行
// 使用绝对路径，避免相对路径如 ../encoding/json
import (
     "strings"
     "fmt"
)
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
6.函数可以作为其他函数的参数进行传递，然后在其他函数内调用执行，一般称为回调
package main

    import "fmt"

    func main() {
        callback(2, Add)
    }

    func Add(x, y int) {
        fmt.Println(x+y)
    }

    func callback(z int, f func(int, int)) {
        f(z, 5)
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
    }(1,2) //匿名函数以及传参形式,采用（）是直接去调用函数
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
//注意，函数里面如果对a,b等非指针进行修改，只是修改ab的副本值，实际的值没有发生改变
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
    return func(x int) int{ //返回一个函数
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

##### 3.2 for 循环

go语言没有while go 语言的循环是使用for实现

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
3.2.6 for的死循环，相当于其他语言的while循环
for ture{
    
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
        fmt.Println("a is string"
    default:
        fmt.Println("can not reconized a's type")
    }    
         /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C" //可以为多个值
      default: grade = "D"  
   }
 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
 switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }
        
  

}
 

```

##### 3.4select 

A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations.

一个select语句用来选择哪个case中的发送或接收操作可以被立即执行。它类似于switch语句，但是它的case涉及到channel有关的I/O操作。

```go
select {
  case communication clause  :
    statement(s);    
  case communication clause  :
    statement(s);
  // 你可以定义任意数量的 case 
  default: // 可选 
    statement(s);
}
每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
```



#### 4. 数组，切片，map

##### 4.1 数组

数组是存储一组相同类型的数据，数组的长度不可变

```go
4.1.1 数组声明
var arrInt [8]int 声明一个长度为8,元素类型为int的数组
4.1.1 数组初始化
arrInt=[8]int{1,2,3,4,5,6,7,8}
4.1.3 数组声明和初始
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
可以通过...系统来帮我确定长度
arrint :=[...] int{1,2,3,4,5,6,7,8}
4.1.4 数组赋值
var arrInt [8]int = [8]int{}
arrInt[0]=1
fmt.Println(arrInt)
4.1.5 数组遍历
var arrInt [8]int = [8]int{1,2,3,4,5,6,7,8}
for index,val = range arrInt{
    fmt.Printf("arrInt[%d] is %d",index,val)
}

for i := 1; i < len(arrInt); i++ {
		fmt.Println(arrInt[i])
} 
注：数组不能改变长度
```

##### 4.2Map

map是一堆未排序键值对的集合，map的key必须为可比较类型,比如 **==** 或 **!=**，map查找比线性查找快,但慢于索引查找(数组，切片)

格式 var name map[keytype]valuetype

```go
var a map[int]int
var b = map[string]string{}
c := map[string]bool{}
d := make(map[string]int)
```

###### 4.2.1map的初始化

map类似于数组和切片,可以在定义时直接指定初始值 如:

```go
var a = map[string]string{"name": "sutao", "age": "26", "sex": "男"}
```

###### 4.2.2map元素访问

上边定义并初始化了一个map,那么如何访问一个map中的元素呢？

可以通过键(key)来访问指定的value 如:

```go
var a = map[string]string{"name": "sutao", "age": "26", "sex": "男"}
fmt.Println(a["name"],a["age"])
```

###### 4.2.3修改map中的元素

和数组一样,可以使用key来修改map中的元素

```go
var a = map[string]string{"name": "sutao", "age": "26", "sex": "男"}
fmt.Println(a)
a["name"] = "李四"
fmt.Println(a)
```

###### 4.2.4新增元素

与slice不同,map元素的新增不需要使用copy()函数 可以直接 `map[key] = value` 的方式增加元素 如:

```
package main

import (
    "fmt"
)

func main() {
    var a = map[string]string{"name": "wangchao", "age": "32", "sex": "男"}
    a["girlfriend"] = "苍老师"
    fmt.Println(a)
}
```

###### 4.2.5删除元素

比slice相比,map提供delete()函数进行元素的删除
格式 `delete(map,key)` 如下例:

```
package main

import (
    "fmt"
)

func main() {
    var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
    delete(a, "name")
    fmt.Println(a)
}
```

###### 4.2.6map是引用类型

map是引用类型,来看一个简单的例子

```
func main() {
    var a = map[int]int{1: 1, 2: 2}
    b := a
    a[1] = 123456
    fmt.Println(a, b)
}
```

以上程序会输出 map[1:123456 2:2] map[1:123456 2:2]

我们再传递进函数里面测试下,关于函数的定义后面我们会讲,此处只是一个demo

```
func main() {
    var a = map[int]int{1: 1, 2: 2}
    b := a
    test(a)
    fmt.Println(a, b)
}

func test(a map[int]int) {
    a[1] = 111
}
```

以上程序会输出 map[2:2 1:111] map[1:111 2:2] 或者 map[1:111 2:2] map[1:111 2:2]

##### 4.3切片

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

- 基于数组或者slice生成一个slice的时候,新的slice和原来数组/slice 的底层数组是同一个

- 基于数组或者slice生成slice的cap=原来对应的数组长度-现在slice的第一个元素对应的索引

- slice 作为参数传递的是 副本,但是对应的数组的指针不变

- 扩容规则:
  在一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下
  简称原容量）的 2 倍。
  但是，当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25
  倍作为新容量的基准（以下新容量基准）

  对于`slice`有几个有用的内置函数：

  - `len` 获取`slice`的长度
  - `cap` 获取`slice`的最大容量
  - `append` 向`slice`里面追加一个或者多个元素，然后返回一个和`slice`一样类型的`slice`
  - `copy` 函数`copy`从源`slice`的`src`中复制元素到目标`dst`，并且返回复制的元素的个数

  注：`append`函数会改变`slice`所引用的数组的内容，从而影响到引用同一数组的其它`slice`。 但当`slice`中没有剩余空间（即`(cap-len) == 0`）时，此时将动态分配新的数组空间。返回的`slice`数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的`slice`则不受影响。

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```



```go
// 声明一个数组
var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
// 声明两个slice
var aSlice, bSlice []byte

// 演示一些简便操作
aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

// 从slice中获取slice
aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
```

#### 5.字符串

##### 5.1 字符串

由于和其他语言相比，字符串在 Go 语言中有着自己特殊的实现，因此在这里需要被特别提出来

Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。

```go
5.1.1 Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串。
package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	fmt.Println(str)
}
输出hello world
5.1.2 字符串遍历
package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}
输出:68 65 6c 6c 6f 20 77 6f 72 6c 64 21 为hello world Unicode UTF-8 编码的结果
ps:golang是使用UTF8进行编码，一个中文是占3个字节
    str := "你好"
	fmt.Println(len(str))//6
如果是打印字符串每个字符如果用for i循环，如果遇到中文会出现乱码情况
	theme := "开始start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}
原因：中文3个字节造成，解决方法用for each循环
for _, s := range theme {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}
分析range循环，内部使用rune的值，来进行整合，到时不会出现上面 for i乱码的情况

package main

import (
	"fmt"
)

func main() {
	str := "hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
	for _, val := range str {
		fmt.Printf("%c", val)
	}
}
输出结果：
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
上面的程序获取字符串的每一个字符，虽然看起来是合法的，但却有一个严重的 bug。让我拆解这个代码来看看我们做错了什么。
package main

import (
	"fmt"
)

func printByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}

func printChar(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")
}

func main() {
	str := "hello world!"
	printByte(str)
	printChar(str)
	str = "Señor"
	printByte(str)
	printChar(str)
}
打印结果：
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
53 65 c3 b1 6f 72 
SeÃ±or
在上面程序的第 28 行，我们尝试输出 Señor 的字符，但却输出了错误的 S e Ã ± o r。 为什么程序分割 Hello World 时表现完美，但分割 Señor 就出现了错误呢？这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码占用了 c3 和 b1 两个字节。它的 UTF-8 编码占用了两个字节 c3 和 b1。而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。那么我们该怎么办呢？rune 能帮我们解决这个难题。
```

##### 5.2 rune

rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示

```go
package main

import (
	"fmt"
)

func printByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Printf("\n")
}

func printChar(str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	fmt.Printf("\n")
}

func main() {
	str := "hello world!"
	printByte(str)
	printChar(str)
	str = "Señor"
	printByte(str)
	printChar(str)
}
输出结果:
68 65 6c 6c 6f 20 77 6f 72 6c 64 21 
hello world!
53 65 c3 b1 6f 72 
Señor

```

##### 5.3 for range 循环

上述问题可以通过for rang循环解决

```go
package main

import (
	"fmt"
)

func printStr(str string) {
	for index, runes := range str {
		fmt.Printf("%c bytes is %d\n", runes, index)
	}
}

func main() {

	str := "Señor"

	printStr(str)
}
输出结果：
S bytes is 0
e bytes is 1
ñ bytes is 2
o bytes is 4
r bytes is 5
可以看到ñ 占了两个字节

用字节切片构造字符串
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}
打印结果：
Café

如果我们把 16 进制换成对应的 10 进制值会怎么样呢
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{67, 97, 102, 195, 169} 
    str := string(byteSlice)
    fmt.Println(str)
}
打印结果：
Café

用 rune 切片构造字符串
package main

import (  
    "fmt"
)

func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str := string(runeSlice)
    fmt.Println(str)
}
打印结果：
Señor

```

##### 5.4 字符串的长度

```go
utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。
package main

import (
	"fmt"
	"unicode/utf8"
)

func lenth(str string){
	fmt.Printf("the lenth of %s is %d\n",str,utf8.RuneCountInString(str))
}

func main(){
	str := "Señor"
	lenth(str)
	str = "chun"
	lenth(str)
}

len(str) 也可以获取字节的长度
```

##### 5.5  字符串不可变

字符串是不可变的,一旦被创建它将无法修改

```go
package main

import (
	"fmt"
)

func changString(str string) string{
	str[0] = 'a'
	return str
}

func main(){
	str := "hello"
	fmt.Println(changString(str))
}
试图修改字符串：
# xiao.com/golang/writefile/lesson1
./lesson1.go:8:9: cannot assign to str[0]
exit status 2
Process exiting with code: 1
```

可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。

```go
package main

import (
	"fmt"
)

func changStr(s []rune) string{
	s[0] = 'a'
	return string(s)
}

func main(){
	str := "hello"
	fmt.Println(changStr([]rune(str)))
}
此例实现了对字符串的修改
终端输出：
aello
```

#### 6. 指针

指针是一种存储变量内存地址（Memory Address）的变量。

##### 6.1 指针的声明

```go
6.1.1
var t *T //声明一个指向T类型的指针变量

示例：
package main

import (
	"fmt"
)

func main(){
	var a *int
	var b int = 5
	a = &b
	fmt.Printf("the type of a is %T\n",a)
	fmt.Println("The adrres of a is", a)
    fmt.Println("The value of a is", *a)
}
终端输出结果:
Type of a is *int  
address of b is 0x1040a124
The value of is 5

6.1.2 指针的零值：
指针的零值（Zero Value）为nil
package main

import (
	"fmt"
)

func main(){
	a := 15
	var b *int
	if b == nil {
		fmt.Println("b is",b)
		b = &a
		fmt.Println("b is",b)
	}
}
终端输出:
b is <nil>
b is 0xc0000ac010
```

##### 6.2 指针解引用

指针解引用是可以通过指针获取到指针指向变量的值

```go
6.2.1 示例
package main

import (
	"fmt"
)

func main(){
	var b int = 23
	var a *int = &b
	fmt.Println("the address of b is",a)
	fmt.Println("the value of b is",*a)
}

终端输出:
the address of b is 0xc000016090
the value of b is 23

6.2.2 利用指针解引用来修改变量的值
package main

import (
	"fmt"
)

func main(){
	var b int = 24
	var a *int = &b
	*a++
	fmt.Println("the value of b is",b)
}
终端输出：
the value of b is 25
```

##### 6.3 指针传参

指针可以作为参数传递给变量

```go
6.3.1 示例1
package main

import (
	"fmt"
)

func testPointer(t *int){
	*t *= 2 
}

func main(){
	b := 2
	testPointer(&b)
	fmt.Println(b)
}
终端输出：
API server listening at: 127.0.0.1:47030
4
6.3.2 示例2
这里使用了数组指针
package main

import (
	"fmt"
)

func testSplprt(t *[3]int){
	for i := 0;i<len(*t);i++ {
		t[i] +=1
	}
	
}

func main(){
	var arr [3]int=[3]int{90,91,92}
	testSplprt(&arr)
	fmt.Println(arr)
}
终端输出：
API server listening at: 127.0.0.1:3011
[91 92 93]
//这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。

使用切片
package main

import (
	"fmt"
)

func testArr(t []int){
	for i :=0;i<len(t);i++{
		t[i] += 1 
	}
}

func main(){
	var arr [3]int=[3]int{90,91,92}
	testArr(arr[:])
	fmt.Println(arr)
}
终端输出：
API server listening at: 127.0.0.1:43329
[91 92 93]
```

##### 6.5  不支持指针运算

Go 并不支持其他语言（例如 C）中的指针运算。

```go
//这是一个错误案例

package main

import (
	"fmt"
)

func main(){
	b := [...]int{109,110,111}
	p = &b
	p++
}
```



#### 7.  面向对象

##### 7.1 结构体

结构体（struct）是用户自定义的类型，它代表若干字段的集合，可以用于描述一个实体对象，类似java中的class，是golang面向对象编程的基础类型。
结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type）。在 C++ 它也存在，并且名字也是 struct，在面向对象的编程语言中，跟一个无方法的轻量级类一样。

Go语言对关键字的增加非常吝啬，其中没有`private`、`protected`、`public`这样的关键 字。要使某个符号对其他包(`package`)可见(即可以访问)，需要将该符号定义为以大写字母开头

##### 7.1.1 结构体声明

结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。

例如，一个职员有 `firstName`、`lastName` 和 `age` 三个属性，而把这些属性组合在一个结构体 `employee` 中就很合理。

```go
7.1.1示例1
package main

import (
	"fmt"
)

type employee struct{
	fistname string
	lastname string
	age int
}
//可以简写 属于同一类型的字段
type Employee struct {
    firsname,lastename string
    age,sex int
}

```

###### 7.1.2  匿名结构体（Anonymous Structure）

声明结构体时也可以不用声明一个新类型，这样的结构体类型称为匿名结构体

```go 
package main

import (
	"fmt"
)

func main(){
	var employee struct {
		firstname,lastname string
		age,sex int
	}
	employee.firstname = "w"
	employee.lastname = "c"
	employee.age = 31
	employee.sex = 1
	fmt.Println(employee)
}
```

###### 7.1.3 命名的结构体 

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}
```

​          上面的结构体 Employee 称为 命名的结构体（Named Structure）。我们创建了名为 Employee 的新类型，而它可以用于创建 Employee 类型的结构体变量。

###### 7.1.4 结构体零值

当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值。

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	fmt.Println(cai)
}
终端输出
API server listening at: 127.0.0.1:9978
{  0 0}
firstname lastname 为""
sex age 为 0
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	cai.firstname = "c"
	cai.lastname = "xx"
	fmt.Println(cai)
}
终端输出
API server listening at: 127.0.0.1:19042
{c xx 0 0}
由于sex age 没有赋值，所以就采用0
```

###### 7.1.5 结构体字段访问

结构体字段访问方式为结构体变量.字段名

```go
package main

import (
	"fmt"
)

type Employee struct{
	firstname,lastname string
	sex,age int
}

func main(){
	var cai Employee
	cai.firstname = "c"
	cai.lastname = "xx"
	fmt.Println(cai)
}
```

###### 7.1.6 结构体指针

结构体指针是指向结构体的指针，反问结构体指针指向变量的值的字段时类似于（*a)解引用 可以直接

```go
package main

import (
	"fmt"
)

type Employee struct{
	Name string
	Sex,Age int
}

func main(){
	var employee *Employee = &Employee{
		Name: "hello",
		Sex: 1,
		Age: 31,
	}
	fmt.Println(employee.Name)//指针解引用struct 是指针的时候，编译器会把employees.field识别为（*employees).field
}
终端输出：
API server listening at: 127.0.0.1:46671
hello
```

###### 7.1.7 匿名字段

当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）

```go
package main

import (
	"fmt"
)

type person struct{
	string
	int
}

func main(){
	var people person=person{
		"hello",
		11,
	}
	fmt.Println(people.string)
	fmt.Println(people.int)
}
终端输出:
API server listening at: 127.0.0.1:31998
hello
11
```

虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。比如在上面的 `Person` 结构体里，虽说字段是匿名的，但 Go 默认这些字段名是它们各自的类型。所以 `Person` 结构体有两个名为 `string` 和 `int` 的字段。

```go
package main

import (
	"fmt"
)

type person struct{
	string
	int
}

func main(){
	var people person=person{
		"hello",
		11,
	}
	fmt.Println(people.string)
	fmt.Println(people.int)
}
终端输出：
API server listening at: 127.0.0.1:31998
hello
11
```

###### 7.1.8 嵌套结构体

结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	Addr Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		Addr: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e)
}
终端输出:
API server listening at: 127.0.0.1:39086
{Hello 31 2 {ChengDu}}
```

###### 7.1.9 提升字段（Promoted Fields）

如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		Addr: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City) //用外层结构体直接访问City字段
}
终端输出：
ChengDu
```

但是应提升结构体字段冲突：

```go
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City) //用外层结构体直接访问City字段
}
终端输出:
# xiao.com/golang/writefile/lesson1
./lesson1.go:35:15: ambiguous selector e.City
exit status 2
Process exiting with code: 1

出现报错的原因：
School,Address里面都有一个City,此时go程序不知道需要去获取那个结构体的City,因此就发生了冲突
冲突的解决方案1
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.School.City) //用内部的结构体加上字段名反问直接访问City字段
}
终端输出:
Wuhan
第二种
package main

import  (
	"fmt"
)

type Address struct{
	City string
}

type School struct{
	City string
}

type Employee struct{
	Name string
	Age int
	Sex int
	City string
	School
	Address
}

func main(){
	var e Employee=Employee{
		Name: "Hello",
		Age: 31,
		Sex: 2,
		City: "Beijing",
		School: School{
			City: "Wuhan",
		},
		Address: Address{
			City: "ChengDu",
		},
	}
	fmt.Println(e.City)
}
终端输出：
API server listening at: 127.0.0.1:39228
Beijing
也由此我们得出一个结论，当外层的结构体和提升字段中包含同样名称字段时，优先级时外层结构体优于提升字段
```

###### 7.1.10 导出结构体和字段

如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。

```go
package person

type Student struct {
	Name     string
	Age, Sex int
	score    float32
}
//在文件夹person的文件person.go 新建一个结构体



//在main.go中访问
package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
		score: 100.00,//注释掉这行程序就能正常运行
	}
	fmt.Println(xiaoxiao)
}
终端输出:
# xiao.com/golang/writefile/lesson1
./main.go:14:3: unknown field 'score' in struct literal of type person.Student
exit status 2
Process exiting with code: 1
因为score字段是小写开头，所以我们不能直接访问。

```

###### 7.1.11结构体相等性

结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。

```go
package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
	}
	var xiaochao person.Student
	xiaochao.Name = "xiaoxiao"
	xiaochao.Age = 18
	xiaochao.Sex = 2
	if xiaoxiao == xiaochao {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
终端输出:
API server listening at: 127.0.0.1:13069
true
```

###### 7.1.12 tag

tag 是用于序列化struct 字段的，解析成tag相映的名称

```go
package person

type Student struct {
	Name    string `ymal:"name" json:"name" ini:"name"`
	Age   int
	Sex  int
	score    float32
}

package main

import (
	"fmt"
	"xiao.com/golang/writefile/person"
	"encoding/json"
)


func main(){
	var xiaoxiao person.Student = person.Student{
		Name: "xiaoxiao",
		Age: 18,
		Sex: 2,
	}
	var xiaochao person.Student
	xiaochao.Name = "xiaoxiao"
	xiaochao.Age = 18
	xiaochao.Sex = 2
	if xiaoxiao == xiaochao {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
	data,_:= json.Marshal(xiaoxiao)
	fmt.Println(string(data))
}
输出：
API server listening at: 127.0.0.1:24945
true
{"name":"xiaoxiao","Age":18,"Sex":2}
//这是一个简单的应用。
```



##### 7.2 结构体方法

方法其实就是一个函数，在 `func` 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。

```go
package main

import (
	"fmt"
)

type Employee struct{
	Name string
	Money float32
}

func (e Employee)Show(){
	fmt.Printf("%#v",e)
}

func main(){
	cai := Employee{
		Name: "cai",
		Money: 15000.00,
	}
	cai.Show() 
}

通过 变量.方法 的形式调用了方法
```

既然我们可以使用函数写出相同的程序，那么为什么我们需要方法？这有着几个原因，让我们一个个的看看。

- [Go 不是纯粹的面向对象编程语言](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
- 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。假设我们有一个 `Square` 和 `Circle` 结构体。可以在 `Square` 和 `Circle` 上分别定义一个 `Area` 方法。见下面的程序。

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct{
	length int
	width int
}

type Circle struct{
	radius float64
}

func (r Rectangle)Area() int{
	return r.length * r.width
}

func (c Circle)Area() float64{
	return math.Pi * c.radius * c.radius
}

func main(){
	r := Rectangle{
		length: 10,
		width: 5,
	}
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of rectangle %d\n",r.Area())
	fmt.Printf("Area of circle %f",c.Area())
}
终端输出：
API server listening at: 127.0.0.1:16689
Area of rectangle 50
Area of circle 452.389342

从上可以看出每个结构体都可以有自己的area计算方法。
```

###### 7.2.1 指针接收器与值接收器

到目前为止，我们只看到了使用值接收器的方法。还可以创建使用指针接收器的方法。值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。让我们用下面的程序来帮助理解这一点。

```go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct{
	length int
	width int
}

type Circle struct{
	radius float64
}

func (r *Rectangle)Area() int{
	return r.length * r.width
}

func (c *Circle)Area() float64{
	return math.Pi * c.radius * c.radius
}

func main(){
	r := Rectangle{
		length: 10,
		width: 5,
	}
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of rectangle %d\n",r.Area())
    fmt.Printf("Area of circle %f",(&c).Area())
}
终端输出:
API server listening at: 127.0.0.1:38075
Area of rectangle 50
Area of circle 452.389342

可以看到(&c).Area() 和 c.Area()等价
```

什么时候使用指针接收器，什么时候使用值接收器？

一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

在其他的所有情况，值接收器都可以被使用。

建议使用指针接收器，除非有特殊的要求

###### 7.2.2 匿名字段的方法

属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。

```go
package main

import (
	"fmt"
)

type address struct{
	city string
	state string
}

func (a *address) fullAdress(){
	fmt.Printf("Full address: %s，%s",a.city,a.state)
}

type person struct{
	firsrName string
	lastName string
	address
}

func main(){
	p := person{
		firsrName: "hello",
		lastName: "world",
		address: address{
			city: "chengdu",
			state: "gaoxin",
		},
	}
	p.fullAdress()
}

终端输出：
API server listening at: 127.0.0.1:18025
Full address: chengdu，gaoxin
```

在方法中使用值接收器 与 在函数中使用值参数

当一个函数有一个值参数，它只能接受一个值参数。

当一个方法有一个值接收器，它可以接受值接收器和指针接收器。

```go
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	area(r)
	r.area()
}
终端输出:
# xiao.com/golang/writefile/lesson1
./main.go:25:6: cannot use r (type *rectangle) as type rectangle in argument to area
exit status 2
Process exiting with code: 1

原因是area函数只能接收一个值类型。
正确的写法：
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	//area(r)
	area(*r)
	r.area()
}
终端输出：
API server listening at: 127.0.0.1:15773
Area Function result: 50
Area Function result: 50
反之亦然
函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器。
package main

import (
	"fmt"
)

type rectangle struct{
	length int
	width int
}

func area(r *rectangle){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func (r *rectangle)area(){
	fmt.Printf("Area Function result: %d\n",r.length * r.width)
}

func main(){
	r := &rectangle{
		length:  10,
		width: 5,
	}
	area(r)
	// area(*r) 不能使用值类型
	r.area()
}
终端输出：
API server listening at: 127.0.0.1:23433
Area Function result: 50
Area Function result: 50

```

###### 7.2.3在非结构体上的方法

到目前为止，我们只在结构体类型上定义方法。也可以在非结构体类型上定义方法，但是有一个问题。为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 `main` 包中，因此它们是可以运行的。

```go
package main

import (
	"fmt"
)

type Myint int

func (a Myint)add(b Myint) Myint{
	return a + b
}

func main() {
	num1 := Myint(5)
	num2 := Myint(10)
	sum := num1.add(num2)
	fmt.Println("sum is",sum)
}

终端输出:
API server listening at: 127.0.0.1:38977
sum is 15


示例2
package main

func (a int) add(b int) {
}

func main() {

}
在上面程序的第 3 行，我们尝试把一个 add 方法添加到内置的类型 int。这是不允许的，因为 add 方法的定义和 int 类型的定义不在同一个包中。该程序会抛出编译错误 cannot define new methods on non-local type int
```

##### 7.3 接口

接口就是一组方法的签名

任何实现了该接口方法的结构体，就实现了该接口

###### 7.3.1 接口的声明与实现

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
	Talk()
}

type Dog struct{
	name string
	kind string
}

func (d Dog)Eat(){
	fmt.Printf("I am eat\n")
}

func (d Dog)Talk(){
	fmt.Printf("I am %s,I am %s",d.name,d.kind)
}

func main(){
	var a Animal
	d := Dog{
		name: "maomao",
		kind: "samoye",
	}
	a = d
	a.Eat()
	a.Talk()
}
终端输出:
API server listening at: 127.0.0.1:36648
I am eat
I am maomao,I am samoye
```

###### 7.3.2 接口的实际用途

前面的例子教我们创建并实现了接口，但还没有告诉我们接口的实际用途。在上面的程序里，如果我们使用 `name.FindVowels()`，而不是 `v.FindVowels()`，程序依然能够照常运行，但接口并没有体现出实际价值。

因此，我们现在讨论一下接口的实际应用场景。

我们编写一个简单程序，根据公司员工的个人薪资，计算公司的总支出。为了简单起见，我们假定支出的单位都是美元。

```go
package main

import (
	"fmt"
)

type SalaryCalculator interface{
	CalculateSalary() int
}

type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

func (p Permanent) CalculateSalary() int{
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int{
	return c.basicpay
}

func totalExpense(s []SalaryCalculator){
	expense := 0
	for _,v := range s{
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total EXpense per Month $%d",expense)
}

func main(){
	pemp1 := Permanent{1,5000,20}
	pemp2 := Permanent{2,6000,30}
	cemp1 := Contract{3,3000}
	empolyees := []SalaryCalculator{pemp1,pemp2,cemp1}
	totalExpense(empolyees)
}
终端输出:
API server listening at: 127.0.0.1:29236
Total EXpense per Month $14050
```

###### 7.3.3 接口的内部表示

我们可以把接口看作内部的一个元组 `(type, value)`。 `type` 是接口底层的具体类型（Concrete Type），而 `value` 是具体类型的值

```go
package main

import (
	"fmt"
)

type Test interface{
	Tester()
}

type myFloat float64

func (m myFloat)Tester(){
	fmt.Println(m)
}

func describe(t Test){
	fmt.Printf("Interface type %T value %v\n",t,t)
}

func main(){
	var t Test
	f := myFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
Test 接口只有一个方法 Tester()，而 MyFloat 类型实现了该接口。在第 24 行，我们把变量 f（MyFloat 类型）赋值给了 t（Test 类型）。现在 t 的具体类型为 MyFloat，而 t 的值为 89.7。第 17 行的 describe 函数打印出了接口的具体类型和值。该程序输出：
API server listening at: 127.0.0.1:22370
Interface type main.myFloat value 89.7
89.7
```

###### 7.3.4 空接口

没有包含方法的接口称为空接口。空接口表示为 `interface{}`。由于空接口没有方法，因此所有类型都实现了空接口。

空接口的应用非常广泛

```go
var a []interface{} // 这个切片可以存储任意类型的值
示例1：特殊的切片和map
package main

import (
	"fmt"
)

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	fmt.Println(arr)
}
终端输出:
API server listening at: 127.0.0.1:17679
[hello 1 3.14]

package main

import (
	"fmt"
)

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	fmt.Println(arr)
	var maps map[interface{}]interface{} = map[interface{}]interface{}{
		1:"hello",
		"hello":1,
		3.14: "Pi",
		"Pi":3.14,
	}
	fmt.Println(maps)
}
终端输出:
API server listening at: 127.0.0.1:43578
[hello 1 3.14]
map[3.14:Pi 1:hello Pi:3.14 hello:1]

示例2:实现接收任意类型参数的函数
package main

import (
	"fmt"
)

func test(i interface{}){
	fmt.Println(i)
}

func main(){
	var arr []interface{}=[]interface{}{"hello",1,3.14}
	var maps map[interface{}]interface{} = map[interface{}]interface{}{
		1:"hello",
		"hello":1,
		3.14: "Pi",
		"Pi":3.14,
	}
	test(arr)
	test(maps)
}
终端输出:
API server listening at: 127.0.0.1:39922
[hello 1 3.14]
map[3.14:Pi 1:hello Pi:3.14 hello:1]
```

###### 7.3.5类型断言

类型断言用于提取接口的底层值（Underlying Value）。

在语法 `i.(T)` 中，接口 `i` 的具体类型是 `T`，该语法用于获得接口的底层值。

一段代码胜过千言。下面编写个关于类型断言的程序。

```go
package main

import (
	"fmt"
)

func assert(i interface{}){
	s := i.(int)
	fmt.Println(s)
}

func main(){
	var s interface{} = 56
	assert(s)
}
终端输出:
API server listening at: 127.0.0.1:3834
56

断言类型错误是，package main

import (
	"fmt"
)

func assert(i interface{}){
	s,ok := i.(int)
	fmt.Println(s，ok)
}

func main(){
	var s interface{} = "hello"
	assert(s)
}

推荐的类型断言实现:
API server listening at: 127.0.0.1:8129
0 false
```

###### 7.3.6 类型选择

类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。

类型选择的语法类似于类型断言。类型断言的语法是 `i.(T)`，而对于类型选择，类型 `T` 由关键字 `type` 代替。下面看看程序是如何工作的。

```go
package main

import (
	"fmt"
)

func findType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n",i.(int))
	case string:
		fmt.Printf("I am an string and my value %s\n",i.(string))
	default:
		fmt.Printf("Unknow type\n")
	}
}

func main(){
	findType("Naveen")
	findType(77)
	findType(89.98)
}
终端输出:
API server listening at: 127.0.0.1:30192
I am an string and my value Naveen
I am a int and my value is 77
Unknow type
```

还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。

```go
package main

import (
	"fmt"
)

type person struct{

}
type Des interface{
	describe()
}

func (p person)describe(){

}

func findType(i interface{}){
	switch i.(type) {
	case int:
		fmt.Printf("I am a int and my value is %d\n",i.(int))
	case string:
		fmt.Printf("I am an string and my value %s\n",i.(string))
	case Des:
		fmt.Printf("I am an Des\n")
	default:
		fmt.Printf("Unknow type\n")
	}
}

func main(){
	findType("Naveen")
	findType(77)
	findType(89.98)
	var p person
	findType(p)
}
```

###### 7.3.7 指针接收者和值接受者

指针接收者是指针实现了这个接口，值接受者是值实现的这个接口

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
	Talk()
}

type Dog struct{
	name string
}

type Cat struct{
	name string
}

func (d *Dog)Eat(){
	fmt.Println("dog is eatting")
}

func (d *Dog)Talk(){
	fmt.Printf("I am dog,my name is %s\n",d.name)
}

func (c Cat)Eat(){
	fmt.Println("cat is eatting")
}

func (c Cat)Talk(){
	fmt.Printf("I am cat,my name is %s\n",c.name)
}

func main(){
	d := Dog{
		name: "Lily",
	}
	c := Cat{
		name: "Lucy",
	}

	var a1,a2 Animal
	a1 = &d //是d的指针实现了这个接口所以需要传入一个接口，如果a1=d就是错误的
	a1.Eat()
	a1.Talk()
	a2 = c
	a2.Eat()
	a2.Talk()
}

终端输出:
dog is eatting
I am dog,my name is Lily
cat is eatting
I am cat,my name is Lucy
```

###### 7.3.8 多接口

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
}

type Pet interface{
	peiban()
}

type Dog struct{
	Name string
}


func (d *Dog)Eat(){
	fmt.Println("Dog is eatting")
}

func (d *Dog)peiban(){
	fmt.Println("Dog peiban people")
}

func main(){
	var a Animal
	var p Pet
	var d *Dog = &Dog{
		Name: "maomao",
	}
	a = d
	p = d
	a.Eat()
	p.peiban()
}
终端输出:
API server listening at: 127.0.0.1:24037
Dog is eatting
Dog peiban people
```

###### 7.3.9 接口嵌套

尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。

```go
package main

import (
	"fmt"
)

type Animal interface{
	Eat()
}

type Pet interface{
	peiban()
}

type Lovely interface{
	Animal
	Pet
} 

type Dog struct{
	Name string
}


func (d *Dog)Eat(){
	fmt.Println("Dog is eatting")
}

func (d *Dog)peiban(){
	fmt.Println("Dog peiban people")
}

func main(){
	var l Lovely
	var d *Dog = &Dog{
		Name: "maomao",
	}
	l = d
	l.Eat()
	l.peiban()
}
终端输出:
API server listening at: 127.0.0.1:19039
Dog is eatting
Dog peiban people

Lovely就是一个新接口，包含了Animal和Pet的接口签名方法
Dog的指针实现了Animal和Pet接口，所以他就实现了Lovely接口
```

###### 7.3.10 接口零值

```go
package main

import (
	"fmt"
)

type Test interface{
	Test()
}

type A struct{

}

func (a A)Test(){
	fmt.Println(true)
}

func main(){
	var t Test
	if t == nil {
		a := A{}
		t = a
		a.Test()
	}
}

终端输出:
API server listening at: 127.0.0.1:21973
true
```

#### 8. IO

##### 8.1 格式化输⼊

 从终端获取⽤户的输⼊

fmt.Scanf(format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和

 格式化输出⼀致

fmt.Scan(a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，空格和换⾏符

 作为分隔符

fmt.Scanln(a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，空格作为分隔符，

遇到换⾏符结束

```go
import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanf("%s%s",&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//遇到换行符事结束用户输入

package main

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scan(&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//输入完成后结束输出
package main

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanln(&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//遇到换行符就结束用户输入

import (
	"fmt"
)

var str1,str2 string

func testInput(){
	fmt.Scanf("%s\n%s",&str1,&str2)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
}
//\n可以占位\n所以可以隔行输入
```

##### 8.2 从字符串格式化输入

fmt.Sscanf(str, format string, a…interface{}): 格式化输⼊，空格作为分隔符，占位符和

 格式化输出⼀致

fmt.Sscan(str string, a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，

空格和换⾏符作为分隔符

fmt.Sscanln(str string, a …interface{}): 从终端获取⽤户输⼊，存储在**Scanln**中的参数⾥，

空格作为分隔符，遇到换⾏符结束

```go
package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscanf(str,"%s%s%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出：
API server listening at: 127.0.0.1:20797
hello
world
!


package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanf(str,"%s%s%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
hello

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanf(str,"%s\n%s\n%s",&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:18615
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:34727
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}

终端输出:
API server listening at: 127.0.0.1:10148
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\tworld\t!"

func testInput(){
	fmt.Sscan(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}


package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello world !"

func testInput(){
	fmt.Sscanln(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:48954
hello
world
!

package main

import (
	"fmt"
)

var str1,str2,str3 string
var str = "hello\nworld\n!"

func testInput(){
	fmt.Sscanln(str,&str1,&str2,&str3)	
}


func main(){
	testInput()
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
终端输出:
API server listening at: 127.0.0.1:42289
hello

```

##### 8.3 格式化输出

fmt.Sprintf(format string, a…interface{}): 格式化并返回字符串

fmt.Sprintln(a …interface{}): 把零个或多个变量按空格进⾏格式化并换⾏，返回字符串

fmt.Sprint(a …interface{}): 把零个或多个变量按空格进⾏格式化，返回字符串

```go
package main

import (
	"fmt"
)

func testSprint(){
	str := fmt.Sprintln("hello world!")
	fmt.Println(str)
	str = fmt.Sprintf("%s hello world!","hehe")
	fmt.Println(str)
	str = fmt.Sprint("hello world!")
	fmt.Println(str)
}

func main(){
	testSprint()
}
终端输出:
API server listening at: 127.0.0.1:9921
hello world!

hehe hello world!
hello world!


```

##### 8.4 终端输⼊输出

终端其实是⼀个⽂件 

终端相关⽂件的实例

os.Stdin：标准输⼊的⽂件实例，类型为*File

os.Stdout：标准输出的⽂件实例，类型为*File

os.Stderr：标准错误输出的⽂件实例，类型为*File

```go
package main

import (
	"fmt"
	"os"
)

var str1,str2 string

func testOs(){
	fmt.Fscanf(os.Stdin,"%s%s",&str1,&str2)
	fmt.Fprintf(os.Stdout,"%s %s\n",str1,str2)
}

func main(){
	testOs()
}


```

##### 8.5 格式化输出

fmt.Fprintf(fifile, format string, a…interface{}): 格式化输出，并写⼊到⽂件中

fmt.Fprintln(fifile, a …interface{}): 把零个或多个变量写⼊到⽂件中， 并换⾏

fmt.Fprint(fifile, a …interface{}): 把零个或多个变量写⼊到⽂件

```go
import (
	"fmt"
	"os"
)

func testPrint(){
	var str1,str2 string = "hello world","hehe"
	fmt.Fprint(os.Stdout,str1," ",str2,"\n")//不能换行所以要加换行符
	fmt.Fprintf(os.Stdout,"%s %s\n",str1,str2) //格式化输出到文件
	fmt.Fprintln(os.Stdout,str1,str2)//格式化输出并换行
}

func main(){
	testPrint()
}

终端输出:
API server listening at: 127.0.0.1:40053
hello world hehe
hello world hehe
hello world hehe
```

##### 8.7 带缓冲区的读写

```go
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println("please input some things:")
	str,err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Fprintf(os.Stdout,"you input is %s\n",str)
	}
}

终端输出:
please input some things:
hello world
you input is hello world
```

##### 8.8 命令行参数获取

```go
package main

import (
	"fmt"
	"os"
)

func testArgs(){
	if len(os.Args) >1 {
		fmt.Println(len(os.Args))
		for i := 1;i < len(os.Args);i++{
			fmt.Println(os.Args[i])
		}
	}
}

func main(){
	testArgs()
}
终端输出:
hello
-p
11
```

使⽤flag包获取命令⾏参数

```go
package main

import (
	"fmt"
	"flag"
)

var length int
var charset string

func ParseArgs(){
	flag.IntVar(&length,"l",16,"l 生成密码的长度")
	flag.StringVar(&charset,"s","num",`s 制定密码⽣成的字符集, 
	num:只使⽤数字[0-9], 
	char:只使⽤英⽂字母[a-zA-Z], 
	mix: 使⽤数字和字母，
	advance:使⽤数字、字母以及特殊字符`)
	flag.Parse()
}


func main(){
	ParseArgs()
	fmt.Println(length)
	fmt.Println(charset)
}


```

urfave/cli包的使⽤

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "Test"
	app.Action = func(c *cli.Context) error{
		fmt.Println("Hello friend!") 
 		return nil
	}
	app.Run(os.Args)
}
```

获取命令⾏参数

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "Test"
	app.Action = func(c *cli.Context) error{
		fmt.Printf("Hello friend! %q\n",c.Args().Get(0)) 
 		return nil
	}
	app.Run(os.Args)
}

打包：
./main 222
终端输出:
Hello friend! "222"
```

获取选项参数

```go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main(){
	var langueage string
	var recusive bool
	app := cli.NewApp()
	app.Name = "Test"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "language,l",
			Value: "Enlish",
			Usage: "language for the greeting",
			Destination: &langueage,
		},
		cli.BoolFlag{
			Name: "recusive,r",
			Usage: "recusive for the greeting",
			Destination: &recusive,
		},
	}
	app.Action = func(c *cli.Context) error{
		var cmd string
		fmt.Println(c.NArg())
		if c.NArg() > 0{
			cmd = c.Args()[0]
			fmt.Printf("cmd is %s\n",cmd)
		}
		fmt.Println("recusive is ", recusive) 
 		fmt.Println("language is ", langueage) 
 		return nil
	}
	app.Run(os.Args)
}
编译：
./main -l chinese -r
终端输出:
0
recusive is  true
language is  chinses
```

#### 9. 文件读写

##### 9.1 文件打开和读写

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file := "/home/golang/log/test.log"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	defer f.Close()
	var b [1024]byte
	var str []byte
	for {
		n, err := f.Read(b[:])
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}
		str = append(str, b[:n]...)
	}
	fmt.Println(string(str))
}

终端输出:
file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
file is open 32222
file is open 32222

file is open 32222

2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
2020-07-22 10:29:23.008file is open 32222
2020-07-22 10:30:10.033: file is open 32222
2020-07-22 10:49:09.488: DEBUG: file is open 32222
2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
```

文件读取, file.Read和file.ReadAt。读到文件末尾返回:io.EOF

##### 9.2 bufio读取文件

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "/home/golang/log/test.log"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	for {
		context, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("the input is %s", context)
	}

}
终端输出:
the input is file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
the input is file is open 32222
the input is file is open 32222
the input is 
the input is file is open 32222
the input is 
the input is 2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
the input is 2020-07-22 10:29:23.008file is open 32222
the input is 2020-07-22 10:30:10.033: file is open 32222
the input is 2020-07-22 10:49:09.488: DEBUG: file is open 32222
the input is 2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
the input is 2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
the input is 2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
the input is 2020-07-22 11:12:31.926: DEBUG: (/home/golang/lesson/logger/file_logger.go:xiao.com/golang/lesson/logger.(*Filelog).DeBug:72) file is open 32222
the input is 2020-07-22 11:13:02.336: DEBUG: (/home/golang/lesson/logger/file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:17:18.379: DEBUG: (file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:19:48.674: DEBUG: (file_test.go:logger.TestFileLog:14) file is open 32222
the input is 2020-07-22 11:22:01.664: DEBUG: (file_test.go: logger.TestFileLog: 14) file is open 32222
the input is 2020-07-22 11:23:02.104: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) file is open 32222
the input is 2020-07-22 14:24:33.765: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) file is open 32222
the input is 2020-07-22 14:26:50.36: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) ERROR: code 32222
the input is 2020-07-22 14:27:40.152: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-22 15:40:46.841: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-22 16:02:14.879: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
the input is 2020-07-23 14:41:41.522: DEBUG: (filelogs.go:  function:logger.Debug line:20) init logger success
the input is 2020-07-23 14:43:44.555: DEBUG: (filemain.go:  function:main.main line:33) init logger success
the input is 2020-07-23 18:10:53.816: DEBUG: (filemain.go:  function:main.main line:35) init logger success
```

整个文件读取

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/home/golang/log/test.log"
	bufReader, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", string(bufReader))
}
终端输出:

```

读取压缩文件

```go
package main

import (
	"fmt"
	"compress/gzip"
	"bufio"
	"os"
)

func main(){
	file := "/home/golang/log.tar.gz"
	var r *bufio.Reader
	fi,err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s",err)
		return
	}
	defer fi.Close()
	fz,err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open gzip failed, err: %v\n", err)
		return
	}
	r = bufio.NewReader(fz)
	for {
		line,err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line)
	}

}
终端输出:

```

##### 9.3 文件写入



```go
   O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
package main

import (
	"fmt"
	"os"
)

func main(){
	file := "/home/golang/log/test.txt"
	var str string = "hello xiaoxiao" 
	f,err := os.OpenFile(file,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err != nil {
		fmt.Fprintf(os.Stderr,"error:%v",err)
		return
	}
	defer f.Close()
	n,err := f.Write([]byte(str))
	if err != nil {
		fmt.Fprintf(os.Stderr,"error:%v",err)
	}
	fmt.Println(n)
}
终端输出：
14
```

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file := "/home/golang/log/test.txt"
	f,err := os.OpenFile(file,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("ERROR:%s",err)
		return
	}
	defer f.Close()
	var str string = "hello world\n"
	neWrite := bufio.NewWriter(f)
	neWrite.WriteString(str)
	neWrite.Flush()
}
终端输出:
cat file
hello xiaoxiao
hello world
```

写入整个文件的示例

```go
import (
	"fmt"
	"os"

	"io/ioutil"
)

func main(){
	file := "/home/golang/log/test.txt"
	var str string = "hello world\n" 
	err := ioutil.WriteFile(file,[]byte(str),0644)
	if err != nil {
		panic(err.Error())
	}
}
终端输出：
cat file
hello world
由此可见会覆盖文件的内容
```

文件拷贝

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile()(writing int64,err error){
	src := "/home/golang/log/test.txt"
	des := "/home/golang/log/test1.txt" 
	sr,ok := os.Open(src)
	if ok != nil {
		err = ok
		return
	}
	defer sr.Close()
	dest,ok := os.OpenFile(des,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0600)
	if ok != nil{		
		err = ok
		return
	}
	defer dest.Close()
	return io.Copy(dest,sr)
}

func main(){
	fmt.Println(copyFile())
}
终端输出:
12 nil
```

cat 命令实现

```go

package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func cat(){
	if len(os.Args) < 2{
		fmt.Printf("can not find file\n")
		return
	}
	for i := 1;i <len(os.Args);i++{
		file := os.Args[i]
		data,err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("cat %s ERROR: %s",file,err)
			continue
		}
		fmt.Println(file)
		fmt.Println(string(data))
	}
}

func main(){
	cat()
}
执行
cat 
终端输出:
/home/golang/log/test.txt
hello world

/home/golang/log/test.log
file is open 32222file is open 32222file is open 32222file is open 32222file is open 32222
file is open 32222
file is open 32222

file is open 32222

2020-07-22 10:27:44.506file is open 322222020-07-22 10:29:15.998file is open 32222
2020-07-22 10:29:23.008file is open 32222
2020-07-22 10:30:10.033: file is open 32222
2020-07-22 10:49:09.488: DEBUG: file is open 32222
2020-07-22 10:58:47.505: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991)file is open 32222
2020-07-22 10:59:49.187: DEBUG: (/usr/local/go/src/testing/testing.go:testing.tRunner:991) file is open 32222
2020-07-22 11:12:02.319: DEBUG: (/home/golang/lesson/logger/util.go:xiao.com/golang/lesson/logger.GetLineInfo:8) file is open 32222
2020-07-22 11:12:31.926: DEBUG: (/home/golang/lesson/logger/file_logger.go:xiao.com/golang/lesson/logger.(*Filelog).DeBug:72) file is open 32222
2020-07-22 11:13:02.336: DEBUG: (/home/golang/lesson/logger/file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
2020-07-22 11:17:18.379: DEBUG: (file_test.go:xiao.com/golang/lesson/logger.TestFileLog:14) file is open 32222
2020-07-22 11:19:48.674: DEBUG: (file_test.go:logger.TestFileLog:14) file is open 32222
2020-07-22 11:22:01.664: DEBUG: (file_test.go: logger.TestFileLog: 14) file is open 32222
2020-07-22 11:23:02.104: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) file is open 32222
2020-07-22 14:24:33.765: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) file is open 32222
2020-07-22 14:26:50.36: DEBUG: (filefile_logger.go:  function:logger.(*Filelog).DeBug line:82) ERROR: code 32222
2020-07-22 14:27:40.152: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-22 15:40:46.841: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-22 16:02:14.879: DEBUG: (filefile_test.go:  function:logger.TestFileLog line:14) ERROR: code 32222
2020-07-23 14:41:41.522: DEBUG: (filelogs.go:  function:logger.Debug line:20) init logger success
2020-07-23 14:43:44.555: DEBUG: (filemain.go:  function:main.main line:33) init logger success
2020-07-23 18:10:53.816: DEBUG: (filemain.go:  function:main.main line:35) init logger success
```

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)


func cat(buf *bufio.Reader){
	for {
		line,err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
    
}

func Parse(){
	flag.Parse()
	if flag.NArg() < 2{
		fmt.Printf("can not find file\n")
		return
	}
	for i := 1;i <len(os.Args);i++{
		file := os.Args[i]
		f,err := os.Open(file)
		if err != nil {
			fmt.Printf("cat %s ERROR: %s",file,err)
			return
		}
		defer f.Close()	    	
		cat(bufio.NewReader(f))		
	}
}

func main(){
	Parse()
	
}
```

##### 9.4 defer解析

defer原理分析

​			return x

​			返回值= x

​			RET指令

defer原理

​		返回值= x

​		RET指令

​		运行defer

```go
package main

import (
	"fmt"
)


func testDefer1()(x int){
	x = 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer1())
}

终端输出:
API server listening at: 127.0.0.1:33190
6
//已经设置了返回值为x,defer的时候x的值改变

package main

import (
	"fmt"
)


func testDefer2()(int){
	x := 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer2())
}
终端输出:
API server listening at: 127.0.0.1:13897
5

package main

import (
	"fmt"
)


func testDefer2()(y int){
	x := 5
	defer func(){
		x += 1
	}() 
	return x
}

func main(){
	fmt.Println(testDefer2())
}
终端输出:
API server listening at: 127.0.0.1:30377
5
```

#### 10.反射

在计算机科学领域，反射是指一类应用，它们能够自描述和自控制。也就是说，这类应用通过采用某种机制来实现对自己行为的描述（self-representation）和监测（examination），并能根据自身行为的状态和结果，调整或修改应用所描述行为的状态和相关的语义。

每种语言的反射模型都不同，并且有些语言根本不支持反射。Golang语言实现了反射，反射机制就是在运行时动态的调用对象的方法和属性，官方自带的reflect包就是反射相关的，只要包含这个包就可以使用。

Golang的gRPC也是通过反射实现的。

- 变量包括（type, value）两部分

  - 理解这一点就知道为什么nil != nil了

- type 包括 static type和concrete type. 简单来说 static type是你在编码是看见的类型(如int、string)，concrete type是runtime系统看见的类型

- 类型断言能否成功，取决于变量的concrete type，而不是static type. 因此，一个 reader变量如果它的concrete type也实现了write方法的话，它也可以被类型断言为writer.

- Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型包含该两个指针,一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。

- ```
  ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
  TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
  已知类型后转换为其对应的类型的做法如下，直接通过Interface方法然后强制转换：
  	realValue := value.Interface().(已知的类型)
      说明
      转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
      转换的时候，要区分是指针还是指
      也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
      
      
   
  ```

```go
package main

import (
	"fmt"
	"os"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefect(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			fmt.Fprintf(os.Stdout, "%s=%v\n", field.Tag.Get("ini"), v.Field(i))
		}
	default:
		fmt.Printf("unkonwn type")
	}
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefect(student)
}
中断输出：
API server listening at: 127.0.0.1:33535
name=chun
sex=2
```

##### 10.1 获取value和type

reflect 包提供反射机制

`reflect.Type` 表示 `interface{}` 的具体类型，而 `reflect.Value` 表示它的具体值。`reflect.TypeOf()` 和 `reflect.ValueOf()` 两个函数可以分别返回 `reflect.Type` 和 `reflect.Value`。这两种类型是我们创建查询生成器的基础

```go
1.reflect.ValueOf()

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectValue(a interface{}) {
	v := reflect.ValueOf(a) //获取value
	t := v.Type() //通过v.Type()来获取Type
	fmt.Printf("the type is %s\n", t)
	fmt.Printf("the value is %v\n", v)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectValue(student)
}
终端输出:
API server listening at: 127.0.0.1:33121
the type is main.Student
the value is {chun 2}

2.reflect.TypeOf()
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectType(a interface{}) {
	t := reflect.TypeOf(a) //获取type
	fmt.Printf("the type is %s\n", t)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectType(student)
}
终端输出:
API server listening at: 127.0.0.1:34516
the type is main.Student
```

##### 10.2 reflect.Kind

`reflect` 包中还有一个重要的类型：[`Kind`](https://golang.org/pkg/reflect/#Kind)。

在反射包中，`Kind` 和 `Type` 的类型可能看起来很相似，但在下面程序中，可以很清楚地看出它们的不同之处。

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectKind(a interface{}) {
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	k := v.Kind() //获取type
	fmt.Printf("the type is %s\n", t)
	fmt.Printf("the kind is %s\n", k)
}

func main() {
	var student Student = Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectKind(student)
}
终端输出:
API server listening at: 127.0.0.1:24274
the type is main.Student
the kind is struct

```

我想你应该很清楚两者的区别了。`Type` 表示 `interface{}` 的实际类型（在这里是 **`main.Order`**)，而 `Kind` 表示该类型的特定类别（在这里是 **`struct`**）。

```go
const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
```

##### 10.3 通过反射设置变量值

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectSetValue(a interface{}) {
	t := reflect.ValueOf(a)
	t.Elem().FieldByName("Name").SetString("chunyan") //修改字段的值
}

func main() {
	var student *Student = &Student{
		Name: "chun",
		Sex:  2,
	}
	testRefectSetValue(student)
	fmt.Println(*student)
}
终端输出:
API server listening at: 127.0.0.1:35319
{chunyan 2}
```

```go 
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectSetValue(a interface{}) {
	t := reflect.ValueOf(a)
	t.Elem().SetInt(10)
}

func main() {
	a := 14
	testRefectSetValue(&a)
	fmt.Println(a)
}
终端输出:
API server listening at: 127.0.0.1:44567
10
```

以上的情况有个问题,一段代码来展示一下

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectSetValue(a interface{}) {
	t := reflect.ValueOf(a)
	t.Elem().SetFloat64(10) //假如我不知道a的类型，用float
}

func main() {
	a := 14
	testRefectSetValue(&a)
	fmt.Println(a)
}
以上会触发panic
终端输出:
API server listening at: 127.0.0.1:39021
panic: reflect: call of reflect.Value.SetFloat on int Value

goroutine 1 [running]:
reflect.Value.SetFloat(0x4bb560, 0xc0000ae010, 0x182, 0x4024000000000000)
	/usr/local/go/src/reflect/value.go:1590 +0x130
main.testRefectSetValue(0x4b6bc0, 0xc0000ae010)
	/home/golang/writefile/lesson1/main.go:15 +0xa9
main.main()
	/home/golang/writefile/lesson1/main.go:20 +0x6f
Process exiting with code: 
```

改写一下:

```go
package main

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRefectSetValue(a interface{}) {

	v := reflect.ValueOf(a)
	t := v.Type().Kind()
	if t != reflect.Ptr {
		fmt.Println("input is not address!")
		return
	}
	k := v.Elem().Kind()
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.Elem().SetInt(10)
	case reflect.Float32, reflect.Float64:
		v.Elem().SetFloat(14.3)
	default:
		fmt.Println("unkown type")
	}
}

func main() {
	a := 14
	testRefectSetValue(a)
	testRefectSetValue(&a)
	fmt.Println(a)
	b := 11.23
	testRefectSetValue(&b)
	fmt.Println(b)
	s := "hello"
	testRefectSetValue(&s)
	fmt.Println(s)
}
终端输出:
API server listening at: 127.0.0.1:8751
input is not address!
10
14.3
unkown type
hello
```

##### 10.4 结构体的反射

```go
10.4.1 获取结构体的类型和值
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRelectStruct(a interface{}) {
	t := reflect.TypeOf(a) //获取实际的类型
	v := reflect.ValueOf(a) //获取实际的值
	fmt.Println(t)
	fmt.Println(v)
}

func main() {
	a := Student{
		Name: "chun",
		Sex:  2,
	}
	testRelectStruct(a)
}
终端输出:
API server listening at: 127.0.0.1:38478
main.Student
{chun 2}

10.4.2 获取字段和字段的值
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRelectStruct(a interface{}) {
	t := reflect.TypeOf(a)
	if t.Name() != "Student" { //t.Name() 获取类型的名字
		fmt.Printf("%s is not Type Student", t.Name())
		return
	}
	v := reflect.ValueOf(a)
	for i := 0; i < v.NumField(); i++ {
		Tfield, Vfield := t.Field(i), v.Field(i) //Tfield 是a 反射字段类型相关，Vfield 是a 反射字段值相关
		fmt.Printf("field of %s is %s,value of field %s is %v\n", Tfield.Name, Tfield.Type.Kind(), Tfield.Name, Vfield)
	}
}

func main() {
	a := Student{
		Name: "chun",
		Sex:  2,
	}
	testRelectStruct(a)
}
终端输出:
API server listening at: 127.0.0.1:16675
field of Name is string,value of field Name is chun
field of Sex is int,value of field Sex is 2

10.4.3 修改字段值
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func testRelectStruct(a interface{}) {
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Ptr {
		fmt.Printf("the agrs must address")
		return
	}
	if t.Elem().Name() != "Student" { //t.Name() 获取类型的名字
		fmt.Printf("%s is not Type Student", t.Name())
		return
	}
	v := reflect.ValueOf(a)
	for i := 0; i < v.Elem().NumField(); i++ {
		Tfield, Vfield := t.Elem().Field(i), v.Elem().Field(i) //Tfield 是a 反射字段类型相关，Vfield 是a 反射字段值相关
		k := Tfield.Type.Kind()
		switch k {
		case reflect.String:
			Vfield.SetString("chunyan")
		case reflect.Int:
			Vfield.SetInt(2)
		default:
		}
	}
}

func main() {
	a := &Student{
		Name: "chun",
		Sex:  2,
	}
	testRelectStruct(a)
	fmt.Println(*a)
}

10.4.5 获取方法
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func (s *Student) Study(course string) {
	fmt.Printf("%s is study %s", s.Name, course)
}

func (s *Student) Score(course string, score float64) {
	fmt.Printf("%s,score of %s is %f", s.Name, course, score)
}

func testRelectStruct(a interface{}) {
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Ptr {
		fmt.Printf("the agrs must address")
		return
	}
	if t.Elem().Name() != "Student" { //t.Name() 获取类型的名字
		fmt.Printf("%s is not Type Student", t.Name())
		return
	}
	for i := 0; i < t.NumMethod(); i++ {
		met := t.Method(i)
		fmt.Printf("%v,%v\n", met.Name, met.Type)
	}
}

func main() {
	a := &Student{
		Name: "chun",
		Sex:  2,
	}
	testRelectStruct(a)
}
终端输出:
API server listening at: 127.0.0.1:22754
Score,func(*main.Student, string, float64)
Study,func(*main.Student, string)

10.4.6 方法调用
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `ini:"name"`
	Sex  int    `ini:"sex"`
}

func (s *Student) Study(course string) {
	fmt.Printf("%s is study %s\n", s.Name, course)
}

func (s *Student) Score(course string, score float64) {
	fmt.Printf("%s,score of %s is %f\n", s.Name, course, score)
}

func testRelectStruct(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	if t.Kind() != reflect.Ptr {
		fmt.Printf("the agrs must address")
		return
	}
	if t.Elem().Name() != "Student" { //t.Name() 获取类型的名字
		fmt.Printf("%s is not Type Student", t.Name())
		return
	}
	met1 := v.MethodByName("Study")

	var args1 []reflect.Value
	var course string = "English"
	args1 = append(args1, reflect.ValueOf(course))
	met1.Call(args1)

	met2 := v.MethodByName("Score")
	var score float64 = 98.5
	args1 = append(args1, reflect.ValueOf(score))
	met2.Call(args1)
}

func main() {
	a := &Student{
		Name: "chun",
		Sex:  2,
	}
	testRelectStruct(a)
}
终端输出:

```

##### 10.5 反射tag

```go
package main

import (
	"fmt"
	"os"
	"reflect"
)

type Student struct {
	Name string `yaml:"name"`
	Sex  int    `yaml:"sex"`
}

func SexToString(sex int) string {
	switch sex {
	case reflect.ValueOf(1):
		return "boy"
	case reflect.ValueOf(2):
		return "girl"
	default:
		return "boy"
	}
}

func testReflectStruct(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	if t.Elem().Name() != "Student" {
		fmt.Fprintf(os.Stderr, "%v is not Student type", a)
		return
	}
	var info map[string]interface{} = map[string]interface{}{}
	for i := 0; i < t.Elem().NumField(); i++ {
		var val interface{}
		field := t.Elem().Field(i)
		val = v.Elem().Field(i)
		key := field.Tag.Get("yaml")
		info[key] = fmt.Sprintf("%v", val)
	}
	fmt.Printf("%v", info)
}

func main() {
	stu := Student{
		Name: "chun",
		Sex:  2,
	}
	testReflectStruct(&stu)
}

终端输出：
API server listening at: 127.0.0.1:24685
map[name:chun sex:2]
```

#### 11.测试

##### 11.1 单元测试

testing

A. testing包提供了自动化测试相关的框架

B. 支持单元测试和压力测试

```go
package main

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a int = 10
	var b int = 20
	c := Add(a, b)
	if c != 30 {
		t.Fatalf("%d + %d = %d", a, b, c)
	}
	t.Logf("%d", c)
}

func TestSub(t *testing.T) {
	var a int = 20
	var b int = 10
	c := Sub(a, b)
	if c != 10 {
		t.Fatalf("%d - %d = %d", a, b, c)
	}
	t.Logf("%d", c)
}

运行：go test -v
终端输出:
golang@python:~/lesson/lesson19/unit_test$ go test -v
=== RUN   TestAdd
    TestAdd: add_test.go:14: 30
--- PASS: TestAdd (0.00s)
=== RUN   TestSub
    TestSub: add_test.go:24: 10
--- PASS: TestSub (0.00s)
PASS
ok      xiao.com/golang/lesson/lesson19/unit_test       0.002s
```

##### 11.2 压力测试

基准测试或压力测试必须以 Benchmark开头，并且只有参数，

类型是*Testing.B

```go
package main

import (
	"testing"
)

func BenchmarkAdd(t *testing.B) {
	for i := 0; i < 100000; i++ {
		var a int = 10
		var b int = 20
		c := Add(a, b)
		if c != 30 {
			t.Fatalf("%d + %d = %d is wrong", a, b, c)
		}
		t.Logf("%d", c)
	}
}

运行go test -bench .:
golang@python:~/lesson/lesson19/unit_test$ go test -bench .
goos: linux
goarch: amd64
pkg: xiao.com/golang/lesson/lesson19/unit_test
BenchmarkAdd-2            526819              2073 ns/op
--- BENCH: BenchmarkAdd-2
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
        ... [output truncated]
PASS
ok      xiao.com/golang/lesson/lesson19/unit_test       1.125s
```



##### 11.3 go test 命令



```go
package main

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a int = 10
		var d int = 20
		c := Add(a, d)
		if c != 30 {
			b.Fatalf("%d + %d = %d is wrong", a, d, c)
		}
		b.Logf("%d", c)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 10
		d := 20
		c := Sub(d, a)
		if c != 10 {
			b.Fatalf("%d - %d = %d is wrong", a, d, c)
		}
		b.Logf("%d", c)
	}
}
go test -bench BenchmarkSub
终端输出:
golang@python:~/lesson/lesson19/unit_test$ go test -bench BenchmarkSub 
goos: linux
goarch: amd64
pkg: xiao.com/golang/lesson/lesson19/unit_test
BenchmarkSub-2            493924              2110 ns/op
--- BENCH: BenchmarkSub-2
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
        ... [output truncated]
PASS
ok      xiao.com/golang/lesson/lesson19/unit_test       1.072s

go test -bench . 测试所用例
golang@python:~/lesson/lesson19/unit_test$ go test -bench .
goos: linux
goarch: amd64
pkg: xiao.com/golang/lesson/lesson19/unit_test
BenchmarkAdd-2            563209              4177 ns/op
--- BENCH: BenchmarkAdd-2
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
    bench_test.go:15: 30
        ... [output truncated]
BenchmarkSub-2            474064              2125 ns/op
--- BENCH: BenchmarkSub-2
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
    bench_test.go:27: 10
        ... [output truncated]
PASS
ok      xiao.com/golang/lesson/lesson19/unit_test       4.388s
```

go test file file 单元测试某某文件

##### 11.4 dlv调试

A. go get github.com/derekparker/delve/cmd/dlv

B. 默认安装到GOPATH的bin目录下

C. 把安装目录设置到PATH环境变量中



```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		var i int
		var curTime time.Time
		time.Sleep(5 * time.Second)
		i++
		curTime = time.Now()
		fmt.Printf("run %d count,cur time；%v\n", i, curTime)
	}
}

Delve使用
A. dlv命令， dlv debug 包的路径或源代码路径 dlv attach 进程ID
B. Dlv会编译我们的程序，然后进入调试界面
C. 进入调试界面后，就可以一步一步的执行的我们代码了
b 设置断点
c 开始调试
next 
step
p 打印变量
exit 退出
goruntine 查看携程 多线程调试
goruntine 携程名称

package main

import (
	"fmt"
	"time"
)

func isPerm(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func test1(c chan int) {
	var i int
	for {
		i++
		result := isPerm(i)
		if !result {
			c <- i
		}
	}
}

func test2(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	var c chan int = make(chan int)
	go test1(c)
	go test2(c)

	time.Sleep(time.Hour)
}
```

#### 12  panic defer recover

##### 12.1 panic 是程序崩溃

```go
package main

import (
	"fmt"
)

func main() {
	var b int = 10
	if b != 5 {
		panic("%d is not 10")
	}
	fmt.Printf("%d", b)
}
终端输出:
API server listening at: 127.0.0.1:13505
panic: %d is not 10

goroutine 1 [running]:
main.main()
	/home/golang/lesson/lesson20/panic_test/main.go:10 +0x44
Process exiting with code: 0
```

当程序运行出现严重逻辑错误时，我们需要触发panic,来终止程序，不让他继续运行

##### 12.2 defer

defer 前面的章节已经有介绍，panic允许defer执行后再panic。

```go
package main

import (
	"fmt"
)

func main() {
	var b int = 10
	defer func() {
		fmt.Println("input is ", b)
	}()
	if b != 5 {
		panic("b is not 5")
	}
}




API server listening at: 127.0.0.1:30062
input is  10
panic: b is not 5

goroutine 1 [running]:
main.main()
	/home/golang/lesson/lesson20/panic_test/main.go:13 +0xa2
Process exiting with code: 0
```

##### 12.3 recover

recover 恢复panic,使程序继续运行，recover必须再defer 的func中，recover的返回值为panic的返回值。

```go
package main

import (
	"fmt"
)

func defFunc(b int) {
	fmt.Println("input is ", b)
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic that is: ", r)
	}
}

func main() {
	var b int = 10
	defer defFunc(b)
	if b != 5 {
		panic("b is not 5")
	}
}
终端输出:
API server listening at: 127.0.0.1:17739
input is  10
Recovered from panic that is:  b is not 5
Process exiting with code: 0
```

#### 13 多线程

##### 13.1 并发和并行

并发：同一时间段类执行多个操作。

并行：同一时刻执行多个操作。

多线程

A. 线程是由操作系统进行管理，也就是处于内核态。

B. 线程之间进行切换，需要发生用户态到内核态的切换。

C. 当系统中运行大量线程，系统会变的非常慢。

D. 用户态的线程，支持大量线程创建。也叫协程或goroutine。

##### 13.2 goroutine

goroutine是go的一个轻量级的并发方式，成为协程，以关键字go 加函数方法名

```go
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("the goroutine")
}

func main() {
	go hello() //go 开启一个新的进程，主进程并不会等待协程等待完毕再退出，所以需要等待时间延迟退出。
	fmt.Println("the main")
	time.Sleep(time.Second)
}
终端输出:
API server listening at: 127.0.0.1:25262
the main
the goroutine
```

```go
import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello this is %d", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	time.Sleep(time.Second)
}
终端输出:
API server listening at: 127.0.0.1:9102
hello this is %d 3
hello this is %d 9
hello this is %d 0
hello this is %d 4
hello this is %d 1
hello this is %d 5
hello this is %d 2
hello this is %d 6
hello this is %d 7
hello this is %d 8

可以看到并发的特点是无序的
```

启动多个goroutine

```go
package main

import (
	"fmt"
	"time"
)

func printInt() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printChar() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printInt() //启动一个goroutine,每250 毫秒打印一次数字
	go printChar() //启动另一个goroutine,每400毫秒打印一次字节
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("main terminated")
}
终端输出:
API server listening at: 127.0.0.1:14728
0 a 1 2 b 3 c 4 d 
main terminated
```

##### 13.2 runtime

多核控制

A. 通过runtime包进行多核设置

B. GOMAXPROCS设置当前程序运行时占用的cpu核数

C. NumCPU获取当前系统的cpu核数

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func cale() {
	for {
		i++
	}
}

func main() {
	cpu := runtime.NumCPU() //获取CPU的个数
	fmt.Printf("cpu: %d\n", cpu)

	runtime.GOMAXPROCS(2) //设置CPU的使用个数
	for i := 0; i < 10; i++ {
		go cale()
	}
	time.Sleep(time.Minute)
}
终端输出:
API server listening at: 127.0.0.1:7402
cpu: 2
再看看cpu的使用情况呢，会发生什么呢
```

##### 13.4 Goroutine原理浅析

A. 一个操作系统线程对应用户态多个goroutine

B. 同时使用多个操作系统线程

C. 操作系统线程对goroutine是多对多关系，即M:N

模型抽象

A. 操作系统线程： M 内核线程

B. 用户态线程（goroutine）: G

C. 上下文对象：P （相当于一个处理器）

![image-20200818153802109](F:/111/golang/assets/image-20200818153802109.png)

goroutine调度

![image-20200818153908364](F:/111/golang/assets/image-20200818153908364.png)

以上这个图讲的是两个线程(内核线程)的情况。一个**M**会对应一个内核线程，一个**M**也会连接一个上下文**P**，一个上下文**P**相当于一个“处理器”，一个上下文连接一个或者多个Goroutine。为了运行goroutine，线程必须保存上下文。

上下文P(Processor)的数量在启动时设置为`GOMAXPROCS`环境变量的值或通过运行时函数`GOMAXPROCS()`。通常情况下，在程序执行期间不会更改。上下文数量固定意味着只有固定数量的线程在任何时候运行Go代码。我们可以使用它来调整Go进程到个人计算机的调用，例如4核PC在4个线程上运行Go代码。

图中P正在执行的`Goroutine`为蓝色的；处于待执行状态的`Goroutine`为灰色的，灰色的`Goroutine`形成了一个队列`runqueues`。

Go语言里，启动一个goroutine很容易：go function 就行，所以每有一个go语句被执行，runqueue队列就在其末尾加入一个goroutine，一旦上下文运行goroutine直到调度点，它会从其runqueue中弹出goroutine，设置堆栈和指令指针并开始运行goroutine。

一个很简单的例子就是系统调用`sysall`，一个线程肯定不能同时执行代码和系统调用被阻塞，这个时候，此线程M需要放弃当前的上下文环境P，以便可以让其他的`Goroutine`被调度执行。

![image-20200818162049569](F:/111/golang/assets/image-20200818162049569.png)

如上图左图所示，M0中的G0执行了syscall，然后就创建了一个M1(也有可能来自线程缓存)，（转向右图）然后M0丢弃了P，等待syscall的返回值，M1接受了P，将·继续执行`Goroutine`队列中的其他`Goroutine`。

当系统调用syscall结束后，M0会“偷”一个上下文，如果不成功，M0就把它的Gouroutine G0放到一个全局的runqueue中，将自己置于线程缓存中并进入休眠状态。全局runqueue是各个P在运行完自己的本地的Goroutine runqueue后用来拉取新goroutine的地方。P也会周期性的检查这个全局runqueue上的goroutine，否则，全局runqueue上的goroutines可能得不到执行而饿死。

#### 14 channel

信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收。

##### 14.1 信道的声明

所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。

`chan T` 表示 `T` 类型的信道。

信道的零值为 `nil`。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 `make` 来定义信道。

```go
package main

import (
	"fmt"
	"time"
)

func inputChan(c chan int) {
	c <- 100 //通过信道接收值
}

func outputChan(c chan int) {
	data := <-c //通过信道发送值
	fmt.Println(data)
}

func main() {
	var c chan int //未初始化的channel 是零值nil
	if c == nil{
		c = make(chan int)
        fmt.Printf("type c is %T",c)
	}
	go inputChan(c)
	go outputChan(c)
	time.Sleep(time.Second)
}
终端输出:
API server listening at: 127.0.0.1:19571
type c is chan int
100
```

通过信道进行发送和接收

`c <-  100`

`data := <- c`

信道旁的箭头方向指定了是发送数据还是接收数据。

在第一行，箭头对于 `a` 来说是向外指的，因此我们读取了信道 `a` 的值，并把该值存储到变量 `data`。

在第二行，箭头指向了 `a`，因此我们在把数据写入信道 `a`。

发送与接收默认是阻塞的

发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 `Go` 协程从信道读取到数据，才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。

信道的这种特性能够帮助` Go` 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。

前面我们写过一个`goroutine`的代码，需要用到`time.Sleep()`来等待协程结束

```go
package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("hello goroutine")
}

func main(){
	go hello()
	fmt.Println("main funcion end")
	time.Sleep(time.Second)
}
终端输出:
API server listening at: 127.0.0.1:42299
main funcion end
hello goroutine
```

需要使用休眠来等待`goroutine`执行结束。不然主线程不会等待`goroutine`执行完就退出，然后导致`goroutine`也跟着退出。

现在我们使用channel来改写代码

```go
import (
	"fmt"
)

func hello(c chan bool){
	fmt.Println("hello goroutine")
	c <- true
}

func main(){
	c := make(chan bool)
	go hello(c)
	fmt.Println("main function end")
	<-c
}
终端输出:
API server listening at: 127.0.0.1:28527
main function end
hello goroutine
```

由于channel是双向阻塞的，在`<-c` 没有取到值时，它会一直阻塞，等到`c <- true` 写入信道以后，`<-c` 读取到值了，程序就不在阻塞。

再写一个列子来更好的理解channel的阻塞

```go
package main

import (
	"fmt"
	"time"
)

func hello(c chan bool){
	fmt.Println("hello goroutine")
	time.Sleep(2 * time.Second)
	fmt.Println("2 second later")
	c <- true
}

func main(){
	c := make(chan bool)
	go hello(c)
	<-c
	fmt.Println("main function end")
	
}
终端输出:
先输出：hello goroutine
2秒后再输出：2 second later
然后把true写入信道
main线程读取到信道中的值
打印：main function end
main执行完毕，程序结束
```

再写一个例子我们来更好的理解`channel`。

```go
package main

import (
	"fmt"
)

func squreFunc(squre chan int,num int){
	var sum int
	for num != 0 {
		d :=num % 10
		sum += d *d
		num /= 10
	}
	squre <- sum
}

func cubesFunc(cube chan int,num int){
	var sum int
	for num != 0 {
		d :=num % 10
		sum += d * d *d
		num /= 10
	}
	cube <- sum
}

func main(){
	num := 144
	squre := make(chan int)
	cube := make(chan int)
	go squreFunc(squre,num)
	go cubesFunc(cube,num)
	squres,cubes := <-squre,<-cube
	fmt.Println("Final output",squres+cubes)
}
```

这个例子中，我们求一个数每位数的平方的和，和每位数上的立方的和，再将两位数上的和相加，协程`go squreFunc(squre,num)`，`go cubesFunc(cube,num)`分别执行求平方和和求立方方和，并将值写入写入信道`squre`，`cube`。`squres,cubes := <-squre,<-cube`，`squres`和`cubes`分别读取信道`squre`，`cube`，`squres + cubes`后打印

`终端输出`：

`API server listening at: 127.0.0.1:25540
Final output 82`

##### 14.2 死锁

使用信道需要考虑的一个重点是死锁。当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁。

同理，当有 Go 协程等着从一个信道接收数据时，我们期望其他的 Go 协程会向该信道写入数据，要不然程序就会触发 panic。

```go
package main

import (
	_"fmt"
)

func main(){
	var c chan int = make(chan int)
	<-c 
}
```

由于程序只有读取`channel`  `c` 没有其他协程写入`channel`  `c` 的操作，就会形成死锁，触犯`panic`

`终端输出`：

`fatal error: all goroutines are asleep - deadlock!`

`goroutine 1 [chan receive]:
main.main()
	/home/golang/writefile/lesson1/main.go:9 +0x52
Process exiting with code: 0`

同样如果只有写入，没有读取也会死锁

```go
package main

import (
	_"fmt"
)

func main(){
	var c chan int = make(chan int)
	c <- 5
}
```

程序只往信道里面写入了数据，但是没有协程去读取信道，因此也会形成死锁，触犯panic

`终端输出`：

`fatal error: all goroutines are asleep - deadlock!`

`goroutine 1 [chan send]:
main.main()
	/home/golang/writefile/lesson1/main.go:9 +0x55
Process exiting with code: 0`

如果在同一个协程里面写入和读取也会形成死锁，make(chan int),是一个不带缓冲区的信道，长度为0,也就是没有容量给它在一个线程种写入和读取

```go
package main

import (
	_"fmt"
)

func main(){
	var c chan int = make(chan int)
	c <- 5
	<-c
}

```

`终端输出`：

`API server listening at: 127.0.0.1:29224
fatal error: all goroutines are asleep - deadlock!`

`goroutine 1 [chan send]:
main.main()
	/home/golang/writefile/lesson1/main.go:9 +0x55
Process exiting with code: 0`

##### 14.3 单向信道

我们目前讨论的信道都是双向信道，即通过信道既能发送数据，又能接收数据。其实也可以创建单向信道，这种信道只能发送或者接收数据。

```go
package main

import (
	"fmt"
)

func sendData(sendch chan<-int){
	sendch <- 10
}

func main(){
	var sendch chan int = make(chan <- int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

```

上面程序的第 10 行，我们创建了唯送（Send Only）信道 `sendch`。`chan<- int` 定义了唯送信道，因为箭头指向了 `chan`。在第 12 行，我们试图通过唯送信道接收数据，于是编译器报错：

```
main.go:11: invalid operation: <-sendch (receive from send-only type chan<- int)
```

一切都很顺利，只不过一个不能读取数据的唯送信道究竟有什么意义呢？

这就需要用到信道转换（Channel Conversion）了。把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。

```
package main

import (
	"fmt"
)

func sendData(sendch chan<-int){
	sendch <- 10
}

func main(){
	var sendch chan int = make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}
```

在上述程序的第 `12` 行，我们创建了一个双向信道 `cha1`。在第 13 行 `cha1` 作为参数传递给了 `sendData` 协程。在第 57行，函数 `sendData` 里的参数 `sendch chan<- int` 把 `cha1` 转换为一个唯送信道。于是该信道在 `sendData` 协程里是一个唯送信道，而在 Go 主协程里是一个双向信道。该程序最终打印输出 `10`。

`终端输出`：

`API server listening at: 127.0.0.1:23730
10`

##### 14.4 for range 遍历信道

数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。

当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。

```
v, ok := <- ch
```

上面的语句里，如果成功接收信道所发送的数据，那么 `ok` 等于 true。而如果 `ok` 等于 false，说明我们试图读取一个关闭的通道。从关闭的信道读取到的值会是该信道类型的零值。例如，当信道是一个 `int` 类型的信道时，那么从关闭的信道读取的值将会是 `0`

```go
package main

import (
	"fmt"
)

func producer(chnl chan int){
	for i := 0;i < 10;i++{
		chnl <- i
	}
	close(chnl)
}

func main(){
	ch := make(chan int)
	go producer(ch)
	for i := 0;i<10;i++ {
		data,ok := <-ch
		if !ok {
			break
		}
		fmt.Println(data)
	}
}
```

在上述的程序中，`producer` 协程会从 0 到 9 写入信道 `chn1`，然后关闭该信道。主函数有一个无限的 for 循环（第 16 行），使用变量 `ok`（第 18 行）检查信道是否已经关闭。如果 `ok` 等于 false，说明信道已经关闭，于是退出 for 循环。如果 `ok` 等于 true，会打印出接收到的值和 `ok` 的值。

```
API server listening at: 127.0.0.1:26491
0
1
2
3
4
5
6
7
8
9
```

for range

for range 循环用于在一个信道关闭之前，从信道接收数据。

接下来我们使用 for range 循环重写上面的代码。

```go
package main

import (
	"fmt"
)

func producter(chnl chan int){
	for i := 0;i < 10;i++{
		chnl <- i
	}
	close(chnl)
}

func main(){
	ch := make(chan int)
	go producter(ch)
	for data := range ch{
		fmt.Println(data)
	}
}
```

在第 16 行，for range 循环从信道 `ch` 接收数据，直到该信道关闭。一旦关闭了 `ch`，循环会自动结束。该程序会输出：

```
API server listening at: 127.0.0.1:39171
0
1
2
3
4
5
6
7
8
9
```

还记得那个cube和spure的例子吗？有一部分代码重复，现在我们把这段代码抽离出来

```go
package main

import (
	"fmt"
)

func dilter(chnl chan int,num int){
	for num != 0{
		dil := num % 10
		chnl <- dil
		num /= 10  
	}
	close(chnl)
}//以前的公用部分获取每位上的数字，现在被抽离成公用部分。

func squreFunc(squres chan int,num int){
	chal := make(chan int)
	sum := 0
	go dilter(chal,num)//启动一个协程往信道里面写入num的每位的数字
	for i := range chal{
		sum += i * i
	}
	squres <- sum
}

func cubeFunc(cube chan int,num int){
	chal := make(chan int)
	sum := 0
	go dilter(chal,num)//启动一个协程往信道里面写入num的每位的数字
	for i := range chal{
		sum += i * i * i
	}
	cube <- sum
}

func main(){
	num := 144
	squre := make(chan int)
	cube := make(chan int)
	go squreFunc(squre,num)
	go cubeFunc(cube,num)
	squres,cubes := <-squre,<- cube
	fmt.Println("Final output",squres+cubes)
}
```

上述程序里的 `dilter 函数，包含了获取一个数的每位数的逻辑，并且 `squareFunc` 和 `cubeFunc` 两个函数并发地调用了 `digits`。当计算完数字里面的每一位数时，第 13 行就会关闭信道。`squareFunc` 和 `cubeFunc` 两个协程使用 for range 循环分别监听了它们的信道，直到该信道关闭。程序的其他地方不变，该程序同样会输出：

`终端输出`:

```
API server listening at: 127.0.0.1:41615
Final output 162
```

##### 14.5 带缓冲区的channel

###### 14.5.1 声明

我们讨论的主要是无缓冲信道。我们在信道的教程里详细讨论了，无缓冲信道的发送和接收过程是阻塞的。

我们还可以创建一个有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。

通过向 `make` 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。

```go
ch := make(chan type, capacity)
```

要让一个信道有缓冲，上面语法中的 `capacity` 应该大于 0。无缓冲信道的容量默认为 0，因此我们在上面创建信道时，省略了容量参数。

我们开始编写代码，创建一个缓冲信道。

```go
package main

import (
	"fmt"
)

func main (){
	ch := make(chan int,3)
	ch <- 0
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

上面的代码创建了一个容量为3的`channel` `ch`,然后往`ch` 里面写入0 ，1，2，然后第一个fmt.Println读取数据0,并打印，第二个fmt.Println读取数据1，并打印，第三个fmt.Println读取数据3，并打印，带缓冲的信道是先进先出。

```
API server listening at: 127.0.0.1:25202
0
1
2
```

示例二，编写以协程写入数据。go 主程负责读取数据

```go
package main

import (
	"fmt"
	"time"
)

func worker(ch chan int){
    for i := 0;i < cap(ch);i++{
        ch <- i
        time.Sleep(time.Second)
    }
}

func main (){
	ch := make(chan int,3)
	go worker(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

在上面的程序中，第 17 行在 Go 主协程中创建了容量为 2 的缓冲信道 `ch`，而第 178行把 `ch` 传递给了 worker 协程。接下来 Go 主协程休眠了两秒。在这期间，`write` 协程在并发地运行。`write` 协程有一个 for 循环，依次向信道 `ch` 写入 0～4。而缓冲信道的容量为 2，因此 `write` 协程里立即会向 `ch` 写入 0 和 1，接下来发生阻塞，直到 `ch` 内的值被读取。因此，该程序立即打印出下面两行

```
API server listening at: 127.0.0.1:40901
0 输出0
1 等待1秒后输出 1
2 再等待1苗输出 2
```

从这么多例子我们看出，信道读取放在启动协程的协程内，这样就能保证信道的阻塞当前线程

我们再写一个例子来更好的理解缓冲的信道的阻塞。

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func worker(ch chan int){
	for i := 0; i < 5;i++{
		ch <- i
		fmt.Fprintf(os.Stdout,"the ch write value %d\n",i)
	}
	close(ch)
}

func main (){
	ch := make(chan int,2)
	go worker(ch)
	time.Sleep(2 * time.Second)
	for i := range ch {
		fmt.Printf("read value %d\n",i)
		time.Sleep(2 * time.Second)
	}
}
```

程序再18行创建了一个容量为2的带缓冲区信道，启动协程worker()往ch里面写数据，由于信道容量只有2，所有每次只能写入2条int数据。此时主协程等待两秒，

开始读取信道ch中的第一条数据，读取以后，信道的长度变为1，即可以再写入一条数据2，在等待两秒，重复上面的步骤再读取一条数据，在写入一条数据。直到写入操作停止。

```
API server listening at: 127.0.0.1:41879
the ch write value 0
the ch write value 1
read value 0
the ch write value 2
read value 1
the ch write value 3
read value 2
the ch write value 4
read value 3
read value 4
```

###### 14.5.2 死锁

当给我们向一个带缓冲区的信道写入数据，如果在其他协程没有读取的信道值的时候，写入超出了信道的容量，写入发生了阻塞。

我们举个例子来看：

```go
package main

import (
	"fmt"
)

func main(){
	ch := make(chan int,2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

终端输出上可以很清晰的看到报错在程序的11行，由于信道ch的容量已满，有没有程序读取他的值，在写入3的话就会触发死锁

```
API server listening at: 127.0.0.1:36649
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/golang/writefile/lesson1/main.go:11 +0x9b
Process exiting with code: 0
```

同样的如果读取超过了信道的容量也会发射阻塞，因为读取不到数据，也会触发死锁。

```go
package main

import (
	"fmt"
)

func main(){
	ch := make(chan int,2)
	ch <- 1
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

在如下的终端输出可以看到时报错的是11行，由于ch信道内只有一条数据，已经被第一个<-ch读取。所以信道没有数据，11行再去读取就会报错：

```
API server listening at: 127.0.0.1:41441
1
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/golang/writefile/lesson1/main.go:11 +0x12a
Process exiting with code: 0
```

###### 14.5.3 长度和容量

我们在使用 `make` 函数创建缓冲信道的时候会指定容量大小。

缓冲信道的长度是指信道中当前排队的元素个数。

可以通过len()来获取channel的长度

可以通过cap()来获取channel的容量

```go
package main

import (
	"fmt"
)

func main(){
	ch := make(chan int,10)
    fmt.Printf("the len of ch is %d\n",len(ch))
	fmt.Printf("the cap of ch is %d\n",cap(ch))
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("the len of ch is %d\n",len(ch))
	fmt.Printf("the cap of ch is %d\n",cap(ch))
	<-ch
    fmt.Printf("the len of ch is %d\n",len(ch))
	fmt.Printf("the cap of ch is %d\n",cap(ch))
}
```

程序的第8行创建了一个容量为10的channel ch，此时它的长度是0，容量是10，接着我们写入三条数据，此时长度变成3，容量还是10，接着读取一条数据，此时长度变成2，容量依旧是10。这里可以总结两点：1.信道容量是初始化时固定的，2.一旦数据从信道中被读取，数据就会被信道弹出。

```
API server listening at: 127.0.0.1:22502
the len of ch is 0
the cap of ch is 10
the len of ch is 3
the cap of ch is 10
the len of ch is 2
the cap of ch is 10
```

###### 14.5.4 WaitGroup

`WaitGroup` 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。Go 主协程需要等待这 3 个协程执行结束后，才会终止。这就可以用 `WaitGroup` 来实现。

```go
import (
	"fmt"
	"sync"
	"time"
)

func process(i int,wg *sync.WaitGroup){
	fmt.Printf("worker is start %d\n",i)
	time.Sleep(time.Second)
	fmt.Printf("worker is end %d\n",i)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	for i := 0;i<3;i++{
		wg.Add(1)
		go process(i,&wg)
	}
	wg.Wait()
}
```

waitgroup 是一个结构体类型，我们在第 15 行创建了 `WaitGroup` 类型的变量，其初始值为零值。`WaitGroup` 使用计数器来工作。当我们调用 `WaitGroup` 的 `Add` 并传递一个 `int` 时，`WaitGroup` 的计数器会加上 `Add` 的传参。要减少计数器，可以调用 `WaitGroup` 的 `Done()` 方法。`Wait()` 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。

上述程序里，for 循环迭代了 3 次，我们在循环内调用了 `wg.Add(1)`（第 17 行）。因此计数器变为 3。for 循环同样创建了 3 个 `process` 协程，然后在第 20 行调用了 `wg.Wait()`，确保 Go 主协程等待计数器变为 0。在第 13 行，`process` 协程内调用了 `wg.Done`，可以让计数器递减。一旦 3 个子协程都执行完毕（即 `wg.Done()` 调用了 3 次），那么计数器就变为 0，于是主协程会解除阻塞。

**在第 21 行里，传递 `wg` 的地址是很重要的。如果没有传递 `wg` 的地址，那么每个 Go 协程将会得到一个 `WaitGroup` 值的拷贝，因而当它们执行结束时，`main` 函数并不会知道**。

```
API server listening at: 127.0.0.1:29851
worker is start 2
worker is start 1
worker is start 0
worker is end 0
worker is end 2
worker is end 1
```

###### 14.5.5 工作池的实现

缓冲信道的重要应用之一就是实现`工作池`。

一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。

我们会使用缓冲信道来实现工作池。我们工作池的任务是计算所输入数字的每一位的和。例如，如果输入 234，结果会是 9（即 2 + 3 + 4）。向工作池输入的是一列伪随机数。

我们工作池的核心功能如下：

- 创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
- 将作业添加到该输入型缓冲信道中。
- 作业完成后，再将结果写入一个输出型缓冲信道。
- 从输出型缓冲信道读取并打印结果。

我们会逐步编写这个程序，让代码易于理解。我们写一个程序来计算一个数的每一位的和。

首先我们需要两个struct，Job,和result来分别表示作业和结果

```go
type Job struct {
	Id, Number int
}

type Result struct{
	job Job
	SumofNum int
}
```

接下来我们需要一个创建一个两个信道，用于Job的写入，读取，和result的写入，读取

```go
var jobChan chan Job = make(chan Job,10)
var resChan chan Result = make(chan Result,10)
```

接着我们写一个计算和的函数用于计算各位数的计算，并返回这个计算结果

```go
func digist(num int) (sum int){
	for num != 0 {
		dist := num % 10
		sum += dist
		num /= 10
	}
	return sum
} 
```

接下来我们需要创建一个工作协程。

```go
func worker(wg *sync.WaitGroup){
	for job := range jobChan{
		result := Result{Myjob: job,ResNum:disgit(job.Number)}
		resChan  <- result
	}
	wg.Done()
}
```

上面的函数创建了一个工作者（Worker），读取 `jobs` 信道的数据，根据当前的 `job` 和 `digits` 函数的返回值，创建了一个 `Result` 结构体变量，然后将结果写入 `results` 缓冲信道。`worker` 函数接收了一个 `WaitGroup` 类型的 `wg` 作为参数，当所有的 `jobs` 完成的时候，调用了 `Done()` 方法。

有了工作协程，我们就需要创建一个workerpool

```go
func createWorkerPool(count int){
	var wg sync.WaitGroup
	for i := 0; i < count;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resChan)
}
```

上面函数的参数是需要创建的工作协程的数量。在创建 Go 协程之前，它调用了 `wg.Add(1)` 方法，于是 `WaitGroup` 计数器递增。接下来，我们创建工作协程，并向 `worker` 函数传递 `wg` 的地址。创建了需要的工作协程后，函数调用 `wg.Wait()`，等待所有的 Go 协程执行完毕。所有协程完成执行之后，函数会关闭 `results` 信道。因为所有协程都已经执行完毕，于是不再需要向 `results` 信道写入数据了。

再接下来我们需要一个携程函数来为我们创建Job,来分配到工作者

```
func creatJob(count int){
    for i := 0;i < count;i++{
        job := Job {Id: i,Number: rand.Intn(999)}
        jobChan <- job
    }
    close(jobChan)
}
```

上面的 `createJob` 函数接收所需创建的作业数量作为输入参数，生成了最大值为 998 的伪随机数，并使用该随机数创建了 `Job` 结构体变量。这个函数把 for 循环的计数器 `i` 作为 id，最后把创建的结构体变量写入 `jobsChan` 信道。当写入所有的 `job` 时，它关闭了 `jobChan` 信道。

然后我们需要一个函数来打印result的内容，并用一个channel来阻塞，防止主进程提前退出

```go
func printResult(done chan bool){
	for res := range resChan{
		fmt.Printf("jobid%d: job num %d,the sum is %d",res.Myjob.Id,res.Myjob.Number,res.ResNum)
	}
	done <- true
}
```

最后完成完整的代码如下

```go 
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id, Number int
}

type Result struct {
	Myjob  Job
	ResNum int
}

var jobChan chan Job = make(chan Job, 10)
var resChan chan Result = make(chan Result, 10)

func disgit(num int) (sum int) {
	for num != 0 {
		dig := num % 10
		sum += dig
		num /= 10
	}
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobChan {
		result := Result{Myjob: job, ResNum: disgit(job.Number)}
		resChan <- result
	}
	wg.Done()
}

func createWorkerPool(count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resChan)
}

func creatJob(count int) {
	for i := 0; i < count; i++ {
		job := Job{Id: i, Number: rand.Intn(999)}
		jobChan <- job
	}
	close(jobChan)
}

func printResult(done chan bool) {
	for res := range resChan {
		fmt.Printf("jobid%d: job num %d,the sum is %d\n", res.Myjob.Id, res.Myjob.Number, res.ResNum)
	}
	done <- true
}

func main() {
	rand.Seed(time.Now().Unix())
	count := 100
	var done chan bool = make(chan bool)
	go creatJob(count)
	go createWorkerPool(count)
	go printResult(done)
	<-done
}

```

终端输出：

```
API server listening at: 127.0.0.1:48785
jobid0: job num 619,the sum is 16
jobid1: job num 812,the sum is 11
jobid2: job num 56,the sum is 11
jobid3: job num 118,the sum is 10
jobid4: job num 943,the sum is 16
jobid5: job num 798,the sum is 24
jobid6: job num 553,the sum is 13
jobid7: job num 688,the sum is 22
jobid8: job num 435,the sum is 12
jobid9: job num 331,the sum is 7
jobid10: job num 470,the sum is 11
jobid11: job num 42,the sum is 6
jobid12: job num 340,the sum is 7
jobid13: job num 917,the sum is 17
jobid14: job num 935,the sum is 17
jobid15: job num 808,the sum is 16
jobid16: job num 997,the sum is 25
jobid17: job num 710,the sum is 8
jobid18: job num 921,the sum is 12
jobid19: job num 390,the sum is 12
jobid20: job num 461,the sum is 11
jobid21: job num 230,the sum is 5
jobid22: job num 595,the sum is 19
jobid23: job num 965,the sum is 20
jobid24: job num 95,the sum is 14
jobid25: job num 728,the sum is 17
jobid26: job num 689,the sum is 23
jobid27: job num 481,the sum is 13
jobid28: job num 693,the sum is 18
jobid29: job num 976,the sum is 22
jobid30: job num 115,the sum is 7
jobid31: job num 882,the sum is 18
jobid32: job num 938,the sum is 20
jobid33: job num 216,the sum is 9
jobid34: job num 692,the sum is 17
jobid35: job num 14,the sum is 5
jobid36: job num 522,the sum is 9
jobid37: job num 49,the sum is 13
jobid38: job num 983,the sum is 20
jobid39: job num 216,the sum is 9
jobid40: job num 938,the sum is 20
jobid41: job num 199,the sum is 19
jobid42: job num 959,the sum is 23
jobid43: job num 919,the sum is 19
jobid44: job num 901,the sum is 10
jobid45: job num 880,the sum is 16
jobid46: job num 849,the sum is 21
jobid47: job num 382,the sum is 13
jobid48: job num 229,the sum is 13
jobid49: job num 371,the sum is 11
jobid50: job num 902,the sum is 11
jobid51: job num 520,the sum is 7
jobid52: job num 221,the sum is 5
jobid53: job num 7,the sum is 7
jobid54: job num 352,the sum is 10
jobid55: job num 122,the sum is 5
jobid56: job num 932,the sum is 14
jobid57: job num 968,the sum is 23
jobid58: job num 69,the sum is 15
jobid59: job num 377,the sum is 17
jobid60: job num 281,the sum is 11
jobid61: job num 959,the sum is 23
jobid62: job num 694,the sum is 19
jobid63: job num 216,the sum is 9
jobid64: job num 63,the sum is 9
jobid65: job num 56,the sum is 11
jobid66: job num 596,the sum is 20
jobid67: job num 5,the sum is 5
jobid68: job num 225,the sum is 9
jobid69: job num 781,the sum is 16
jobid70: job num 190,the sum is 10
jobid71: job num 856,the sum is 19
jobid72: job num 290,the sum is 11
jobid73: job num 213,the sum is 6
jobid74: job num 431,the sum is 8
jobid75: job num 78,the sum is 15
jobid76: job num 906,the sum is 15
jobid77: job num 454,the sum is 13
jobid78: job num 63,the sum is 9
jobid79: job num 374,the sum is 14
jobid80: job num 206,the sum is 8
jobid81: job num 935,the sum is 17
jobid82: job num 652,the sum is 13
jobid83: job num 608,the sum is 14
jobid84: job num 580,the sum is 13
jobid85: job num 576,the sum is 18
jobid86: job num 767,the sum is 20
jobid87: job num 681,the sum is 15
jobid88: job num 382,the sum is 13
jobid89: job num 400,the sum is 4
jobid90: job num 576,the sum is 18
jobid91: job num 57,the sum is 12
jobid92: job num 627,the sum is 15
jobid93: job num 24,the sum is 6
jobid94: job num 72,the sum is 9
jobid95: job num 365,the sum is 14
jobid96: job num 512,the sum is 8
jobid97: job num 874,the sum is 19
jobid98: job num 756,the sum is 18
jobid99: job num 933,the sum is 15
Process exiting with code: 0
```

#### 15 select 和线程安全

##### 15.1 select

多channel场景
select语义介绍和使用
A. 多个channel同时需要读取或写入，怎么办？
B. 串行操作？ NONONO

举一个例子

```go
package main

import (
	"fmt"
	"time"
)

func server1(chnl chan string) {
	time.Sleep(time.Second * 6)
	fmt.Println("server1 begin")
	chnl <- "from server1"
}

func server2(chnl chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println("server2 begin")
	chnl <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	s1 := <-output1
	fmt.Println(s1)
	s2 := <-output2
	fmt.Println(s2)
}

```

由于我们的代码是从上往下执行，启动server1和server2两个协程时，我们可以看到server2休眠三秒后，将数据写入到了信道，server1则休眠了6秒才将数据写入信道，但是主协程是从上至下执行，所以在`s1 := <- output1`进行了阻塞，等待数据写入信道，然后读取到数据，才放弃阻塞，因此导致了s2 读取时间也变成6秒。所以终端输出是这样

`终端输出`：

```
API server listening at: 127.0.0.1:29757
server2 begin
server1 begin
from server1
from server2
```

假使在高并发的场景中，这样会导致后面准备好返回数据的信道被阻塞，假如server1和server2都同时对某个用户提供数据返回，那就可能造成性能降低，因此select登场

**select**

A. 同时监听一个或多个channel，直到其中一个channel ready 
B. 如果其中多个channel同时ready，随机选择一个进行操作。
C. 语法和switch case有点类似，代码可读性更好。

来写一个实例

```go
package main

import (
	"fmt"
	"time"
)

func server1(chnl chan string) {
	time.Sleep(time.Second * 6)
	chnl <- "from server1"
}

func server2(chnl chan string) {
	time.Sleep(time.Second * 3)
	chnl <- "from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	select { //监听ch1,和ch2
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	}
}

```

从上面我们可以看到select 监听了ch1,和ch2,由于ch1延迟了6秒才会写入数据，ch2,延迟三秒写入数据，因此select 在监听到了ch2写入数据，即<-ch2 ready,因此执行了data := <-ch2的分支

```
API server listening at: 127.0.0.1:7158
from server2
```

在多个分支都处于ready状态的时候，随机选择一个分支执行，我们再来改改上面的代码

```go
package main

import (
	"fmt"
	"time"
)

func server1(chnl chan string) {
	chnl <- "from server1"
}

func server2(chnl chan string) {
	chnl <- "from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	time.Sleep(time.Millisecond * 100)
	select {
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	}
}

```

我们运行五次，程序的输出结果分别是

```
第一次是from server1
第二次是from server1
第三次是from server1
第四次和第五次是from server2
```

select 还可以用于检查信道是否被写满的



##### 15.2 线程安全

现实例子

A. 多个goroutine同时操作一个资源，这个资源又叫临界区

B. 现实生活中的十字路口，通过红路灯实现线程安全

C. 火车上的厕所，通过互斥锁来实现线程安全线程安全介绍

实际例子， x = x +1

A. 先从内存中取出x的值

B. CPU进行计算，x+1

C. 然后把x+1的结果存储在内存中

![](D:/assets/image-20200824104701687.png)





##### 15.2.1 互斥锁

```go


package main

import (
	"fmt"
	"sync"
)

var x int

func add(wg *sync.WaitGroup,m *sync.Mutex){
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0;i < 10000;i++{
		wg.Add(1)
		go add(&wg,&m)
	}
	wg.Wait()
	fmt.Println(x)
}
```

终端输出：

```
API server listening at: 127.0.0.1:40741
10000
```

##### 15.2.2 读写锁

A. 读多写少的场景

B. 分为两种角色，读锁和写锁

C. 当一个goroutine获取写锁之后，其他的goroutine获取写锁或读锁都会等待

D. 当一个goroutine获取读锁之后，其他的goroutine获取写锁都会等待, 但其他

goroutine获取读锁时，都会继续获得锁



```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var rwmutex sync.RWMutex

func read(wg *sync.WaitGroup,i int){
	rwmutex.RLock()
	fmt.Println(i)		
	rwmutex.RUnlock()
	time.Sleep(time.Second)
	wg.Done()
}

func write(wg *sync.WaitGroup){
	rwmutex.Lock()
	x +=1
	rwmutex.Unlock()
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go write(&wg)
	wg.Wait()
	for i :=0;i<10;i++{
		wg.Add(1)
		go read(&wg,i)
	}
	wg.Wait()
}
```

##### 15.2.3 原子操作

原子操作

A. 加锁代价比较耗时，需要上下文切换

B. 针对基本数据类型，可以使用原子操作保证线程安全

C. 原子操作在用户态就可以完成，因此性能比互斥锁要高

![image-20200824142849849](D:/assets/image-20200824142849849.png)

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int32
var mutex sync.Mutex

func add(wg *sync.WaitGroup){
	for i := 0; i <10000;i++{
		// mutex.Lock()
		// x += 1
		// mutex.Unlock()
		atomic.AddInt32(&x,1)
	}
	wg.Done()
}

func main(){
	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	for i :=0;i < 10;i++{
		wg.Add(1)
		go add(&wg)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println(x)
	fmt.Println((end - start)/1000)
}
```

终端输出：

```
API server listening at: 127.0.0.1:9690
100000
2308
```

