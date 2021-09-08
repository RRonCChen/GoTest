package main

import "fmt"

type point struct {
	x int
	y int
}

//pointer receiver 透過指標接受器將方法bind到struct上,更動會反應到原變數上
func (p *point) setPoint(x, y int) {
	p.x = x
	p.y = y
}

//value recevier
func (p point) getPoint() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}

func main() {
	var p1 point
	fmt.Printf("p1 : %#v \n", p1)
	p1.setPoint(10, 20)
	fmt.Printf("p1 : %#v \n", p1)

	//不管是pointer recevier或是value recevier, Go會直接轉換可直接呼叫bind的方法
	p2 := point{}
	p3 := &point{}
	p2.setPoint(10, 20)
	p3.setPoint(20, 30)
	fmt.Println("p2 :", p2.getPoint())
	fmt.Println("p3 :", p3.getPoint())
}
