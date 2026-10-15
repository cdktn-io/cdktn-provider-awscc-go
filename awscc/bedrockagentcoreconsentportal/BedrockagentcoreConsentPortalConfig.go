// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreconsentportal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreConsentPortalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#execution_role_arn BedrockagentcoreConsentPortal#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#idp_config BedrockagentcoreConsentPortal#idp_config}.
	IdpConfig *BedrockagentcoreConsentPortalIdpConfig `field:"required" json:"idpConfig" yaml:"idpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#name BedrockagentcoreConsentPortal#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#sources BedrockagentcoreConsentPortal#sources}.
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#description BedrockagentcoreConsentPortal#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to assign to the consent portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#tags BedrockagentcoreConsentPortal#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

