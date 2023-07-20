package functions

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/purawaktra/semeru2-go/utils"
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateSalt(length int) string {
	instance := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[instance.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateSHA1Hash(data string) (string, error) {
	h := sha1.New()
	n, err := h.Write([]byte(data))
	if err != nil {
		utils.Error(err, "GenerateSHA1Hash", "")
		return "", err
	}
	if n < 1 {
		utils.Error(err, "GenerateSHA1Hash", "")
		return "", errors.New("cannot generate SHA1Hash")
	}
	hashBytes := h.Sum(nil)
	hash := fmt.Sprintf("%x", hashBytes)
	return hash, nil
}
