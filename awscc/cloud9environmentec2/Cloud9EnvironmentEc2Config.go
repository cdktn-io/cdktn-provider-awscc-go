// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloud9environmentec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Cloud9EnvironmentEc2Config struct {
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
	// The number of minutes until the running instance is shut down after the environment was last used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#automatic_stop_time_minutes Cloud9EnvironmentEc2#automatic_stop_time_minutes}
	AutomaticStopTimeMinutes *float64 `field:"optional" json:"automaticStopTimeMinutes" yaml:"automaticStopTimeMinutes"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#connection_type Cloud9EnvironmentEc2#connection_type}
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#description Cloud9EnvironmentEc2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#image_id Cloud9EnvironmentEc2#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The type of instance to connect to the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#instance_type Cloud9EnvironmentEc2#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#name Cloud9EnvironmentEc2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the environment owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#owner_arn Cloud9EnvironmentEc2#owner_arn}
	OwnerArn *string `field:"optional" json:"ownerArn" yaml:"ownerArn"`
	// Any AWS CodeCommit source code repositories to be cloned into the development environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#repositories Cloud9EnvironmentEc2#repositories}
	Repositories interface{} `field:"optional" json:"repositories" yaml:"repositories"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#subnet_id Cloud9EnvironmentEc2#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/cloud9_environment_ec2#tags Cloud9EnvironmentEc2#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

