// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsidmappingtable


type CleanroomsIdMappingTableInputReferenceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_id_mapping_table#input_reference_arn CleanroomsIdMappingTable#input_reference_arn}.
	InputReferenceArn *string `field:"required" json:"inputReferenceArn" yaml:"inputReferenceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_id_mapping_table#manage_resource_policies CleanroomsIdMappingTable#manage_resource_policies}.
	ManageResourcePolicies interface{} `field:"required" json:"manageResourcePolicies" yaml:"manageResourcePolicies"`
}

