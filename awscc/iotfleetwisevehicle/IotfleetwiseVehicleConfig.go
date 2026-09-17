// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisevehicle

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseVehicleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#decoder_manifest_arn IotfleetwiseVehicle#decoder_manifest_arn}.
	DecoderManifestArn *string `field:"required" json:"decoderManifestArn" yaml:"decoderManifestArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#model_manifest_arn IotfleetwiseVehicle#model_manifest_arn}.
	ModelManifestArn *string `field:"required" json:"modelManifestArn" yaml:"modelManifestArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#name IotfleetwiseVehicle#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#association_behavior IotfleetwiseVehicle#association_behavior}.
	AssociationBehavior *string `field:"optional" json:"associationBehavior" yaml:"associationBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#attributes IotfleetwiseVehicle#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#state_templates IotfleetwiseVehicle#state_templates}.
	StateTemplates interface{} `field:"optional" json:"stateTemplates" yaml:"stateTemplates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_vehicle#tags IotfleetwiseVehicle#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

