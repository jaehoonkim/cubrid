package cubrid


import (
	"database/sql"
	"testing"
	"fmt"
	"log"
	//"unsafe"
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
		var s_name, f_name string
		rows.Scan(&s_name, &f_name)

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
*/
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
*/
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

	var game_date CCI_DATE
	var host_year, event_code, athlete_code, stadium_code, nation_code, medal string
	rows.Scan(&host_year, &event_code, &athlete_code, &stadium_code, &nation_code, &medal, &game_date)
	
	fmt.Printf("host_year:%s, event_code:%s, athlete_code:%s, stadium_code:%s, nation_code:%s, medal:%s, game_date:%d,%d,%d\n", host_year, event_code, athlete_code, stadium_code, nation_code, medal, game_date.yr(), game_date.mon(), game_date.day())
}
*/

/*
	table name : tbl_bitn
	column
	idx : integer
	bitn : BIT_VARYING
*/
func TestStmtQueryBind_bit(t *testing.T) {
	db, err := sql.Open("cubrid", "127.0.0.1/33000/testdb/dba/1234")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	if db.Driver() == nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...0")
	//stmt, err := db.Prepare("select * from tbl_bit")
	stmt, err := db.Prepare("select * from tbl_bitn")

	defer stmt.Close()
	if err != nil {
		t.Fatal(err)
	}
	//log.Println("TestPrepare: test...1")
	rows, err := stmt.Query()
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

	var buf CCI_BIT
	var idx int

	rows.Scan(&idx,&buf)
	fmt.Printf("idx : %d, size:%d, buf: %s\n", idx, buf.size(), buf.buf())

	//if rows.Next() == false {
	//	t.Fatal(err)
	//}

	//rows.Scan(&idx,&buf)
	//fmt.Printf("idx : %d, size:%d, buf: %x\n", idx, buf.size(), buf.buf())

}

