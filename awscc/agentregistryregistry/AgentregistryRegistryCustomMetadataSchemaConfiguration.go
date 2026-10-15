// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryCustomMetadataSchemaConfiguration struct {
	// The default JSON Schema that applies to record types without a specific override.
	//
	// Supported property types are string, string with an enum constraint, string with a uri format, and boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/agentregistry_registry#default_schema AgentregistryRegistry#default_schema}
	DefaultSchema *string `field:"optional" json:"defaultSchema" yaml:"defaultSchema"`
	// A list of per-record-type schema overrides.
	//
	// When a record's type matches an override, that override's schema is used instead of the default schema for validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/agentregistry_registry#record_type_schema_overrides AgentregistryRegistry#record_type_schema_overrides}
	RecordTypeSchemaOverrides interface{} `field:"optional" json:"recordTypeSchemaOverrides" yaml:"recordTypeSchemaOverrides"`
}

