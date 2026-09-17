// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotbillinggroup


type IotBillingGroupTags struct {
	// Tag key (1-128 chars). No 'aws:' prefix. Allows: [A-Za-z0-9 _.:/=+-].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_billing_group#key IotBillingGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag value (1-256 chars). No 'aws:' prefix. Allows: [A-Za-z0-9 _.:/=+-].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_billing_group#value IotBillingGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

