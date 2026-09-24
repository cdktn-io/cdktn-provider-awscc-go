// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcorePolicyConfig struct {
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
	// The definition structure for policies. Encapsulates different policy formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#definition BedrockagentcorePolicy#definition}
	Definition *BedrockagentcorePolicyDefinition `field:"required" json:"definition" yaml:"definition"`
	// The customer-assigned immutable name for the policy. Must be unique within the policy engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#name BedrockagentcorePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the policy engine which contains this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#policy_engine_id BedrockagentcorePolicy#policy_engine_id}
	PolicyEngineId *string `field:"required" json:"policyEngineId" yaml:"policyEngineId"`
	// A human-readable description of the policy's purpose and functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#description BedrockagentcorePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the policy contributes to the enforce decision returned to Gateway.
	//
	// LOG_ONLY policies are still evaluated but their decisions are observed only, allowing customers to validate a policy against real traffic before promoting it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#enforcement_mode BedrockagentcorePolicy#enforcement_mode}
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
	// The validation mode for the policy. Determines how Cedar analyzer validation results are handled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_policy#validation_mode BedrockagentcorePolicy#validation_mode}
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

