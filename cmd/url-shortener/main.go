package main

import (
	"fmt"

	"github.com/VadimRight/Go_WebApp/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
