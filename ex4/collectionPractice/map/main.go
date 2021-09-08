package main

import "fmt"

func main() {
	//TestMap1()
	TestMap2()
}

func TestMap1() {

	var testMap map[string]string //未初始化不可賦值
	fmt.Printf("testMap 未初始化零值為nil : %#v \n", testMap)

	users := map[int]string{
		1: "Ron",
		2: "Lynn",
	}
	fmt.Println("users[1] : ", users[1])

}

func TestMap2() {

	users := map[int]string{
		1: "Ron",
		2: "Lynn",
	}

	users[3] = "Ronn"

	for key, value := range users {
		fmt.Printf("users[%d] = %s \n", key, value)
	}

	delete(users, 3) // 刪除users[3]

	user, exist := users[3] //可透過回傳檢查key是否存在
	if exist {
		fmt.Println("users[3] : ", user)
	}

}
