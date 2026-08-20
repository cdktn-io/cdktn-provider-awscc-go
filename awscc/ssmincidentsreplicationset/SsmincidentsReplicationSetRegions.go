// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsreplicationset


type SsmincidentsReplicationSetRegions struct {
	// The ReplicationSet regional configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmincidents_replication_set#region_configuration SsmincidentsReplicationSet#region_configuration}
	RegionConfiguration *SsmincidentsReplicationSetRegionsRegionConfiguration `field:"optional" json:"regionConfiguration" yaml:"regionConfiguration"`
	// The AWS region name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmincidents_replication_set#region_name SsmincidentsReplicationSet#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

