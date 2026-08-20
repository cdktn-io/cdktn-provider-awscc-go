// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeartifactpackagegroup


type CodeartifactPackageGroupOriginConfiguration struct {
	// The origin configuration that is applied to the package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codeartifact_package_group#restrictions CodeartifactPackageGroup#restrictions}
	Restrictions *CodeartifactPackageGroupOriginConfigurationRestrictions `field:"optional" json:"restrictions" yaml:"restrictions"`
}

