package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"strings"

	"github.com/eliassama/go-utils/convert"
)

// GetHashCode 生成 HASH 字符串
func GetHashCode(hashAdo hash.Hash, items ...string) string {
	var itemStr strings.Builder
	for _, item := range items {
		itemStr.WriteString(item)
		itemStr.WriteString(":")
		hashAdo.Write(convert.StringToBytes(item))
	}
	hashAdo.Write(convert.StringToBytes(itemStr.String()))

	return hex.EncodeToString(hashAdo.Sum(nil))
}

// GetSHA256 生成 SHA256 哈希值
func GetSHA256(items ...string) string {
	return GetHashCode(sha256.New(), items...)
}

// GetSHA512 生成 SHA512 哈希值
func GetSHA512(items ...string) string {
	return GetHashCode(sha512.New(), items...)
}

// GetMD5 生成 md5 哈希值
func GetMD5(items ...string) string {
	return GetHashCode(md5.New(), items...)
}

// GetHmac 使用自定义 hash 方法生成 Hmac 密钥哈希值
func GetHmac(hashFn func() hash.Hash, key string, items ...string) string {
	return GetHashCode(hmac.New(hashFn, []byte(key)), items...)
}

// GetHmacMd5 生成 md5 Hmac 密钥哈希值
func GetHmacMd5(key string, items ...string) string {
	return GetHmac(md5.New, key, items...)
}

// GetHmacSHA256 生成 SHA256 Hmac 密钥哈希值
func GetHmacSHA256(key string, items ...string) string {
	return GetHmac(sha256.New, key, items...)
}

// GetHmacSHA512 生成 SHA512 Hmac 密钥哈希值
func GetHmacSHA512(key string, items ...string) string {
	return GetHmac(sha512.New, key, items...)
}

// GetBase64StdEncode base64 标准方式编码
func GetBase64StdEncode(data string) string {
	return base64.StdEncoding.EncodeToString(convert.StringToBytes(data))
}

// GetBase64StdDecode base64 标准方式解码
func GetBase64StdDecode(data string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(data)
	return convert.BytesToString(res), err
}

// GetBase64URLEncode base64 URL和文件名安全编码
func GetBase64URLEncode(data string) string {
	return base64.URLEncoding.EncodeToString(convert.StringToBytes(data))
}

// GetBase64URLDecode base64 URL和文件名安全解码
func GetBase64URLDecode(data string) (string, error) {
	res, err := base64.URLEncoding.DecodeString(data)
	return convert.BytesToString(res), err
}

// GetBase64RawStdEncode base64 无=填充编码
func GetBase64RawStdEncode(data string) string {
	return base64.RawStdEncoding.EncodeToString(convert.StringToBytes(data))
}

// GetBase64RawStdDecode base64 无=填充解码
func GetBase64RawStdDecode(data string) (string, error) {
	res, err := base64.RawStdEncoding.DecodeString(data)
	return convert.BytesToString(res), err
}
