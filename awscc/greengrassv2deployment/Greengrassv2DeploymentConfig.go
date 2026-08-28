// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2DeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#target_arn Greengrassv2Deployment#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#components Greengrassv2Deployment#components}.
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#deployment_name Greengrassv2Deployment#deployment_name}.
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#deployment_policies Greengrassv2Deployment#deployment_policies}.
	DeploymentPolicies *Greengrassv2DeploymentDeploymentPolicies `field:"optional" json:"deploymentPolicies" yaml:"deploymentPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#iot_job_configuration Greengrassv2Deployment#iot_job_configuration}.
	IotJobConfiguration *Greengrassv2DeploymentIotJobConfiguration `field:"optional" json:"iotJobConfiguration" yaml:"iotJobConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#parent_target_arn Greengrassv2Deployment#parent_target_arn}.
	ParentTargetArn *string `field:"optional" json:"parentTargetArn" yaml:"parentTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#tags Greengrassv2Deployment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

