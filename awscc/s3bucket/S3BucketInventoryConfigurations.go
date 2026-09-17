// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketInventoryConfigurations struct {
	// Contains information about where to publish the inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketInventoryConfigurationsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// If set to ``True``, an inventory list is generated. If set to ``False``, no inventory list is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#enabled S3Bucket#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// If set to ``All``, the list includes all the object versions, which adds the version-related fields ``VersionId``, ``IsLatest``, and ``DeleteMarker`` to the list. If set to ``Current``, the list does not contain these version-related fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#included_object_versions S3Bucket#included_object_versions}
	IncludedObjectVersions *string `field:"optional" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Contains the optional fields that are included in the inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#optional_fields S3Bucket#optional_fields}
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the inventory filter prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies the schedule for generating inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#schedule_frequency S3Bucket#schedule_frequency}
	ScheduleFrequency *string `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
}

