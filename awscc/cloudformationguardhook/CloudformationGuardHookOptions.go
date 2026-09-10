// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookOptions struct {
	// Specifies the S3 location of input parameter files for your Guard rules.
	//
	// You can specify either a single S3 location or an array of up to 10 S3 locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_guard_hook#input_params CloudformationGuardHook#input_params}
	InputParams *string `field:"optional" json:"inputParams" yaml:"inputParams"`
}

