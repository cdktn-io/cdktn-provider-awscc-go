// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#aws_managed_client_application_reference DatazoneConnection#aws_managed_client_application_reference}.
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/datazone_connection#user_managed_client_application_client_id DatazoneConnection#user_managed_client_application_client_id}.
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

