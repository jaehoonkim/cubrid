package cubrid

import (
	"database/sql"
	"testing"
	"fmt"
)

func TestCubrid(t *testing.T) {
	fmt.Println("TestCubrid")
	db, err := sql.Open("cubrid", "localhost/33000/demodb/dba/1212123")
	if err != nil {
		t.Fatal(err)
	} 

	defer db.Close()
}

func TestPrepare(t *testing.T) {
	fmt.Println("TestPrepare0")
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	fmt.Println("TestPrepare1")
	stmt, err := db.Prepare("select * from code")
	fmt.Println("TestPrepare2")
	defer stmt.Close()
	if err != nil {
		fmt.Println("TestPrepare3")
		t.Fatal(err)
	}
	fmt.Println("TestPreparei end")
}
