// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationVolumesEbsConfiguration struct {
	// Whether to encrypt the volume. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#encrypted BedrockagentcoreCapacityProvider#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of IOPS to provision. Only valid for gp3, io1, and io2 volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#iops BedrockagentcoreCapacityProvider#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#kms_key_id BedrockagentcoreCapacityProvider#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The logical name of the volume, used to reference it when mounting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#name BedrockagentcoreCapacityProvider#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The size of the volume in GiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#size_gi_b BedrockagentcoreCapacityProvider#size_gi_b}
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// Optional EBS snapshot ID to initialize the volume from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#snapshot_id BedrockagentcoreCapacityProvider#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput in MiB/s. Only valid for gp3 volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#throughput BedrockagentcoreCapacityProvider#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type. Defaults to gp3 if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_capacity_provider#volume_type BedrockagentcoreCapacityProvider#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

