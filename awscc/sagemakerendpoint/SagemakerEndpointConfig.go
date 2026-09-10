// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfig struct {
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
	// The name of the endpoint configuration for the SageMaker endpoint. This is a required property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#endpoint_config_name SagemakerEndpoint#endpoint_config_name}
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
	// Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#deployment_config SagemakerEndpoint#deployment_config}
	DeploymentConfig *SagemakerEndpointDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The name of the SageMaker endpoint. This name must be unique within an AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#endpoint_name SagemakerEndpoint#endpoint_name}
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Specifies a list of variant properties that you want to exclude when updating an endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#exclude_retained_variant_properties SagemakerEndpoint#exclude_retained_variant_properties}
	ExcludeRetainedVariantProperties interface{} `field:"optional" json:"excludeRetainedVariantProperties" yaml:"excludeRetainedVariantProperties"`
	// When set to true, retains all variant properties for an endpoint when it is updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#retain_all_variant_properties SagemakerEndpoint#retain_all_variant_properties}
	RetainAllVariantProperties interface{} `field:"optional" json:"retainAllVariantProperties" yaml:"retainAllVariantProperties"`
	// When set to true, retains the deployment configuration during endpoint updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#retain_deployment_config SagemakerEndpoint#retain_deployment_config}
	RetainDeploymentConfig interface{} `field:"optional" json:"retainDeploymentConfig" yaml:"retainDeploymentConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint#tags SagemakerEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

