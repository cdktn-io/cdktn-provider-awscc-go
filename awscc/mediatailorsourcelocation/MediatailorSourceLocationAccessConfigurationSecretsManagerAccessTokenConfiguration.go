// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation


type MediatailorSourceLocationAccessConfigurationSecretsManagerAccessTokenConfiguration struct {
	// <p>The name of the HTTP header used to supply the access token in requests to the source location.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_source_location#header_name MediatailorSourceLocation#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// <p>The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_source_location#secret_arn MediatailorSourceLocation#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// <p>The AWS Secrets Manager <a href="https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html">SecretString</a> key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_source_location#secret_string_key MediatailorSourceLocation#secret_string_key}
	SecretStringKey *string `field:"optional" json:"secretStringKey" yaml:"secretStringKey"`
}

