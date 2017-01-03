package util

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

const (
	// WhoisDomain provide domain to be used for whois
	WhoisDomain = "whois-servers.net"
	// WhoisPort provide port to use
	WhoisPort = "43"
)

// Whois lookup domain info
func Whois(domain string, servers ...string) (result string, err error) {
	result, err = query(domain, servers...)
	if err != nil {
		return
	}

	start := strings.Index(result, "Whois Server:")
	if start == -1 {
		return
	}

	start += 13
	end := strings.Index(result[start:], "\n")
	server := strings.Trim(strings.Replace(result[start:start+end], "\r", "", -1), " ")
	tmpResult, err := query(domain, server)
	if err != nil {
		return
	}

	result += tmpResult

	return
}

func query(domain string, servers ...string) (result string, err error) {
	var server string
	if len(servers) == 0 || servers[0] == "" {
		domains := strings.SplitN(domain, ".", 2)
		if len(domains) != 2 {
			err = fmt.Errorf("Domain %s is invalid.", domain)
			return
		}
		server = domains[1] + "." + WhoisDomain

		check := strings.SplitN(domain, ".", 2)
		if len(check) != 0 {
			if check[len(check)-1] == "id" {
				server = "whois.id"
			}
		}
	} else {
		server = servers[0]
	}
	check := strings.SplitN(domain, ".", 3)
	if len(check) != 0 {
		if check[len(check)-1] == "id" {
			server = "whois.id"
		} else {
			server = fmt.Sprintf("%s.%s", check[len(check)-1], WhoisDomain)
		}
	}
	fmt.Printf("Using %s for domain: %s\n", server, domain)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, WhoisPort), time.Second*30)
	if err != nil {
		fmt.Printf("Error dialing whois server: %s\n", err.Error())
		return
	}

	conn.Write([]byte(domain + "\r\n"))
	var buffer []byte
	buffer, err = ioutil.ReadAll(conn)
	if err != nil {
		return
	}

	conn.Close()
	result = string(buffer[:])

	return
}
