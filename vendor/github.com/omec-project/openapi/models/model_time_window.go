// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to retrieve background data transfer policies from the PCF and to update the PCF with the background data transfer policy selected by the NF service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TimeWindow struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime string `json:"startTime" yaml:"startTime" bson:"startTime" mapstructure:"StartTime"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime string `json:"stopTime" yaml:"stopTime" bson:"stopTime" mapstructure:"StopTime"`
}