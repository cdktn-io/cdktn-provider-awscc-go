// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the connection type. Must be prefixed with REST-.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#connection_type GlueConnectionType#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// Configuration for HTTP request and response handling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#rest_configuration GlueConnectionType#rest_configuration}
	RestConfiguration *GlueConnectionTypeRestConfiguration `field:"required" json:"restConfiguration" yaml:"restConfiguration"`
	// Configuration that defines the base URL and additional request parameters needed during connection creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#connection_properties GlueConnectionType#connection_properties}
	ConnectionProperties *GlueConnectionTypeConnectionProperties `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Configuration that defines supported authentication types and required properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#connector_authentication_configuration GlueConnectionType#connector_authentication_configuration}
	ConnectorAuthenticationConfiguration *GlueConnectionTypeConnectorAuthenticationConfiguration `field:"optional" json:"connectorAuthenticationConfiguration" yaml:"connectorAuthenticationConfiguration"`
	// A description of the connection type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#description GlueConnectionType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The integration type for the connection. Currently only REST is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#integration_type GlueConnectionType#integration_type}
	IntegrationType *string `field:"optional" json:"integrationType" yaml:"integrationType"`
	// Tags to assign to the connection type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#tags GlueConnectionType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

