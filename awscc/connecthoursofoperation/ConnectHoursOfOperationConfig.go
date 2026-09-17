// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectHoursOfOperationConfig struct {
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
	// Configuration information for the hours of operation: day, start time, and end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#config ConnectHoursOfOperation#config}
	Config interface{} `field:"required" json:"config" yaml:"config"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#instance_arn ConnectHoursOfOperation#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#name ConnectHoursOfOperation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#time_zone ConnectHoursOfOperation#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// List of child hours of operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#child_hours_of_operations ConnectHoursOfOperation#child_hours_of_operations}
	ChildHoursOfOperations interface{} `field:"optional" json:"childHoursOfOperations" yaml:"childHoursOfOperations"`
	// The description of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#description ConnectHoursOfOperation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more hours of operation overrides assigned to an hour of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#hours_of_operation_overrides ConnectHoursOfOperation#hours_of_operation_overrides}
	HoursOfOperationOverrides interface{} `field:"optional" json:"hoursOfOperationOverrides" yaml:"hoursOfOperationOverrides"`
	// List of parent hours of operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#parent_hours_of_operations ConnectHoursOfOperation#parent_hours_of_operations}
	ParentHoursOfOperations interface{} `field:"optional" json:"parentHoursOfOperations" yaml:"parentHoursOfOperations"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_hours_of_operation#tags ConnectHoursOfOperation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

