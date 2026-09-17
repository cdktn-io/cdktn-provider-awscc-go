// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryVisualReferences struct {
	// Canary run id to be used as base reference for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#base_canary_run_id SyntheticsCanary#base_canary_run_id}
	BaseCanaryRunId *string `field:"optional" json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// List of screenshots used as base reference for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#base_screenshots SyntheticsCanary#base_screenshots}
	BaseScreenshots interface{} `field:"optional" json:"baseScreenshots" yaml:"baseScreenshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#browser_type SyntheticsCanary#browser_type}.
	BrowserType *string `field:"optional" json:"browserType" yaml:"browserType"`
}

