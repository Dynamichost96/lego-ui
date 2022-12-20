package main

import (
	"encoding/json"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"log"
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
		log.Println("Content was:", inputCustomerNumber.Text)
	})

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
	//inBox := container.New(layout.NewCenterLayout(), layout.NewSpacer(), inputCustomerNumber, layout.NewSpacer(), saveButton, layout.NewSpacer())

	//hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	//vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), inBox, widget.NewSeparator(), factText)
	//test := container.NewContainerW

	testbox := container.New(layout.NewVBoxLayout(), title)
	testbox2 := container.New(layout.NewVBoxLayout(), testbox, customerNumber, inputCustomerNumber, apiKey, inputAPIKey, apiPass, inputAPIPass, getCertButton)
	testbox4 := container.NewWithoutLayout(logOutBox)
	//testbox7 := container.NewWithoutLayout(list)

	//testbox5 := container.New(layout.NewHBoxLayout(), testbox7, widget.NewSeparator(), testbox2, testbox4)
	testbox6 := container.New(layout.NewHBoxLayout(), testbox4, testbox2)
	testbox7 := container.New(layout.NewHBoxLayout(), testbox4, testbox2)
	tabItem1 := container.NewTabItem("DNS1", testbox6)
	tabItem2 := container.NewTabItem("DNS2", testbox7)
	tabs := container.NewAppTabs(tabItem1, tabItem2)
	win.SetContent(tabs)

	win.ShowAndRun()
}
