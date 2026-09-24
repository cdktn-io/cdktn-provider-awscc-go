// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob


type DatasyncLocationAzureBlobFederatedIdentityAzureOidc struct {
	// Specifies the client ID of the Microsoft Entra (Azure AD) identity that DataSync uses to obtain an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_azure_blob#client_id DatasyncLocationAzureBlob#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Specifies the Microsoft Entra (Azure AD) tenant ID that the identity belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_azure_blob#tenant_id DatasyncLocationAzureBlob#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

