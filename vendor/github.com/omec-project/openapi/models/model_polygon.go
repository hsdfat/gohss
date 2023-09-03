// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Polygon struct {
	Shape     SupportedGadShapes        `json:"shape" yaml:"shape" bson:"shape"`
	PointList []GeographicalCoordinates `json:"pointList" yaml:"pointList" bson:"pointList"`
}
