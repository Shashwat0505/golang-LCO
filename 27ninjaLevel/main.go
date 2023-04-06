package main

import (
	"fmt"
)

var x int
var y string
var z bool

type dog int

func main() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v\t %v\t %v\t", x, y, z)
	fmt.Println(s)
	var x dog
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
	age := 0
	for age < 20 {
		fmt.Println(age)
		age++
	}

	x1 := []int{2, 3, 4}
	for _, v := range x1 {
		fmt.Println(v)
	}
	x1 = append(x1, 5, 65, 7, 8)
	fmt.Println(x1)
	x2 := []int{22, 44, 66, 77}
	x1 = append(x1, x2...)
	fmt.Println(x1)
	x1 = append([]int{0, 00, 000}, x1...)
	fmt.Println(x1)

	// y1 := make([]int, 3)
	// fmt.Println(len(y1))

	// for i := 0; i < 3; i++ {
	// 	var temp int
	// 	fmt.Println("Please Enter number")
	// 	fmt.Scan(&temp)
	// 	y1[i] = temp
	// }
	// fmt.Println(y1)

	m1 := map[int]string{
		1: "James",
		2: "Harry",
	}
	fmt.Println(m1)
	delete(m1, 2)
	fmt.Println(m1)
	arr11 := []int{1, 2, 3, 4, 5}
	for _, v := range arr11 {
		fmt.Println(v)
	}

	fmt.Printf("%d is value , and type is %T\n", arr11, arr11)

	sl11 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sl11)
	sl11 = append(sl11[:3], sl11[6:]...)
	fmt.Println(sl11)

	sl12 := []string{"James", "Bond", "Shaken,not stirred"}
	sl13 := []string{"Miss", "Moneypenny", "Helloooo,James"}
	sl14 := [][]string{sl12, sl13}
	fmt.Println(sl14)

	m11 := map[string][]string{
		"bond":   []string{"action", "fight"},
		"potter": []string{"magic", "action"},
	}
	fmt.Println(m11)

	type person struct {
		first string
		last  string
	}
	type secretAgent struct {
		person
		ltk bool
	}
	p1 := person{
		first: "Tom",
		last:  "Riddle",
	}
	fmt.Println(p1)
	agent := secretAgent{
		person: person{first: "James", last: "Bond"},
		ltk:    true,
	}
	fmt.Println(agent)

	p11 := struct {
		first string
		age   int
	}{
		first: "Harry",
		age:   23,
	}
	fmt.Println(p11)

	type p55 struct {
		name    string
		flavors []string
	}
	person22 := p55{
		name: "jerry",
		flavors: []string{
			"chocolate",
			"Mango",
		},
	}
	for _, v := range person22.flavors {
		fmt.Println(v)
	}
	person33 := p55{
		name: "Tom",
		flavors: []string{
			"vanila",
			"chickoo",
		},
	}
	for _, v := range person33.flavors {
		fmt.Println(v)
	}

	m55 := map[string]p55{
		"Tom": p55{name: "Tom", flavors: []string{
			"vanila",
			"chickoo",
		}},
	}
	fmt.Println(m55)

	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	t1 := truck{
		vehicle:   vehicle{doors: 12, color: "black"},
		fourWheel: true,
	}
	fmt.Println(t1)
	s1 := sedan{
		vehicle: vehicle{doors: 4, color: "red"},
		luxury:  false,
	}
	fmt.Println(s1)
	s44 := struct {
		name   string
		hasCar bool
	}{
		name:   "Tom",
		hasCar: false,
	}
	fmt.Println(s44)

}
