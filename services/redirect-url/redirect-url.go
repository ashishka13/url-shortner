package redirecturl

import "url-shortner/datalayer"

type RedirectURLInterface interface {
	GenerateRedirectURL(tinyURL string) string
}

type RedirectURLService struct {
	Datalayer datalayer.TinyURLDatalayerInterfce
}

func NewRedirectURLInterface(urlMap map[string]string) RedirectURLInterface {
	return &RedirectURLService{
		Datalayer: datalayer.NewDatalayerInterface(urlMap),
	}
}

func (s *RedirectURLService) GenerateRedirectURL(tinyURL string) (redirectURL string) {
	// originalURL := s.Datalayer.GetOriginalURL(tinyURL)

	return redirectURL
}
