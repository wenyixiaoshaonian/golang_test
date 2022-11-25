package main

import (
	"fmt"
	"runtime"
	"sync"
	"container/list"
)


func mslice() {

	var k,y = 100,111.1
	x :=10
	var a []int = []int{11,22,33}
	var b [4][2]int = [4][2]int{1:{0:11},3:{1:22}}
	fmt.Println(a[0]);
	fmt.Println(a[len(a)-1])
	fmt.Println(len(a))
	fmt.Printf("Hello world! %d %f %d\n",x,y,k)
	for i,v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _,v := range a {
		fmt.Printf("%d\n", v)
	}

	for i,v := range b {
		fmt.Printf("%d %d\n", i, v)
	}

	fmt.Println(a[1:])
	var c []int = a
	fmt.Println(c[:])

	var ac []int

	fmt.Println(ac,len(ac))
	ac = append(ac,[]int{1,2,3}...)
	fmt.Println(ac,len(ac))
	
	ac = append(ac[:2],append([]int{4,5,6},ac[2:]...)...)
	fmt.Println(ac,len(ac))

	copy(a,ac)
	fmt.Println(a,len(a))

	ac = append(ac[:2],ac[4:]...)
	fmt.Println(ac,len(ac))

	ac = ac[:len(ac)-2]
	fmt.Println(ac,len(ac))

	for index :=1; index < len(ac);index++ {
		fmt.Printf("Index: %d Value: %d\n", index, ac[index])
	}
}

func mmap() {
	var scene []string
	var mb map[int] int
	ma := make(map[int] int,100)
	mc := map[string]int{"one" : 1 , "two" : 2,"three" : 3}
	// delete(mc,"one")
	ma[1] = 11
	ma[2] = 22
	mc = make(map[string] int,3)		//清空
	fmt.Println(mb,len(mb))
	fmt.Println(ma,len(ma))
	fmt.Println(mc,len(mc))

	for i,v := range mc {
		fmt.Printf("%s %d\n", i, v)
		scene = append(scene,i)
	}
	fmt.Println(scene,len(scene))
}

type queKey struct {
	name string
	six  int
}

type Profile struct {
	name string
	age  int
	adr  string
	six  int
}

var index = make(map[queKey]*Profile)

func buildIndex(list []*Profile) {
	for _, data := range list {
		key := queKey {
			name : data.name,
			six : data.six,
		}
		index[key] = data
	}
}

func find_map(names string,sixs int) {
	key := queKey{
		name : names,
		six : sixs,
	}
	result,ok:= index[key]
	if ok {
		fmt.Println(result)
		fmt.Println(ok)
	} else {
		fmt.Println("info can not find...")
	}
}

func find_easy(list []*Profile,name string,six int) {
	for _,data := range list {
		if data.name == name && data.six == six {
			fmt.Println(data)
			return
		}
	}
	fmt.Println("info can not find...")
}
func find() {
	list := []*Profile {
		{name : "xiaoming",age : 18,adr : "china",six : 1},
		{name : "xiaohua" ,age : 19,adr : "china",six : 0},
		{name : "xiaozhi" ,age : 21,adr : "USA"  ,six : 1},
		{name : "xiaojun" ,age : 23,adr : "china",six : 1},
		{name : "xiaocao" ,age : 26,adr : "china",six : 1},
		{name : "xiaorou" ,age : 25,adr : "china",six : 0},
	}
	// fmt.Println(list)

	// find_easy(list,"xiaocao",1)

	buildIndex(list)
	find_map("xiaocao",0)
}

// 开启一段并发代码
func wrfunc() {
	runtime.GOMAXPROCS(2)
	m := make(map[int]int)
	// 开启一段并发代码
	go func() {
		// 不停地对map进行写入
		for {
			m[1] = 1
		}
	}()
	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			_ = m[1]
		}
	}()
	// 无限循环, 让并发程序在后台执行
	for {
	}
}

// 开启一段并发代码
func wrfunc_sync() {
	runtime.GOMAXPROCS(2)
	var scene sync.Map
	// 开启一段并发代码
	go func() {
		// 不停地对map进行写入
		for {
			scene.Store("greece", 97)
		}
	}()
	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			scene.Load("greece")
		}
	}()
	// 无限循环, 让并发程序在后台执行
	for {
	}
}

func lf() {
	l := list.New()

	element := l.PushBack("first")
	l.PushFront("111")
	l.PushBack("second")
	l.InsertAfter("third",element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func fnil() {
	var ac []int
	fmt.Printf("%#v \n",ac)
}

func ffor() {
	var i int
	JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
		fmt.Println(j)
	}
	for i<10 {
		i++
		fmt.Println(i)
	}
}

func fchal() {
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func fstr() {
	var str = "hello 你好"
	for key, value := range str {
    	fmt.Printf("key:%d value:%c\n", key, value)
}
}

func fswitch() {
	a := "hello"

	switch a {
	case "hello":
			fmt.Println(1)
	case "world":
			fmt.Println(1)
	default:
			fmt.Println(0)	
	}
}

func retu(a,b int) (c,d int) {
	c = a + b
	d = a - b

	return a + b , a - b
	// return 
}

type inter interface {
	mythod(interface{}) string
	mythod2(interface{}) int
}

type mystru struct {
	name string
	age int
}

func (my mystru) mythod(p interface{}) string {
	var ok bool
	my.name,ok = p.(string)
	if ok {
		fmt.Println(">>>===222")
	}else {
		fmt.Println(">>>===111")
	}
	return my.name
}

func (my mystru) mythod2(p interface{}) int {
	my.age = p.(int)
	return my.age
}

func Adder(a int) func(int) int{
    var x int
    return func(d int) int{
        x+=d
		a++
        return a
    }
}

func pr (args ... interface{}) {
	prin(args...)
}
func prin (args ... interface{}) {
	for _,arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
	}
	}
}

func deferf () {
	fmt.Println("defer begin")
    // 将defer放入延迟调用栈
    defer fmt.Println(1)
    defer fmt.Println(2)
    // 最后一个放入, 位于栈顶, 最先调用
    defer fmt.Println(3)
    fmt.Println("defer end")
}

func ad (b int) (res int) {
	if b < 3 {
		res = 1
	} else {
		res += ad(b-1) + ad(b-2) 
	}
	return
}
func adda (a int) {
	for i := 1;i <= a; i++ {
		val := ad(i)
		fmt.Println(val)
	}
}

func fpanic() {
	defer err := recover()
	switch err.(type) {
	case runtime.Error: // 运行时错误
		fmt.Println("runtime error:", err)
	default: // 非运行时错误
		fmt.Println("error:", err)
	}	
	
	panic("crash")

	
}