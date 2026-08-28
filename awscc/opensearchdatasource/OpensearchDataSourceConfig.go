// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdatasource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchDataSourceConfig struct {
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
	// The type of data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearch_data_source#data_source_type OpensearchDataSource#data_source_type}
	DataSourceType *OpensearchDataSourceDataSourceType `field:"required" json:"dataSourceType" yaml:"dataSourceType"`
	// The name of the OpenSearch Service domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearch_data_source#domain_name OpensearchDataSource#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearch_data_source#name OpensearchDataSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearch_data_source#description OpensearchDataSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

