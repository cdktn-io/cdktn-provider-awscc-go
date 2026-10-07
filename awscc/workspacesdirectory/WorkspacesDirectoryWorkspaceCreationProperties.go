// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryWorkspaceCreationProperties struct {
	// The identifier of the default security group to apply to WorkSpaces when they are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#custom_security_group_id WorkspacesDirectory#custom_security_group_id}
	CustomSecurityGroupId *string `field:"optional" json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// The organizational unit (OU) in the directory for the WorkSpace machine accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#default_ou WorkspacesDirectory#default_ou}
	DefaultOu *string `field:"optional" json:"defaultOu" yaml:"defaultOu"`
	// Specifies whether to automatically assign an Elastic public IP address to WorkSpaces in this directory by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#enable_internet_access WorkspacesDirectory#enable_internet_access}
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Specifies whether maintenance mode is enabled for WorkSpaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#enable_maintenance_mode WorkspacesDirectory#enable_maintenance_mode}
	EnableMaintenanceMode interface{} `field:"optional" json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Indicates the IAM role ARN of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#instance_iam_role_arn WorkspacesDirectory#instance_iam_role_arn}
	InstanceIamRoleArn *string `field:"optional" json:"instanceIamRoleArn" yaml:"instanceIamRoleArn"`
	// Specifies whether WorkSpace users are local administrators on their WorkSpaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#user_enabled_as_local_administrator WorkspacesDirectory#user_enabled_as_local_administrator}
	UserEnabledAsLocalAdministrator interface{} `field:"optional" json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}

