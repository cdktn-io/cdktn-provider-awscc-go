// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsdeliverysource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsDeliverySourceConfig struct {
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
	// The unique name of the Log source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/logs_delivery_source#name LogsDeliverySource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A map of key-value pairs to configure the delivery source.
	//
	// Both keys and values must be between 1 and 255 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/logs_delivery_source#delivery_source_configuration LogsDeliverySource#delivery_source_configuration}
	DeliverySourceConfiguration *map[string]*string `field:"optional" json:"deliverySourceConfiguration" yaml:"deliverySourceConfiguration"`
	// The type of logs being delivered.
	//
	// Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/logs_delivery_source#log_type LogsDeliverySource#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// The ARN of the resource that will be sending the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/logs_delivery_source#resource_arn LogsDeliverySource#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The tags that have been assigned to this delivery source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/logs_delivery_source#tags LogsDeliverySource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

