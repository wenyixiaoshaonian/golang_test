package main

import (
	"fmt"
	"sort"
)

type std struct {
	name string
	age int
	six int
}
type stds struct {
	city string
	school string
	int
	student std
	student2 std
}
func structf() {
	var std1 std
	std1.name = "xiaohua"
	std1.age = 18
	std1.six = 1
	fmt.Println(std1)
	fmt.Printf("std1 adr : %p\n",&std1)
	std2 := new(std)
	std2.name = "xiaoming"
	std2.age = 23
	std2.six = 0
	fmt.Println(*std2)
	fmt.Printf("std2 adr : %p\n",std2)
	// std3 := std{
	// 	name : "xiaozhi",
	// 	age  : 24,
	// 	six  : 1,
	// }
	// fmt.Println(std3)
	// std4 := std{
	// 	"xiaoquan",
	// 	27,
	// 	0,
	// }
	// fmt.Println(std4)
	std5 := &std{
		"xiaoquan2",
		37,
		0,
	}
	fmt.Println(std5)
	fmt.Printf("std5 adr : %p\n",std5)
	msg := struct{
		id int
		age int
	} {
		1,23,
	}
	fmt.Println(msg)
}

func tests () {
	var studs stds
	studs.city = "hangzhou"
	studs.school = "beishida"
	studs.int = 23
	studs.student.name = "zhiming"
	studs.student.age = 18
	studs.student.six = 0
	studs.student2.name = "zhiming2"
	studs.student2.age = 28
	studs.student2.six = 1
	fmt.Println(studs)

	std4 := stds{
		city:"zhuzhou",
		school:"beishida",
		int:0,
		student:std{
			"chunjiao",
			16,
			1,
		},
		student2:std{
			"chunjiao2",
			26,
			0,
		},
	}
	fmt.Println(std4)
}

type My_inferface interface {
	Write_Data (data interface{} ) error
	Read_Data (data interface{} ) error
}

type Computer struct {
	name string
}

func (c Computer) Write_Data(data interface{}) error {
	fmt.Println(data,c.name)
	return nil
}

type Mystring []string

func (m Mystring) Len() int {
	return len(m)
}

func (m Mystring) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m Mystring) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func sort_int() {
	names := sort.IntSlice{
		34,
		12,
		234,
		555,
		1,
	}
	sort.Sort(names)
	fmt.Println(names)
}

func sort_int_q() {
	names := []int{
		34,
		12,
		234,
		555,
		12,
	}
	sort.Ints(names)
	fmt.Println(names)
}
