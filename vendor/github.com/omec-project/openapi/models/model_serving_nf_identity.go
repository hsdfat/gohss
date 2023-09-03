// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ServingNfIdentity struct {
	ServNfInstId string       `json:"servNfInstId,omitempty" yaml:"servNfInstId" bson:"servNfInstId" mapstructure:"ServNfInstId"`
	Guami        *Guami       `json:"guami,omitempty" yaml:"guami" bson:"guami" mapstructure:"Guami"`
	AnGwAddr     *AnGwAddress `json:"anGwAddr,omitempty" yaml:"anGwAddr" bson:"anGwAddr" mapstructure:"AnGwAddr"`
}