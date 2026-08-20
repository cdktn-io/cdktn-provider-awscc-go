// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsMlflowProperties struct {
	// The ARN of the MLflow tracking server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_connection#tracking_server_arn DatazoneConnection#tracking_server_arn}
	TrackingServerArn *string `field:"optional" json:"trackingServerArn" yaml:"trackingServerArn"`
}

