package main

import (
	"fmt"

	"github.com/smirnova-daria/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
