// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsglobalcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsGlobalClusterConfig struct {
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
	// The deletion protection setting for the new global database.
	//
	// The global database can't be deleted when deletion protection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#deletion_protection RdsGlobalCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).
	// If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#engine RdsGlobalCluster#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The life cycle type of the global cluster.
	//
	// You can use this setting to enroll your global cluster into Amazon RDS Extended Support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#engine_lifecycle_support RdsGlobalCluster#engine_lifecycle_support}
	EngineLifecycleSupport *string `field:"optional" json:"engineLifecycleSupport" yaml:"engineLifecycleSupport"`
	// The version number of the database engine to use.
	//
	// If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#engine_version RdsGlobalCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#global_cluster_identifier RdsGlobalCluster#global_cluster_identifier}
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The Amazon Resource Name (ARN) to use as the primary cluster of the global database.
	//
	// This parameter is optional. This parameter is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#source_db_cluster_identifier RdsGlobalCluster#source_db_cluster_identifier}
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The storage encryption setting for the new global database cluster.
	//
	// If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#storage_encrypted RdsGlobalCluster#storage_encrypted}
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_global_cluster#tags RdsGlobalCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

