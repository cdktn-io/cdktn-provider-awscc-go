// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataTelemetrySinkConfigTelemetrySinkDataKinesisDataStreamData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#kinesis_data_stream_arn GroundstationConfig#kinesis_data_stream_arn}.
	KinesisDataStreamArn *string `field:"optional" json:"kinesisDataStreamArn" yaml:"kinesisDataStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#kinesis_role_arn GroundstationConfig#kinesis_role_arn}.
	KinesisRoleArn *string `field:"optional" json:"kinesisRoleArn" yaml:"kinesisRoleArn"`
}

