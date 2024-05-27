package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO : init logger slog
	// TODO : init storage sqlite
	// TODO : init logger chi
	// TODO : run server
}
