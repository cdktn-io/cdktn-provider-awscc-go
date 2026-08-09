// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedomain


type DatazoneDomainSingleSignOn struct {
	// The ARN of the AWS Identity Center instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_domain#idc_instance_arn DatazoneDomain#idc_instance_arn}
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The type of single sign-on in Amazon DataZone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_domain#type DatazoneDomain#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The single sign-on user assignment in Amazon DataZone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_domain#user_assignment DatazoneDomain#user_assignment}
	UserAssignment *string `field:"optional" json:"userAssignment" yaml:"userAssignment"`
}

