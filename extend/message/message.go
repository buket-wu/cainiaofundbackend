package message

type Drive interface {
	Send(acc, content string) error
}

func NewMessageDrive(driveType uint) Drive {
	return nil
}
