package maketinyurl

import (
	"url-shortner/datalayer"
	randomstring "url-shortner/services/random-string"
)

type MakeTinyURLInterface interface {
	ConvertURLToTinyURL(url string) string
}

type MakeTinyURLService struct {
	RandomString randomstring.RandomStringInterface
	Datalayer    datalayer.TinyURLDatalayerInterfce
}

func NewMakeTinyURLInterface(urlMap map[string]string) MakeTinyURLInterface {
	return &MakeTinyURLService{
		Datalayer:    datalayer.NewDatalayerInterface(urlMap),
		RandomString: randomstring.NewRandomStringInterface(),
	}
}

func (s *MakeTinyURLService) ConvertURLToTinyURL(url string) (tinyURL string) {
	tinyURL = s.RandomString.GenerateRandomStringGivenLength(5)

	s.Datalayer.SaveTinyURL(url, tinyURL)

	return tinyURL
}
