// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendflywheel


type ComprehendFlywheelTaskConfigDocumentClassificationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_flywheel#labels ComprehendFlywheel#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_flywheel#mode ComprehendFlywheel#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

