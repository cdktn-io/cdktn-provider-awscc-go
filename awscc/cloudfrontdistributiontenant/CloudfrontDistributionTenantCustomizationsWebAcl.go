// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizationsWebAcl struct {
	// The action for the WAF web ACL customization.
	//
	// You can specify ``override`` to specify a separate WAF web ACL for the distribution tenant. If you specify ``disable``, the distribution tenant won't have WAF web ACL protections and won't inherit from the multi-tenant distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_distribution_tenant#action CloudfrontDistributionTenant#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The Amazon Resource Name (ARN) of the WAF web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_distribution_tenant#arn CloudfrontDistributionTenant#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

