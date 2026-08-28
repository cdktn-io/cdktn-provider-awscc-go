// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointDynamoDbSettings struct {
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the iam:PassRole action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dms_endpoint#service_access_role_arn DmsEndpoint#service_access_role_arn}
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

