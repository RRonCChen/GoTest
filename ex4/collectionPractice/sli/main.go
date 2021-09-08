package main

import "fmt"

func TestSli1() {

	var sli []int

	fmt.Println("sli : ", sli)
	for i := 0; i < 5; i++ {
		sli = append(sli, i) // 使用append來加入新元素
	}
	fmt.Println("sli : ", sli)
}

func TestSli2() {
	sli := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("sli 0 to 2 : ", sli[:3])
}

func TestSli3() {
	sli := make([]int, 10)

	fmt.Println(" len(sli) : ", len(sli))
	fmt.Println(" cap(sli) : ", cap(sli))
	fmt.Println(" sli : ", sli)

	for i := 0; i < 11; i++ {
		sli = append(sli, i)
	}

	fmt.Println(" len(sli) : ", len(sli))
	fmt.Println(" cap(sli) : ", cap(sli))
	fmt.Println(" sli : ", sli)

}

func TestSli4() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:]

	fmt.Println("s1[3] s2[3] : ", s1[3], s2[3])

	s2 = append(s2, 1) //超過原本隱藏陣列的容量，發生陣列置換
	s1[3] = 100
	fmt.Println("s1[3] s2[3] : ", s1[3], s2[3])

	s3 := []int{1, 2, 3, 4, 5}
	s4 := make([]int, len(s3), cap(s3))
	copied := copy(s4, s3) // 複製s3切片元素給s4,並回傳複製數量
	s3[3] = 100            //s1與s2沒有linked
	fmt.Print("s3[3] s4[3] : ", s3[3], s4[3])
	fmt.Printf(" 複製了%v個元素\n", copied)

	s5 := []int{1, 2, 3, 4, 5}
	s6 := append([]int{}, s5...) //把s5得元素加入到s6裡面
	s5[3] = 100                  //s5與s6沒有linked
	fmt.Println("s5[3] s6[3] : ", s5[3], s6[3])

}

func main() {
	TestSli4()
}
