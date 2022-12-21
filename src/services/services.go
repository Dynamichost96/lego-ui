package services

import(
	"github.com/bujuhu/lego-ui/userLogService"
	"github.com/bujuhu/lego-ui/config"
)

//The SingletonServices Struct is used, to simply pass long living, single instance services, everywhere where they are needed
type SingletonServices struct {
	UserLog userLogService.UserLogService
	UserConfig config.Config
}