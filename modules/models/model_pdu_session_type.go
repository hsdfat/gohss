// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PduSessionType string

// List of PduSessionType
const (
	PduSessionType_IPV4         PduSessionType = "IPV4"
	PduSessionType_IPV6         PduSessionType = "IPV6"
	PduSessionType_IPV4_V6      PduSessionType = "IPV4V6"
	PduSessionType_UNSTRUCTURED PduSessionType = "UNSTRUCTURED"
	PduSessionType_ETHERNET     PduSessionType = "ETHERNET"
)