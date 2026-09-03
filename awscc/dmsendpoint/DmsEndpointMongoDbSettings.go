// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointMongoDbSettings struct {
	// The authentication mechanism you use to access the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#auth_mechanism DmsEndpoint#auth_mechanism}
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// The MongoDB database name. This setting isn't used when AuthType is set to "no".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#auth_source DmsEndpoint#auth_source}
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// The authentication type you use to access the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#auth_type DmsEndpoint#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The database name on the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#database_name DmsEndpoint#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Indicates the number of documents to preview to determine the document organization.
	//
	// Use this setting when NestingLevel is set to "one".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#docs_to_investigate DmsEndpoint#docs_to_investigate}
	DocsToInvestigate *string `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID. Use this setting when NestingLevel is set to "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#extract_doc_id DmsEndpoint#extract_doc_id}
	ExtractDocId *string `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Specifies either document or table mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#nesting_level DmsEndpoint#nesting_level}
	NestingLevel *string `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// The password for the user account you use to access the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#password DmsEndpoint#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port value for the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#port DmsEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#secrets_manager_access_role_arn DmsEndpoint#secrets_manager_access_role_arn}
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MongoDB endpoint connection details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#secrets_manager_secret_id DmsEndpoint#secrets_manager_secret_id}
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The name of the server on the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#server_name DmsEndpoint#server_name}
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The user name you use to access the MongoDB source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dms_endpoint#username DmsEndpoint#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

