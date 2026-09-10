// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryVisualReferencesBaseScreenshots struct {
	// List of coordinates of rectangles to be ignored during visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/synthetics_canary#ignore_coordinates SyntheticsCanary#ignore_coordinates}
	IgnoreCoordinates *[]*string `field:"optional" json:"ignoreCoordinates" yaml:"ignoreCoordinates"`
	// Name of the screenshot to be used as base reference for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/synthetics_canary#screenshot_name SyntheticsCanary#screenshot_name}
	ScreenshotName *string `field:"optional" json:"screenshotName" yaml:"screenshotName"`
}

