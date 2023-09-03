// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type GlobalRanNodeId struct {
	PlmnId  *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	N3IwfId string  `json:"n3IwfId,omitempty" yaml:"n3IwfId" bson:"n3IwfId"`
	GNbId   *GNbId  `json:"gNbId,omitempty" yaml:"gNbId" bson:"gNbId"`
	NgeNbId string  `json:"ngeNbId,omitempty" yaml:"ngeNbId" bson:"ngeNbId"`
}
