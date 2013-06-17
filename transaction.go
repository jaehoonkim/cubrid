package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import "fmt"

type cubridTx struct {
	c *cubridConn
}
func (tx *cubridTx) Commit() error {
	var conn_handle C.int
	conn_handle = tx.c.con
	var err C.int
	var cci_error C.T_CCI_ERROR

	err = C.cci_end_tran(conn_handle, C.CCI_TRAN_COMMIT, &cci_error)
	if err < 0 {
		var err_discon C.T_CCI_ERROR
		C.cci_disconnect(conn_handle, &err_discon)
		return fmt.Errorf("commit error : %d", cci_error.err_code)
	}
	return nil
}

func (tx *cubridTx) Rollback() error {
	var conn_handle C.int
	conn_handle = tx.c.con
	var err C.int
	var cci_error C.T_CCI_ERROR

	err = C.cci_end_tran(conn_handle, C.CCI_TRAN_ROLLBACK, &cci_error)
	if err < 0 {
		var err_discon C.T_CCI_ERROR
		C.cci_disconnect(conn_handle, &err_discon)
		return fmt.Errorf("commit error : %d", cci_error.err_code)
	}
	return nil
}

