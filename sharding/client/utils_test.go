package client

import (
	"math/rand"
	"reflect"
	//"runtime"
	"testing"
)

var testbody interface{}

func buildblob(size int64) []byte {

	tempbody := make([]byte, size)
	for i := int64(0); i < size; i++ {
		tempbody[i] = byte(rand.Int())

	}

	return tempbody

}
func TestConvertInterface(t *testing.T) {
	var slice interface{}
	slice = []interface{}{0, 1, 2, 3, 4, 5}
	convertedValue, err := convertInterface(slice, reflect.Slice)
	if err != nil {
		t.Fatalf("Error: %v %v", err, convertedValue)
	}

}

func TestSerializeblob(t *testing.T) {

	blob := buildblob(20)

	serializedblob, err := serializeBlob(blob)

	if err != nil {
		t.Fatalf("Error Serializing blob:%v %v", err, serializedblob)
	}
	test := &testbody
	//runtime.Breakpoint()
	err2 := Deserializebody(serializedblob, &testbody)
	if err2 != nil {
		t.Fatalf("Error Serializing blob:%v", err2)
	}

	if !reflect.DeepEqual(blob, testbody) {

		t.Fatalf("Error Serializing blob with %v %v %v", blob, test, &testbody)
	}

}
