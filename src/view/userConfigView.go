package view

import (
	"fyne.io/fyne/v2/layout"
	"image/color"
	"net/url"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bujuhu/lego-ui/services"
)

func UserConfigView(services services.SingletonServices) *fyne.Container {
	labelEmail := canvas.NewText("E-Mail Address", color.White)
	labelEmail.Alignment = fyne.TextAlignLeading

	inputEmail := widget.NewEntry()
	inputEmail.SetPlaceHolder("Enter text...")
	inputEmail.Text = services.UserConfig.GetEmail()
	inputEmail.OnChanged = func(value string) {
		log.Println("Email set to", value)
		services.UserConfig.SetEmail(value)
	}

	inputTOS := widget.NewCheck("I have read and agree to the Terms of Service above", func(value bool) {
		log.Println("Check set to", value)
		services.UserConfig.SetTOSAgree(value)
	})
	URL, _ := url.Parse("https://community.letsencrypt.org/tos")
	tosText := widget.NewHyperlink("LetsEncrypt Terms of Service", URL)

	inputTOS.Checked = services.UserConfig.GetTOSAgree()

	return container.New(layout.NewPaddedLayout(), container.New(layout.NewVBoxLayout(), Title("Settings"), labelEmail, inputEmail, tosText, inputTOS))
}