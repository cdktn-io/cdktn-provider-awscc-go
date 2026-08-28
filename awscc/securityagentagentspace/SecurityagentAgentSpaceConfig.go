// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentAgentSpaceConfig struct {
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
	// Name of the agent space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#aws_resources SecurityagentAgentSpace#aws_resources}
	AwsResources *SecurityagentAgentSpaceAwsResources `field:"optional" json:"awsResources" yaml:"awsResources"`
	// Details of code review settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#code_review_settings SecurityagentAgentSpace#code_review_settings}
	CodeReviewSettings *SecurityagentAgentSpaceCodeReviewSettings `field:"optional" json:"codeReviewSettings" yaml:"codeReviewSettings"`
	// Description of the agent space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#description SecurityagentAgentSpace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Integrated Resources configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#integrated_resources SecurityagentAgentSpace#integrated_resources}
	IntegratedResources interface{} `field:"optional" json:"integratedResources" yaml:"integratedResources"`
	// Identifier of the KMS key used to encrypt data.
	//
	// Can be a key ID, key ARN, alias name, or alias ARN. If not specified, an AWS managed key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#kms_key_id SecurityagentAgentSpace#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Tags for the agent space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#tags SecurityagentAgentSpace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// List of target domain identifiers registered with the agent space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#target_domain_ids SecurityagentAgentSpace#target_domain_ids}
	TargetDomainIds *[]*string `field:"optional" json:"targetDomainIds" yaml:"targetDomainIds"`
}

