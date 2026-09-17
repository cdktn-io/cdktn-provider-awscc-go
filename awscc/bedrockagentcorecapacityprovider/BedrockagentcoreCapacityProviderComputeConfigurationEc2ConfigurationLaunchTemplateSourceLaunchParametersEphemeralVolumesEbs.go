// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParametersEphemeralVolumesEbs struct {
	// The index of the EBS card. Applies to instances with multiple EBS cards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#ebs_card_index BedrockagentcoreCapacityProvider#ebs_card_index}
	EbsCardIndex *float64 `field:"optional" json:"ebsCardIndex" yaml:"ebsCardIndex"`
	// Indicates whether the EBS volume is encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#encrypted BedrockagentcoreCapacityProvider#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#iops BedrockagentcoreCapacityProvider#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier of the customer managed KMS key to use for EBS encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#kms_key_id BedrockagentcoreCapacityProvider#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ID of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#snapshot_id BedrockagentcoreCapacityProvider#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput to provision for a gp3 volume, in MiB/s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#throughput BedrockagentcoreCapacityProvider#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The rate at which the volume is initialized after creation, in MiB/s.
	//
	// Supported only for volumes created from snapshots. If the snapshot is enabled for fast snapshot restore and a volume initialization rate is also specified, the volume is initialized at the specified rate instead of by fast snapshot restore. Valid range: 100-300 MiB/s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#volume_initialization_rate BedrockagentcoreCapacityProvider#volume_initialization_rate}
	VolumeInitializationRate *float64 `field:"optional" json:"volumeInitializationRate" yaml:"volumeInitializationRate"`
	// The size of the volume, in GiBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#volume_size BedrockagentcoreCapacityProvider#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type. Defaults to gp3 if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#volume_type BedrockagentcoreCapacityProvider#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

