package util

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomInt64() int64 {
	return rand.Int63n(1000000000000000)
}

func RandomInt32() int32 {
	return rand.Int31n(1000000000)
}

func RandomDirection() string {
	in := []string{"RE", "DF", "OD"}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}

func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func RandomEmail(login string) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	email := login + "@" + string(b) + ".com"
	return email
}

func RandomPhone() sql.NullString {
	const chars = "123456789"
	b := make([]byte, 14)
	for i := range b {
		if i == 0 {
			b[i] = '('
		} else if i == 4 {
			b[i] = ')'
		} else if i == 5 {
			b[i] = ' '
		} else if i == 9 {
			b[i] = '-'
		} else {
			b[i] = chars[rand.Intn(len(chars))]
		}
	}
	return sql.NullString{String: string(b), Valid: true}
}

func RandomStringNull(length int) sql.NullString {
	return sql.NullString{String: RandomString(length), Valid: true}
}

func RandomMoney() float32 {
	return rand.Float32()
}

func RandomDate() time.Time {
	return time.Now().AddDate(0, 0, -1*int(RandomInt(0, 365)))
}

func RandomHashMD5(text string) string {
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
