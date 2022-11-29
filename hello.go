package main

// import te "test/test_lib"
import "fmt"

type Any interface {}
func init(){
	fmt.Println("Call init")
}

const(
	identifier1 int = iota * 2
    identifier2
    identifier3
)
func main() {
	// var my = mystru{"xiaoming",18}
	// fmt.Println(my.mythod("xiaohua"))
	// fmt.Println(my.mythod2(20))
	// var a = Adder(1)
	// fmt.Println(a(0))
	// fmt.Println(a(4))
	// pr(1,2,3,"cas",5)
	// tests()
	// c := Computer{"apple"}
	// c.Write_Data("xiaoming") 
	// names := Mystring{
    //     "3. Triple Kill",
    //     "5. Penta Kill",
    //     "2. Double Kill",
    //     "4. Quadra Kill",
    //     "1. First Blood",
    // }
    // // 使用sort包进行排序
    // sort.Sort(names)
    // // 遍历打印结果
    // for _, v := range names {
    //         fmt.Printf("%s\n", v)
    // }
	// sort_int_q()

	// var a Any
	// a = 1
	// fmt.Println(a)

	// web_server()
	// te.Hello_test()
	// var name string
	// var age byte
	// var score float32
	// fmt.Println("请输入姓名 年龄 成绩")
	// fmt.Scanf("%s %d %f", &name, &age, &score)
	// fmt.Println("Name = ", name, "Age =", age, "Score =", score)
	// fmt.Printf("Name = %v Age = %v Score = %v", name, age, score)

	// fmt.Println(identifier1)
	// fmt.Println(identifier2)
	// fmt.Println(identifier3)
	
	// std := te.Getstd("xiaohua")
	// std.Getname()

	var person Personer
	stu := &Student {
		"xiaoming1",
		34,
		9.119,
	}
	person = stu
	person.Eat()
	person.Sleep(8)
}
