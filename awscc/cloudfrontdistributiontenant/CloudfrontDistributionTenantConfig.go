// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontDistributionTenantConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the multi-tenant distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#distribution_id CloudfrontDistributionTenant#distribution_id}
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// The domains associated with the distribution tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#domains CloudfrontDistributionTenant#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// The name of the distribution tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#name CloudfrontDistributionTenant#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the connection group for the distribution tenant.
	//
	// If you don't specify a connection group, CloudFront uses the default connection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#connection_group_id CloudfrontDistributionTenant#connection_group_id}
	ConnectionGroupId *string `field:"optional" json:"connectionGroupId" yaml:"connectionGroupId"`
	// Customizations for the distribution tenant.
	//
	// For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#customizations CloudfrontDistributionTenant#customizations}
	Customizations *CloudfrontDistributionTenantCustomizations `field:"optional" json:"customizations" yaml:"customizations"`
	// Indicates whether the distribution tenant is in an enabled state. If disabled, the distribution tenant won't serve traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#enabled CloudfrontDistributionTenant#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that represents the request for the Amazon CloudFront managed ACM certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#managed_certificate_request CloudfrontDistributionTenant#managed_certificate_request}
	ManagedCertificateRequest *CloudfrontDistributionTenantManagedCertificateRequest `field:"optional" json:"managedCertificateRequest" yaml:"managedCertificateRequest"`
	// A list of parameter values to add to the resource.
	//
	// A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#parameters CloudfrontDistributionTenant#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A complex type that contains zero or more ``Tag`` elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#tags CloudfrontDistributionTenant#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

