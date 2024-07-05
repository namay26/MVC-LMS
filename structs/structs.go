package structs

type User struct {
	Userid   float64
	Username string
	IsAdmin  bool `default:"false"`
	Pass     []byte
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
