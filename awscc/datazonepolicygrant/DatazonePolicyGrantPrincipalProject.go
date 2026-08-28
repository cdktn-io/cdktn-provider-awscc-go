// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipalProject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#project_designation DatazonePolicyGrant#project_designation}.
	ProjectDesignation *string `field:"optional" json:"projectDesignation" yaml:"projectDesignation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#project_grant_filter DatazonePolicyGrant#project_grant_filter}.
	ProjectGrantFilter *DatazonePolicyGrantPrincipalProjectProjectGrantFilter `field:"optional" json:"projectGrantFilter" yaml:"projectGrantFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#project_identifier DatazonePolicyGrant#project_identifier}.
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
}

