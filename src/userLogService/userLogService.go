package userLogService


//Simple struct that allows some log messages to be displayed to the user
type UserLogService struct {

	//Can be Called to Send Messages and Set to recieve them
	Write func(string)
}

func New() UserLogService {
	return UserLogService{ Write: func(msg string) {} }
}