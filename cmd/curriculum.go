package main

import (
	"fmt"

	"github.com/gbrancods/go-curriculum-web/router"
)

func main() {
	if err := router.HandleRequest(); err != nil {
		fmt.Printf("Error lintening: %v", err)
	}
}
