// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfiguration struct {
	// A map of entity configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#entity_configurations GlueConnectionType#entity_configurations}
	EntityConfigurations interface{} `field:"optional" json:"entityConfigurations" yaml:"entityConfigurations"`
	// Configuration that defines how to make requests to endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#global_source_configuration GlueConnectionType#global_source_configuration}
	GlobalSourceConfiguration *GlueConnectionTypeRestConfigurationGlobalSourceConfiguration `field:"optional" json:"globalSourceConfiguration" yaml:"globalSourceConfiguration"`
	// Configuration for the validation endpoint. Only supports RequestMethod and RequestPath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#validation_endpoint_configuration GlueConnectionType#validation_endpoint_configuration}
	ValidationEndpointConfiguration *GlueConnectionTypeRestConfigurationValidationEndpointConfiguration `field:"optional" json:"validationEndpointConfiguration" yaml:"validationEndpointConfiguration"`
}

