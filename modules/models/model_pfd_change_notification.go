// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nnef_PFDmanagement Sevice API
 *
 * Packet Flow Description Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PfdChangeNotification struct {
	ApplicationId string       `json:"applicationId" yaml:"applicationId" bson:"applicationId" mapstructure:"ApplicationId"`
	RemovalFlag   bool         `json:"removalFlag,omitempty" yaml:"removalFlag" bson:"removalFlag" mapstructure:"RemovalFlag"`
	PartialFlag   bool         `json:"partialFlag,omitempty" yaml:"partialFlag" bson:"partialFlag" mapstructure:"PartialFlag"`
	Pfds          []PfdContent `json:"pfds,omitempty" yaml:"pfds" bson:"pfds" mapstructure:"Pfds"`
}