package main

import "fmt"

func (h Human) SayHello() {
	fmt.Printf("Всем привет, моё ФИО: %s %s %s, мне %d лет!\n", h.FirstName, h.LastName, h.MiddleName, h.Age)
}

func (h Human) IsAdult() {
	if h.Age < 18 {
		fmt.Println("Несовершеннолетний")
	} else {
		fmt.Println("Совершеннолетний")
	}
}

func (a Action) DoSomething() {
	fmt.Printf("%s %s %s выполняет действие в роли %s\n", a.LastName, a.FirstName, a.MiddleName, a.Role)
}

func main() {
	user := Action{
		Human: Human{
			LastName:   "Тестов",
			FirstName:  "Тест",
			MiddleName: "Тестович",
			Age:        17,
		},
		Role: "Wildberries разработчик",
	}

	user.SayHello()
	user.IsAdult()
	user.DoSomething()
}
