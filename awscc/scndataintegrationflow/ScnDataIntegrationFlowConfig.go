// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataintegrationflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScnDataIntegrationFlowConfig struct {
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
	// The Amazon Web Services Supply Chain instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#instance_id ScnDataIntegrationFlow#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the DataIntegrationFlow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#name ScnDataIntegrationFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source configurations for the DataIntegrationFlow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#sources ScnDataIntegrationFlow#sources}
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// The DataIntegrationFlow target parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#target ScnDataIntegrationFlow#target}
	Target *ScnDataIntegrationFlowTarget `field:"required" json:"target" yaml:"target"`
	// The DataIntegrationFlow transformation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#transformation ScnDataIntegrationFlow#transformation}
	Transformation *ScnDataIntegrationFlowTransformation `field:"required" json:"transformation" yaml:"transformation"`
	// The tags for the DataIntegrationFlow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/scn_data_integration_flow#tags ScnDataIntegrationFlow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

