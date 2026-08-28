// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubstandard


type SecurityhubStandardDisabledStandardsControls struct {
	// A user-defined reason for changing a control's enablement status in a specified standard.
	//
	// If you are disabling a control, then this property is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_standard#reason SecurityhubStandard#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The Amazon Resource Name (ARN) of the control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_standard#standards_control_arn SecurityhubStandard#standards_control_arn}
	StandardsControlArn *string `field:"optional" json:"standardsControlArn" yaml:"standardsControlArn"`
}

