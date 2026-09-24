// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightcustomization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightCustomizationConfig struct {
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
	// The ARN of the theme applied by default in the QuickSight console for this namespace.
	//
	// May be an AWS-managed starter theme such as arn:{Partition}:quicksight::aws:theme/MIDNIGHT or a theme owned by this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_customization#default_theme QuicksightCustomization#default_theme}
	DefaultTheme *string `field:"required" json:"defaultTheme" yaml:"defaultTheme"`
	// The QuickSight namespace the customization applies to.
	//
	// One customization exists per (account, region, namespace), so this is create-only: changing it addresses a different resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_customization#namespace QuicksightCustomization#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Tags applied to the customization.
	//
	// QuickSight rejects any key prefixed aws: or quicksight:, so CloudFormation system tags are not propagated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_customization#tags QuicksightCustomization#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

