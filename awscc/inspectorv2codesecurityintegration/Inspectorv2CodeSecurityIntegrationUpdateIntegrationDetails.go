// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2codesecurityintegration


type Inspectorv2CodeSecurityIntegrationUpdateIntegrationDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_code_security_integration#github Inspectorv2CodeSecurityIntegration#github}.
	Github *Inspectorv2CodeSecurityIntegrationUpdateIntegrationDetailsGithub `field:"optional" json:"github" yaml:"github"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_code_security_integration#gitlab_self_managed Inspectorv2CodeSecurityIntegration#gitlab_self_managed}.
	GitlabSelfManaged *Inspectorv2CodeSecurityIntegrationUpdateIntegrationDetailsGitlabSelfManaged `field:"optional" json:"gitlabSelfManaged" yaml:"gitlabSelfManaged"`
}

