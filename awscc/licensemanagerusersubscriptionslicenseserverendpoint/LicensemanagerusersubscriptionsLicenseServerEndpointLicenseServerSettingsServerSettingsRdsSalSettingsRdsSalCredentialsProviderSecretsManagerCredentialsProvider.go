// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint


type LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsRdsSalSettingsRdsSalCredentialsProviderSecretsManagerCredentialsProvider struct {
	// The identifier of the Secrets Manager secret that holds the license server administrator credentials.
	//
	// The name of the secret must begin with the prefix license-manager-user-.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#secret_id LicensemanagerusersubscriptionsLicenseServerEndpoint#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

