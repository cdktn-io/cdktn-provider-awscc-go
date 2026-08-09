// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package uxcaccountcustomization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UxcAccountCustomizationConfig struct {
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
	// The color theme assigned to the account for visual identification in the AWS Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/uxc_account_customization#account_color UxcAccountCustomization#account_color}
	AccountColor *string `field:"optional" json:"accountColor" yaml:"accountColor"`
	// A list of AWS region identifiers visible to the account in the AWS Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/uxc_account_customization#visible_regions UxcAccountCustomization#visible_regions}
	VisibleRegions *[]*string `field:"optional" json:"visibleRegions" yaml:"visibleRegions"`
	// A list of AWS service identifiers visible to the account in the AWS Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/uxc_account_customization#visible_services UxcAccountCustomization#visible_services}
	VisibleServices *[]*string `field:"optional" json:"visibleServices" yaml:"visibleServices"`
}

