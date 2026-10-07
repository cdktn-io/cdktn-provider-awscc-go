// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizerendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComprehendEntityRecognizerEndpointConfig struct {
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
	// The desired number of inference units to be used by the model.
	//
	// Each inference unit represents throughput of 100 characters per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#desired_inference_units ComprehendEntityRecognizerEndpoint#desired_inference_units}
	DesiredInferenceUnits *float64 `field:"required" json:"desiredInferenceUnits" yaml:"desiredInferenceUnits"`
	// The name of the endpoint. The name must be unique within the AWS Region and account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#endpoint_name ComprehendEntityRecognizerEndpoint#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to trained custom models encrypted with a customer managed key (ModelKmsKeyId).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#data_access_role_arn ComprehendEntityRecognizerEndpoint#data_access_role_arn}
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The Amazon Resource Name (ARN) of the flywheel to which the endpoint is attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#flywheel_arn ComprehendEntityRecognizerEndpoint#flywheel_arn}
	FlywheelArn *string `field:"optional" json:"flywheelArn" yaml:"flywheelArn"`
	// The Amazon Resource Name (ARN) of the entity recognizer model to which the endpoint is attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#model_arn ComprehendEntityRecognizerEndpoint#model_arn}
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// Tags associated with the endpoint being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer_endpoint#tags ComprehendEntityRecognizerEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

