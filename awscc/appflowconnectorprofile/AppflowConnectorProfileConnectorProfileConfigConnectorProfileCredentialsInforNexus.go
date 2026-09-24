// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsInforNexus struct {
	// The Access Key portion of the credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#access_key_id AppflowConnectorProfile#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// The encryption keys used to encrypt data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#datakey AppflowConnectorProfile#datakey}
	Datakey *string `field:"optional" json:"datakey" yaml:"datakey"`
	// The secret key used to sign requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#secret_access_key AppflowConnectorProfile#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The identi?er for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#user_id AppflowConnectorProfile#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

