package cubrid

import (
	"database/sql"
	"testing"
//	"fmt"
	"log"
)
/*
func TestCubrid(t *testing.T) {
	fmt.Println("TestCubrid")
	db, err := sql.Open("cubrid", "localhost/33000/demodb/dba/1212123")
	if err != nil {
		t.Fatal(err)
	} 

	defer db.Close()
}
*/
func TestPrepare(t *testing.T) {
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	stmt, err := db.Prepare("select * from code")
	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("TestPrepare: test...0")
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	log.Println("TestPrepare: test...1")
	if rows.Next() == false {
		log.Println(err)
		log.Println("=======================")
		t.Fatal(err)
	}
}
