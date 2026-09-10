// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipalDomainUnit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_policy_grant#domain_unit_designation DatazonePolicyGrant#domain_unit_designation}.
	DomainUnitDesignation *string `field:"optional" json:"domainUnitDesignation" yaml:"domainUnitDesignation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_policy_grant#domain_unit_grant_filter DatazonePolicyGrant#domain_unit_grant_filter}.
	DomainUnitGrantFilter *DatazonePolicyGrantPrincipalDomainUnitDomainUnitGrantFilter `field:"optional" json:"domainUnitGrantFilter" yaml:"domainUnitGrantFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_policy_grant#domain_unit_identifier DatazonePolicyGrant#domain_unit_identifier}.
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
}

