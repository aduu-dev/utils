package dyndns

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	services2 "aduu.dev/utils/dns/dyndns/services"
)

func generateNextIP(prior net.IP) net.IP {
	if prior == nil {
		prior = net.IPv4(byte(rand.Intn(200)), byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)))
	}

	prior[0] = prior[0] + 1

	fmt.Println("Test-IP: " + prior.String())

	return prior
}

func doSkip() bool {
	return true
}

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	wantIP = generateNextIP(wantIP)
}

func myTeardownFunction() {

}

var wantIP net.IP

func TestService(t *testing.T) {
	t.SkipNow()

	t.Run("FlexDNS", func(t *testing.T) {
		testService(t, services2.FlexDNS)
	})

}

func testService(t *testing.T, service services2.Service) {
	for _, hostname := range service.Hostnames {
		testGood(t, service, hostname)
		time.Sleep(time.Second * 10)
		testBadAuth(t, service, hostname)
		time.Sleep(time.Second * 10)
		testBadDomain(t, service)
		time.Sleep(time.Second * 10)
		testNoHost(t, service)
	}
}

func testGood(t *testing.T, service services2.Service, hostname string) {
	ip, err := Service{service.UpdateAddress, service.Username, service.Password}.Update(hostname, wantIP)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(ip)
}

func testBadAuth(t *testing.T, service services2.Service, hostname string) {
	ip, err := Service{service.UpdateAddress, "bogus", service.Password}.Update(hostname, wantIP)
	if err != ErrAuth {
		t.Error(err)
	}
	t.Log(ip)
	ip, err = Service{service.UpdateAddress, service.Username, "bogus"}.Update(hostname, wantIP)
	if err != ErrAuth {
		t.Error(err)
	}
	t.Log(ip)
}

func testBadDomain(t *testing.T, service services2.Service) {
	ip, err := Service{service.UpdateAddress, service.Username, service.Password}.Update("bogus", wantIP)
	if err != ErrDomain {
		t.Error(err)
	}
	t.Log(ip)
}

func testNoHost(t *testing.T, service services2.Service) {
	ip, err := Service{service.UpdateAddress, service.Username, service.Password}.Update("bogus.com", wantIP)
	if err != ErrNoHost {
		t.Error(err)
	}
	t.Log(ip)
}
