package pokeapi

import (
	"net/http"
	"time"

	"github.com/MimiValsi/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokeapi.Cache
}

func NewClient(timeout time.Duration, cacheTimeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokeapi.NewCache(cacheTimeout),
	}
}
