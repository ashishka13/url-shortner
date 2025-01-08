package services

import (
	randomstring "url-shortner/services/random-string"
	maketinyurl "url-shortner/services/tiny-url"
)

type Services struct {
	RandomString randomstring.RandomStringInterface
	TinyURL      maketinyurl.MakeTinyURLInterface
}

func InitServices(urlMap map[string]string) Services {
	s := Services{
		RandomString: randomstring.NewRandomStringInterface(),
		TinyURL:      maketinyurl.NewMakeTinyURLInterface(urlMap),
	}

	return s
}
