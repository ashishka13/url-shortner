package datalayer

type Datalayer struct {
	TinyURLDatabase TinyURLDatalayerInterfce
}

func InitDatalayer(urlMap map[string]string) Datalayer {
	d := Datalayer{
		TinyURLDatabase: NewDatalayerInterface(urlMap),
	}

	return d
}
