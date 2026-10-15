// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentintegration


type SecurityagentIntegrationProviderInput struct {
	// The input required to create a Bitbucket integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/securityagent_integration#bitbucket SecurityagentIntegration#bitbucket}
	Bitbucket *SecurityagentIntegrationProviderInputBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// The input required to create a Confluence integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/securityagent_integration#confluence SecurityagentIntegration#confluence}
	Confluence *SecurityagentIntegrationProviderInputConfluence `field:"optional" json:"confluence" yaml:"confluence"`
	// The input required to create a GitHub integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/securityagent_integration#github SecurityagentIntegration#github}
	Github *SecurityagentIntegrationProviderInputGithub `field:"optional" json:"github" yaml:"github"`
	// The input required to create a GitLab integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/securityagent_integration#gitlab SecurityagentIntegration#gitlab}
	Gitlab *SecurityagentIntegrationProviderInputGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
}

