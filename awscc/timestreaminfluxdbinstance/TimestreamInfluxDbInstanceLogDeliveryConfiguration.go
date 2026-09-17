// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbinstance


type TimestreamInfluxDbInstanceLogDeliveryConfiguration struct {
	// S3 configuration for sending logs to customer account from the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_influx_db_instance#s3_configuration TimestreamInfluxDbInstance#s3_configuration}
	S3Configuration *TimestreamInfluxDbInstanceLogDeliveryConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

