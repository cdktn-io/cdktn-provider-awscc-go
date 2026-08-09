// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsendpoint


type S3OutpostsEndpointFailedReason struct {
	// The failure code, if any, for a create or delete endpoint operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_endpoint#error_code S3OutpostsEndpoint#error_code}
	ErrorCode *string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// Additional error details describing the endpoint failure and recommended action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_endpoint#message S3OutpostsEndpoint#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
}

