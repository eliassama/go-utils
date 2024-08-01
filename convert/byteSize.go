package convert

import (
	"math"
	"math/big"
)

//goland:noinspection ALL
const (
	Byte = float64(iota)
	KB
	MB
	GB
	TB
	PB
)

var Units = map[float64]string{
	0: "Byte",
	1: "KB",
	2: "MB",
	3: "GB",
	4: "TB",
	5: "PB",
}

// ByteSize bytes转换为指定单位
//
//goland:noinspection GoUnusedExportedFunction
func ByteSize(bytes float64, places uint, units ...float64) (float64, string) {
	var level float64
	var negative bool

	if bytes < 0 {
		negative = true
		bytes = -bytes
	}

	if units != nil && len(units) > 0 {
		level = units[0]
	} else {
		level = math.Floor(math.Log(bytes) / math.Log(1024))
	}

	value, _ := big.NewFloat(0).Quo(big.NewFloat(bytes), big.NewFloat(math.Pow(1024, level))).Float64()

	if negative {
		return -RoundFloat(value, places), Units[level]
	} else {
		return RoundFloat(value, places), Units[level]
	}
}
