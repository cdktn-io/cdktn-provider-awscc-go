// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettings struct {
	// The execution role for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#execution_role SagemakerDomain#execution_role}
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Indicates whether auto-mounting of an EFS volume is supported for the user profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#auto_mount_home_efs SagemakerDomain#auto_mount_home_efs}
	AutoMountHomeEfs *string `field:"optional" json:"autoMountHomeEfs" yaml:"autoMountHomeEfs"`
	// The CodeEditor app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#code_editor_app_settings SagemakerDomain#code_editor_app_settings}
	CodeEditorAppSettings *SagemakerDomainDefaultUserSettingsCodeEditorAppSettings `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#custom_file_system_configs SagemakerDomain#custom_file_system_configs}.
	CustomFileSystemConfigs interface{} `field:"optional" json:"customFileSystemConfigs" yaml:"customFileSystemConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#custom_posix_user_config SagemakerDomain#custom_posix_user_config}.
	CustomPosixUserConfig *SagemakerDomainDefaultUserSettingsCustomPosixUserConfig `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Defines which Amazon SageMaker application users are directed to by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#default_landing_uri SagemakerDomain#default_landing_uri}
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// The JupyterLab app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings SagemakerDomain#jupyter_lab_app_settings}
	JupyterLabAppSettings *SagemakerDomainDefaultUserSettingsJupyterLabAppSettings `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// The Jupyter server's app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#jupyter_server_app_settings SagemakerDomain#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerDomainDefaultUserSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings SagemakerDomain#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// A collection of settings that apply to an RSessionGateway app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#r_session_app_settings SagemakerDomain#r_session_app_settings}
	RSessionAppSettings *SagemakerDomainDefaultUserSettingsRSessionAppSettings `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// A collection of settings that configure user interaction with the RStudioServerPro app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#r_studio_server_pro_app_settings SagemakerDomain#r_studio_server_pro_app_settings}
	RStudioServerProAppSettings *SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#security_groups SagemakerDomain#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The sharing settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#sharing_settings SagemakerDomain#sharing_settings}
	SharingSettings *SagemakerDomainDefaultUserSettingsSharingSettings `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// Default storage settings for a space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#space_storage_settings SagemakerDomain#space_storage_settings}
	SpaceStorageSettings *SagemakerDomainDefaultUserSettingsSpaceStorageSettings `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Indicates whether the Studio experience is available to users. If not, users cannot access Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#studio_web_portal SagemakerDomain#studio_web_portal}
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// Studio settings.
	//
	// If these settings are applied on a user level, they take priority over the settings applied on a domain level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_domain#studio_web_portal_settings SagemakerDomain#studio_web_portal_settings}
	StudioWebPortalSettings *SagemakerDomainDefaultUserSettingsStudioWebPortalSettings `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
}

