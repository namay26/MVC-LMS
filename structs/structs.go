package structs

type User struct {
	Userid   int
	Username string
	IsAdmin  bool
	Pass     string
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Genre    string
	Quantity int
}

type ListBooks struct {
	Books []Book
}
