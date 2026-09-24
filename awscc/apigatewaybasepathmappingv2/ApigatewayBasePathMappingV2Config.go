// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaybasepathmappingv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayBasePathMappingV2Config struct {
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
	// The Arn of an AWS::ApiGateway::DomainNameV2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apigateway_base_path_mapping_v2#domain_name_arn ApigatewayBasePathMappingV2#domain_name_arn}
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// The ID of the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apigateway_base_path_mapping_v2#rest_api_id ApigatewayBasePathMappingV2#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The base path name that callers of the API must provide in the URL after the domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apigateway_base_path_mapping_v2#base_path ApigatewayBasePathMappingV2#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The name of the API's stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apigateway_base_path_mapping_v2#stage ApigatewayBasePathMappingV2#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

