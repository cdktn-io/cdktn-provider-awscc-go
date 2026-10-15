// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint


type LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsRdsSalSettingsRdsSalCredentialsProvider struct {
	// Credentials held in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#secrets_manager_credentials_provider LicensemanagerusersubscriptionsLicenseServerEndpoint#secrets_manager_credentials_provider}
	SecretsManagerCredentialsProvider *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsRdsSalSettingsRdsSalCredentialsProviderSecretsManagerCredentialsProvider `field:"optional" json:"secretsManagerCredentialsProvider" yaml:"secretsManagerCredentialsProvider"`
}

