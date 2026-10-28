// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDeploymentConfig struct {
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
	// The name of the Blue/Green Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#blue_green_deployment_name RdsDeployment#blue_green_deployment_name}
	BlueGreenDeploymentName *string `field:"required" json:"blueGreenDeploymentName" yaml:"blueGreenDeploymentName"`
	// The ARN of the source production database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#source RdsDeployment#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The tags assigned to the Blue/Green Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#tags RdsDeployment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The allocated storage size in GiB for the target DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_allocated_storage RdsDeployment#target_allocated_storage}
	TargetAllocatedStorage *float64 `field:"optional" json:"targetAllocatedStorage" yaml:"targetAllocatedStorage"`
	// The DB cluster parameter group for the target Aurora DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_db_cluster_parameter_group_name RdsDeployment#target_db_cluster_parameter_group_name}
	TargetDbClusterParameterGroupName *string `field:"optional" json:"targetDbClusterParameterGroupName" yaml:"targetDbClusterParameterGroupName"`
	// The DB instance class for target RDS DB instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_db_instance_class RdsDeployment#target_db_instance_class}
	TargetDbInstanceClass *string `field:"optional" json:"targetDbInstanceClass" yaml:"targetDbInstanceClass"`
	// The DB parameter group for the target DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_db_parameter_group_name RdsDeployment#target_db_parameter_group_name}
	TargetDbParameterGroupName *string `field:"optional" json:"targetDbParameterGroupName" yaml:"targetDbParameterGroupName"`
	// The DB engine version for the target database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_engine_version RdsDeployment#target_engine_version}
	TargetEngineVersion *string `field:"optional" json:"targetEngineVersion" yaml:"targetEngineVersion"`
	// The Provisioned IOPS value for the target DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_iops RdsDeployment#target_iops}
	TargetIops *float64 `field:"optional" json:"targetIops" yaml:"targetIops"`
	// The storage throughput value for the target DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_storage_throughput RdsDeployment#target_storage_throughput}
	TargetStorageThroughput *float64 `field:"optional" json:"targetStorageThroughput" yaml:"targetStorageThroughput"`
	// The storage type for the target DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#target_storage_type RdsDeployment#target_storage_type}
	TargetStorageType *string `field:"optional" json:"targetStorageType" yaml:"targetStorageType"`
	// Whether to upgrade the storage file system configuration on the target database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment#upgrade_target_storage_config RdsDeployment#upgrade_target_storage_config}
	UpgradeTargetStorageConfig interface{} `field:"optional" json:"upgradeTargetStorageConfig" yaml:"upgradeTargetStorageConfig"`
}

