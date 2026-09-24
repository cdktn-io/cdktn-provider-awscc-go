// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package docdbdbclusterparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DocdbDbClusterParameterGroupConfig struct {
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
	// The description for the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/docdb_db_cluster_parameter_group#description DocdbDbClusterParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB cluster parameter group family name (e.g. docdb5.0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/docdb_db_cluster_parameter_group#family DocdbDbClusterParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// An object containing key-value pairs of parameters to set for the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/docdb_db_cluster_parameter_group#parameters DocdbDbClusterParameterGroup#parameters}
	Parameters *string `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the DB cluster parameter group.
	//
	// If omitted, CloudFormation generates a unique name. The name is stored as lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/docdb_db_cluster_parameter_group#name DocdbDbClusterParameterGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/docdb_db_cluster_parameter_group#tags DocdbDbClusterParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

