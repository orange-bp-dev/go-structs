package main

import "fmt"

//コロンやカンマは使わない！！
type person struct {
	firstName string
	lastName string
}

func main() {
	// プロパティ名を書かなくても前から順に当てはめることは出来るが、バグの原因になるので非推奨
	ryoka := person{firstName: "Ryoka", lastName: "Sato"}
	fmt.Println(ryoka)
}

