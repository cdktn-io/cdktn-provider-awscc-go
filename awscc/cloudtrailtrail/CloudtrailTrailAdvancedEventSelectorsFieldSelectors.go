// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrailtrail


type CloudtrailTrailAdvancedEventSelectorsFieldSelectors struct {
	// An operator that includes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#ends_with CloudtrailTrail#ends_with}
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of Field.
	//
	// This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#equals CloudtrailTrail#equals}
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// A field in an event record on which to filter events to be logged.
	//
	// Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#field CloudtrailTrail#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#not_ends_with CloudtrailTrail#not_ends_with}
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#not_equals CloudtrailTrail#not_equals}
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#not_starts_with CloudtrailTrail#not_starts_with}
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudtrail_trail#starts_with CloudtrailTrail#starts_with}
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

