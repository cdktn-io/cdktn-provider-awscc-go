// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on an instance.
	//
	// To exclude accelerator-enabled instance types, set ``Max`` to ``0``.
	//  Default: No minimum or maximum limits
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#accelerator_count Ec2LaunchTemplate#accelerator_count}
	AcceleratorCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// +  For instance types with AWS devices, specify ``amazon-web-services``.
	//   +  For instance types with AMD devices, specify ``amd``.
	//   +  For instance types with Habana devices, specify ``habana``.
	//   +  For instance types with NVIDIA devices, specify ``nvidia``.
	//   +  For instance types with Xilinx devices, specify ``xilinx``.
	//
	//  Default: Any manufacturer
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#accelerator_manufacturers Ec2LaunchTemplate#accelerator_manufacturers}
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// The accelerators that must be on the instance type.
	//
	// +  For instance types with NVIDIA A10G GPUs, specify ``a10g``.
	//   +  For instance types with NVIDIA A100 GPUs, specify ``a100``.
	//   +  For instance types with NVIDIA H100 GPUs, specify ``h100``.
	//   +  For instance types with AWS Inferentia chips, specify ``inferentia``.
	//   +  For instance types with NVIDIA GRID K520 GPUs, specify ``k520``.
	//   +  For instance types with NVIDIA K80 GPUs, specify ``k80``.
	//   +  For instance types with NVIDIA M60 GPUs, specify ``m60``.
	//   +  For instance types with AMD Radeon Pro V520 GPUs, specify ``radeon-pro-v520``.
	//   +  For instance types with NVIDIA T4 GPUs, specify ``t4``.
	//   +  For instance types with NVIDIA T4G GPUs, specify ``t4g``.
	//   +  For instance types with Xilinx VU9P FPGAs, specify ``vu9p``.
	//   +  For instance types with NVIDIA V100 GPUs, specify ``v100``.
	//
	//  Default: Any accelerator
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#accelerator_names Ec2LaunchTemplate#accelerator_names}
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The minimum and maximum amount of total accelerator memory, in MiB.  Default: No minimum or maximum limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#accelerator_total_memory_mi_b Ec2LaunchTemplate#accelerator_total_memory_mi_b}
	AcceleratorTotalMemoryMiB *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiB `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// The accelerator types that must be on the instance type.
	//
	// +  For instance types with GPU accelerators, specify ``gpu``.
	//   +  For instance types with FPGA accelerators, specify ``fpga``.
	//   +  For instance types with inference accelerators, specify ``inference``.
	//
	//  Default: Any accelerator type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#accelerator_types Ec2LaunchTemplate#accelerator_types}
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The instance types to apply your specified attributes against.
	//
	// All other instance types are ignored, even if they match your specified attributes.
	//  You can use strings with one or more wild cards, represented by an asterisk (``*``), to allow an instance type, size, or generation. The following are examples: ``m5.8xlarge``, ``c5*.*``, ``m5a.*``, ``r*``, ``*3*``.
	//  For example, if you specify ``c5*``,Amazon EC2 will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ``m5a.*``, Amazon EC2 will allow all the M5a instance types, but not the M5n instance types.
	//   If you specify ``AllowedInstanceTypes``, you can't specify ``ExcludedInstanceTypes``.
	//   Default: All instance types
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#allowed_instance_types Ec2LaunchTemplate#allowed_instance_types}
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types must be included, excluded, or required.
	//
	// +  To include bare metal instance types, specify ``included``.
	//   +  To require only bare metal instance types, specify ``required``.
	//   +  To exclude bare metal instance types, specify ``excluded``.
	//
	//  Default: ``excluded``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#bare_metal Ec2LaunchTemplate#bare_metal}
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
	//
	// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide*.
	//  Default: No minimum or maximum limits
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#baseline_ebs_bandwidth_mbps Ec2LaunchTemplate#baseline_ebs_bandwidth_mbps}
	BaselineEbsBandwidthMbps *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#baseline_performance_factors Ec2LaunchTemplate#baseline_performance_factors}.
	BaselinePerformanceFactors *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactors `field:"optional" json:"baselinePerformanceFactors" yaml:"baselinePerformanceFactors"`
	// Indicates whether burstable performance T instance types are included, excluded, or required.
	//
	// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html).
	//   +  To include burstable performance instance types, specify ``included``.
	//   +  To require only burstable performance instance types, specify ``required``.
	//   +  To exclude burstable performance instance types, specify ``excluded``.
	//
	//  Default: ``excluded``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#burstable_performance Ec2LaunchTemplate#burstable_performance}
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// The CPU manufacturers to include.
	//
	// +  For instance types with Intel CPUs, specify ``intel``.
	//   +  For instance types with AMD CPUs, specify ``amd``.
	//   +  For instance types with AWS CPUs, specify ``amazon-web-services``.
	//
	//   Don't confuse the CPU manufacturer with the CPU architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	//   Default: Any manufacturer
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#cpu_manufacturers Ec2LaunchTemplate#cpu_manufacturers}
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// The instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk (``*``), to exclude an instance type, size, or generation. The following are examples: ``m5.8xlarge``, ``c5*.*``, ``m5a.*``, ``r*``, ``*3*``.
	//  For example, if you specify ``c5*``,Amazon EC2 will exclude the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ``m5a.*``, Amazon EC2 will exclude all the M5a instance types, but not the M5n instance types.
	//   If you specify ``ExcludedInstanceTypes``, you can't specify ``AllowedInstanceTypes``.
	//   Default: No excluded instance types
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#excluded_instance_types Ec2LaunchTemplate#excluded_instance_types}
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// The current generation instance types are recommended for use. Current generation instance types are typically the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide*.
	//  For current generation instance types, specify ``current``.
	//  For previous generation instance types, specify ``previous``.
	//  Default: Current and previous generation instance types
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#instance_generations Ec2LaunchTemplate#instance_generations}
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide*.
	//   +  To include instance types with instance store volumes, specify ``included``.
	//   +  To require only instance types with instance store volumes, specify ``required``.
	//   +  To exclude instance types with instance store volumes, specify ``excluded``.
	//
	//  Default: ``included``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#local_storage Ec2LaunchTemplate#local_storage}
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// The type of local storage that is required.
	//
	// +  For instance types with hard disk drive (HDD) storage, specify ``hdd``.
	//   +  For instance types with solid state drive (SSD) storage, specify ``ssd``.
	//
	//  Default: ``hdd`` and ``ssd``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#local_storage_types Ec2LaunchTemplate#local_storage_types}
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.
	//  The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//  If you set ``TargetCapacityUnitType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is based on the per vCPU or per memory price instead of the per instance price.
	//   Only one of ``SpotMaxPricePercentageOverLowestPrice`` or ``MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`` can be specified. If you don't specify either, Amazon EC2 will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ``999999``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#max_spot_price_as_percentage_of_optimal_on_demand_price Ec2LaunchTemplate#max_spot_price_as_percentage_of_optimal_on_demand_price}
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// The minimum and maximum amount of memory per vCPU, in GiB.  Default: No minimum or maximum limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#memory_gi_b_per_v_cpu Ec2LaunchTemplate#memory_gi_b_per_v_cpu}
	MemoryGiBPerVCpu *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpu `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// The minimum and maximum amount of memory, in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#memory_mi_b Ec2LaunchTemplate#memory_mi_b}
	MemoryMiB *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiB `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).  Default: No minimum or maximum limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#network_bandwidth_gbps Ec2LaunchTemplate#network_bandwidth_gbps}
	NetworkBandwidthGbps *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// The minimum and maximum number of network interfaces.  Default: No minimum or maximum limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#network_interface_count Ec2LaunchTemplate#network_interface_count}
	NetworkInterfaceCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// [Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.
	//  The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//  To turn off price protection, specify a high value, such as ``999999``.
	//  This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html).
	//   If you set ``TargetCapacityUnitType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//   Default: ``20``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#on_demand_max_price_percentage_over_lowest_price Ec2LaunchTemplate#on_demand_max_price_percentage_over_lowest_price}
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must support hibernation for On-Demand Instances.  This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html).  Default: ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#require_hibernate_support Ec2LaunchTemplate#require_hibernate_support}
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price.
	//
	// The identified Spot price is the Spot price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified Spot price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose Spot price exceeds your specified threshold.
	//  The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//  If you set ``TargetCapacityUnitType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//  This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html).
	//   Only one of ``SpotMaxPricePercentageOverLowestPrice`` or ``MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`` can be specified. If you don't specify either, Amazon EC2 will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ``999999``.
	//   Default: ``100``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#spot_max_price_percentage_over_lowest_price Ec2LaunchTemplate#spot_max_price_percentage_over_lowest_price}
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum amount of total local storage, in GB.  Default: No minimum or maximum limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#total_local_storage_gb Ec2LaunchTemplate#total_local_storage_gb}
	TotalLocalStorageGb *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_launch_template#v_cpu_count Ec2LaunchTemplate#v_cpu_count}
	VCpuCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCount `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

