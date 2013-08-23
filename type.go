package cubrid

/*
#include "cas_cci.h"
*/
import "C"

//import "log"

type AUTOCOMMIT_MODE int
const (
	AUTOCOMMIT_FALSE AUTOCOMMIT_MODE = iota
	AUTOCOMMIT_TRUE
)

const (
	TRAN_COMMIT = 1
	TRAN_ROLLBACK = 2
)

type GCI_A_TYPE int
const (
	_ = iota
	A_TYPE_STR GCI_A_TYPE = iota + 1
	A_TYPE_INT
	A_TYPE_FLOAT
	A_TYPE_DOUBLE
	A_TYPE_BIT
	A_TYPE_DATE
	A_TYPE_SET
	A_TYPE_BIGINT
	A_TYPE_BLOB
	A_TYPE_CLOB
	A_TYPE_LAST
)

type GCI_U_TYPE int
const (
	U_TYPE_FIRST GCI_U_TYPE = 0
	U_TYPE_UNKNOWN GCI_U_TYPE = 0
	U_TYPE_NULL GCI_U_TYPE = 0
	
	U_TYPE_CHAR GCI_U_TYPE = 1
	U_TYPE_STRING GCI_U_TYPE = 2
	U_TYPE_NCHAR GCI_U_TYPE = 3
	U_TYPE_VARCHAR GCI_U_TYPE = 4
	U_TYPE_BIT GCI_U_TYPE = 5
	U_TYPE_VARBIT GCI_U_TYPE = 6
	U_TYPE_NUMERIC GCI_U_TYPE = 7
	U_TYPE_INT GCI_U_TYPE = 8
	U_TYPE_SHORT GCI_U_TYPE = 9
	U_TYPE_MONETARY GCI_U_TYPE = 10
	U_TYPE_FLOAT GCI_U_TYPE = 11
	U_TYPE_DOUBL GCI_U_TYPE = 12
	U_TYPE_DATE GCI_U_TYPE = 13
	U_TYPE_TIME GCI_U_TYPE = 14
	U_TYPE_TIMESTAMP GCI_U_TYPE = 15
	U_TYPE_SET GCI_U_TYPE = 16
	U_TYPE_MULTISET GCI_U_TYPE = 17
	U_TYPE_SEQUENCE GCI_U_TYPE = 18
	U_TYPE_OBJECT GCI_U_TYPE = 19
	U_TYPE_RESULTSET GCI_U_TYPE = 20
	U_TYPE_BIGINT GCI_U_TYPE = 21
	U_TYPE_DATETIME GCI_U_TYPE = 22
	U_TYPE_BLOB GCI_U_TYPE = 23
	U_TYPE_CLOB GCI_U_TYPE = 24
	U_TYPE_ENUM GCI_U_TYPE = 25

	U_TYPE_LAST GCI_U_TYPE = U_TYPE_ENUM


)

const GCI_BIND_PTR int = 1
/*
type CCI_ERROR struct {
	Err_code int
	Err_msg string
}
*/
type GCI_ERROR struct {
	Err_code int
	Err_msg string
}

type GCI_COL_INFO struct {
	u_type GCI_U_TYPE
	is_non_null string
	scale int16
	precision int
	col_name string
	real_attr string
	class_name string
	default_value string
	is_auto_increment bool
	is_unique_key bool
	is_primary_key bool
	is_foreign_key bool
	is_reverse_index bool
	is_reverse_unique bool
	is_shared bool


}

type GCI_CUBRID_STMT int
const (
	CUBRID_STMT_ALTER_CLASS GCI_CUBRID_STMT = iota
	CUBRID_STMT_ALTER_SERIAL
	CUBRID_STMT_COMMIT_WORK
	CUBRID_STMT_REGISTER_DATABASE
	CUBRID_STMT_CREATE_CLASS
	CUBRID_STMT_CREATE_INDEX
	CUBRID_STMT_CREATE_TRIGGER
	CUBRID_STMT_CREATE_SERIAL
	CUBRID_STMT_DROP_DATABASE
	CUBRID_STMT_DROP_CLASS
	CUBRID_STMT_DROP_INDEX
	CUBRID_STMT_DROP_LABEL
	CUBRID_STMT_DROP_TRIGGER
	CUBRID_STMT_DROP_SERIAL
	CUBRID_STMT_EVALUATE
	CUBRID_STMT_RENAME_CLASS
	CUBRID_STMT_ROLLBACK_WORK
	CUBRID_STMT_GRANT
	CUBRID_STMT_REVOKE
	CUBRID_STMT_STATISTICS
	CUBRID_STMT_INSERT
	CUBRID_STMT_SELECT
	CUBRID_STMT_UPDATE
	CUBRID_STMT_DELETE
	CUBRID_STMT_CALL
	CUBRID_STMT_GET_ISO_LVL
	CUBRID_STMT_GET_TIMEOUT
	CUBRID_STMT_GET_OPT_LVL
	CUBRID_STMT_SET_OPT_LVL
	CUBRID_STMT_SCOPE
	CUBRID_STMT_GET_TRIGGER
	CUBRID_STMT_SET_TRIGGER
	CUBRID_STMT_SAVEPOINT
	CUBRID_STMT_PREPARE
	CUBRID_STMT_ATTACH
	CUBRID_STMT_USE
	CUBRID_STMT_REMOVE_TRIGGER
	CUBRID_STMT_RENAME_TRIGGER
	CUBRID_STMT_ON_LDB
	CUBRID_STMT_GET_LDB
	CUBRID_STMT_SET_LDB
	CUBRID_STMT_GET_STATS
	CUBRID_STMT_CREATE_USER
        CUBRID_STMT_DROP_USER
	CUBRID_STMT_ALTER_USER
)

type CCI_DATE struct {
	_DATE C.T_CCI_DATE
}

/*
 cubrid manual
 - 비트열은 0과 1로 이루어진 이진 값의 순열(sequence) 이다.
 - 2진수 형식으로 사용할 때에는 다음과 같이 문자 B뒤에 0과 1로 이루어진 문자열을 붙이거나, 
   0b 뒤에 값을 붙여 표현한다.
   ex) B'1010'
       0b1010
 - 16진수 형식은 대문자 X 뒤에 0-9 그리고 A-F 문자로 이루어진 문자열을 붙이거나
   0x 뒤에 값을 붙여 표현한다.
   ex) X'a'
       0xA
*/
type CCI_BIT struct {
	_BIT C.T_CCI_BIT
}

type CCI_SET struct {
	//size C.int
	size int
	buf []string
}

type CCI_BLOB struct {
	_BLOB []byte
}

type CCI_CLOB struct {
	_CLOB string
}

/*-----------------------------------------------*/
func (date *CCI_DATE) Yr() uint {
	return uint(date._DATE.yr)
}

func (date *CCI_DATE) Mon() uint {
	return uint(date._DATE.mon)
}

func (date *CCI_DATE) Day() uint {
	return uint(date._DATE.day)
}

func (date *CCI_DATE) Hh() uint {
	return uint(date._DATE.hh)
}

func (date *CCI_DATE) Mm() uint {
	return uint(date._DATE.mm)
}

func (date *CCI_DATE) Ss() uint {
	return uint(date._DATE.ss)
}

func (date *CCI_DATE) Ms() uint {
	return uint(date._DATE.ms)
}

/**************************************/
func (bit *CCI_BIT) Size() int {
	return int(bit._BIT.size)
}

func (bit *CCI_BIT) Buf() string {
	return C.GoStringN(bit._BIT.buf, bit._BIT.size)
}

/**************************************/
func (set *CCI_SET) Buf(idx int) string {
	return set.buf[idx]
}

func (set *CCI_SET) Size() int {
	return set.size
}

func (set *CCI_SET) MakeBuf(size int) {
	set.size = size
	set.buf = make([]string, size)
}

func (set *CCI_SET) SetBuf(idx int, buf string) {
	set.buf[idx] = buf
}

/*************************************/
func (blob *CCI_BLOB) Buf() []byte {
	//log.Printf("type:%x", blob._BLOB)
	return blob._BLOB
}

/*************************************/
func (clob *CCI_CLOB) Buf() string {
	//log.Printf("type:%s", clob._CLOB)
	return clob._CLOB
}
/*
func (clob *CCI_CLOB) getBytes() []byte {
	return ([]byte)(unsafe.Pointer(clob._CLOB))
}
*/
