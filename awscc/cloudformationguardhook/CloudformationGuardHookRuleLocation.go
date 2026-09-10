// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationguardhook


type CloudformationGuardHookRuleLocation struct {
	// S3 uri of Guard files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_guard_hook#uri CloudformationGuardHook#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// S3 object version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudformation_guard_hook#version_id CloudformationGuardHook#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

