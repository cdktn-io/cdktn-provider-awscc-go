// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResources struct {
	// Azure DevOps repository capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#azure_dev_ops_capabilities SecurityagentAgentSpace#azure_dev_ops_capabilities}
	AzureDevOpsCapabilities *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesAzureDevOpsCapabilities `field:"optional" json:"azureDevOpsCapabilities" yaml:"azureDevOpsCapabilities"`
	// Azure DevOps repository details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#azure_dev_ops_repository SecurityagentAgentSpace#azure_dev_ops_repository}
	AzureDevOpsRepository *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesAzureDevOpsRepository `field:"optional" json:"azureDevOpsRepository" yaml:"azureDevOpsRepository"`
	// Bitbucket repository capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#bitbucket_capabilities SecurityagentAgentSpace#bitbucket_capabilities}
	BitbucketCapabilities *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketCapabilities `field:"optional" json:"bitbucketCapabilities" yaml:"bitbucketCapabilities"`
	// Bitbucket repository details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#bitbucket_repository SecurityagentAgentSpace#bitbucket_repository}
	BitbucketRepository *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesBitbucketRepository `field:"optional" json:"bitbucketRepository" yaml:"bitbucketRepository"`
	// Confluence document capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#confluence_capabilities SecurityagentAgentSpace#confluence_capabilities}
	ConfluenceCapabilities *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilities `field:"optional" json:"confluenceCapabilities" yaml:"confluenceCapabilities"`
	// Confluence document details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#confluence_document SecurityagentAgentSpace#confluence_document}
	ConfluenceDocument *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocument `field:"optional" json:"confluenceDocument" yaml:"confluenceDocument"`
	// GitHub repository capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#git_hub_capabilities SecurityagentAgentSpace#git_hub_capabilities}
	GitHubCapabilities *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubCapabilities `field:"optional" json:"gitHubCapabilities" yaml:"gitHubCapabilities"`
	// GitHub repository details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#git_hub_repository SecurityagentAgentSpace#git_hub_repository}
	GitHubRepository *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitHubRepository `field:"optional" json:"gitHubRepository" yaml:"gitHubRepository"`
	// GitLab repository capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#git_lab_capabilities SecurityagentAgentSpace#git_lab_capabilities}
	GitLabCapabilities *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabCapabilities `field:"optional" json:"gitLabCapabilities" yaml:"gitLabCapabilities"`
	// GitLab repository details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_agent_space#git_lab_repository SecurityagentAgentSpace#git_lab_repository}
	GitLabRepository *SecurityagentAgentSpaceIntegratedResourcesProviderResourcesGitLabRepository `field:"optional" json:"gitLabRepository" yaml:"gitLabRepository"`
}

