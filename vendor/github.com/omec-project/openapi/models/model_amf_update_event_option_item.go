// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type AmfUpdateEventOptionItem struct {
	Op    string     `json:"op" bson:"op" `
	Path  string     `json:"path" bson:"path" `
	Value *time.Time `json:"value" bson:"value" `
}
