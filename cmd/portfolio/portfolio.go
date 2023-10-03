package main

import (
	"fmt"

	"github.com/igab-dev/go-curriculum-web/pkg/router"
)

func main() {
	if err := router.HandleRequest(); err != nil {
		fmt.Printf("Error lintening: %v", err)
	}
}
