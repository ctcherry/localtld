package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func main() {
	tld := os.Getenv("TLD")
	ipStr := os.Getenv("IP")
	listenAddrPort := os.Getenv("LISTEN")

	if ipStr == "" {
		ipStr = "127.0.0.1"
	}
	answerIP := net.ParseIP(ipStr)

	if answerIP == nil {
		log.Fatal("Unable to parse IP " + ipStr)
	}

	if listenAddrPort == "" {
		listenAddrPort = "127.0.0.1:10053"
	}

	oneAnswer := oneAnswerResolver{answerIP: answerIP}

	dns.Handle(tld+".", oneAnswer)

	server := &dns.Server{Addr: listenAddrPort, Net: "udp"}
	fmt.Println("Listening on " + listenAddrPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type oneAnswerResolver struct {
	answerIP net.IP
}

func (o oneAnswerResolver) ServeDNS(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)

	m.Extra = make([]dns.RR, 1)
	m.Extra[0] = &dns.A{
		Hdr: dns.RR_Header{
			Name:   m.Question[0].Name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    1800,
		},
		A: o.answerIP,
	}
	err := w.WriteMsg(m)
	if err != nil {
		log.Println("Error writing DNS response: " + err.Error())
	}
}
