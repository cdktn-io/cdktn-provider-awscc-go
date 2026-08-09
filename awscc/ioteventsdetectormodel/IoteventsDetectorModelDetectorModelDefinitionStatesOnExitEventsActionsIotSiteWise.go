// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWise struct {
	// The ID of the asset that has the specified property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_detector_model#asset_id IoteventsDetectorModel#asset_id}
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_detector_model#entry_id IoteventsDetectorModel#entry_id}
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_detector_model#property_alias IoteventsDetectorModel#property_alias}
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_detector_model#property_id IoteventsDetectorModel#property_id}
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// The value to send to the asset property. This value contains timestamp, quality, and value (TQV) information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotevents_detector_model#property_value IoteventsDetectorModel#property_value}
	PropertyValue *IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEventsActionsIotSiteWisePropertyValue `field:"optional" json:"propertyValue" yaml:"propertyValue"`
}

