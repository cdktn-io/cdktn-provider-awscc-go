// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore


type HealthlakeFhirDatastoreIdentityProviderConfiguration struct {
	// Type of Authorization Strategy. The two types of supported Authorization strategies are SMART_ON_FHIR_V1 and AWS_AUTH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_fhir_datastore#authorization_strategy HealthlakeFhirDatastore#authorization_strategy}
	AuthorizationStrategy *string `field:"optional" json:"authorizationStrategy" yaml:"authorizationStrategy"`
	// Flag to indicate if fine-grained authorization will be enabled for the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_fhir_datastore#fine_grained_authorization_enabled HealthlakeFhirDatastore#fine_grained_authorization_enabled}
	FineGrainedAuthorizationEnabled interface{} `field:"optional" json:"fineGrainedAuthorizationEnabled" yaml:"fineGrainedAuthorizationEnabled"`
	// The Amazon Resource Name (ARN) of the Lambda function that will be used to decode the access token created by the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_fhir_datastore#idp_lambda_arn HealthlakeFhirDatastore#idp_lambda_arn}
	IdpLambdaArn *string `field:"optional" json:"idpLambdaArn" yaml:"idpLambdaArn"`
	// The JSON metadata elements for identity provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_fhir_datastore#metadata HealthlakeFhirDatastore#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

