package crypto

import (
	"strconv"
	"testing"
)

func TestGetSHA256(t *testing.T) {
	t.Log("GetSHA256 Test Start")
	for idx := 0; idx < 2; idx++ {
		t.Log(GetSHA256("test"))
		t.Log(GetSHA256("test", strconv.FormatInt(int64(idx), 10)))
	}
	t.Log("GetSHA256 Test Done")
}

func TestGetSHA512(t *testing.T) {
	t.Log("GetSHA512 Test Start")
	for idx := 0; idx < 2; idx++ {
		t.Log(GetSHA512("test"))
		t.Log(GetSHA512("test", strconv.FormatInt(int64(idx), 10)))
	}
	t.Log("GetSHA512 Test Done")
}

func TestGetMD5(t *testing.T) {
	t.Log("GetMD5 Test Start")
	for idx := 0; idx < 2; idx++ {
		t.Log(GetMD5("test"))
		t.Log(GetMD5("test", strconv.FormatInt(int64(idx), 10)))
	}
	t.Log("GetMD5 Test Done")
}

func TestGetHmac(t *testing.T) {
	t.Log("GetHmac Test Start")
	for idx := 0; idx < 2; idx++ {
		t.Log(GetHmacMd5("test", "asad"))
		t.Log(GetHmacMd5("test", strconv.FormatInt(int64(idx), 10)))
		t.Log(GetHmacSHA256("test", "asad"))
		t.Log(GetHmacSHA256("test", strconv.FormatInt(int64(idx), 10)))
		t.Log(GetHmacSHA512("test", "asad"))
		t.Log(GetHmacSHA512("test", strconv.FormatInt(int64(idx), 10)))
	}
	t.Log("GetHmac Test Done")
}

func TestGetBase64StdEncode(t *testing.T) {
	t.Log("GetBase64StdEncode Test Start")
	str := "omg! what's happend? emmm... look-up~ () he is 1, i cost 1000$, *&*"
	encode := GetBase64StdEncode(str)
	decode, err := GetBase64StdDecode(encode)
	t.Logf("str: %s, encode: %s, decode: %s", str, encode, decode)

	if err != nil {
		t.Error(err)
	}

	if decode != str {
		t.Error("GetBase64StdEncode decode != str")
	}

	t.Log("GetBase64StdEncode Test Done")
}

func TestGetBase64URLEncode(t *testing.T) {
	t.Log("GetBase64URLEncode Test Start")
	str := "omg! what's happend? emmm... look-up~ () he is 1, i cost 1000$, *&*"
	encode := GetBase64URLEncode(str)
	decode, err := GetBase64URLDecode(encode)
	t.Logf("str: %s, encode: %s, decode: %s", str, encode, decode)

	if err != nil {
		t.Error(err)
	}

	if decode != str {
		t.Error("GetBase64URLEncode decode != str")
	}

	t.Log("GetBase64URLEncode Test Done")
}

func TestGetBase64RawStdEncode(t *testing.T) {
	t.Log("GetBase64RawStdEncode Test Start")
	str := "omg! what's happend? emmm... look-up~ () he is 1, i cost 1000$, *&*"
	encode := GetBase64RawStdEncode(str)
	decode, err := GetBase64RawStdDecode(encode)
	t.Logf("str: %s, encode: %s, decode: %s", str, encode, decode)

	if err != nil {
		t.Error(err)
	}

	if decode != str {
		t.Error("GetBase64RawStdEncode decode != str")
	}

	t.Log("GetBase64RawStdEncode Test Done")
}
