package view

import (
	"io"
	"fmt"
	"bufio"
	"fyne.io/fyne/v2/layout"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bujuhu/lego-ui/dnsViewModels/netcup"
	"github.com/bujuhu/lego-ui/services"
)

//Graphical User Interface configuration and initalisation
type MainWindow struct {
	services services.SingletonServices
	data []string
}

func New(services services.SingletonServices, data []string) MainWindow {
	return MainWindow { services: services, data: data }
}

func (mw MainWindow) Show(a fyne.App) {
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

	logOut := canvas.NewText("Logging Output", color.White)
	logOut.Alignment = fyne.TextAlignLeading

	logOutBox := widget.NewMultiLineEntry()

	//Setting up Log
	reader, writer := io.Pipe()
	mw.services.UserLog.Writer = writer
	go func() {
		r := bufio.NewReaderSize(reader, 4*1024)
			line, isPrefix, err := r.ReadLine()
			for err == nil && !isPrefix {
				s := string(line)
				logOutBox.SetText(logOutBox.Text + s + "\n")
				line, isPrefix, err = r.ReadLine()
			}
			if isPrefix {
				fmt.Println("buffer size to small")
				return
			}
			if err != io.EOF {
				fmt.Println(err)
				return
			}
    }()

	logOutBox.Resize(fyne.Size{400, 1000})
	logOutBox.Move(fyne.Position{400, 0})

	//TODO derzeit noch nicht ready liste aller certs für den angegebenen DNS
	list := widget.NewList(
		func() int {
			return len(mw.data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(mw.data[i])
		})
	list.Resize(fyne.Size{200, 1000})
	list.Move(fyne.Position{0, 0})

	overhead := container.New(layout.NewVBoxLayout(), title)

	//Invoke partial View here to get DNS provider specific form
	inputForm := netcup.PartialView(mw.services)
	logOutput := container.NewWithoutLayout(logOutBox)

	dns1Form := container.New(layout.NewHBoxLayout(), logOutput, container.New(layout.NewVBoxLayout(), overhead, inputForm))
	dns2Form := container.New(layout.NewHBoxLayout(), logOutput, container.New(layout.NewVBoxLayout(), overhead, inputForm))

	tabItem1 := container.NewTabItem("DNS1", dns1Form)
	tabItem2 := container.NewTabItem("DNS2", dns2Form)
	tabs := container.NewAppTabs(tabItem1, tabItem2)

	win.SetContent(tabs)
	win.ShowAndRun()
}
