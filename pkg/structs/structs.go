package structs

type User struct {
	Userid   float64
	Username string
	IsAdmin  bool `default:"false"`
	Pass     []byte
}

type ListUsers struct {
	Users []User
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

type BorrowHistory struct {
	BookId      int
	Title       string
	Author      string
	Genre       string
	UserID      int
	Username    string
	RequestDate string
	ReturnDate  string
	AcceptDate  string
	Status      string
}

type ListBookReq struct {
	BorrowHistory []BorrowHistory
}

type PageMessage struct {
	Message interface{}
}

type Datasent struct {
	Results interface{} `default:"nil"`
	Message PageMessage `default:"nil"`
}
