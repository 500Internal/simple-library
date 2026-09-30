package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

type Reader struct {
	ID        *int
	FirstName string
	LastName  string
	IsActive  bool
}

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}

	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n",
			reader.FirstName, reader.LastName)
		return
	}

	b.IsIssued = true
	b.ReaderID = reader.ID

	reader.AssignBook(b)

	fmt.Printf("Книга %s была выдана читателю %s %s\n",
		b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}

	b.IsIssued = false
	b.ReaderID = nil
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

func (r *Reader) AssignBook(book *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s\n",
		r.FirstName, r.LastName, book)
}
