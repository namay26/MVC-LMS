package tests

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/namay26/MVC-LMS/pkg/model"
)

type PayloadData struct {
	id  string
	err error
}

var payloads = []PayloadData{
	{
		id:  "4",
		err: nil,
	},
	{
		id:  "",
		err: strconv.ErrSyntax,
	},
	{
		id:  "abc",
		err: strconv.ErrSyntax,
	},
	{
		id:  "0",
		err: sql.ErrNoRows,
	},
}

func TestGetBook(t *testing.T) {

	db, _ := model.Connect(false)
	for _, book := range payloads {
		_, err := model.GetBook(db, book.id)
		if err != nil {
			if err != book.err {
				t.Fatal(err.Error(), book.err)
			}
		}
	}
}
