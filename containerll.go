package main

import "fmt"

type Member struct {
	Id    int
	Nama  string
	Point float32
}

type MemberLL struct {
	Container Member
	Next      *MemberLL
}

func main() {
	head := MemberLL{}

	member1 := Member{
		Id:    1,
		Nama:  "Kurniawan",
		Point: 0,
	}

	member2 := Member{
		Id:    2,
		Nama:  "Aan",
		Point: 100,
	}

	data1 := MemberLL{
		Container: member1,
		Next:      nil,
	}

	data2 := MemberLL{
		Container: member2,
		Next:      nil,
	}

	head.Next = &data1
	head.Next.Next = &data2

	fmt.Println("isi head : ", head.Next.Next)
}
