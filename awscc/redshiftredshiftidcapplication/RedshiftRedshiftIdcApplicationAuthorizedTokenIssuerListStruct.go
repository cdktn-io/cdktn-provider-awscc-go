// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftredshiftidcapplication


type RedshiftRedshiftIdcApplicationAuthorizedTokenIssuerListStruct struct {
	// The list of audiences for the authorized token issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#authorized_audiences_list RedshiftRedshiftIdcApplication#authorized_audiences_list}
	AuthorizedAudiencesList *[]*string `field:"optional" json:"authorizedAudiencesList" yaml:"authorizedAudiencesList"`
	// The ARN for the authorized token issuer for integrating Amazon Redshift with IDC Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application#trusted_token_issuer_arn RedshiftRedshiftIdcApplication#trusted_token_issuer_arn}
	TrustedTokenIssuerArn *string `field:"optional" json:"trustedTokenIssuerArn" yaml:"trustedTokenIssuerArn"`
}

