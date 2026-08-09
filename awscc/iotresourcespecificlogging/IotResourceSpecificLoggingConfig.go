// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotresourcespecificlogging

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotResourceSpecificLoggingConfig struct {
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
	// The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_resource_specific_logging#log_level IotResourceSpecificLogging#log_level}
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The target name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_resource_specific_logging#target_name IotResourceSpecificLogging#target_name}
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID, or EVENT_TYPE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_resource_specific_logging#target_type IotResourceSpecificLogging#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

