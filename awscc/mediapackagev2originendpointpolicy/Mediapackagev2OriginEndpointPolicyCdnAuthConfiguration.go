// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpointpolicy


type Mediapackagev2OriginEndpointPolicyCdnAuthConfiguration struct {
	// <p>The ARN for the secret in Secrets Manager that your CDN uses for authorization to access the endpoint.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint_policy#cdn_identifier_secret_arns Mediapackagev2OriginEndpointPolicy#cdn_identifier_secret_arns}
	CdnIdentifierSecretArns *[]*string `field:"optional" json:"cdnIdentifierSecretArns" yaml:"cdnIdentifierSecretArns"`
	// <p>The ARN for the IAM role that gives MediaPackage read access to Secrets Manager and KMS for CDN authorization.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackagev2_origin_endpoint_policy#secrets_role_arn Mediapackagev2OriginEndpointPolicy#secrets_role_arn}
	SecretsRoleArn *string `field:"optional" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

