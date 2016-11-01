package main

import "fmt"

type person struct {
	name string
	friends []string
}

type friendship struct {
	person1 string
	person2 string
}


func main() {
	akash := person{name:"akash"}
	andy := person{name:"andy"}

	andy.makeFriendsWith(akash)
	andy.listFriends()

}

func (p *person) makeFriendsWith(friendee person)  {
	p.friends = append(p.friends, friendee.name)
}

func (p person) listFriends() {
	fmt.Println(p.friends)
}

