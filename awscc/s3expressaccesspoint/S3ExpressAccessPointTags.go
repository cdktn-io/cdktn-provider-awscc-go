// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressaccesspoint


type S3ExpressAccessPointTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3express_access_point#key S3ExpressAccessPoint#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3express_access_point#value S3ExpressAccessPoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

