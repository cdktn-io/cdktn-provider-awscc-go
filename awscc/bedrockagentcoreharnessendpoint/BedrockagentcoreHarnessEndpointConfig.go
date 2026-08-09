// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharnessendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessEndpointConfig struct {
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
	// The name of the endpoint. Must start with a letter and contain only alphanumeric characters and underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness_endpoint#endpoint_name BedrockagentcoreHarnessEndpoint#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The ID of the harness that the endpoint belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness_endpoint#harness_id BedrockagentcoreHarnessEndpoint#harness_id}
	HarnessId *string `field:"required" json:"harnessId" yaml:"harnessId"`
	// The description of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness_endpoint#description BedrockagentcoreHarnessEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to apply to the harness endpoint resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness_endpoint#tags BedrockagentcoreHarnessEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The harness version that the endpoint points to and serves invocations from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_harness_endpoint#target_version BedrockagentcoreHarnessEndpoint#target_version}
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

