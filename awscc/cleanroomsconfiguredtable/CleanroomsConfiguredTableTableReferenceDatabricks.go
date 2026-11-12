// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReferenceDatabricks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#connection_type CleanroomsConfiguredTable#connection_type}.
	ConnectionType *CleanroomsConfiguredTableTableReferenceDatabricksConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#schema_name CleanroomsConfiguredTable#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#table_name CleanroomsConfiguredTable#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cleanrooms_configured_table#workspace_url CleanroomsConfiguredTable#workspace_url}.
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

