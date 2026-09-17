// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcloudconnector


type SsmCloudConnectorConfigurationAzureConfiguration struct {
	// The Azure AD application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#application_id SsmCloudConnector#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Azure AD tenant ID. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#tenant_id SsmCloudConnector#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The display name of the Azure AD application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#application_display_name SsmCloudConnector#application_display_name}
	ApplicationDisplayName *string `field:"optional" json:"applicationDisplayName" yaml:"applicationDisplayName"`
	// The targets for the cloud connector. If omitted, the entire tenant is targeted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#targets SsmCloudConnector#targets}
	Targets *SsmCloudConnectorConfigurationAzureConfigurationTargets `field:"optional" json:"targets" yaml:"targets"`
	// The display name of the Azure AD tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssm_cloud_connector#tenant_display_name SsmCloudConnector#tenant_display_name}
	TenantDisplayName *string `field:"optional" json:"tenantDisplayName" yaml:"tenantDisplayName"`
}

