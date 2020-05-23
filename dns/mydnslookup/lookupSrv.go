package mydnslookup

import (
	"bytes"
	"context"
	"net"
	"text/template"

	"github.com/miekg/dns"
)

var fqdnTemplate = template.Must(template.New("").Parse("_{{.Service}}._{{.Proto}}.{{.KubernetesService}}.{{.SearchPath}}."))

type fqdnSettings struct {
	Service           string
	Proto             string
	KubernetesService string
	SearchPath        string
}

func LookupSRVMiekgGrpc(context context.Context, service, proto, name string) (cname string, addrs []*net.SRV, err error) {
	return LookupSRVMiekg(service, proto, name)
}

func LookupSRVMiekg(service, proto, name string) (cname string, addrs []*net.SRV, err error) {
	config, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Servers: %#v\n", config.Servers)
	c := new(dns.Client)
	m := new(dns.Msg)

	settings := fqdnSettings{
		Service:           service,
		Proto:             proto,
		KubernetesService: name,
		SearchPath:        config.Search[0],
	}

	writer := new(bytes.Buffer)
	if err := fqdnTemplate.Execute(writer, settings); err != nil {
		panic(err)
	}
	usedAddress := writer.String()

	//fmt.Println("UsedAddress: " + usedAddress)
	m.SetQuestion(usedAddress, dns.TypeSRV)
	//m.SetQuestion(usedAddress, dns.TypeCNAME)

	m.RecursionDesired = true
	r, _, err := c.Exchange(m, config.Servers[0]+":53")
	//r, _, err := c.Exchange(m, "127.0.0.1:53")
	if err != nil {
		return
	}
	if r.Rcode != dns.RcodeSuccess {
		return
	}

	for _, a := range r.Answer {
		srv, ok := a.(*dns.SRV)
		if !ok {
			cnme, ok := a.(*dns.CNAME)
			if !ok {
				panic("record is not srv and not cname")
			}
			cname = cnme.Target
		}

		addrs = append(addrs, &net.SRV{
			Target:   srv.Target,
			Port:     srv.Port,
			Priority: srv.Priority,
			Weight:   srv.Weight,
		})
	}
	return
}
