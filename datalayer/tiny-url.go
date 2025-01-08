package datalayer

import "log"

type TinyURLDatalayerInterfce interface {
	GetOriginalURL(url string) string
	SaveTinyURL(url, tinyURL string)
}

type TinyURLDatalayer struct {
	GlobalURLMap map[string]string
}

func NewDatalayerInterface(urlMap map[string]string) TinyURLDatalayerInterfce {
	return &TinyURLDatalayer{GlobalURLMap: urlMap}
}

func (d *TinyURLDatalayer) GetOriginalURL(url string) (tinyURL string) {
	tinyURL = d.GlobalURLMap[url]
	return tinyURL
}

func (d *TinyURLDatalayer) SaveTinyURL(url, tinyURL string) {
	log.Println("Tiny URL saved")
	d.GlobalURLMap[url] = tinyURL
}
