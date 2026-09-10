// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagergrant

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerGrantConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#allowed_operations LicensemanagerGrant#allowed_operations}.
	AllowedOperations *[]*string `field:"optional" json:"allowedOperations" yaml:"allowedOperations"`
	// Name for the created Grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#grant_name LicensemanagerGrant#grant_name}
	GrantName *string `field:"optional" json:"grantName" yaml:"grantName"`
	// Home region for the created grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#home_region LicensemanagerGrant#home_region}
	HomeRegion *string `field:"optional" json:"homeRegion" yaml:"homeRegion"`
	// License Arn for the grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#license_arn LicensemanagerGrant#license_arn}
	LicenseArn *string `field:"optional" json:"licenseArn" yaml:"licenseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#principals LicensemanagerGrant#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#status LicensemanagerGrant#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/licensemanager_grant#tags LicensemanagerGrant#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

