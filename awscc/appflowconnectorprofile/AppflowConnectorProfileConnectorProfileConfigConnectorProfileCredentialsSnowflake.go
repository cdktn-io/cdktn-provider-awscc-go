// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSnowflake struct {
	// The password that corresponds to the username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_connector_profile#password AppflowConnectorProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_connector_profile#username AppflowConnectorProfile#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

