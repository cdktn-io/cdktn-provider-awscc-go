// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchDomainConfig struct {
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
	// The identity providers configured for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain#identity_providers CloudwatchDomain#identity_providers}
	IdentityProviders *[]*string `field:"required" json:"identityProviders" yaml:"identityProviders"`
	// A name that identifies the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain#name CloudwatchDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identity provider configuration for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain#identity_provider_configuration CloudwatchDomain#identity_provider_configuration}
	IdentityProviderConfiguration *CloudwatchDomainIdentityProviderConfiguration `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_domain#tags CloudwatchDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

