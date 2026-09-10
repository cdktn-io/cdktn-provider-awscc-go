// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttestcase


type ConnectTestCaseEntryPointVoiceCallEntryPointParameters struct {
	// The destination phonenumber of the EntryPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_test_case#destination_phone_number ConnectTestCase#destination_phone_number}
	DestinationPhoneNumber *string `field:"optional" json:"destinationPhoneNumber" yaml:"destinationPhoneNumber"`
	// The flow id used for the TestCase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_test_case#flow_id ConnectTestCase#flow_id}
	FlowId *string `field:"optional" json:"flowId" yaml:"flowId"`
	// The source phonenumber of the EntryPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_test_case#source_phone_number ConnectTestCase#source_phone_number}
	SourcePhoneNumber *string `field:"optional" json:"sourcePhoneNumber" yaml:"sourcePhoneNumber"`
}

