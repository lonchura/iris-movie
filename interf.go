package main

import "fmt"

type Phone interface {
	call()
}

type Nokia struct {
}

type IPhone struct {
}

func (phone Nokia) call() {
	fmt.Printf("I am Nokia!\n")
}

func (phone IPhone) call() {
	fmt.Printf("I am IPhone!\n")
}

func phoneInfo(phone *Phone) {
	phone.call()
}

func main()  {
	var phone Phone

	phone = Nokia{}
	phoneInfo(&phone)

	phone = IPhone{}
	phoneInfo(&phone)
}