// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationobjectstorage


type DatasyncLocationObjectStorageFederatedIdentityExternalIdentity struct {
	// Specifies the Google Cloud workload identity federation configuration that DataSync uses to obtain an access token for your Google Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#google_oidc DatasyncLocationObjectStorage#google_oidc}
	GoogleOidc *DatasyncLocationObjectStorageFederatedIdentityExternalIdentityGoogleOidc `field:"optional" json:"googleOidc" yaml:"googleOidc"`
}

