// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchIntegrationConfig struct {
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
	// The type of integration. Immutable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#integration_type CloudwatchIntegration#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The customer-supplied name of the integration. Unique per account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#name CloudwatchIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Write-only API-key credential material for the integration. Stored encrypted and never returned on read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#credential CloudwatchIntegration#credential}
	Credential *CloudwatchIntegrationCredential `field:"optional" json:"credential" yaml:"credential"`
	// Type-specific configuration attributes (max 50 entries).
	//
	// EXTERNAL_AGENT requires 'catalogId'. On update, supplied attributes merge into the existing map; omitting a key does not remove it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#integration_attributes CloudwatchIntegration#integration_attributes}
	IntegrationAttributes *map[string]*string `field:"optional" json:"integrationAttributes" yaml:"integrationAttributes"`
	// IAM role ARN the integration assumes.
	//
	// Required for AWS_INTEGRATION (a valid, service-assumable IAM role ARN). Not used by AWS_CONFIG_SLREC or EXTERNAL_AGENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#role_arn CloudwatchIntegration#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags associated with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#tags CloudwatchIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

