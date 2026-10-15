// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint


type LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettings struct {
	// The settings of a license server, by server type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#server_settings LicensemanagerusersubscriptionsLicenseServerEndpoint#server_settings}
	ServerSettings *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettingsServerSettings `field:"optional" json:"serverSettings" yaml:"serverSettings"`
	// The type of license server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#server_type LicensemanagerusersubscriptionsLicenseServerEndpoint#server_type}
	ServerType *string `field:"optional" json:"serverType" yaml:"serverType"`
}

