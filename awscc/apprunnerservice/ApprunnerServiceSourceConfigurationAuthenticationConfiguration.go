// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationAuthenticationConfiguration struct {
	// Access Role Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_service#access_role_arn ApprunnerService#access_role_arn}
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// Connection Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_service#connection_arn ApprunnerService#connection_arn}
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

