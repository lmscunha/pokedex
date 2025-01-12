package pokeapi

import (
	"net/http"
	"time"

	"github.com/lmscunha/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	http  http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		http: http.Client{
			Timeout: timeout,
		},
	}
}
