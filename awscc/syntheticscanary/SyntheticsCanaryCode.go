// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#blueprint_types SyntheticsCanary#blueprint_types}.
	BlueprintTypes *[]*string `field:"optional" json:"blueprintTypes" yaml:"blueprintTypes"`
	// List of Lambda layers to attach to the canary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#dependencies SyntheticsCanary#dependencies}
	Dependencies interface{} `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#handler SyntheticsCanary#handler}.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#s3_bucket SyntheticsCanary#s3_bucket}.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#s3_key SyntheticsCanary#s3_key}.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#s3_object_version SyntheticsCanary#s3_object_version}.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#script SyntheticsCanary#script}.
	Script *string `field:"optional" json:"script" yaml:"script"`
}

