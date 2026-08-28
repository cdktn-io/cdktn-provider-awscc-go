// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatapipelinePipelineConfig struct {
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
	// The name of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#name DatapipelinePipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether to validate and start the pipeline or stop an active pipeline.
	//
	// By default, the value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#activate DatapipelinePipeline#activate}
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// A description of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#description DatapipelinePipeline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameter objects used with the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#parameter_objects DatapipelinePipeline#parameter_objects}
	ParameterObjects interface{} `field:"optional" json:"parameterObjects" yaml:"parameterObjects"`
	// The parameter values used with the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#parameter_values DatapipelinePipeline#parameter_values}
	ParameterValues interface{} `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// The objects that define the pipeline.
	//
	// These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#pipeline_objects DatapipelinePipeline#pipeline_objects}
	PipelineObjects interface{} `field:"optional" json:"pipelineObjects" yaml:"pipelineObjects"`
	// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
	//
	// For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datapipeline_pipeline#pipeline_tags DatapipelinePipeline#pipeline_tags}
	PipelineTags interface{} `field:"optional" json:"pipelineTags" yaml:"pipelineTags"`
}

