// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateData struct {
	// The block device mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#block_device_mappings Ec2LaunchTemplate#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The Capacity Reservation targeting option.
	//
	// If you do not specify this parameter, the instance's Capacity Reservation preference defaults to ``open``, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#capacity_reservation_specification Ec2LaunchTemplate#capacity_reservation_specification}
	CapacityReservationSpecification *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// The CPU options for the instance. For more information, see [Optimize CPU options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#cpu_options Ec2LaunchTemplate#cpu_options}
	CpuOptions *Ec2LaunchTemplateLaunchTemplateDataCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The credit option for CPU usage of the instance. Valid only for T instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#credit_specification Ec2LaunchTemplate#credit_specification}
	CreditSpecification *Ec2LaunchTemplateLaunchTemplateDataCreditSpecification `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Indicates whether to enable the instance for stop protection.
	//
	// For more information, see [Enable stop protection for your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-stop-protection.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#disable_api_stop Ec2LaunchTemplate#disable_api_stop}
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// If you set this parameter to ``true``, you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set ``InstanceInitiatedShutdownBehavior`` to ``terminate``, you can terminate the instance by running the shutdown command from the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#disable_api_termination Ec2LaunchTemplate#disable_api_termination}
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#ebs_optimized Ec2LaunchTemplate#ebs_optimized}
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Deprecated.
	//
	// Amazon Elastic Graphics reached end of life on January 8, 2024. For workloads that require graphics acceleration, we recommend that you use Amazon EC2 G4ad, G4dn, or G5 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#elastic_gpu_specifications Ec2LaunchTemplate#elastic_gpu_specifications}
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance.
	//
	// Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.
	//  You cannot specify accelerators from different generations in the same request.
	//   Starting April 15, 2023, AWS will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#elastic_inference_accelerators Ec2LaunchTemplate#elastic_inference_accelerators}
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	//
	// For more information, see [What is Nitro Enclaves?](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the *Nitro Enclaves User Guide*.
	//  You can't enable AWS Nitro Enclaves and hibernation on the same instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#enclave_options Ec2LaunchTemplate#enclave_options}
	EnclaveOptions *Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	//
	// This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html). For more information, see [Hibernate your Amazon EC2 instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#hibernation_options Ec2LaunchTemplate#hibernation_options}
	HibernationOptions *Ec2LaunchTemplateLaunchTemplateDataHibernationOptions `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// The name or Amazon Resource Name (ARN) of an IAM instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#iam_instance_profile Ec2LaunchTemplate#iam_instance_profile}
	IamInstanceProfile *Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	//
	// Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.
	//  Valid formats:
	//   +   ``ami-0ac394d6a3example``
	//   +   ``resolve:ssm:parameter-name``
	//   +   ``resolve:ssm:parameter-name:version-number``
	//   +   ``resolve:ssm:parameter-name:label``
	//
	//  For more information, see [Use a Systems Manager parameter to find an AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html#using-systems-manager-parameter-to-find-AMI) in the *Amazon Elastic Compute Cloud User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#image_id Ec2LaunchTemplate#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Default: ``stop``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#instance_initiated_shutdown_behavior Ec2LaunchTemplate#instance_initiated_shutdown_behavior}
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#instance_market_options Ec2LaunchTemplate#instance_market_options}
	InstanceMarketOptions *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with these attributes.
	//  You must specify ``VCpuCount`` and ``MemoryMiB``. All other attributes are optional. Any unspecified optional attribute is set to its default.
	//  When you specify multiple attributes, you get instance types that satisfy all of the specified attributes. If you specify multiple values for an attribute, you get instance types that satisfy any of the specified values.
	//  To limit the list of instance types from which Amazon EC2 can identify matching instance types, you can use one of the following parameters, but not both in the same request:
	//   +   ``AllowedInstanceTypes`` - The instance types to include in the list. All other instance types are ignored, even if they match your specified attributes.
	//   +   ``ExcludedInstanceTypes`` - The instance types to exclude from the list, even if they match your specified attributes.
	//
	//   If you specify ``InstanceRequirements``, you can't specify ``InstanceType``.
	//  Attribute-based instance type selection is only supported when using Auto Scaling groups, EC2 Fleet, and Spot Fleet to launch instances. If you plan to use the launch template in the [launch instance wizard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-instance-wizard.html), or with the [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) API or [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) AWS CloudFormation resource, you can't specify ``InstanceRequirements``.
	//   For more information, see [Attribute-based instance type selection for EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html), [Attribute-based instance type selection for Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html), and [Spot placement score](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#instance_requirements Ec2LaunchTemplate#instance_requirements}
	InstanceRequirements *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type.
	//
	// For more information, see [Amazon EC2 instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide*.
	//  If you specify ``InstanceType``, you can't specify ``InstanceRequirements``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#instance_type Ec2LaunchTemplate#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User Provided Kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#kernel_id Ec2LaunchTemplate#kernel_id}
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair.
	//
	// You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).
	//   If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#key_name Ec2LaunchTemplate#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The license configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#license_specifications Ec2LaunchTemplate#license_specifications}
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The maintenance options of your instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#maintenance_options Ec2LaunchTemplate#maintenance_options}
	MaintenanceOptions *Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// The metadata options for the instance.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#metadata_options Ec2LaunchTemplate#metadata_options}
	MetadataOptions *Ec2LaunchTemplateLaunchTemplateDataMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// The monitoring for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#monitoring Ec2LaunchTemplate#monitoring}
	Monitoring *Ec2LaunchTemplateLaunchTemplateDataMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The network interfaces for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#network_interfaces Ec2LaunchTemplate#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The placement for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#placement Ec2LaunchTemplate#placement}
	Placement *Ec2LaunchTemplateLaunchTemplateDataPlacement `field:"optional" json:"placement" yaml:"placement"`
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries should be handled.
	//
	// For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#private_dns_name_options Ec2LaunchTemplate#private_dns_name_options}
	PrivateDnsNameOptions *Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// The ID of the RAM disk.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#ram_disk_id Ec2LaunchTemplate#ram_disk_id}
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// The IDs of the security groups.
	//
	// You can specify the IDs of existing security groups and references to resources created by the stack template.
	//  If you specify a network interface, you must specify any security groups as part of the network interface instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#security_group_ids Ec2LaunchTemplate#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The names of the security groups.
	//
	// For a nondefault VPC, you must use security group IDs instead.
	//  If you specify a network interface, you must specify any security groups as part of the network interface instead of using this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#security_groups Ec2LaunchTemplate#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The tags to apply to the resources that are created during instance launch.
	//
	// To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	//  To tag the launch template itself, use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#tag_specifications Ec2LaunchTemplate#tag_specifications}
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The user data to make available to the instance.
	//
	// You must provide base64-encoded text. User data is limited to 16 KB. For more information, see [Run commands on your Amazon EC2 instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) in the *Amazon EC2 User Guide*.
	//  If you are creating the launch template for use with BATCH, the user data must be provided in the [MIME multi-part archive format](https://docs.aws.amazon.com/https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive). For more information, see [Amazon EC2 user data in launch templates](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_launch_template#user_data Ec2LaunchTemplate#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

