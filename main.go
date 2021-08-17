package main

import "fmt"

//コロンやカンマは使わない！！
type person struct {
	firstName string
	lastName  string
	// contact -> contactInfoの場合contactInfoのみでOK
	contact contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// プロパティ名を書かなくても前から順に当てはめることは出来るが、バグの原因になるので非推奨
	ryoka := person{firstName: "Ryoka", lastName: "Sato"}

	var chiharu person
	chiharu.firstName = "Chiharu"
	chiharu.lastName = "Yamamoto"

	masatoshi := person{
		firstName: "Masatoshi",
		lastName:  "Hamada",
		contact: contactInfo{
			email:   "hamada@downtown.com",
			zipCode: 2222222,
		},
	}
	fmt.Println(ryoka)
	fmt.Println(chiharu)

	// &variable -> 変数が格納されているメモリの場所を示す
	// *pointer -> あるメモリ内に入っている変数の値を示す

	//masatoshiはperson型で、updateNameのレシーバーの*person型とは異なるが、
	//Goの言語仕様として、受け取ったperson型の変数を*person型に自動変換してくれるらしい。
	masatoshi.updateName("Hamachon")
	//これでは名前が更新されていない！ => ポインタの概念
	masatoshi.print()
}

// *person -> 型であり、さっき出てきたメモリ内の変数のことではない。
// pointerToPersonというのは*person型の値
func (pointerToPerson *person) updateName(newFirstName string) {
	//これはメモリ内の変数の中身
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// まとめ
// address -> value ==>  *address
// value -> address ==> &value
