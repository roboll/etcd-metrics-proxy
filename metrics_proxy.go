package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type config struct {
	port               int
	upstreamHost       string
	upstreamPort       int
	upstreamServerName string
	etcdCA             string
	etcdCert           string
	etcdKey            string
}

func initFlags(c *config) {
	flag.IntVar(&c.port, "port", 2381, "Port to bind to.")
	flag.StringVar(&c.upstreamHost, "upstream-host", "localhost", "The upstream etcd host.")
	flag.IntVar(&c.upstreamPort, "upstream-port", 2379, "The upstream etcd port.")
	flag.StringVar(&c.upstreamServerName, "upstream-server-name", "localhost", "The upstream tls server name.")
	flag.StringVar(&c.etcdCA, "etcd-ca", "", "The CA file for etcd tls.")
	flag.StringVar(&c.etcdCert, "etcd-cert", "", "The cert file for etcd tls.")
	flag.StringVar(&c.etcdKey, "etcd-key", "", "The key file for etcd tls.")
}

func validateFlags(c *config) {
	if len(c.etcdCA) == 0 {
		log.Fatal("--etcd-ca=<ca-file> is required")
	}
	if len(c.etcdCert) == 0 {
		log.Fatal("--etcd-cert=<cert-file> is required")
	}
	if len(c.etcdKey) == 0 {
		log.Fatal("--etcd-key=<key-file> is required")
	}
}

func main() {
	c := config{}
	initFlags(&c)
	flag.Parse()
	validateFlags(&c)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%d", c.upstreamHost, c.upstreamPort),
	})

	pool := x509.NewCertPool()
	capem, err := ioutil.ReadFile(c.etcdCA)
	if err != nil {
		log.Fatal(err)
	}
	if !pool.AppendCertsFromPEM(capem) {
		log.Fatal("error: failed to add ca to cert pool")
	}

	cert, err := tls.LoadX509KeyPair(c.etcdCert, c.etcdKey)
	if err != nil {
		log.Fatal(err)
	}

	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cert},
			ServerName:   c.upstreamServerName,
		},
	}

	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		log.Printf("server: proxy metrics request to etcd")
		director(req)
	}

	server := http.NewServeMux()
	server.Handle("/metrics", proxy)

	addr := fmt.Sprintf(":%d", c.port)
	log.Printf("server: listening on %s\n", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatal(err)
	}
}
