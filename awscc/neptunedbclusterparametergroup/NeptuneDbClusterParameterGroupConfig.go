// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunedbclusterparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptuneDbClusterParameterGroupConfig struct {
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
	// Provides the customer-specified description for this DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_db_cluster_parameter_group#description NeptuneDbClusterParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_db_cluster_parameter_group#family NeptuneDbClusterParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_db_cluster_parameter_group#parameters NeptuneDbClusterParameterGroup#parameters}
	Parameters *string `field:"required" json:"parameters" yaml:"parameters"`
	// Provides the name of the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_db_cluster_parameter_group#name NeptuneDbClusterParameterGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of tags for the cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_db_cluster_parameter_group#tags NeptuneDbClusterParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

