// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsS3Properties struct {
	// Specifies whether to register the S3 Access Grant location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#register_s3_access_grant_location DatazoneConnection#register_s3_access_grant_location}
	RegisterS3AccessGrantLocation interface{} `field:"optional" json:"registerS3AccessGrantLocation" yaml:"registerS3AccessGrantLocation"`
	// The Amazon S3 Access Grant location ID that's part of the Amazon S3 properties of a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#s3_access_grant_location_id DatazoneConnection#s3_access_grant_location_id}
	S3AccessGrantLocationId *string `field:"optional" json:"s3AccessGrantLocationId" yaml:"s3AccessGrantLocationId"`
	// The Amazon S3 URI that's part of the Amazon S3 properties of a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#s3_uri DatazoneConnection#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

