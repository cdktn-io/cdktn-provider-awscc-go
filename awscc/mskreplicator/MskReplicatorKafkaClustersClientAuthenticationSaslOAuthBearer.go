// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearer struct {
	// Details for SASL/OAUTHBEARER using standard client_credentials grant. Mutually exclusive with iamJwtBearer and clientCredentialsAssertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#client_credentials MskReplicator#client_credentials}
	ClientCredentials *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials `field:"optional" json:"clientCredentials" yaml:"clientCredentials"`
	// Details for SASL/OAUTHBEARER using client credentials grant with JWT client assertion (RFC 7521/7523). Mutually exclusive with clientCredentials and iamJwtBearer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#client_credentials_assertion MskReplicator#client_credentials_assertion}
	ClientCredentialsAssertion *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertion `field:"optional" json:"clientCredentialsAssertion" yaml:"clientCredentialsAssertion"`
	// Details for SASL/OAUTHBEARER using JWT Bearer assertion grant (RFC 7523). Mutually exclusive with clientCredentials and clientCredentialsAssertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#iam_jwt_bearer MskReplicator#iam_jwt_bearer}
	IamJwtBearer *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer `field:"optional" json:"iamJwtBearer" yaml:"iamJwtBearer"`
	// OAuth scope to request. Included in the token request if provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#scope MskReplicator#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// How client credentials are sent to the identity provider (POST, BASIC, or NONE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#token_endpoint_authentication_method MskReplicator#token_endpoint_authentication_method}
	TokenEndpointAuthenticationMethod *string `field:"optional" json:"tokenEndpointAuthenticationMethod" yaml:"tokenEndpointAuthenticationMethod"`
	// Secrets Manager ARN containing a custom CA certificate for the identity provider.
	//
	// Required only if the identity provider uses a private CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#token_endpoint_tls_certificate_arn MskReplicator#token_endpoint_tls_certificate_arn}
	TokenEndpointTlsCertificateArn *string `field:"optional" json:"tokenEndpointTlsCertificateArn" yaml:"tokenEndpointTlsCertificateArn"`
	// The HTTPS URL of the OAuth token endpoint that vends OAuth Bearer tokens per RFC 6749.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#token_endpoint_url MskReplicator#token_endpoint_url}
	TokenEndpointUrl *string `field:"optional" json:"tokenEndpointUrl" yaml:"tokenEndpointUrl"`
}

