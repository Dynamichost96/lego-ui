package certificateService

import (
	"fmt"
	"log"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
	"github.com/bujuhu/lego-ui/services"
)

//CertificateService acts as an intermediate between View(Models) and the Lego Library
func Main(services services.SingletonServices, dnsProviderConfig *netcup.Config, domain string) {
	log.SetOutput(services.UserLog.Writer)
	config := lego.NewConfig(&services.UserConfig)

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Print(err.Error())
	}

	//Configure DNS Provider Settings
	provider, err := netcup.NewDNSProviderConfig(dnsProviderConfig)
	if err != nil {
		log.Print(err.Error())
	}
	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		log.Print(err.Error())
	}

	//Register User with LetsEncrpt
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: services.UserConfig.GetTOSAgree()})
	if err != nil {
		log.Print(err.Error())
	}
	services.UserConfig.Registration = reg

	// Request Certificates
	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Print(err.Error())
	}

	// TODO save results
	fmt.Printf("%#v\n", certificates)
}
