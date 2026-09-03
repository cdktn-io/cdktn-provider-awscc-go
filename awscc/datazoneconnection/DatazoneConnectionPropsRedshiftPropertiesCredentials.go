// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsRedshiftPropertiesCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_connection#secret_arn DatazoneConnection#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The username and password to be used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_connection#username_password DatazoneConnection#username_password}
	UsernamePassword *DatazoneConnectionPropsRedshiftPropertiesCredentialsUsernamePassword `field:"optional" json:"usernamePassword" yaml:"usernamePassword"`
}

