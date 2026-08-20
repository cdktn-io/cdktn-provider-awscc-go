// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2ConnectorConfig struct {
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
	// Display name for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_connector#name Inspectorv2Connector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Provider-specific configuration including regions and scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_connector#provider_configuration Inspectorv2Connector#provider_configuration}
	ProviderConfiguration *Inspectorv2ConnectorProviderConfiguration `field:"required" json:"providerConfiguration" yaml:"providerConfiguration"`
	// The cloud provider for this connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_connector#provider_name Inspectorv2Connector#provider_name}
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Optional description of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_connector#description Inspectorv2Connector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to apply to the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_connector#tags Inspectorv2Connector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

