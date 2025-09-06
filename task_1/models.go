package main

type Human struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
}

type Action struct {
	Human
	Role string
}
