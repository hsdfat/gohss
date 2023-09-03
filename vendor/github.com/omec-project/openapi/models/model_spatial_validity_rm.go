// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// this data type is defined in the same way as the SpatialValidity data type, but with the OpenAPI nullable property set to true
type SpatialValidityRm struct {
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList" yaml:"presenceInfoList" bson:"presenceInfoList" mapstructure:"PresenceInfoList"`
}