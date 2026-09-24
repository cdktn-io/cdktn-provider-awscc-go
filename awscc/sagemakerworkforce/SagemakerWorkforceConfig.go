// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerWorkforceConfig struct {
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
	// The name of the private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#workforce_name SagemakerWorkforce#workforce_name}
	WorkforceName *string `field:"required" json:"workforceName" yaml:"workforceName"`
	// The configuration of an Amazon Cognito workforce.
	//
	// A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#cognito_config SagemakerWorkforce#cognito_config}
	CognitoConfig *SagemakerWorkforceCognitoConfig `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// The IP address type for the workforce. IPv4 only or dualstack (IPv4 and IPv6).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#ip_address_type SagemakerWorkforce#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The configuration of an OIDC Identity Provider (IdP) private workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#oidc_config SagemakerWorkforce#oidc_config}
	OidcConfig *SagemakerWorkforceOidcConfig `field:"optional" json:"oidcConfig" yaml:"oidcConfig"`
	// A list of IP address ranges used to access your training data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#source_ip_config SagemakerWorkforce#source_ip_config}
	SourceIpConfig *SagemakerWorkforceSourceIpConfig `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// An array of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#tags SagemakerWorkforce#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The VPC configuration for the workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_workforce#workforce_vpc_config SagemakerWorkforce#workforce_vpc_config}
	WorkforceVpcConfig *SagemakerWorkforceWorkforceVpcConfig `field:"optional" json:"workforceVpcConfig" yaml:"workforceVpcConfig"`
}

