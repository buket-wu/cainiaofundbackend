package message

import "github.com/sirupsen/logrus"

const (
	DriveTypeWechatDefault = 0
)

type Drive interface {
	Send(acc, content string) error
}

func NewMessageDrive(driveType uint) Drive {
	return &DefaultDrive{Num: driveType}
}

// DefaultDrive 默认发消息驱动，只写一个log
type DefaultDrive struct {
	Num uint // 驱动编号
}

func (d *DefaultDrive) Send(acc, content string) error {
	logrus.Infof("send remind to user; acc:%v; content:%v", acc, content)

	return nil
}
