// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationobjectstorage


type DatasyncLocationObjectStorageFederatedIdentityExternalIdentityGoogleOidc struct {
	// The name of the Google Cloud workload identity pool that DataSync federates with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#identity_pool_name DatasyncLocationObjectStorage#identity_pool_name}
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// The name of the OIDC identity provider configured in the Google Cloud workload identity pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#identity_provider_name DatasyncLocationObjectStorage#identity_provider_name}
	IdentityProviderName *string `field:"optional" json:"identityProviderName" yaml:"identityProviderName"`
	// The human-readable Google Cloud project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#project_name DatasyncLocationObjectStorage#project_name}
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The numeric Google Cloud project ID, as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#project_number DatasyncLocationObjectStorage#project_number}
	ProjectNumber *string `field:"optional" json:"projectNumber" yaml:"projectNumber"`
}

