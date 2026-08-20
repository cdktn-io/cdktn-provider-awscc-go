// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotthingtype


type IotThingTypeThingTypeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_thing_type#mqtt_5_configuration IotThingType#mqtt_5_configuration}.
	Mqtt5Configuration *IotThingTypeThingTypePropertiesMqtt5Configuration `field:"optional" json:"mqtt5Configuration" yaml:"mqtt5Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_thing_type#searchable_attributes IotThingType#searchable_attributes}.
	SearchableAttributes *[]*string `field:"optional" json:"searchableAttributes" yaml:"searchableAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_thing_type#thing_type_description IotThingType#thing_type_description}.
	ThingTypeDescription *string `field:"optional" json:"thingTypeDescription" yaml:"thingTypeDescription"`
}

