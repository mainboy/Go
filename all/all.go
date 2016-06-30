
package main

import "fmt"

// 定义变量
var a int = 10
var b float32 = 3.14
var c string = "hello world"

// 定义数组
var array = [5]int{1,2,3,4,5}

// 数组的切片操作
var assign [5]int

// 定义二维数组
const(
    x = 4
    y = 4
)

var d2 [x][y]int

// 定义结构
type student struct{
    name string
    age int
    sex string
}

// 定义矩形结构体
type rect struct{
    width int
    height int 
}

func (r *rect)area() int {
    return r.width * r.height
}

func (r *rect)perimeter() int {
    return 2 * (r.width + r.height)
}

// 定义错误信息结构体
type errmsg struct{
    arg int 
    msg string
}

func (e *errmsg) ERROR() {
    if e.arg < 0 || e.arg > 255{
	fmt.Println("越界了！！！")
    } else{
	fmt.Println(e.arg, "success", e.msg)
    }
	
}

func main(){
    // 变量测试
    fmt.Printf("%d %f %s\n",a,b,c)
    fmt.Println(a,b,c)
    
    // 数组测试
    fmt.Println(array)
    
    // 数组切片操作
    assign = [5]int{0,1,2,3,4}
    assign := array[0:5]
    fmt.Println(assign)

    // 二维数组测试
    var i,j int
    for i = 0; i < x; i++ {
	for j = 0; j < y; j++{
	    d2[i][j] = i+j
	}
    }
    fmt.Println(d2)

    // 测试if分支
    var num int = 25
    if num%2==1{
	fmt.Println("奇数")
    }else{
	fmt.Println("偶数")
    }

    // 测试while循环，Go没有while
    var top int = 5
    for i = 0; i < top ; i++ {
	fmt.Println("cur =",i)
    }

    // 输入输出测试
    var in int
    fmt.Scanf("%d",&in)
    fmt.Println(in)

    // switch测试
    var swit int = 4
    switch swit {
    case 1:
	fmt.Println("hello1")
    case 2:
	fmt.Println("hello2")
    case 3,4:
	fmt.Println("hello world!")
    default:
	fmt.Printf("default!")
    }

    // map测试
    m := make(map[string]int)
    m["one"] = 1
    m["two"] = 2
    m["three"] = 3
    fmt.Println(m)

    fmt.Println(m["three"])

    delete(m,"two")
    fmt.Println(m)

    for key,val := range m{
	fmt.Println(key,"-->",val)
    }

    // go指针
    var it int = 2
    var ptr *int = &it
    fmt.Println("整数",it,"指针地址",ptr,"指针指向的数值",*ptr)
    
    *ptr = 3
    fmt.Println("整数",it,"指针地址",ptr,"指针指向的数值",*ptr)
    
    it = 4   
    fmt.Println("整数",it,"指针地址",ptr,"指针指向的数值",*ptr)

    // 内存分配
    /*
     * new(T)返回*T, 没有初始化；
     * make(T, args)返回T，初始化；仅用于创建切片，map和管道
     **/
    // new用法
    var p *[]int = new([]int)
    fmt.Println(p)
    *p = make([]int,10,10)
    fmt.Println(p)
    fmt.Println((*p)[3])
    (*p)[3] = 4
    fmt.Println(p)
    // make用法
    v := make([]int,10,10)
    fmt.Println(v)

    // 函数
    fmt.Println(min(1,2))

    // 函数返回多个值
    sv, se := dist("two")
    fmt.Println(sv, se)

    /* 注意sv,se之前已经定义过，无需再定义，直接使用= */
    sv, se = dist("four")
    fmt.Println(sv, se)

    // 函数递归
    fmt.Println(fib(5))

    // 函数返回不定参数
    sum(1,2,3)
    nums := []int{1,2,3,4,5}
    sum(nums...)

    // 函数闭包
    nextNumFunc := nextNum()
    for i := 0; i < 10; i++{
	fmt.Println(nextNumFunc())
    }

    // go的结构
    stuA := student{"严康", 22, "男"}
    fmt.Println(stuA)
    
    stuB := &stuA
    fmt.Println(stuB)

    stuB.age = 23
    fmt.Println(stuB)
    fmt.Println(stuA)
 
    // go的结构体方法
    r := rect{width:4, height:5}
    fmt.Println("面积: ", r.area())
    fmt.Println("周长: ", r.perimeter())
    
    // 错误处理
    for _,i := range []int{-2, 1, 256}{
	test := errmsg{i, "hello world"}
	test.ERROR()
    }

}

func min(a int, b int) int {
    if a < b {
	return a
    }
    return b
}

func dist(key string) (int, bool) {
    m := map[string]int{"one": 1, "two": 2, "three": 3}

    var errno bool
    var val int

    val, errno = m[key]

    return val, errno
}

func fib(n int) int {
    if n == 1{
	return 1
    }
    return n*fib(n-1)
}

func sum(nums ...int) {
    fmt.Println(nums, " ")
    total := 0
    /* 用_过滤掉下标*/
    for _, num := range nums {
	total += num
    }

    fmt.Println(total)
}

func nextNum() func() int {
    var i int = 1
    var j int = 1
    return func() int{
	var tmp int = i+j
	i, j = j, tmp
	return tmp
    }
}
