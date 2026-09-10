// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchserviceenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchServiceEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_service_environment#capacity_limits BatchServiceEnvironment#capacity_limits}.
	CapacityLimits interface{} `field:"required" json:"capacityLimits" yaml:"capacityLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_service_environment#service_environment_type BatchServiceEnvironment#service_environment_type}.
	ServiceEnvironmentType *string `field:"required" json:"serviceEnvironmentType" yaml:"serviceEnvironmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_service_environment#service_environment_name BatchServiceEnvironment#service_environment_name}.
	ServiceEnvironmentName *string `field:"optional" json:"serviceEnvironmentName" yaml:"serviceEnvironmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_service_environment#state BatchServiceEnvironment#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_service_environment#tags BatchServiceEnvironment#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

