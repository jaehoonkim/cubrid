package cubrid

import "fmt"

type cubridTx struct {
	c *cubridConn
}
func (tx *cubridTx) Commit() error {
	var conn_handle int
	conn_handle = tx.c.con
	var err int
	var cci_error CCI_ERROR

	err, cci_error = gci_end_tran(conn_handle, TRAN_COMMIT)
	if err < 0 {
		var err_discon CCI_ERROR
		_, err_discon = gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", cci_error.err_code)
	}
	return nil
}

func (tx *cubridTx) Rollback() error {
	var conn_handle int
	conn_handle = tx.c.con
	var err int
	var cci_error CCI_ERROR

	err, cci_error = gci_end_tran(conn_handle, TRAN_ROLLBACK)
	if err < 0 {
		var err_discon CCI_ERROR
		_, err_discon = gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", cci_error.err_code)
	}
	return nil
}

