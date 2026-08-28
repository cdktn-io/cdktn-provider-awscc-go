// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptuneglobalcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptuneGlobalClusterConfig struct {
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
	// Whether deletion protection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#deletion_protection NeptuneGlobalCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The name of the database engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#engine NeptuneGlobalCluster#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version number of the database engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#engine_version NeptuneGlobalCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier of the global database cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#global_cluster_identifier NeptuneGlobalCluster#global_cluster_identifier}
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The Amazon Resource Name (ARN) of an existing Neptune DB cluster to use as the primary cluster of the new global database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#source_db_cluster_identifier NeptuneGlobalCluster#source_db_cluster_identifier}
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Whether the global database cluster is storage encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#storage_encrypted NeptuneGlobalCluster#storage_encrypted}
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/neptune_global_cluster#tags NeptuneGlobalCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

