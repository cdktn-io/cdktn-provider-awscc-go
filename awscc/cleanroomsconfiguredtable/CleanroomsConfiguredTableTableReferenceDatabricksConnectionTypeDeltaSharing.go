// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReferenceDatabricksConnectionTypeDeltaSharing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#delta_share_endpoint CleanroomsConfiguredTable#delta_share_endpoint}.
	DeltaShareEndpoint *string `field:"optional" json:"deltaShareEndpoint" yaml:"deltaShareEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#share_name CleanroomsConfiguredTable#share_name}.
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
}

