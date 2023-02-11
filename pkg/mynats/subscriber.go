package mynats

import (
	"WEB_REST_exm0302/pkg/service"
	"github.com/sirupsen/logrus"
)

type SubsNats struct {
	services *service.Service
}

func NewSubsNats(services *service.Service) *SubsNats {
	return &SubsNats{services: services}
}

func (sn *SubsNats) GetAndWriteJson() {
	//TODO добавить параметры вызова subscriber
	jsonFormNats, errSubscriber := Subscriber(sn)
	if errSubscriber != nil {

		logrus.Fatalf("Ошибка subscriber: %s", errSubscriber.Error())
	}
	logrus.Println(jsonFormNats)

	return
}

//func (sn *SubsNats) GetJson() WEB_REST_exm0302.Json {
//
//	return jsonFormNats
//}
