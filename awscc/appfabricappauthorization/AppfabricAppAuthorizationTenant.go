// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization


type AppfabricAppAuthorizationTenant struct {
	// The display name of the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_app_authorization#tenant_display_name AppfabricAppAuthorization#tenant_display_name}
	TenantDisplayName *string `field:"required" json:"tenantDisplayName" yaml:"tenantDisplayName"`
	// The ID of the application tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_app_authorization#tenant_identifier AppfabricAppAuthorization#tenant_identifier}
	TenantIdentifier *string `field:"required" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}

