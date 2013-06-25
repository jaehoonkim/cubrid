package cubrid

/*
#include "cas_cci.h"
*/
import "C"

type CCI_DATE struct {
	_DATE C.T_CCI_DATE
}

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

