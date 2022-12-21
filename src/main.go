package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/bujuhu/lego-ui/view"
	"github.com/bujuhu/lego-ui/config"
	"github.com/bujuhu/lego-ui/userLogService"
	"github.com/bujuhu/lego-ui/services"
	"net/http"
	"time"
)

var client *http.Client

var data = []string{"Cert1", "Cert2", "Cert3"}


//Entrypoint of the Application
func main() {
	var services = services.SingletonServices {
		UserConfig: config.Config {
			Email: "you@yours.com",
			Key:   config.NewPrivateKey(),
			TOSAgreed: true,
		},
		UserLog: userLogService.New(),
	}

	client = &http.Client{Timeout: 10 * time.Second}
	services.UserLog.Write("test")
	a := app.New()
	view.New(services, data).Show(a)
}
