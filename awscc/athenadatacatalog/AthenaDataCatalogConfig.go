// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenadatacatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaDataCatalogConfig struct {
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
	// The name of the data catalog to create.
	//
	// The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#name AthenaDataCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	//
	// FEDERATED is a federated catalog for which Athena creates the connection and the Lambda function for you based on the parameters that you pass.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#type AthenaDataCatalog#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The type of connection for a FEDERATED data catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#connection_type AthenaDataCatalog#connection_type}
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// A description of the data catalog to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#description AthenaDataCatalog#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Text of the error that occurred during data catalog creation or deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#error AthenaDataCatalog#error}
	Error *string `field:"optional" json:"error" yaml:"error"`
	// Specifies the Lambda function or functions to use for creating the data catalog.
	//
	// This is a mapping whose values depend on the catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#parameters AthenaDataCatalog#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The status of the creation or deletion of the data catalog.
	//
	// LAMBDA, GLUE, and HIVE data catalog types are created synchronously. Their status is either CREATE_COMPLETE or CREATE_FAILED. The FEDERATED data catalog type is created asynchronously.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#status AthenaDataCatalog#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of comma separated tags to add to the data catalog that is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_data_catalog#tags AthenaDataCatalog#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

