// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#accelerator_count Ec2SpotFleet#accelerator_count}.
	AcceleratorCount *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#accelerator_manufacturers Ec2SpotFleet#accelerator_manufacturers}.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#accelerator_names Ec2SpotFleet#accelerator_names}.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#accelerator_total_memory_mi_b Ec2SpotFleet#accelerator_total_memory_mi_b}.
	AcceleratorTotalMemoryMiB *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsAcceleratorTotalMemoryMiB `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#accelerator_types Ec2SpotFleet#accelerator_types}.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#allowed_instance_types Ec2SpotFleet#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#bare_metal Ec2SpotFleet#bare_metal}.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#baseline_ebs_bandwidth_mbps Ec2SpotFleet#baseline_ebs_bandwidth_mbps}.
	BaselineEbsBandwidthMbps *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#baseline_performance_factors Ec2SpotFleet#baseline_performance_factors}.
	BaselinePerformanceFactors *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactors `field:"optional" json:"baselinePerformanceFactors" yaml:"baselinePerformanceFactors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#burstable_performance Ec2SpotFleet#burstable_performance}.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#cpu_manufacturers Ec2SpotFleet#cpu_manufacturers}.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#excluded_instance_types Ec2SpotFleet#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#instance_generations Ec2SpotFleet#instance_generations}.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#local_storage Ec2SpotFleet#local_storage}.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#local_storage_types Ec2SpotFleet#local_storage_types}.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#max_spot_price_as_percentage_of_optimal_on_demand_price Ec2SpotFleet#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#memory_gi_b_per_v_cpu Ec2SpotFleet#memory_gi_b_per_v_cpu}.
	MemoryGiBPerVCpu *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryGiBPerVCpu `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#memory_mi_b Ec2SpotFleet#memory_mi_b}.
	MemoryMiB *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#network_bandwidth_gbps Ec2SpotFleet#network_bandwidth_gbps}.
	NetworkBandwidthGbps *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#network_interface_count Ec2SpotFleet#network_interface_count}.
	NetworkInterfaceCount *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#on_demand_max_price_percentage_over_lowest_price Ec2SpotFleet#on_demand_max_price_percentage_over_lowest_price}.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#require_encryption_in_transit Ec2SpotFleet#require_encryption_in_transit}.
	RequireEncryptionInTransit interface{} `field:"optional" json:"requireEncryptionInTransit" yaml:"requireEncryptionInTransit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#require_hibernate_support Ec2SpotFleet#require_hibernate_support}.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#spot_max_price_percentage_over_lowest_price Ec2SpotFleet#spot_max_price_percentage_over_lowest_price}.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#total_local_storage_gb Ec2SpotFleet#total_local_storage_gb}.
	TotalLocalStorageGb *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_spot_fleet#v_cpu_count Ec2SpotFleet#v_cpu_count}.
	VCpuCount *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsVCpuCount `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

