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

	result := sqlmock.NewResult(0,1)
	mock.ExpectExec("SELECT @@version").WillReturnResult(result)

	row,startTime,duration,err := executeSQL(db, "SELECT @@version")
	if err != nil {
    	fmt.Printf("unexpected error: %s", err)
    	t.Fail()
	}

	t.Logf("\n%d, %s, %s\n", row, startTime, duration)
}