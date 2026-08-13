// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionConfig struct {
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
	// The identifier of the domain in which the connection is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#domain_identifier DatazoneConnection#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#name DatazoneConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS Location of project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#aws_location DatazoneConnection#aws_location}
	AwsLocation *DatazoneConnectionAwsLocation `field:"optional" json:"awsLocation" yaml:"awsLocation"`
	// The configurations of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#configurations DatazoneConnection#configurations}
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The description of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#description DatazoneConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the trusted identity propagation is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#enable_trusted_identity_propagation DatazoneConnection#enable_trusted_identity_propagation}
	EnableTrustedIdentityPropagation interface{} `field:"optional" json:"enableTrustedIdentityPropagation" yaml:"enableTrustedIdentityPropagation"`
	// The identifier of the environment in which the connection is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#environment_identifier DatazoneConnection#environment_identifier}
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The identifier of the project in which the connection should be created. If.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#project_identifier DatazoneConnection#project_identifier}
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#props DatazoneConnection#props}.
	Props *DatazoneConnectionProps `field:"optional" json:"props" yaml:"props"`
	// The scope of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_connection#scope DatazoneConnection#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

