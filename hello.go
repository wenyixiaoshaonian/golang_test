package main

import (
	"fmt"
	"sort"
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
	names := Mystring{
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
    }
    // 使用sort包进行排序
    sort.Sort(names)
    // 遍历打印结果
    for _, v := range names {
            fmt.Printf("%s\n", v)
    }
}
