// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppflowConnectorConfig struct {
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
	// Contains information about the configuration of the connector being registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector#connector_provisioning_config AppflowConnector#connector_provisioning_config}
	ConnectorProvisioningConfig *AppflowConnectorConnectorProvisioningConfig `field:"required" json:"connectorProvisioningConfig" yaml:"connectorProvisioningConfig"`
	// The provisioning type of the connector. Currently the only supported value is LAMBDA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector#connector_provisioning_type AppflowConnector#connector_provisioning_type}
	ConnectorProvisioningType *string `field:"required" json:"connectorProvisioningType" yaml:"connectorProvisioningType"`
	// The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector#connector_label AppflowConnector#connector_label}
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// A description about the connector that's being registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector#description AppflowConnector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

