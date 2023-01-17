package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/bujuhu/lego-ui/dnsViewModels/netcup"
	"github.com/bujuhu/lego-ui/services"
)

func CertificateForm(services services.SingletonServices) *fyne.Container {
	
	overhead := container.New(layout.NewVBoxLayout(), Title("Get Your TLS Certificates"))
	inputForm := netcup.PartialView(services)

	return container.New(layout.NewHBoxLayout(), LogBox(services), container.New(layout.NewVBoxLayout(), overhead, inputForm))
}