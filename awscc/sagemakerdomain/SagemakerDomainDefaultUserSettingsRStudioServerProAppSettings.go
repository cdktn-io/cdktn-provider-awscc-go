// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings struct {
	// Indicates whether the current user has access to the RStudioServerPro app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#access_status SagemakerDomain#access_status}
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// The level of permissions that the user has within the RStudioServerPro app.
	//
	// This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#user_group SagemakerDomain#user_group}
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}

