package main

import "fmt"

type user struct {
	name   string
	age    int
	gender string
}

type point struct {
	x int
	y int
}

type name string

type size struct {
	width  int
	length int
}

type dot struct {
	name
	point
	size
}

func main() {
	//TestStruct1()
	//TestStruct2()
	TestStruct3()
}

func TestStruct1() {
	user1 := user{
		"Ron",
		27,
		"male",
	}

	fmt.Printf("user1 : %#v \n", user1)

	user1.age = 28
	fmt.Printf("設定user1的age內容 ： %#v \n", user1)
}

func TestStruct2() {
	//匿名結構
	point1 := struct {
		x int
		y int
	}{
		10,
		20,
	}

	point2 := point{
		10, 20,
	}

	fmt.Printf("point1 : %#v \n", point1)
	fmt.Printf("point2 : %#v \n", point2)

	fmt.Println("potin1 == point2  :", point1 == point2)
}

func TestStruct3() {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "dot2"
	dot2.x = 10
	dot2.y = 10
	dot2.width = 20
	dot2.length = 20

	fmt.Printf("dot1 : %#v \n", dot1)
	fmt.Printf("dot2 : %#v \n", dot2)

}
