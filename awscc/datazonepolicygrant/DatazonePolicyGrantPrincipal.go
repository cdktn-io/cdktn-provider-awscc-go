// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#domain_unit DatazonePolicyGrant#domain_unit}.
	DomainUnit *DatazonePolicyGrantPrincipalDomainUnit `field:"optional" json:"domainUnit" yaml:"domainUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#group DatazonePolicyGrant#group}.
	Group *DatazonePolicyGrantPrincipalGroup `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#project DatazonePolicyGrant#project}.
	Project *DatazonePolicyGrantPrincipalProject `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_policy_grant#user DatazonePolicyGrant#user}.
	User *DatazonePolicyGrantPrincipalUser `field:"optional" json:"user" yaml:"user"`
}

