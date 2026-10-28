// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint


type LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsRdsSalSettings struct {
	// The source of the credentials that License Manager uses to reach the license server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#rds_sal_credentials_provider LicensemanagerusersubscriptionsLicenseServerEndpoint#rds_sal_credentials_provider}
	RdsSalCredentialsProvider *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettingsRdsSalSettingsRdsSalCredentialsProvider `field:"optional" json:"rdsSalCredentialsProvider" yaml:"rdsSalCredentialsProvider"`
}

