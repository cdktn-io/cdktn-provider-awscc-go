// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessDataSourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#application_id QbusinessDataSource#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#configuration QbusinessDataSource#configuration}.
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#display_name QbusinessDataSource#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#index_id QbusinessDataSource#index_id}.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#description QbusinessDataSource#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#document_enrichment_configuration QbusinessDataSource#document_enrichment_configuration}.
	DocumentEnrichmentConfiguration *QbusinessDataSourceDocumentEnrichmentConfiguration `field:"optional" json:"documentEnrichmentConfiguration" yaml:"documentEnrichmentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#media_extraction_configuration QbusinessDataSource#media_extraction_configuration}.
	MediaExtractionConfiguration *QbusinessDataSourceMediaExtractionConfiguration `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#role_arn QbusinessDataSource#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#sync_schedule QbusinessDataSource#sync_schedule}.
	SyncSchedule *string `field:"optional" json:"syncSchedule" yaml:"syncSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#tags QbusinessDataSource#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#vpc_configuration QbusinessDataSource#vpc_configuration}.
	VpcConfiguration *QbusinessDataSourceVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

