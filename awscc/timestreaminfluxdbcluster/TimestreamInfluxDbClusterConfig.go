// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TimestreamInfluxDbClusterConfig struct {
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
	// The allocated storage for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#allocated_storage TimestreamInfluxDbCluster#allocated_storage}
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The bucket for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#bucket TimestreamInfluxDbCluster#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The compute instance of the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#db_instance_type TimestreamInfluxDbCluster#db_instance_type}
	DbInstanceType *string `field:"optional" json:"dbInstanceType" yaml:"dbInstanceType"`
	// The name of an existing InfluxDB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#db_parameter_group_identifier TimestreamInfluxDbCluster#db_parameter_group_identifier}
	DbParameterGroupIdentifier *string `field:"optional" json:"dbParameterGroupIdentifier" yaml:"dbParameterGroupIdentifier"`
	// The storage type of the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#db_storage_type TimestreamInfluxDbCluster#db_storage_type}
	DbStorageType *string `field:"optional" json:"dbStorageType" yaml:"dbStorageType"`
	// Deployment type of the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#deployment_type TimestreamInfluxDbCluster#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Failover mode of the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#failover_mode TimestreamInfluxDbCluster#failover_mode}
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// Configuration for sending logs to customer account from the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#log_delivery_configuration TimestreamInfluxDbCluster#log_delivery_configuration}
	LogDeliveryConfiguration *TimestreamInfluxDbClusterLogDeliveryConfiguration `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// The maintenance schedule for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#maintenance_schedule TimestreamInfluxDbCluster#maintenance_schedule}
	MaintenanceSchedule *TimestreamInfluxDbClusterMaintenanceSchedule `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// The unique name that is associated with the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#name TimestreamInfluxDbCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Network type of the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#network_type TimestreamInfluxDbCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The organization for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#organization TimestreamInfluxDbCluster#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The password for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#password TimestreamInfluxDbCluster#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port number on which InfluxDB accepts connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#port TimestreamInfluxDbCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Attach a public IP to the customer ENI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#publicly_accessible TimestreamInfluxDbCluster#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An arbitrary set of tags (key-value pairs) for this DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#tags TimestreamInfluxDbCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The username for the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#username TimestreamInfluxDbCluster#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#vpc_security_group_ids TimestreamInfluxDbCluster#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A list of EC2 subnet IDs for this InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/timestream_influx_db_cluster#vpc_subnet_ids TimestreamInfluxDbCluster#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

