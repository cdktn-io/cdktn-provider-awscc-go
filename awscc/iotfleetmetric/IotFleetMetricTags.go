// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetmetric


type IotFleetMetricTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#key IotFleetMetric#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_fleet_metric#value IotFleetMetric#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

