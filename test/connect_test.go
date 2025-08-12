package test

import (
	"testing"

	"github.com/raymone2306/taos_native"
)

func TestConnect(t *testing.T) {
	taos := taos_native.NewDriver("192.168.144.130", 6030, "root", "taosdata", "iotsql", "ms")
	err := taos.Connect()
	if err != nil {
		t.Error(err)
	}
	taos.Close()
}

func TestAFConnect(t *testing.T) {
	lineDemo := `meters,tname=t_merters,location=California.SanFrancisco current=10.3000002f64,voltage=219i32,phase=0.31f64 1626006833639
	meters2,tname=t_mergters2,location=California.SanFrancisco current=10.3000002f64,voltage=219i32,phase=0.31f64 1626006833639`
	taos := taos_native.NewDriver("192.168.144.130", 6030, "root", "taosdata", "iot", "ms")
	err := taos.AFConnect()
	defer taos.AFClose()
	if err != nil {
		t.Error(err)
	}
	err = taos.SchemalessInsert(lineDemo)
	if err != nil {
		t.Error(err)
	}
	t.Logf("SchemalessInsert lineDemo:%s", lineDemo)

}
