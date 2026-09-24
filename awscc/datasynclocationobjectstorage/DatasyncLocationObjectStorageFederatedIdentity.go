// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationobjectstorage


type DatasyncLocationObjectStorageFederatedIdentity struct {
	// Specifies the ARN of the AWS Identity and Access Management (IAM) role that DataSync assumes to mint the OIDC token used to authenticate with the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#aws_iam_role DatasyncLocationObjectStorage#aws_iam_role}
	AwsIamRole *string `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// Specifies the external (non-AWS) identity provider that DataSync federates with to access your object storage location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_object_storage#external_identity DatasyncLocationObjectStorage#external_identity}
	ExternalIdentity *DatasyncLocationObjectStorageFederatedIdentityExternalIdentity `field:"optional" json:"externalIdentity" yaml:"externalIdentity"`
}

