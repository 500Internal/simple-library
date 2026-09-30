package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	userID := 1
	user1 := Reader{
		ID:        &userID,
		FirstName: "Пётр",
		LastName:  "Петров",
		IsActive:  true,
	}

	book1 := Book{
		ID:     1,
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	book1.IssueBook(&user1)
	fmt.Println(book1)
	fmt.Println("---")

	reader2ID := 2
	reader2 := Reader{
		ID:        &reader2ID,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}
	book1.IssueBook(&reader2)
	fmt.Println("---")

	user1.IsActive = false
	fmt.Println(user1)
	fmt.Println("---")

	book1.IssueBook(&user1)
	fmt.Println("---")

	book1.ReturnBook()
	fmt.Println(book1)

	notifiers := []Notifier{
		EmailNotifier{EmailAddress: "student@example.com"},
		SMSNotifier{PhoneNumber: "+79991234567"},
	}

	for _, n := range notifiers {
		n.Notify("Ваша книга просрочена!")
	}
}
