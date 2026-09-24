// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob


type DatasyncLocationAzureBlobFederatedIdentity struct {
	// Specifies the ARN of the AWS Identity and Access Management (IAM) role that DataSync assumes to mint the OIDC token used to authenticate with the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_azure_blob#aws_iam_role DatasyncLocationAzureBlob#aws_iam_role}
	AwsIamRole *string `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// Specifies the Microsoft Entra (Azure AD) identity that DataSync federates with to obtain an access token for your Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_azure_blob#azure_oidc DatasyncLocationAzureBlob#azure_oidc}
	AzureOidc *DatasyncLocationAzureBlobFederatedIdentityAzureOidc `field:"optional" json:"azureOidc" yaml:"azureOidc"`
}

