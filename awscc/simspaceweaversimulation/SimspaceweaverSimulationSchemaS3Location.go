// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package simspaceweaversimulation


type SimspaceweaverSimulationSchemaS3Location struct {
	// The Schema S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/simspaceweaver_simulation#bucket_name SimspaceweaverSimulation#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/simspaceweaver_simulation#object_key SimspaceweaverSimulation#object_key}
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

