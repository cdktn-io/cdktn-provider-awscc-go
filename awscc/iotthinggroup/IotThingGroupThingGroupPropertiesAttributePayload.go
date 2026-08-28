// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotthinggroup


type IotThingGroupThingGroupPropertiesAttributePayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_thing_group#attributes IotThingGroup#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

