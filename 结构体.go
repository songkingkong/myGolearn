package main

import "fmt"

/*结构体创建方式1*/
type Movie struct {
	Name   string
	Rating float32
}

/*结构体创建方式1*/
type Movie1 struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "King",
		Rating: 10,
	}
	var c = Movie1{}
	c.Name = "Kong"
	c.Rating = 1.0
	fmt.Println(m, c)
	fmt.Printf("%+v", c)
}
