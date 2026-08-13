// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbinstance


type TimestreamInfluxDbInstanceLogDeliveryConfigurationS3Configuration struct {
	// The bucket name for logs to be sent from the InfluxDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#bucket_name TimestreamInfluxDbInstance#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies whether logging to customer specified bucket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_influx_db_instance#enabled TimestreamInfluxDbInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

