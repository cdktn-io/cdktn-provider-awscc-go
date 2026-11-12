// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputMediaConnectFlows struct {
	// The unique ARN of the MediaConnect Flow being used as a source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#flow_arn MedialiveInput#flow_arn}
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
}

