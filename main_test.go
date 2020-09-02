package main

import (
	"fmt"
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestExecuteSQL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		t.Fail()
	}
	defer db.Close()

	// Just generate some random result - don't really care what it is...
	result := sqlmock.NewRows([]string{"id", "langcode", "title", "link__uri", "view_sidemenu"}).
        AddRow(1, "en", "enTitle", "/en-link", "0").
        AddRow(2, "en", "enTitle2", "/en-link2", "0")
	mock.ExpectQuery("SELECT @@version").WillReturnRows(result)

	row,startTime,duration,err := executeSQL(db, "SELECT @@version")
	if err != nil {
    	fmt.Printf("unexpected error: %s", err)
    	t.Fail()
	}

	t.Logf("\n%d, %s, %s\n", row, startTime, duration)
}