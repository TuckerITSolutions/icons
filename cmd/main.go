package main

import (
	"context"
	"os"

	"github.com/tuckeritsolutions/icons"
)

func main() {
	icon := icons.Close()

	err := icon.Render(context.Background(), os.Stdout)

	if err != nil {
		panic(err)
	}
}
