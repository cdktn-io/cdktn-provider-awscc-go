// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccesstrustprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VerifiedAccessTrustProviderConfig struct {
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
	// The identifier to be used when working with policy rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#policy_reference_name Ec2VerifiedAccessTrustProvider#policy_reference_name}
	PolicyReferenceName *string `field:"required" json:"policyReferenceName" yaml:"policyReferenceName"`
	// Type of trust provider. Possible values: user|device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#trust_provider_type Ec2VerifiedAccessTrustProvider#trust_provider_type}
	TrustProviderType *string `field:"required" json:"trustProviderType" yaml:"trustProviderType"`
	// A description for the Amazon Web Services Verified Access trust provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#description Ec2VerifiedAccessTrustProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The options for device identity based trust providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#device_options Ec2VerifiedAccessTrustProvider#device_options}
	DeviceOptions *Ec2VerifiedAccessTrustProviderDeviceOptions `field:"optional" json:"deviceOptions" yaml:"deviceOptions"`
	// The type of device-based trust provider. Possible values: jamf|crowdstrike.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#device_trust_provider_type Ec2VerifiedAccessTrustProvider#device_trust_provider_type}
	DeviceTrustProviderType *string `field:"optional" json:"deviceTrustProviderType" yaml:"deviceTrustProviderType"`
	// The OpenID Connect details for an oidc -type, user-identity based trust provider for L4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#native_application_oidc_options Ec2VerifiedAccessTrustProvider#native_application_oidc_options}
	NativeApplicationOidcOptions *Ec2VerifiedAccessTrustProviderNativeApplicationOidcOptions `field:"optional" json:"nativeApplicationOidcOptions" yaml:"nativeApplicationOidcOptions"`
	// The OpenID Connect details for an oidc -type, user-identity based trust provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#oidc_options Ec2VerifiedAccessTrustProvider#oidc_options}
	OidcOptions *Ec2VerifiedAccessTrustProviderOidcOptions `field:"optional" json:"oidcOptions" yaml:"oidcOptions"`
	// The configuration options for customer provided KMS encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#sse_specification Ec2VerifiedAccessTrustProvider#sse_specification}
	SseSpecification *Ec2VerifiedAccessTrustProviderSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#tags Ec2VerifiedAccessTrustProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of device-based trust provider. Possible values: oidc|iam-identity-center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_trust_provider#user_trust_provider_type Ec2VerifiedAccessTrustProvider#user_trust_provider_type}
	UserTrustProviderType *string `field:"optional" json:"userTrustProviderType" yaml:"userTrustProviderType"`
}

