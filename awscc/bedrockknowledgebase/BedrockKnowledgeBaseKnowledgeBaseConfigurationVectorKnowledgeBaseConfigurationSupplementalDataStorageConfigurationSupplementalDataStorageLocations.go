// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations struct {
	// An Amazon S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_knowledge_base#s3_location BedrockKnowledgeBase#s3_location}
	S3Location *BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocationsS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
	// Supplemental data storage location type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_knowledge_base#supplemental_data_storage_location_type BedrockKnowledgeBase#supplemental_data_storage_location_type}
	SupplementalDataStorageLocationType *string `field:"optional" json:"supplementalDataStorageLocationType" yaml:"supplementalDataStorageLocationType"`
}

