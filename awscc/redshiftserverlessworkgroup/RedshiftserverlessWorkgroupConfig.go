// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessworkgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftserverlessWorkgroupConfig struct {
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
	// The name of the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#workgroup_name RedshiftserverlessWorkgroup#workgroup_name}
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// The base compute capacity of the workgroup in Redshift Processing Units (RPUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#base_capacity RedshiftserverlessWorkgroup#base_capacity}
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// A list of parameters to set for finer control over a database.
	//
	// Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#config_parameters RedshiftserverlessWorkgroup#config_parameters}
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#enhanced_vpc_routing RedshiftserverlessWorkgroup#enhanced_vpc_routing}
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// The max compute capacity of the workgroup in Redshift Processing Units (RPUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#max_capacity RedshiftserverlessWorkgroup#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The namespace the workgroup is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#namespace_name RedshiftserverlessWorkgroup#namespace_name}
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The custom port to use when connecting to a workgroup.
	//
	// Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#port RedshiftserverlessWorkgroup#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A property that represents the price performance target settings for the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#price_performance_target RedshiftserverlessWorkgroup#price_performance_target}
	PricePerformanceTarget *RedshiftserverlessWorkgroupPricePerformanceTarget `field:"optional" json:"pricePerformanceTarget" yaml:"pricePerformanceTarget"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#publicly_accessible RedshiftserverlessWorkgroup#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The identifier of the recovery point to restore the namespace from.
	//
	// When this resource is first created, the namespace is restored from this recovery point. On subsequent updates, a restore occurs only when RecoveryPointId changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#recovery_point_id RedshiftserverlessWorkgroup#recovery_point_id}
	RecoveryPointId *string `field:"optional" json:"recoveryPointId" yaml:"recoveryPointId"`
	// A list of security group IDs to associate with the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#security_group_ids RedshiftserverlessWorkgroup#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) of the snapshot to restore the namespace from.
	//
	// Specify either SnapshotArn or SnapshotName, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotArn changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#snapshot_arn RedshiftserverlessWorkgroup#snapshot_arn}
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// The name of the snapshot to restore the namespace from.
	//
	// Because snapshot names are unique only within an account, also specify SnapshotOwnerAccount when restoring from a snapshot owned by a different account. Specify either SnapshotName or SnapshotArn, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotName or SnapshotOwnerAccount changes from its previous value. If both values are unchanged or SnapshotName is removed, no restore takes place and existing data is preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#snapshot_name RedshiftserverlessWorkgroup#snapshot_name}
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The AWS account ID that owns the snapshot.
	//
	// Required when restoring from a snapshot shared by another account. Used in combination with SnapshotName. On updates, changing this value while SnapshotName is set triggers a restore from the newly referenced snapshot. If the value is unchanged, no restore takes place and existing data is preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#snapshot_owner_account RedshiftserverlessWorkgroup#snapshot_owner_account}
	SnapshotOwnerAccount *string `field:"optional" json:"snapshotOwnerAccount" yaml:"snapshotOwnerAccount"`
	// A list of subnet IDs the workgroup is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#subnet_ids RedshiftserverlessWorkgroup#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The map of the key-value pairs used to tag the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#tags RedshiftserverlessWorkgroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#track_name RedshiftserverlessWorkgroup#track_name}.
	TrackName *string `field:"optional" json:"trackName" yaml:"trackName"`
	// Definition for workgroup resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#workgroup RedshiftserverlessWorkgroup#workgroup}
	Workgroup *RedshiftserverlessWorkgroupWorkgroup `field:"optional" json:"workgroup" yaml:"workgroup"`
}

