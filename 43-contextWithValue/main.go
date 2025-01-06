package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "John")
	bookName(ctx)
}

func bookName(ctx context.Context) {
	nome := ctx.Value("name")
	fmt.Println(nome)
}
