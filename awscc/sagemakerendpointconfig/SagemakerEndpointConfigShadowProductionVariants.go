// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigShadowProductionVariants struct {
	// Settings for the capacity reservation for the compute instances that SageMaker AI reserves for an endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#capacity_reservation_config SagemakerEndpointConfigA#capacity_reservation_config}
	CapacityReservationConfig *SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfig `field:"optional" json:"capacityReservationConfig" yaml:"capacityReservationConfig"`
	// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#container_startup_health_check_timeout_in_seconds SagemakerEndpointConfigA#container_startup_health_check_timeout_in_seconds}
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// Specifies configuration for a core dump from the model container when the process crashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#core_dump_config SagemakerEndpointConfigA#core_dump_config}
	CoreDumpConfig *SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfig `field:"optional" json:"coreDumpConfig" yaml:"coreDumpConfig"`
	// You can use this parameter to turn on native AWS Systems Manager (SSM) access for a production variant behind an endpoint.
	//
	// By default, SSM access is disabled for all production variants behind an endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#enable_ssm_access SagemakerEndpointConfigA#enable_ssm_access}
	EnableSsmAccess interface{} `field:"optional" json:"enableSsmAccess" yaml:"enableSsmAccess"`
	// Specifies an option from a collection of preconfigured Amazon Machine Image (AMI) images.
	//
	// Each image is configured by AWS with a set of software and driver versions. AWS optimizes these configurations for different machine learning workloads. By selecting an AMI version, you can ensure that your inference environment is compatible with specific software requirements, such as CUDA driver versions, Linux kernel versions, or AWS Neuron driver versions
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#inference_ami_version SagemakerEndpointConfigA#inference_ami_version}
	InferenceAmiVersion *string `field:"optional" json:"inferenceAmiVersion" yaml:"inferenceAmiVersion"`
	// Number of instances to launch initially.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#initial_instance_count SagemakerEndpointConfigA#initial_instance_count}
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#initial_variant_weight SagemakerEndpointConfigA#initial_variant_weight}
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// A list of instance pools for the production variant.
	//
	// Each instance pool specifies an instance type and its priority for provisioning. Use instance pools to configure heterogeneous endpoints that deploy models across multiple instance types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#instance_pools SagemakerEndpointConfigA#instance_pools}
	InstancePools interface{} `field:"optional" json:"instancePools" yaml:"instancePools"`
	// The ML compute instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#instance_type SagemakerEndpointConfigA#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Settings that control the range in the number of instances that the endpoint provisions as it scales up or down to accommodate traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#managed_instance_scaling SagemakerEndpointConfigA#managed_instance_scaling}
	ManagedInstanceScaling *SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScaling `field:"optional" json:"managedInstanceScaling" yaml:"managedInstanceScaling"`
	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#model_data_download_timeout_in_seconds SagemakerEndpointConfigA#model_data_download_timeout_in_seconds}
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#model_name SagemakerEndpointConfigA#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Settings that control how the endpoint routes incoming traffic to the instances that the endpoint hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#routing_config SagemakerEndpointConfigA#routing_config}
	RoutingConfig *SagemakerEndpointConfigShadowProductionVariantsRoutingConfig `field:"optional" json:"routingConfig" yaml:"routingConfig"`
	// The serverless configuration for an endpoint. Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#serverless_config SagemakerEndpointConfigA#serverless_config}
	ServerlessConfig *SagemakerEndpointConfigShadowProductionVariantsServerlessConfig `field:"optional" json:"serverlessConfig" yaml:"serverlessConfig"`
	// The timeout value, in seconds, for provisioning instances for the production variant.
	//
	// When SageMaker encounters an insufficient capacity error while provisioning instances, it retries with the next instance pool (if configured) or waits until the timeout expires. This timeout applies only to capacity provisioning and does not include the time for model download or container startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#variant_instance_provision_timeout_in_seconds SagemakerEndpointConfigA#variant_instance_provision_timeout_in_seconds}
	VariantInstanceProvisionTimeoutInSeconds *float64 `field:"optional" json:"variantInstanceProvisionTimeoutInSeconds" yaml:"variantInstanceProvisionTimeoutInSeconds"`
	// The name of the production variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#variant_name SagemakerEndpointConfigA#variant_name}
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
	// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant.
	//
	// Currently only Amazon EBS gp2 storage volumes are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#volume_size_in_gb SagemakerEndpointConfigA#volume_size_in_gb}
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

