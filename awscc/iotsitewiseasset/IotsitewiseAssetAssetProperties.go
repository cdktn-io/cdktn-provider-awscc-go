// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseasset


type IotsitewiseAssetAssetProperties struct {
	// The property alias that identifies the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#alias IotsitewiseAsset#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// String-friendly customer provided external ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#external_id IotsitewiseAsset#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual UUID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#id IotsitewiseAsset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Customer provided ID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#logical_id IotsitewiseAsset#logical_id}
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
	// The MQTT notification state (ENABLED or DISABLED) for this asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#notification_state IotsitewiseAsset#notification_state}
	NotificationState *string `field:"optional" json:"notificationState" yaml:"notificationState"`
	// The unit of measure (such as Newtons or RPM) of the asset property.
	//
	// If you don't specify a value for this parameter, the service uses the value of the assetModelProperty in the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_asset#unit IotsitewiseAsset#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

