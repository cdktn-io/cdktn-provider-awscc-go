// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeartifactpackagegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodeartifactPackageGroupConfig struct {
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
	// The name of the domain that contains the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#domain_name CodeartifactPackageGroup#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The package group pattern that is used to gather packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#pattern CodeartifactPackageGroup#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The contact info of the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#contact_info CodeartifactPackageGroup#contact_info}
	ContactInfo *string `field:"optional" json:"contactInfo" yaml:"contactInfo"`
	// The text description of the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#description CodeartifactPackageGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The 12-digit account ID of the AWS account that owns the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#domain_owner CodeartifactPackageGroup#domain_owner}
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// The package origin configuration of the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#origin_configuration CodeartifactPackageGroup#origin_configuration}
	OriginConfiguration *CodeartifactPackageGroupOriginConfiguration `field:"optional" json:"originConfiguration" yaml:"originConfiguration"`
	// An array of key-value pairs to apply to the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#tags CodeartifactPackageGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

