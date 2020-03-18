package main

import (
	"fmt"
	"time"
)

// create oop
type Student struct {
	firstName string
	lastName  string
}

// value receiver | Pointer receiver
// value receiver allow create a copy in function Email and dont effect original
// pointer receiver allow work original, add( * )
func (s *Student) FirstName() string {

	s.firstName = "abc234@gmail.com"
	return s.firstName
}

func demoBasic() {

	c := make(chan int)
	go func() {
		c <- 100
	}()

	go func() {
		fmt.Println(<-c)
	}()

	go g2()
	time.Sleep(time.Second)
}

func g2() {
	fmt.Println("2")
}

// cach tao func trong golang
func sayHello() {
	fmt.Println("Say Hello Func")
}
func sum(a int, b int) int {
	return a + b
}
func demoGo() {
	var email = "abc"
	fmt.Print(email)

	// sorhand
	fullName := "Hello World"
	fmt.Print(fullName)

	var number = 10
	fmt.Print(number)
	fmt.Println()
}

func demostructure() {
	s := make([]string, 0)
	s = append(s, "m")
	s = append(s, "n")
	fmt.Println(s)

	//map
	m := make(map[string]int)
	m["key1"] = 90
	fmt.Println(m)

	// arrray
	arr := [2]int{}
	arr[0] = 50
	arr[1] = 22
	fmt.Println(arr)
}

func demostrucure1() {
	st := Student{
		firstName: "join",
		lastName:  "loa",
	}

	e := st.FirstName()
	fmt.Println(e)

	fmt.Println(st.firstName)
}
