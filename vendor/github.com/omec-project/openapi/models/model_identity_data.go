// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type IdentityData struct {
	SupiList []string `json:"supiList,omitempty" bson:"supiList"`
	GpsiList []string `json:"gpsiList,omitempty" bson:"gpsiList"`
}
