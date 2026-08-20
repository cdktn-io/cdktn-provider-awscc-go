// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironmentactions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneEnvironmentActionsConfig struct {
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
	// The name of the environment action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#name DatazoneEnvironmentActions#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Amazon DataZone environment action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#description DatazoneEnvironmentActions#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the Amazon DataZone domain in which the environment would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#domain_identifier DatazoneEnvironmentActions#domain_identifier}
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The identifier of the Amazon DataZone environment in which the action is taking place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#environment_identifier DatazoneEnvironmentActions#environment_identifier}
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The ID of the Amazon DataZone environment action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#identifier DatazoneEnvironmentActions#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The parameters of the environment action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_environment_actions#parameters DatazoneEnvironmentActions#parameters}
	Parameters *DatazoneEnvironmentActionsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

