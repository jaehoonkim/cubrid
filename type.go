package cubrid

/*
#include "cas_cci.h"
*/
import "C"

import "unsafe"
//import "fmt"

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
	_BLOB C.T_CCI_BLOB
}

type CCI_CLOB struct {
	_CLOB C.T_CCI_CLOB
}

/*-----------------------------------------------*/
func (date *CCI_DATE) yr() uint {
	return uint(date._DATE.yr)
}

func (date *CCI_DATE) mon() uint {
	return uint(date._DATE.mon)
}

func (date *CCI_DATE) day() uint {
	return uint(date._DATE.day)
}

func (date *CCI_DATE) hh() uint {
	return uint(date._DATE.hh)
}

func (date *CCI_DATE) mm() uint {
	return uint(date._DATE.mm)
}

func (date *CCI_DATE) ss() uint {
	return uint(date._DATE.ss)
}

func (date *CCI_DATE) ms() uint {
	return uint(date._DATE.ms)
}

/**************************************/
func (bit *CCI_BIT) size() int {
	return int(bit._BIT.size)
}

func (bit *CCI_BIT) buf() string {
	return C.GoStringN(bit._BIT.buf, bit._BIT.size)
}

/**************************************/
func (set *CCI_SET) Buf(idx int) string {
	return set.buf[idx]
}

func (set *CCI_SET) makeBuf(size int) {
	set.size = size
	set.buf = make([]string, size)
}

func (set *CCI_SET) setBuf(idx int, buf string) {
	set.buf[idx] = buf
}

/*************************************/
func (blob *CCI_BLOB) buf() unsafe.Pointer {
	return unsafe.Pointer(blob._BLOB)
}

/*************************************/
func (clob *CCI_CLOB) buf() string {
	return C.GoString((*C.char)(clob._CLOB))
}

