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

type Periodicity string

// List of Periodicity
const (
	Periodicity_YEARLY  Periodicity = "YEARLY"
	Periodicity_MONTHLY Periodicity = "MONTHLY"
	Periodicity_WEEKLY  Periodicity = "WEEKLY"
	Periodicity_DAILY   Periodicity = "DAILY"
	Periodicity_HOURLY  Periodicity = "HOURLY"
)