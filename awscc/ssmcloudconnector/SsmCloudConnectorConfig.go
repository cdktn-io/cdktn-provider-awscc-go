// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcloudconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmCloudConnectorConfig struct {
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
	// The ARN of the AWS Config connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#config_connector_arn SsmCloudConnector#config_connector_arn}
	ConfigConnectorArn *string `field:"required" json:"configConnectorArn" yaml:"configConnectorArn"`
	// The configuration for the cloud connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#configuration SsmCloudConnector#configuration}
	Configuration *SsmCloudConnectorConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The display name of the cloud connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#display_name SsmCloudConnector#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The IAM role ARN used by the cloud connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#role_arn SsmCloudConnector#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the cloud connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#description SsmCloudConnector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to apply to the cloud connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssm_cloud_connector#tags SsmCloudConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

