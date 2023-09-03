// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type GbrQosFlowInformation struct {
	MaxFbrDl            string              `json:"maxFbrDl"`
	MaxFbrUl            string              `json:"maxFbrUl"`
	GuaFbrDl            string              `json:"guaFbrDl"`
	GuaFbrUl            string              `json:"guaFbrUl"`
	NotifControl        NotificationControl `json:"notifControl,omitempty"`
	MaxPacketLossRateDl int32               `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl int32               `json:"maxPacketLossRateUl,omitempty"`
}
