package main

import (
	"beginner/handler"
	"beginner/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api/v1")
	h := handler.New()
	h.Register(v1)
	r.Logger.Fatal(r.Start(":3000"))
}
