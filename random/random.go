package random

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/eliassama/go-utils/convert"

	"github.com/google/uuid"
)

var randomScope = struct {
	int   []string
	lower []string
	upper []string
	mix   []string
}{
	int: []string{
		string('0'), string('1'), string('2'),
		string('3'), string('4'), string('5'),
		string('6'), string('7'), string('8'),
		string('9'),
	},
	lower: []string{
		string('a'), string('b'), string('c'),
		string('d'), string('e'), string('f'),
		string('g'), string('h'), string('i'),
		string('j'), string('k'), string('l'),
		string('m'), string('n'), string('o'),
		string('p'), string('q'), string('r'),
		string('s'), string('t'), string('u'),
		string('v'), string('w'), string('x'),
		string('y'), string('z'),
	},
	upper: []string{
		string('A'), string('B'), string('C'),
		string('D'), string('E'), string('F'),
		string('G'), string('H'), string('I'),
		string('J'), string('K'), string('L'),
		string('M'), string('N'), string('O'),
		string('P'), string('Q'), string('R'),
		string('S'), string('T'), string('U'),
		string('V'), string('W'), string('X'),
		string('Y'), string('Z'),
	},
	mix: []string{
		string('o'), string('O'), string('0'),
		string('i'), string('I'), string('L'),
		string('l'), string('1'), string('q'),
		string('g'), string('9'), string('Q'),
		string('G'),
	},
}

/**
 *
 * 针对于可生成的随机字符串类型，做如下解释：
 *      Int：生成 纯数字 随机符串
 *      Lower：生成 纯小写字母 随机字符串
 *      Upper：生成 纯大写字母 随机字符串
 *      Letter：生成 纯字母（含大写及小写） 随机字符串
 *      IntLower：生成 小写字母 + 数字 随机字符串
 *      IntUpper：生成 大写字母 + 数字 随机字符串
 *      IntLetter：生成 大写字母 + 小写字母 + 数字 随机字符串
 *      Bin：生成 二进制 随机字符串
 *      Oct：生成 八进制 随机字符串
 *      Dec：生成 十进制 随机字符串，可以看作是Int的别名
 *      HexLower：生成 小写的十六进制 随机字符串
 *      HexUpper：生成 大写的十六进制 随机字符串
 *      Hex：生成 包含有大写及小写的十六进制 随机字符串
 */

// Type 可生成的随机字符串类型
var Type = struct {
	Int       int
	Lower     int
	Upper     int
	Letter    int
	IntLower  int
	IntUpper  int
	IntLetter int
	Bin       int
	Oct       int
	Dec       int
	HexLower  int
	HexUpper  int
	Hex       int
}{
	Int:       0,
	Lower:     1,
	Upper:     2,
	Letter:    3,
	IntLower:  4,
	IntUpper:  5,
	IntLetter: 6,
	Bin:       7,
	Oct:       8,
	Dec:       9,
	HexLower:  10,
	HexUpper:  11,
	Hex:       12,
}

// GenerateRandomStr 生成随机字符串
func GenerateRandomStr(randomType int, length int) string {
	var source []string

	switch randomType {
	case Type.Int:
		fallthrough
	case Type.Dec:
		source = append(source, randomScope.int...)
	case Type.Lower:
		source = append(source, randomScope.lower...)

	case Type.Upper:
		source = append(source, randomScope.upper...)

	case Type.Letter:
		source = append(source, randomScope.upper...)
		source = append(source, randomScope.lower...)

	case Type.IntLower:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.lower...)

	case Type.IntUpper:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.upper...)

	case Type.IntLetter:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.upper...)
		source = append(source, randomScope.lower...)

	case Type.HexLower:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.lower[:6]...)

	case Type.HexUpper:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.upper[:6]...)

	case Type.Hex:
		source = append(source, randomScope.int...)
		source = append(source, randomScope.upper[:6]...)
		source = append(source, randomScope.lower[:6]...)

	case Type.Bin:
		source = append(source, randomScope.int[:2]...)
		source = append(source, randomScope.int[:2]...)
		source = append(source, randomScope.int[:2]...)

	case Type.Oct:
		source = append(source, randomScope.int[:8]...)
	}

	var randomStr strings.Builder

	for idx := 0; idx < length; idx++ {
		pos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
		randomStr.WriteString(source[pos.Int64()])
	}

	return randomStr.String()
}

// GenUUID 获取 UUID
func GenUUID() string {
	var id [36 + 8]byte
	var idIndex uint64

	copy(id[:], convert.StringToBytes(uuid.New().String()))
	index := atomic.AddUint64(&idIndex, 1)
	binary.BigEndian.PutUint64(id[36:], index)

	return convert.BytesToString(id[:36]) + strconv.Itoa(int(binary.BigEndian.Uint64(id[:36])))
}
