// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannelflow


type ChimeChannelFlowProcessorsConfigurationLambda struct {
	// Controls how the Lambda function is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chime_channel_flow#invocation_type ChimeChannelFlow#invocation_type}
	InvocationType *string `field:"required" json:"invocationType" yaml:"invocationType"`
	// The ARN of the Lambda message processing function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/chime_channel_flow#resource_arn ChimeChannelFlow#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

