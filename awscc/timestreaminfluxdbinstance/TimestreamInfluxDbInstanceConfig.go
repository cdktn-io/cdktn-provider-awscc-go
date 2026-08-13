// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TimestreamInfluxDbInstanceConfig struct {
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
	// The allocated storage for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#allocated_storage TimestreamInfluxDbInstance#allocated_storage}
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The bucket for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#bucket TimestreamInfluxDbInstance#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The compute instance of the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#db_instance_type TimestreamInfluxDbInstance#db_instance_type}
	DbInstanceType *string `field:"optional" json:"dbInstanceType" yaml:"dbInstanceType"`
	// The name of an existing InfluxDB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#db_parameter_group_identifier TimestreamInfluxDbInstance#db_parameter_group_identifier}
	DbParameterGroupIdentifier *string `field:"optional" json:"dbParameterGroupIdentifier" yaml:"dbParameterGroupIdentifier"`
	// The storage type of the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#db_storage_type TimestreamInfluxDbInstance#db_storage_type}
	DbStorageType *string `field:"optional" json:"dbStorageType" yaml:"dbStorageType"`
	// Deployment type of the InfluxDB Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#deployment_type TimestreamInfluxDbInstance#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Configuration for sending logs to customer account from the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#log_delivery_configuration TimestreamInfluxDbInstance#log_delivery_configuration}
	LogDeliveryConfiguration *TimestreamInfluxDbInstanceLogDeliveryConfiguration `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// The maintenance schedule for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#maintenance_schedule TimestreamInfluxDbInstance#maintenance_schedule}
	MaintenanceSchedule *TimestreamInfluxDbInstanceMaintenanceSchedule `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// The unique name that is associated with the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#name TimestreamInfluxDbInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Network type of the InfluxDB Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#network_type TimestreamInfluxDbInstance#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The organization for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#organization TimestreamInfluxDbInstance#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The password for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#password TimestreamInfluxDbInstance#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port number on which InfluxDB accepts connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#port TimestreamInfluxDbInstance#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Attach a public IP to the customer ENI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#publicly_accessible TimestreamInfluxDbInstance#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#tags TimestreamInfluxDbInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The username for the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#username TimestreamInfluxDbInstance#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#vpc_security_group_ids TimestreamInfluxDbInstance#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A list of EC2 subnet IDs for this InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#vpc_subnet_ids TimestreamInfluxDbInstance#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

