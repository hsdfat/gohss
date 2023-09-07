package hss_models

import "gohss/modules/models"

type SubscriberStore interface {
	Init() error

	GetAuthenSubscriptionData(id string) (*models.AuthenticationSubscription, error)

	UpdateAuthenSubscriptionData(id string, subs *models.AuthenticationSubscription) error

	GetSMSubscriptionData(id string) (*models.SessionManagementSubscriptionData, error)

	GetAMSubscriptionData(id string) (*models.AccessAndMobilitySubscriptionData, error)

	GetLTESubscriptionData(id string) (interface{}, error)

	GetSIPServerName(publicId string) (string, error)

	GetSIPImsi(publicId string) (string, error)

	GetSIPPrivateId(publicId string) (string, error)

	UpdateSIPServerName(publicId string, servername string) error

	GetSIPState(publicId string) (int, error)

	UpdateSIPState(publicId string, state int) error
}
