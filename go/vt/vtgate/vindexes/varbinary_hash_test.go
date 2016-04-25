package vindexes

import (
	"reflect"
	"testing"
)

var varbinaryHash Vindex

func init() {
	varbinaryHash, _ = CreateVindex("varbinaryHash", "vch", nil)
}

func TestVarbinaryHashCost(t *testing.T) {
	if varbinaryHash.Cost() != 1 {
		t.Errorf("Cost(): %d, want 1", varbinaryHash.Cost())
	}
}

func TestVarBinaryMap(t *testing.T) {
	got, err := varbinaryHash.(Unique).Map(nil, []interface{}{[]byte("\x74\x65\x73\x74\x74\x65\x73\x74")})
	if err != nil {
		t.Error(err)
	}
	want := [][]byte{
		[]byte("\x45\x23\x56\x06\x86\xef\x04\x91"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map(): %#v, want %+v", got, want)
	}
}

func TestVarBinaryVerify(t *testing.T) {
	success, err := varbinaryHash.Verify(nil, []byte("\x74\x65\x73\x74\x74\x65\x73\x74"), []byte("\x45\x23\x56\x06\x86\xef\x04\x91"))
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Errorf("Verify(): %+v, want true", success)
	}
}
