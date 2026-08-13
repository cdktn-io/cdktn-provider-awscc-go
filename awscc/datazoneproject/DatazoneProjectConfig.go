// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneproject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneProjectConfig struct {
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
	// The ID of the Amazon DataZone domain in which this project is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#domain_identifier DatazoneProject#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the Amazon DataZone project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#name DatazoneProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Amazon DataZone project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#description DatazoneProject#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the domain unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#domain_unit_id DatazoneProject#domain_unit_id}
	DomainUnitId *string `field:"optional" json:"domainUnitId" yaml:"domainUnitId"`
	// The glossary terms that can be used in this Amazon DataZone project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#glossary_terms DatazoneProject#glossary_terms}
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// The project membership assignments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#membership_assignments DatazoneProject#membership_assignments}
	MembershipAssignments interface{} `field:"optional" json:"membershipAssignments" yaml:"membershipAssignments"`
	// The project category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#project_category DatazoneProject#project_category}
	ProjectCategory *string `field:"optional" json:"projectCategory" yaml:"projectCategory"`
	// The project execution role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#project_execution_role DatazoneProject#project_execution_role}
	ProjectExecutionRole *string `field:"optional" json:"projectExecutionRole" yaml:"projectExecutionRole"`
	// The project profile ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#project_profile_id DatazoneProject#project_profile_id}
	ProjectProfileId *string `field:"optional" json:"projectProfileId" yaml:"projectProfileId"`
	// The project profile version to which the project should be updated.
	//
	// You can only specify the following string for this parameter: latest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#project_profile_version DatazoneProject#project_profile_version}
	ProjectProfileVersion *string `field:"optional" json:"projectProfileVersion" yaml:"projectProfileVersion"`
	// The resource tags of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#resource_tags DatazoneProject#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The user parameters of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datazone_project#user_parameters DatazoneProject#user_parameters}
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

