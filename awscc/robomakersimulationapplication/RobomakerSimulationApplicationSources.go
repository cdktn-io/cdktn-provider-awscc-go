// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package robomakersimulationapplication


type RobomakerSimulationApplicationSources struct {
	// The target processor architecture for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/robomaker_simulation_application#architecture RobomakerSimulationApplication#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The Amazon S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/robomaker_simulation_application#s3_bucket RobomakerSimulationApplication#s3_bucket}
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 object key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/robomaker_simulation_application#s3_key RobomakerSimulationApplication#s3_key}
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

