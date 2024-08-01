package convert

import "testing"

func TestStringBytes(t *testing.T) {
	t.Log("StringBytes Test Start")

	str := "acdc"

	convertByte := StringToBytes(str)

	convertStr := BytesToString(convertByte)

	t.Logf("str: %s, convertByte: %s, convertStr: %s, byte: %s", str, convertByte, convertStr, []byte(str))

	if str != convertStr {
		t.Error("str != convertStr")
	}

	t.Log("StringBytes Test Done")
}
