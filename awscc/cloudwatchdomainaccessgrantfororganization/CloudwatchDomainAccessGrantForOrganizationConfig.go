// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomainaccessgrantfororganization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchDomainAccessGrantForOrganizationConfig struct {
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
	// The ID of the organization domain the grant belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#domain_id CloudwatchDomainAccessGrantForOrganization#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// A name that identifies the access grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#name CloudwatchDomainAccessGrantForOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The permission granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#permission CloudwatchDomainAccessGrantForOrganization#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// The principal receiving the grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#principal CloudwatchDomainAccessGrantForOrganization#principal}
	Principal *CloudwatchDomainAccessGrantForOrganizationPrincipal `field:"required" json:"principal" yaml:"principal"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain_access_grant_for_organization#tags CloudwatchDomainAccessGrantForOrganization#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

