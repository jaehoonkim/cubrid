package cubrid

import (
	"database/sql"
	"testing"
	"fmt"
	"log"
)

func openDb(t *testing.T, dsn string) *sql.DB {
	db, err := sql.Open("cubrid", dsn)
	if err != nil {
		t.Fatal(err)
	}
	return db
}
/*
func TestCubrid(t *testing.T) {
	fmt.Println("TestCubrid")
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	if err != nil {
		t.Fatal(err)
	} 

	defer db.Close()
}
*/
/*
func TestStmtQuery(t *testing.T) {
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
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	//if rows.Next() == false {
	//	t.Fatal(err)
	//}
	
	for rows.Next() == true {
		log.Println("test...0")
		var s_name, f_name string
		rows.Scan(&s_name, &f_name)
		log.Println("test...1")
		fmt.Printf("s : %s, f : %s\n", s_name, f_name)
	}
}
*/
/*
func TestStmtQueryParam(t *testing.T) {
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...0")
	stmt, err := db.Prepare("select * from code where s_name = ?")
	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...1")
	rows, err := stmt.Query("W")
	defer rows.Close()
	if err != nil {
		//log.Println("stmt.Query err")
		log.Println(err)
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...2")
	if rows.Next() == false {
	//	log.Println(err)
	//	log.Println("=======================")
		t.Fatal(err)
	}

	var s_name, f_name string
	rows.Scan(&s_name, &f_name)

	fmt.Printf("s : %s, f : %s\n", s_name, f_name)
}
//*/
/*
func TestStmtQueryBind_int(t *testing.T) {
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...0")
	stmt, err := db.Prepare("select * from athlete where code = ?")
	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...1")
	rows, err := stmt.Query(10999)
	defer rows.Close()
	if err != nil {
		//log.Println("stmt.Query err")
		log.Println(err)
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...2")
	if rows.Next() == false {
	//	log.Println(err)
	//	log.Println("=======================")
		t.Fatal(err)
	}

	var code int
	var name, gender, nation_code, event string
	rows.Scan(&code, &name, &gender, &nation_code, &event)

	fmt.Printf("code:%d, name:%s, gender:%s, nation_code:%s, event:%s\n", code, name, gender, nation_code, event)
}
//*/
/*
func TestStmtQueryBind_date(t *testing.T) {
	db, err := sql.Open("cubrid", "127.0.0.1/33000/demodb/dba/")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...0")
	stmt, err := db.Prepare("select * from game where game_date = ?")
	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...1")
	rows, err := stmt.Query("08/28/2004")
	defer rows.Close()
	if err != nil {
		//log.Println("stmt.Query err")
		log.Println(err)
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...2")
	if rows.Next() == false {
	//	log.Println(err)
	//	log.Println("=======================")
		t.Fatal(err)
	}

	log.Println("scan..before")
	var game_date GCI_DATE
	var host_year, event_code, athlete_code, stadium_code, nation_code, medal string
	rows.Scan(&host_year, &event_code, &athlete_code, &stadium_code, &nation_code, &medal, &game_date)
	
	fmt.Printf("host_year:%s, event_code:%s, athlete_code:%s, stadium_code:%s, nation_code:%s, medal:%s, game_date:%d,%d,%d\n", host_year, event_code, athlete_code, stadium_code, nation_code, medal, game_date.Yr(), game_date.Mon(), game_date.Day())
}
//*/


/*
	table name : tbl_bitn
	column
	idx : integer
	bitn : BIT_VARYING
*/
/*
func TestStmtQueryBind_bit(t *testing.T) {
	db := openDb(t, "127.0.0.1/33000/testdb/dba/1234")
	defer db.Close()
	if db.Driver() == nil {
		t.Fatal(fmt.Errorf("nil driver"))
	}
	//log.Println("TestPrepare: test...0")
	//stmt, err := db.Prepare("select * from tbl_bit")
	stmt, err := db.Prepare("select * from tbl_bitn")

	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	if rows.Next() == false {
		t.Fatal(err)
	}

	var buf GCI_BIT
	var idx int

	rows.Scan(&idx,&buf)
	fmt.Printf("idx : %d, size:%d, buf: %x\n", idx, buf.Size(), buf.Buf())
}
//*/
/*
	table name : tbl_set
	column
	idx : integer
	setn : SET
*/
/*
func TestStmtQueryBind_set(t *testing.T) {
	db := openDb(t, "127.0.0.1/33000/testdb/dba/1234")
	defer db.Close()
	if db.Driver() == nil {
		t.Fatal(fmt.Errorf("nil driver"))
	}
	stmt, err := db.Prepare("select * from set_tbl")

	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	if rows.Next() == false {
		t.Fatal(err)
	}

	//var buf CCI_SET
	var idx int
	var set GCI_SET
	rows.Scan(&idx, &set)
	
	size := Gci_set_size(set)
	d := make([]string, size)
	Gci_set_get(set, 0, A_TYPE_STR)
	if set.Size() > 0 {
		fmt.Printf("idx : %d, %s, %s, %s\n", idx, set.Buf(0), set.Buf(1), set.Buf(2))
	}

	
	rows.Next()
	rows.Scan(&idx, &set)
	fmt.Printf("idx : %d, %s, %s, %s\n", idx, set.Buf(0), set.Buf(1), set.Buf(2))

}
//*/
/*
func TestStmtQueryBind_clob(t *testing.T) {
	db := openDb(t, "127.0.0.1/33000/testdb/dba/1234")
	defer db.Close()
	if db.Driver() == nil {
		t.Fatal(fmt.Errorf("nil driver"))
	}
	stmt, err := db.Prepare("select * from tbl_clob")

	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	if rows.Next() == false {
		t.Fatal(err)
	}

	var idx int
	var clob CCI_CLOB
	rows.Scan(&idx, &clob)
	fmt.Printf("idx : %d, %s\n", idx, clob.Buf())
}
//*/
/*
func TestStmtQueryBind_blob(t *testing.T) {
	db := openDb(t, "127.0.0.1/33000/testdb/dba/1234")
	defer db.Close()
	if db.Driver() == nil {
		t.Fatal(fmt.Errorf("nil driver"))
	}
	stmt, err := db.Prepare("select * from tbl_blob")

	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	if rows.Next() == false {
		t.Fatal(err)
	}

	var idx int
	var blob CCI_BLOB
	rows.Scan(&idx, &blob)
	fmt.Printf("idx : %d, buf : %x\n", idx, blob.Buf())
}
*/
