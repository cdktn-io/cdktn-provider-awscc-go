// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject


type BedrockDataAutomationProjectCustomOutputConfigurationBlueprints struct {
	// ARN of a Blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#blueprint_arn BedrockDataAutomationProject#blueprint_arn}
	BlueprintArn *string `field:"optional" json:"blueprintArn" yaml:"blueprintArn"`
	// Stage of the Blueprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#blueprint_stage BedrockDataAutomationProject#blueprint_stage}
	BlueprintStage *string `field:"optional" json:"blueprintStage" yaml:"blueprintStage"`
	// Blueprint Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_automation_project#blueprint_version BedrockDataAutomationProject#blueprint_version}
	BlueprintVersion *string `field:"optional" json:"blueprintVersion" yaml:"blueprintVersion"`
}

