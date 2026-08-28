// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubstandard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubStandardConfig struct {
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
	// The ARN of the standard that you want to enable.
	//
	// To view a list of available ASH standards and their ARNs, use the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_standard#standards_arn SecurityhubStandard#standards_arn}
	StandardsArn *string `field:"required" json:"standardsArn" yaml:"standardsArn"`
	// Specifies which controls are to be disabled in a standard.   *Maximum*: ``100``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_standard#disabled_standards_controls SecurityhubStandard#disabled_standards_controls}
	DisabledStandardsControls interface{} `field:"optional" json:"disabledStandardsControls" yaml:"disabledStandardsControls"`
}

