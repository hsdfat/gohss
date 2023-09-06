package hss_models

import "gohss/modules/models"

type SubscriberStore interface {
	Init() error

	GetAuthenSubscriptionData(id string) (*models.AuthenticationSubscription, error)

	UpdateAuthenSubscriptionData(id string, subs *models.AuthenticationSubscription) error

	GetSMSubscriptionData(id string) (*models.SessionManagementSubscriptionData, error)

	GetAMSubscriptionData(id string) (*models.AccessAndMobilitySubscriptionData, error)
}
