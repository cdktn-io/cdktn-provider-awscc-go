// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsanalysistemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsAnalysisTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#format CleanroomsAnalysisTemplate#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#membership_identifier CleanroomsAnalysisTemplate#membership_identifier}.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#name CleanroomsAnalysisTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#source CleanroomsAnalysisTemplate#source}.
	Source *CleanroomsAnalysisTemplateSource `field:"required" json:"source" yaml:"source"`
	// The member who can query can provide this placeholder for a literal data value in an analysis template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#analysis_parameters CleanroomsAnalysisTemplate#analysis_parameters}
	AnalysisParameters interface{} `field:"optional" json:"analysisParameters" yaml:"analysisParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#description CleanroomsAnalysisTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#error_message_configuration CleanroomsAnalysisTemplate#error_message_configuration}.
	ErrorMessageConfiguration *CleanroomsAnalysisTemplateErrorMessageConfiguration `field:"optional" json:"errorMessageConfiguration" yaml:"errorMessageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#schema CleanroomsAnalysisTemplate#schema}.
	Schema *CleanroomsAnalysisTemplateSchema `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#source_metadata CleanroomsAnalysisTemplate#source_metadata}.
	SourceMetadata *CleanroomsAnalysisTemplateSourceMetadata `field:"optional" json:"sourceMetadata" yaml:"sourceMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#synthetic_data_parameters CleanroomsAnalysisTemplate#synthetic_data_parameters}.
	SyntheticDataParameters *CleanroomsAnalysisTemplateSyntheticDataParameters `field:"optional" json:"syntheticDataParameters" yaml:"syntheticDataParameters"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_analysis_template#tags CleanroomsAnalysisTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

