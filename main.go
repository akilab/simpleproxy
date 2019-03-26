package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

var (
	basicAuthUser = flag.String("user", "", "Basic User")
	basicAuthPass = flag.String("pass", "", "Basic Password")
	portNumber    = flag.String("port", "8080", "Proxy Port")
)

func main() {
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	if *basicAuthUser != "" && *basicAuthPass != "" {
		log.Println("auth basic")
		auth.ProxyBasic(proxy, "RELM", func(user, pass string) bool {
			log.Println("auth basic do")
			return user == *basicAuthUser && pass == *basicAuthPass
		})
	}
	log.Println("Listen: " + *portNumber)
	log.Fatal(http.ListenAndServe(":"+*portNumber, proxy))
}
