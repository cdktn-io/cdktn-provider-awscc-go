// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventinvokeconfig


type LambdaEventInvokeConfigDestinationConfigOnFailure struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_event_invoke_config#destination LambdaEventInvokeConfig#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

