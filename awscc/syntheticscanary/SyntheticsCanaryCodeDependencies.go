// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryCodeDependencies struct {
	// ARN of the Lambda layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#reference SyntheticsCanary#reference}
	Reference *string `field:"optional" json:"reference" yaml:"reference"`
	// Type of dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#type SyntheticsCanary#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

