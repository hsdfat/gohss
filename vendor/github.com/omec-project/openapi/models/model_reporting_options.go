// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ReportingOptions struct {
	MaxNumOfReports int32      `json:"maxNumOfReports,omitempty" yaml:"maxNumOfReports" bson:"maxNumOfReports" mapstructure:"MaxNumOfReports"`
	Expiry          *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
}
