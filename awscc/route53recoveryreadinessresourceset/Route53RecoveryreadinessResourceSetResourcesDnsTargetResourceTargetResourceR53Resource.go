// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource struct {
	// The DNS target domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_resource_set#domain_name Route53RecoveryreadinessResourceSet#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Resource Record set id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id Route53RecoveryreadinessResourceSet#record_set_id}
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
}

