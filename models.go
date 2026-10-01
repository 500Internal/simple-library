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

func (lib *Library) FindBookByID(id int) (*Book, error) {
	book, ok := lib.Books[id]
	if !ok {
		return nil, fmt.Errorf("книга с ID %d не найдена", id)
	}
	return book, nil
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	reader, ok := lib.Readers[id]
	if !ok {
		return nil, fmt.Errorf("читатель с ID %d не найден", id)
	}
	return reader, nil
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	book.IssueBook(reader)
	return nil
}

func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}
