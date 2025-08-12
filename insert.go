package taos_native

import (
	"fmt"

	"github.com/taosdata/driver-go/v3/common"
)

// sql insert
func (td *TaosDriver) SqlInsert(insertQuery string) error {
	res, err := td.taos.Exec(insertQuery)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// you can check affectedRows here
	fmt.Printf("Successfully inserted %d rows to power.meters.\n", rowsAffected)
	return nil
}

// AF schemaless insert
func (td *TaosDriver) SchemalessInsert(data string) error {

	// insert influxdb line protocol
	err := td.afConn.InfluxDBInsertLinesWithReqID(data, td.cfg.Precision, common.GetReqID(), 0, "tname")
	if err != nil {
		return fmt.Errorf("Failed to insert data with schemaless, data:%s, ErrMessage:%v\n", data, err)
	}

	fmt.Println("Inserted data with schemaless successfully.")
	return nil
}
