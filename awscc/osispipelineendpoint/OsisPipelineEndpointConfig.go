// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipelineendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OsisPipelineEndpointConfig struct {
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
	// The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/osis_pipeline_endpoint#pipeline_arn OsisPipelineEndpoint#pipeline_arn}
	PipelineArn *string `field:"required" json:"pipelineArn" yaml:"pipelineArn"`
	// Configuration settings for the VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/osis_pipeline_endpoint#vpc_options OsisPipelineEndpoint#vpc_options}
	VpcOptions *OsisPipelineEndpointVpcOptions `field:"required" json:"vpcOptions" yaml:"vpcOptions"`
}

