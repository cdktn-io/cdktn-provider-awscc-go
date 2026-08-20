// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbcluster


type TimestreamInfluxDbClusterLogDeliveryConfiguration struct {
	// S3 configuration for sending logs to customer account from the InfluxDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_influx_db_cluster#s3_configuration TimestreamInfluxDbCluster#s3_configuration}
	S3Configuration *TimestreamInfluxDbClusterLogDeliveryConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

