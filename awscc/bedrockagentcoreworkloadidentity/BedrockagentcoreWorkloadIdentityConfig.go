// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreworkloadidentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreWorkloadIdentityConfig struct {
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
	// The name of the workload identity. The name must be unique within your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_workload_identity#name BedrockagentcoreWorkloadIdentity#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of allowed OAuth2 return URLs for resources associated with this workload identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_workload_identity#allowed_resource_oauth_2_return_urls BedrockagentcoreWorkloadIdentity#allowed_resource_oauth_2_return_urls}
	AllowedResourceOauth2ReturnUrls *[]*string `field:"optional" json:"allowedResourceOauth2ReturnUrls" yaml:"allowedResourceOauth2ReturnUrls"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_workload_identity#tags BedrockagentcoreWorkloadIdentity#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

