// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryBrowserConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/synthetics_canary#browser_type SyntheticsCanary#browser_type}.
	BrowserType *string `field:"optional" json:"browserType" yaml:"browserType"`
}

