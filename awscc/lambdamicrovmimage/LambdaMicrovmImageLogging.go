// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageLogging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#cloudwatch LambdaMicrovmImage#cloudwatch}.
	Cloudwatch *LambdaMicrovmImageLoggingCloudwatch `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_microvm_image#disabled LambdaMicrovmImage#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

