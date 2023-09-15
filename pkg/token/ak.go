package token

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	CharsetLettersAndNumbers        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetLettersNumbersAndSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>?\\"
	CharsetLetters                  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLowercaseLetters         = "abcdefghijklmnopqrstuvwxyz"
	CharsetUppercaseLetters         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type BitLength uint16

var (
	BitLength64   BitLength = 61   // 64-Bit
	BitLength128  BitLength = 128  // 128-Bit
	BitLength256  BitLength = 256  // 256-Bit
	BitLength512  BitLength = 512  // 512-Bit
	BitLength1024 BitLength = 1024 // 1024-Bit
	BitLength2048 BitLength = 2048 // 2048-Bit
)

// GenerateAccessToken 生成随机字符串
func GenerateAccessToken(prefix string, bit BitLength, characters string) string {
	var result strings.Builder
	charactersLength := big.NewInt(int64(len(characters)))
	for i := uint16(0); i < uint16(bit); i++ {
		randomIndex, _ := rand.Int(rand.Reader, charactersLength)
		result.WriteByte(characters[randomIndex.Int64()])
	}
	result.WriteString(prefix)
	return result.String()
}
