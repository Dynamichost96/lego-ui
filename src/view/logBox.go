package view

import (
	"io"
	"fmt"
	"bufio"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bujuhu/lego-ui/services"
)

func LogBox(services services.SingletonServices) *fyne.Container {
	logOut := canvas.NewText("Logging Output", color.White)
	logOut.Alignment = fyne.TextAlignLeading

	logOutBox := widget.NewMultiLineEntry()

	//Setting up Log
	reader, writer := io.Pipe()
	services.UserLog.Writer = writer
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
	
	logOutput := container.NewWithoutLayout(logOutBox)
	return logOutput
}