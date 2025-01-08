package randomstring

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type RandomStringInterface interface {
	GenerateRandomStringGivenLength(length int) string
}

type RandomStringService struct{}

func NewRandomStringInterface() RandomStringInterface {
	return &RandomStringService{}
}

func (s *RandomStringService) GenerateRandomStringGivenLength(length int) (randomString string) {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randGen.Intn(len(charset))]
	}
	return string(b)
}
