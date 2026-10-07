// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesDirectoryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Information about the Active Directory config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#active_directory_config WorkspacesDirectory#active_directory_config}
	ActiveDirectoryConfig *WorkspacesDirectoryActiveDirectoryConfig `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// Describes the properties of the certificate-based authentication you want to use with your WorkSpaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#certificate_based_auth_properties WorkspacesDirectory#certificate_based_auth_properties}
	CertificateBasedAuthProperties *WorkspacesDirectoryCertificateBasedAuthProperties `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
	// Indicates whether self-service capabilities are enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#enable_self_service WorkspacesDirectory#enable_self_service}
	EnableSelfService interface{} `field:"optional" json:"enableSelfService" yaml:"enableSelfService"`
	// Endpoint encryption mode that allows you to configure the specified directory between Standard TLS and FIPS 140-2 validated mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#endpoint_encryption_mode WorkspacesDirectory#endpoint_encryption_mode}
	EndpointEncryptionMode *string `field:"optional" json:"endpointEncryptionMode" yaml:"endpointEncryptionMode"`
	// The Amazon Resource Name (ARN) of the identity center instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#idc_instance_arn WorkspacesDirectory#idc_instance_arn}
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The identifiers of the IP access control groups associated with the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#ip_group_ids WorkspacesDirectory#ip_group_ids}
	IpGroupIds *[]*string `field:"optional" json:"ipGroupIds" yaml:"ipGroupIds"`
	// Specifies the configurations of the Microsoft Entra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#microsoft_entra_config WorkspacesDirectory#microsoft_entra_config}
	MicrosoftEntraConfig *WorkspacesDirectoryMicrosoftEntraConfig `field:"optional" json:"microsoftEntraConfig" yaml:"microsoftEntraConfig"`
	// Describes the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#saml_properties WorkspacesDirectory#saml_properties}
	SamlProperties *WorkspacesDirectorySamlProperties `field:"optional" json:"samlProperties" yaml:"samlProperties"`
	// Describes the self-service permissions for a directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#selfservice_permissions WorkspacesDirectory#selfservice_permissions}
	SelfservicePermissions *WorkspacesDirectorySelfservicePermissions `field:"optional" json:"selfservicePermissions" yaml:"selfservicePermissions"`
	// Describes the streaming properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#streaming_properties WorkspacesDirectory#streaming_properties}
	StreamingProperties *WorkspacesDirectoryStreamingProperties `field:"optional" json:"streamingProperties" yaml:"streamingProperties"`
	// The identifiers of the subnets for your virtual private cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#subnet_ids WorkspacesDirectory#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The tags associated with the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#tags WorkspacesDirectory#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether your WorkSpace directory is dedicated or shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#tenancy WorkspacesDirectory#tenancy}
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The type of identity management the user is using.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#user_identity_type WorkspacesDirectory#user_identity_type}
	UserIdentityType *string `field:"optional" json:"userIdentityType" yaml:"userIdentityType"`
	// The device types and operating systems that can be used to access a WorkSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#workspace_access_properties WorkspacesDirectory#workspace_access_properties}
	WorkspaceAccessProperties *WorkspacesDirectoryWorkspaceAccessProperties `field:"optional" json:"workspaceAccessProperties" yaml:"workspaceAccessProperties"`
	// The default values that are used to create WorkSpaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#workspace_creation_properties WorkspacesDirectory#workspace_creation_properties}
	WorkspaceCreationProperties *WorkspacesDirectoryWorkspaceCreationProperties `field:"optional" json:"workspaceCreationProperties" yaml:"workspaceCreationProperties"`
	// Description of the directory to register.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#workspace_directory_description WorkspacesDirectory#workspace_directory_description}
	WorkspaceDirectoryDescription *string `field:"optional" json:"workspaceDirectoryDescription" yaml:"workspaceDirectoryDescription"`
	// The name of the directory to register.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#workspace_directory_name WorkspacesDirectory#workspace_directory_name}
	WorkspaceDirectoryName *string `field:"optional" json:"workspaceDirectoryName" yaml:"workspaceDirectoryName"`
	// Indicates whether the directory's WorkSpace type is personal or pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#workspace_type WorkspacesDirectory#workspace_type}
	WorkspaceType *string `field:"optional" json:"workspaceType" yaml:"workspaceType"`
}

