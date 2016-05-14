package main

import (
	"net/http"
	"fmt"
	"github.com/jaem/nimble"
	"github.com/jaem/kmux"
)

func main() {
	mux := kmux.New()
	mux.GET("/hello/:q/watch", hello)
	mux.GET("/helloinline", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello inline!")
	})

	n := nimble.Default()
	n.UseHandlerFunc(middlewareA)
	n.UseHandlerFunc(middlewareB)
	n.Use(mux)
	n.Run(":3000")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellxxo!")
	ps := kmux.GetHttpParams(r)

	fmt.Println("...." + ps.ByName("q"))

	//bun := hax.GetBundle(c)
	//
	//if value := bun.Get("valueA"); value != nil {
	//	logger.Printy("from helloHandlerFunc, valueA is " + value.(string))
	//}
	//if value := bun.Get("valueB"); value != nil {
	//	logger.Printy("from helloHandlerFunc, valueB is " + value.(string))
	//}
}

func middlewareA(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("[nim.] I am middlewareA")
	//bun := hax.GetBundle(c)
	//bun.Set("valueA", ": from middlewareA")
	next(w, r)
}

func middlewareB(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("[nim.] I am middlewareB")
	//bun := hax.GetBundle(c)
	//bun.Set("valueB", ": from middlewareB")
	next(w, r)
}