// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GrafanaWorkspaceConfig struct {
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
	// These enums represent valid account access types.
	//
	// Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#account_access_type GrafanaWorkspace#account_access_type}
	AccountAccessType *string `field:"required" json:"accountAccessType" yaml:"accountAccessType"`
	// List of authentication providers to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#authentication_providers GrafanaWorkspace#authentication_providers}
	AuthenticationProviders *[]*string `field:"required" json:"authenticationProviders" yaml:"authenticationProviders"`
	// These enums represent valid permission types to use when creating or configuring a Grafana workspace.
	//
	// The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#permission_type GrafanaWorkspace#permission_type}
	PermissionType *string `field:"required" json:"permissionType" yaml:"permissionType"`
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#client_token GrafanaWorkspace#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// List of data sources on the service managed IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#data_sources GrafanaWorkspace#data_sources}
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Description of a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#description GrafanaWorkspace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of Grafana to support in your workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#grafana_version GrafanaWorkspace#grafana_version}
	GrafanaVersion *string `field:"optional" json:"grafanaVersion" yaml:"grafanaVersion"`
	// The user friendly name of a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#name GrafanaWorkspace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration settings for Network Access Control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#network_access_control GrafanaWorkspace#network_access_control}
	NetworkAccessControl *GrafanaWorkspaceNetworkAccessControl `field:"optional" json:"networkAccessControl" yaml:"networkAccessControl"`
	// List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#notification_destinations GrafanaWorkspace#notification_destinations}
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#organizational_units GrafanaWorkspace#organizational_units}
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#organization_role_name GrafanaWorkspace#organization_role_name}
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// Allow workspace admins to install plugins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#plugin_admin_enabled GrafanaWorkspace#plugin_admin_enabled}
	PluginAdminEnabled interface{} `field:"optional" json:"pluginAdminEnabled" yaml:"pluginAdminEnabled"`
	// IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#role_arn GrafanaWorkspace#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// SAML configuration data associated with an AMG workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#saml_configuration GrafanaWorkspace#saml_configuration}
	SamlConfiguration *GrafanaWorkspaceSamlConfiguration `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
	// The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#stack_set_name GrafanaWorkspace#stack_set_name}
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// The list of tags associated with the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#tags GrafanaWorkspace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/grafana_workspace#vpc_configuration GrafanaWorkspace#vpc_configuration}
	VpcConfiguration *GrafanaWorkspaceVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

