// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdanetworkconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaNetworkConnectorConfig struct {
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
	// The network configuration for the connector. Specify a VpcEgressConfiguration to enable outbound traffic routing through your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_network_connector#configuration LambdaNetworkConnector#configuration}
	Configuration *LambdaNetworkConnectorConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// A unique name for the network connector within your account and Region.
	//
	// Must be 1 to 64 alphanumeric characters, hyphens, or underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_network_connector#name LambdaNetworkConnector#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the IAM role that Lambda assumes to manage elastic network interfaces in your VPC.
	//
	// This role must have permissions for ec2:CreateNetworkInterface and related describe operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_network_connector#operator_role LambdaNetworkConnector#operator_role}
	OperatorRole *string `field:"optional" json:"operatorRole" yaml:"operatorRole"`
	// A list of tags to apply to the network connector.
	//
	// Use tags to categorize network connectors for cost allocation, access control, or operational management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_network_connector#tags LambdaNetworkConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

