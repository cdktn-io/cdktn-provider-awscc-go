// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointElasticsearchSettings struct {
	// The endpoint for the OpenSearch cluster.
	//
	// AWS DMS uses HTTPS if a transport protocol (either HTTP or HTTPS) isn't specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#endpoint_uri DmsEndpoint#endpoint_uri}
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// The maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#error_retry_duration DmsEndpoint#error_retry_duration}
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// The maximum percentage of records that can fail to be written before a full load operation stops.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#full_load_error_percentage DmsEndpoint#full_load_error_percentage}
	FullLoadErrorPercentage *float64 `field:"optional" json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the iam:PassRole action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

