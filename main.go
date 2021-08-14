package main

import "fmt"

//コロンやカンマは使わない！！
type person struct {
	firstName string
	lastName string
	// contact -> contactInfoの場合contactInfoのみでOK
	contact contactInfo
}

type contactInfo struct {
	email string
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
		lastName: "Hamada",
		contact: contactInfo{
			email: "hamada@downtown.com",
			zipCode: 2222222,
		},
	}
	fmt.Println(ryoka)
	fmt.Println(chiharu)

	masatoshi.updateName("Hamachon")
	//これでは名前が更新されていない！ => ポインタの概念
	masatoshi.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

