// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage


type LambdaMicrovmImageLoggingCloudwatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#log_group LambdaMicrovmImage#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_microvm_image#log_stream LambdaMicrovmImage#log_stream}.
	LogStream *string `field:"optional" json:"logStream" yaml:"logStream"`
}

