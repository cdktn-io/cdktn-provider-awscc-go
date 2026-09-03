// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttestcase


type ConnectTestCaseEntryPoint struct {
	// The chat entry point parameters for the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_test_case#chat_entry_point_parameters ConnectTestCase#chat_entry_point_parameters}
	ChatEntryPointParameters *ConnectTestCaseEntryPointChatEntryPointParameters `field:"optional" json:"chatEntryPointParameters" yaml:"chatEntryPointParameters"`
	// The type of the Entry Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_test_case#type ConnectTestCase#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The voice call entry point parameters for the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_test_case#voice_call_entry_point_parameters ConnectTestCase#voice_call_entry_point_parameters}
	VoiceCallEntryPointParameters *ConnectTestCaseEntryPointVoiceCallEntryPointParameters `field:"optional" json:"voiceCallEntryPointParameters" yaml:"voiceCallEntryPointParameters"`
}

