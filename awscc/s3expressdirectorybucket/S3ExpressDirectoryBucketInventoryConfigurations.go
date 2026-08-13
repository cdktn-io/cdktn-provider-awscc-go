// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket


type S3ExpressDirectoryBucketInventoryConfigurations struct {
	// Specifies information about where to publish inventory reports for an Amazon S3 Express bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#destination S3ExpressDirectoryBucket#destination}
	Destination *S3ExpressDirectoryBucketInventoryConfigurationsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#enabled S3ExpressDirectoryBucket#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#id S3ExpressDirectoryBucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#included_object_versions S3ExpressDirectoryBucket#included_object_versions}
	IncludedObjectVersions *string `field:"optional" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Contains the optional fields that are included in the inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#optional_fields S3ExpressDirectoryBucket#optional_fields}
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// The prefix that is prepended to all inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#prefix S3ExpressDirectoryBucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies the schedule for generating inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3express_directory_bucket#schedule_frequency S3ExpressDirectoryBucket#schedule_frequency}
	ScheduleFrequency *string `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
}

