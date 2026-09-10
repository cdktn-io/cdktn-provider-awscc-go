// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeartifactpackagegroup


type CodeartifactPackageGroupOriginConfigurationRestrictionsExternalUpstream struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codeartifact_package_group#repositories CodeartifactPackageGroup#repositories}.
	Repositories *[]*string `field:"optional" json:"repositories" yaml:"repositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codeartifact_package_group#restriction_mode CodeartifactPackageGroup#restriction_mode}.
	RestrictionMode *string `field:"optional" json:"restrictionMode" yaml:"restrictionMode"`
}

