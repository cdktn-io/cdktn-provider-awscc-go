// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsdelivery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsDeliveryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ARN of the delivery destination that is associated with this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#delivery_destination_arn LogsDelivery#delivery_destination_arn}
	DeliveryDestinationArn *string `field:"required" json:"deliveryDestinationArn" yaml:"deliveryDestinationArn"`
	// The name of the delivery source that is associated with this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#delivery_source_name LogsDelivery#delivery_source_name}
	DeliverySourceName *string `field:"required" json:"deliverySourceName" yaml:"deliverySourceName"`
	// The field delimiter to use between record fields when the final output format of a delivery is in Plain , W3C , or Raw format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#field_delimiter LogsDelivery#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// The list of record fields to be delivered to the destination, in order.
	//
	// If the delivery's log source has mandatory fields, they must be included in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#record_fields LogsDelivery#record_fields}
	RecordFields *[]*string `field:"optional" json:"recordFields" yaml:"recordFields"`
	// This parameter causes the S3 objects that contain delivered logs to use a prefix structure that allows for integration with Apache Hive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#s3_enable_hive_compatible_path LogsDelivery#s3_enable_hive_compatible_path}
	S3EnableHiveCompatiblePath interface{} `field:"optional" json:"s3EnableHiveCompatiblePath" yaml:"s3EnableHiveCompatiblePath"`
	// This string allows re-configuring the S3 object prefix to contain either static or variable sections.
	//
	// The valid variables to use in the suffix path will vary by each log source. See ConfigurationTemplate$allowedSuffixPathFields for more info on what values are supported in the suffix path for each log source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#s3_suffix_path LogsDelivery#s3_suffix_path}
	S3SuffixPath *string `field:"optional" json:"s3SuffixPath" yaml:"s3SuffixPath"`
	// The tags that have been assigned to this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/logs_delivery#tags LogsDelivery#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

