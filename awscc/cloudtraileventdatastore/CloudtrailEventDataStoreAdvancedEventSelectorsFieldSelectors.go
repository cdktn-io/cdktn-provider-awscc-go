// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore


type CloudtrailEventDataStoreAdvancedEventSelectorsFieldSelectors struct {
	// An operator that includes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#ends_with CloudtrailEventDataStore#ends_with}
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of Field.
	//
	// This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#equals CloudtrailEventDataStore#equals}
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// A field in an event record on which to filter events to be logged.
	//
	// Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#field CloudtrailEventDataStore#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#not_ends_with CloudtrailEventDataStore#not_ends_with}
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#not_equals CloudtrailEventDataStore#not_equals}
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#not_starts_with CloudtrailEventDataStore#not_starts_with}
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudtrail_event_data_store#starts_with CloudtrailEventDataStore#starts_with}
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

