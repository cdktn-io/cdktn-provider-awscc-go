// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentComponentsRunWith struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#posix_user Greengrassv2Deployment#posix_user}.
	PosixUser *string `field:"optional" json:"posixUser" yaml:"posixUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#system_resource_limits Greengrassv2Deployment#system_resource_limits}.
	SystemResourceLimits *Greengrassv2DeploymentComponentsRunWithSystemResourceLimits `field:"optional" json:"systemResourceLimits" yaml:"systemResourceLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_deployment#windows_user Greengrassv2Deployment#windows_user}.
	WindowsUser *string `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

