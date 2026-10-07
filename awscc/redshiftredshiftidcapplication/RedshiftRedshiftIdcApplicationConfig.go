// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftredshiftidcapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftRedshiftIdcApplicationConfig struct {
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
	// The IAM role ARN for the Amazon Redshift IAM Identity Center application instance.
	//
	// It has the required permissions to be assumed and invoke the IDC Identity Center API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#iam_role_arn RedshiftRedshiftIdcApplication#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The display name for the Amazon Redshift IAM Identity Center application instance. It appears in the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#idc_display_name RedshiftRedshiftIdcApplication#idc_display_name}
	IdcDisplayName *string `field:"required" json:"idcDisplayName" yaml:"idcDisplayName"`
	// The Amazon resource name (ARN) of the IAM Identity Center instance where Amazon Redshift creates a new managed application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#idc_instance_arn RedshiftRedshiftIdcApplication#idc_instance_arn}
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The name of the Redshift application in IAM Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#redshift_idc_application_name RedshiftRedshiftIdcApplication#redshift_idc_application_name}
	RedshiftIdcApplicationName *string `field:"required" json:"redshiftIdcApplicationName" yaml:"redshiftIdcApplicationName"`
	// The type of application being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#application_type RedshiftRedshiftIdcApplication#application_type}
	ApplicationType *string `field:"optional" json:"applicationType" yaml:"applicationType"`
	// The token issuer list for the Amazon Redshift IAM Identity Center application instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#authorized_token_issuer_list RedshiftRedshiftIdcApplication#authorized_token_issuer_list}
	AuthorizedTokenIssuerList interface{} `field:"optional" json:"authorizedTokenIssuerList" yaml:"authorizedTokenIssuerList"`
	// The namespace for the Amazon Redshift IAM Identity Center application instance.
	//
	// It determines which managed application verifies the connection token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#identity_namespace RedshiftRedshiftIdcApplication#identity_namespace}
	IdentityNamespace *string `field:"optional" json:"identityNamespace" yaml:"identityNamespace"`
	// A collection of service integrations for the Redshift IAM Identity Center application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#service_integrations RedshiftRedshiftIdcApplication#service_integrations}
	ServiceIntegrations interface{} `field:"optional" json:"serviceIntegrations" yaml:"serviceIntegrations"`
	// A list of tag keys that Redshift Identity Center applications copy to IAM Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#sso_tag_keys RedshiftRedshiftIdcApplication#sso_tag_keys}
	SsoTagKeys *[]*string `field:"optional" json:"ssoTagKeys" yaml:"ssoTagKeys"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#tags RedshiftRedshiftIdcApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

