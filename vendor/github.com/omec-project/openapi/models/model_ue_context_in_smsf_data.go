// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UeContextInSmsfData struct {
	SmsfInfo3GppAccess    *SmsfInfo `json:"smsfInfo3GppAccess,omitempty" yaml:"smsfInfo3GppAccess" bson:"smsfInfo3GppAccess" mapstructure:"SmsfInfo3GppAccess"`
	SmsfInfoNon3GppAccess *SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty" yaml:"smsfInfoNon3GppAccess" bson:"smsfInfoNon3GppAccess" mapstructure:"SmsfInfoNon3GppAccess"`
}
