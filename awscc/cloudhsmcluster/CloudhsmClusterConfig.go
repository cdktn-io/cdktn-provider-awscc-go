// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudhsmcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudhsmClusterConfig struct {
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
	// The type of HSM to use in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#hsm_type CloudhsmCluster#hsm_type}
	HsmType *string `field:"required" json:"hsmType" yaml:"hsmType"`
	// A policy that defines how the service retains backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#backup_retention_policy CloudhsmCluster#backup_retention_policy}
	BackupRetentionPolicy *CloudhsmClusterBackupRetentionPolicy `field:"optional" json:"backupRetentionPolicy" yaml:"backupRetentionPolicy"`
	// The mode to use in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#mode CloudhsmCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The NetworkType to create a cluster with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#network_type CloudhsmCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The identifiers (IDs) of the subnets where the cluster is created. You must specify at least one subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#subnet_ids CloudhsmCluster#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to apply to the CloudHSM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudhsm_cluster#tags CloudhsmCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

