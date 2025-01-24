package main

import "fmt"

type Member struct {
	Id    int
	Nama  string
	Point float32
	Link  *Member
}

func main() {
	member1 := Member{
		Id:    1,
		Nama:  "Kurniawan",
		Point: 0,
		Link:  nil,
	}

	member2 := Member{
		Id:    2,
		Nama:  "Aan",
		Point: 100,
		Link:  nil,
	}

	member1.Link = &member2
	fmt.Println("member 1", member1)
	fmt.Println("nama member 1 :", member1.Nama)
	fmt.Println("member 2", member1.Link)
	fmt.Println("nama member 2 :", member1.Link.Nama)
}
