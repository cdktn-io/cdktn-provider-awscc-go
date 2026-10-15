// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerusersubscriptionslicenseserverendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerusersubscriptionsLicenseServerEndpointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) that identifies the registered identity provider the license server endpoint is created for.
	//
	// For Active Directory, that can be a self-managed Active Directory or an AWS Managed Microsoft AD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#identity_provider_arn LicensemanagerusersubscriptionsLicenseServerEndpoint#identity_provider_arn}
	IdentityProviderArn *string `field:"required" json:"identityProviderArn" yaml:"identityProviderArn"`
	// The settings of the RDS license server to create the endpoint for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#license_server_settings LicensemanagerusersubscriptionsLicenseServerEndpoint#license_server_settings}
	LicenseServerSettings *LicensemanagerusersubscriptionsLicenseServerEndpointLicenseServerSettings `field:"optional" json:"licenseServerSettings" yaml:"licenseServerSettings"`
	// The tags that apply for the license server endpoint. A maximum of 50 tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/licensemanagerusersubscriptions_license_server_endpoint#tags LicensemanagerusersubscriptionsLicenseServerEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

