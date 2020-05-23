package services

// Popular dynamic DNS service URLs.
const (
	DNS_O_Matic = "https://update.dnsomatic.com/nic/update"
	//DynDNS      = "https://members.dyndns.org/nic/update"

	No_IP = "https://dynupdate.no-ip.com/nic/update"
)

type Service struct {
	UpdateAddress string
	Username      string
	Password      string
	Hostnames     []string
}

var (
	FlexDNS = Service{
		UpdateAddress: "https://ddns.do.de/",
		Username:      "DDNS-KD11963-F3086",
		Password:      "BIR2ZlHy875b",
		Hostnames:     []string{"aduu.dev"},
	}

	DynDNS = Service{
		UpdateAddress: "",
		Username:      "DDNS-KD11963-F3086",
		Password:      "BIR2ZlHy875b",
		Hostnames:     []string{"aduu.dev", "api.aduu.dev"},
	}

	DDNS = Service{
		UpdateAddress: "https://www.ddnss.de/upd.php?user=<username>&pwd=<pass>&host=<domain>",
		Username:      "DDNS-KD11963-F3086",
		Password:      "BIR2ZlHy875b",
		Hostnames:     []string{"aduu.dev"},
	}
)
