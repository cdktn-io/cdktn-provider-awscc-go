// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2InstanceConfig struct {
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
	// This property is reserved for internal use.
	//
	// If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#additional_info Ec2Instance#additional_info}
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Indicates whether the instance is associated with a dedicated host.
	//
	// If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#affinity Ec2Instance#affinity}
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// The Availability Zone of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#availability_zone Ec2Instance#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The block device mapping entries that defines the block devices to attach to the instance at launch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#block_device_mappings Ec2Instance#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The CPU options for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#cpu_options Ec2Instance#cpu_options}
	CpuOptions *Ec2InstanceCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#credit_specification Ec2Instance#credit_specification}
	CreditSpecification *Ec2InstanceCreditSpecification `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#disable_api_termination Ec2Instance#disable_api_termination}
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#ebs_optimized Ec2Instance#ebs_optimized}
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// An elastic GPU to associate with the instance. Amazon Elastic Graphics is no longer available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#elastic_gpu_specifications Ec2Instance#elastic_gpu_specifications}
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance. Amazon Elastic Inference is no longer available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#elastic_inference_accelerators Ec2Instance#elastic_inference_accelerators}
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#enclave_options Ec2Instance#enclave_options}
	EnclaveOptions *Ec2InstanceEnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#hibernation_options Ec2Instance#hibernation_options}
	HibernationOptions *Ec2InstanceHibernationOptions `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with.
	//
	// If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#host_id Ec2Instance#host_id}
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// The ARN of the host resource group in which to launch the instances.
	//
	// If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#host_resource_group_arn Ec2Instance#host_resource_group_arn}
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// The IAM instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#iam_instance_profile Ec2Instance#iam_instance_profile}
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	//
	// An AMI ID is required to launch an instance and must be specified here or in a launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#image_id Ec2Instance#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#instance_initiated_shutdown_behavior Ec2Instance#instance_initiated_shutdown_behavior}
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#instance_type Ec2Instance#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#ipv_6_address_count Ec2Instance#ipv_6_address_count}
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#ipv_6_addresses Ec2Instance#ipv_6_addresses}
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The ID of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#kernel_id Ec2Instance#kernel_id}
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#key_name Ec2Instance#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The launch template to use to launch the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#launch_template Ec2Instance#launch_template}
	LaunchTemplate *Ec2InstanceLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The license configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#license_specifications Ec2Instance#license_specifications}
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The metadata options for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#metadata_options Ec2Instance#metadata_options}
	MetadataOptions *Ec2InstanceMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Specifies whether detailed monitoring is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#monitoring Ec2Instance#monitoring}
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The network interfaces to associate with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#network_interfaces Ec2Instance#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The name of an existing placement group that you want to launch the instance into (cluster | partition | spread).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#placement_group_name Ec2Instance#placement_group_name}
	PlacementGroupName *string `field:"optional" json:"placementGroupName" yaml:"placementGroupName"`
	// The options for the instance hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#private_dns_name_options Ec2Instance#private_dns_name_options}
	PrivateDnsNameOptions *Ec2InstancePrivateDnsNameOptions `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#private_ip_address Ec2Instance#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch.
	//
	// If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#propagate_tags_to_volume_on_creation Ec2Instance#propagate_tags_to_volume_on_creation}
	PropagateTagsToVolumeOnCreation interface{} `field:"optional" json:"propagateTagsToVolumeOnCreation" yaml:"propagateTagsToVolumeOnCreation"`
	// The ID of the RAM disk to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#ramdisk_id Ec2Instance#ramdisk_id}
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// The IDs of the security groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#security_group_ids Ec2Instance#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// the names of the security groups. For a nondefault VPC, you must use security group IDs instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#security_groups Ec2Instance#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#source_dest_check Ec2Instance#source_dest_check}
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// The SSM document and parameter values in AWS Systems Manager to associate with this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#ssm_associations Ec2Instance#ssm_associations}
	SsmAssociations interface{} `field:"optional" json:"ssmAssociations" yaml:"ssmAssociations"`
	// [EC2-VPC] The ID of the subnet to launch the instance into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#subnet_id Ec2Instance#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags to add to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#tags Ec2Instance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The tenancy of the instance (if the instance is running in a VPC).
	//
	// An instance with a tenancy of dedicated runs on single-tenant hardware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#tenancy Ec2Instance#tenancy}
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#user_data Ec2Instance#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// The volumes to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_instance#volumes Ec2Instance#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

