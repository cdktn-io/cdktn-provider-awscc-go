// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointDocDbSettings struct {
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when NestingLevel is set to "one".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#docs_to_investigate DmsEndpoint#docs_to_investigate}
	DocsToInvestigate *float64 `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when NestingLevel is set to "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#extract_doc_id DmsEndpoint#extract_doc_id}
	ExtractDocId interface{} `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#nesting_level DmsEndpoint#nesting_level}
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

