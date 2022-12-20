package main

import (
	"encoding/json"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type randomFact struct {
	Text string `json:"text"`
}

func getRandomFact() (randomFact, error) {
	var fact randomFact
	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return randomFact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return randomFact{}, err
	}

	return fact, nil
}

var data = []string{"Cert1", "Cert2", "Cert3"}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("LEGO-UI")
	win.Resize(fyne.NewSize(1000, 1000))

	title := canvas.NewText("Get Your TLS Certificates", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	customerNumber := canvas.NewText("Customer Number", color.White)
	customerNumber.Alignment = fyne.TextAlignLeading

	inputCustomerNumber := widget.NewEntry()
	inputCustomerNumber.SetPlaceHolder("Enter text...")

	apiKey := canvas.NewText("API Key", color.White)
	apiKey.Alignment = fyne.TextAlignLeading

	inputAPIKey := widget.NewEntry()
	inputAPIKey.SetPlaceHolder("Enter text...")

	apiPass := canvas.NewText("Api Password", color.White)
	apiPass.Alignment = fyne.TextAlignLeading

	inputAPIPass := widget.NewPasswordEntry()
	inputAPIPass.SetPlaceHolder("Enter text...")

	logOut := canvas.NewText("Logging Output", color.White)
	logOut.Alignment = fyne.TextAlignLeading

	logOutBox := widget.NewMultiLineEntry()

	logOutBox.SetPlaceHolder("Hier sind die logs lines " + "\n" +
		"Lgoing Data Testoutput" + "\n" + "Lgoing Data Testoutput")

	logOutBox.Resize(fyne.Size{400, 1000})
	logOutBox.Move(fyne.Position{400, 0})

	getCertButton := widget.NewButton("GetCert", func() {
		logOutBox.SetText("Logging Logging Logging" + "\n" + "Logging Logging Logging" + "\n" + "Logging Logging Logging" + "\n" + "Logging Logging Logging" + "\n")
	})

	//TODO derzeit noch nicht ready liste aller certs für den angegebenen DNS
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	list.Resize(fyne.Size{200, 1000})
	list.Move(fyne.Position{0, 0})

	overhead := container.New(layout.NewVBoxLayout(), title)
	inputForm := container.New(layout.NewVBoxLayout(), overhead, customerNumber, inputCustomerNumber, apiKey, inputAPIKey, apiPass, inputAPIPass, getCertButton)
	logOutput := container.NewWithoutLayout(logOutBox)

	dns1Form := container.New(layout.NewHBoxLayout(), logOutput, inputForm)
	dns2Form := container.New(layout.NewHBoxLayout(), logOutput, inputForm)

	tabItem1 := container.NewTabItem("DNS1", dns1Form)
	tabItem2 := container.NewTabItem("DNS2", dns2Form)
	tabs := container.NewAppTabs(tabItem1, tabItem2)

	win.SetContent(tabs)
	win.ShowAndRun()
}
