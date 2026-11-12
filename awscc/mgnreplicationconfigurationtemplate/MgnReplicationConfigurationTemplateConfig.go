// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnreplicationconfigurationtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MgnReplicationConfigurationTemplateConfig struct {
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
	// Whether to associate the default Application Migration Service Security Group with the Replication Configuration Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#associate_default_security_group MgnReplicationConfigurationTemplate#associate_default_security_group}
	AssociateDefaultSecurityGroup interface{} `field:"required" json:"associateDefaultSecurityGroup" yaml:"associateDefaultSecurityGroup"`
	// Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#bandwidth_throttling MgnReplicationConfigurationTemplate#bandwidth_throttling}
	BandwidthThrottling *float64 `field:"required" json:"bandwidthThrottling" yaml:"bandwidthThrottling"`
	// Whether to create a Public IP for the Recovery Instance by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#create_public_ip MgnReplicationConfigurationTemplate#create_public_ip}
	CreatePublicIp interface{} `field:"required" json:"createPublicIp" yaml:"createPublicIp"`
	// The data plane routing mechanism that will be used for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#data_plane_routing MgnReplicationConfigurationTemplate#data_plane_routing}
	DataPlaneRouting *string `field:"required" json:"dataPlaneRouting" yaml:"dataPlaneRouting"`
	// The Staging Disk EBS volume type to be used during replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#default_large_staging_disk_type MgnReplicationConfigurationTemplate#default_large_staging_disk_type}
	DefaultLargeStagingDiskType *string `field:"required" json:"defaultLargeStagingDiskType" yaml:"defaultLargeStagingDiskType"`
	// The type of EBS encryption to be used during replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#ebs_encryption MgnReplicationConfigurationTemplate#ebs_encryption}
	EbsEncryption *string `field:"required" json:"ebsEncryption" yaml:"ebsEncryption"`
	// The instance type to be used for the replication server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#replication_server_instance_type MgnReplicationConfigurationTemplate#replication_server_instance_type}
	ReplicationServerInstanceType *string `field:"required" json:"replicationServerInstanceType" yaml:"replicationServerInstanceType"`
	// The security group IDs that will be used by the replication server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#replication_servers_security_groups_i_ds MgnReplicationConfigurationTemplate#replication_servers_security_groups_i_ds}
	ReplicationServersSecurityGroupsIDs *[]*string `field:"required" json:"replicationServersSecurityGroupsIDs" yaml:"replicationServersSecurityGroupsIDs"`
	// The subnet to be used by the replication staging area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#staging_area_subnet_id MgnReplicationConfigurationTemplate#staging_area_subnet_id}
	StagingAreaSubnetId *string `field:"required" json:"stagingAreaSubnetId" yaml:"stagingAreaSubnetId"`
	// A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#staging_area_tags MgnReplicationConfigurationTemplate#staging_area_tags}
	StagingAreaTags *map[string]*string `field:"required" json:"stagingAreaTags" yaml:"stagingAreaTags"`
	// Whether to use a dedicated Replication Server in the replication staging area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#use_dedicated_replication_server MgnReplicationConfigurationTemplate#use_dedicated_replication_server}
	UseDedicatedReplicationServer interface{} `field:"required" json:"useDedicatedReplicationServer" yaml:"useDedicatedReplicationServer"`
	// The ARN of the EBS encryption key to be used during replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#ebs_encryption_key_arn MgnReplicationConfigurationTemplate#ebs_encryption_key_arn}
	EbsEncryptionKeyArn *string `field:"optional" json:"ebsEncryptionKeyArn" yaml:"ebsEncryptionKeyArn"`
	// Which internet protocol to use for replication of data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#internet_protocol MgnReplicationConfigurationTemplate#internet_protocol}
	InternetProtocol *string `field:"optional" json:"internetProtocol" yaml:"internetProtocol"`
	// Whether to store the Snapshot in the Local Zone of the replication staging area subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#store_snapshot_on_local_zone MgnReplicationConfigurationTemplate#store_snapshot_on_local_zone}
	StoreSnapshotOnLocalZone interface{} `field:"optional" json:"storeSnapshotOnLocalZone" yaml:"storeSnapshotOnLocalZone"`
	// A set of tags to be associated with the Replication Configuration Template resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#tags MgnReplicationConfigurationTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Whether to use a FIPS endpoint when initiating the agent connection to the Application Migration Service service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template#use_fips_endpoint MgnReplicationConfigurationTemplate#use_fips_endpoint}
	UseFipsEndpoint interface{} `field:"optional" json:"useFipsEndpoint" yaml:"useFipsEndpoint"`
}

