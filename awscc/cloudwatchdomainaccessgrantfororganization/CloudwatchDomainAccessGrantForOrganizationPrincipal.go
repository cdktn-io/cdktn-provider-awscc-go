// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomainaccessgrantfororganization


type CloudwatchDomainAccessGrantForOrganizationPrincipal struct {
	// The type of principal the grant is issued to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#principal_type CloudwatchDomainAccessGrantForOrganization#principal_type}
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// Attribute key-value pairs that select the principals the grant applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#principal_attributes CloudwatchDomainAccessGrantForOrganization#principal_attributes}
	PrincipalAttributes interface{} `field:"optional" json:"principalAttributes" yaml:"principalAttributes"`
	// The identifier of the principal, such as an IAM role ARN or an Identity Center user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#principal_id CloudwatchDomainAccessGrantForOrganization#principal_id}
	PrincipalId *string `field:"optional" json:"principalId" yaml:"principalId"`
}

