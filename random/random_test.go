package random

import (
	"testing"
)

func TestGenerateRandomStr(t *testing.T) {
	t.Log("GenerateRandomStr Test Start")

	for idx := 0; idx < 10; idx++ {
		t.Log(GenerateRandomStr(Type.Hex, 20))
	}

	t.Log("GenerateRandomStr Test Done")
}

func TestGenUUID(t *testing.T) {
	t.Log("GenUUID Test Start")

	for idx := 0; idx < 10; idx++ {
		t.Log(GenUUID())
	}

	t.Log("GenUUID Test Done")
}
