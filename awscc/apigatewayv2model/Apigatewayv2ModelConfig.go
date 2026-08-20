// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2model

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ModelConfig struct {
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
	// The API identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_model#api_id Apigatewayv2Model#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_model#name Apigatewayv2Model#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_model#schema Apigatewayv2Model#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The content-type for the model, for example, "application/json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_model#content_type Apigatewayv2Model#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_model#description Apigatewayv2Model#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

