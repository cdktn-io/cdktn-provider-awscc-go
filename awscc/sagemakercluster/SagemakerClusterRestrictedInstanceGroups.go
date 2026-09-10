// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterRestrictedInstanceGroups struct {
	// The number of instances that are currently in the restricted instance group of a SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#current_count SagemakerCluster#current_count}
	CurrentCount *float64 `field:"optional" json:"currentCount" yaml:"currentCount"`
	// The configuration for the restricted instance groups (RIG) environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#environment_config SagemakerCluster#environment_config}
	EnvironmentConfig *SagemakerClusterRestrictedInstanceGroupsEnvironmentConfig `field:"optional" json:"environmentConfig" yaml:"environmentConfig"`
	// The execution role for the instance group to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#execution_role SagemakerCluster#execution_role}
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The number of instances you specified to add to the restricted instance group of a SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#instance_count SagemakerCluster#instance_count}
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The name of the instance group of a SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#instance_group_name SagemakerCluster#instance_group_name}
	InstanceGroupName *string `field:"optional" json:"instanceGroupName" yaml:"instanceGroupName"`
	// The instance storage configuration for the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#instance_storage_configs SagemakerCluster#instance_storage_configs}
	InstanceStorageConfigs interface{} `field:"optional" json:"instanceStorageConfigs" yaml:"instanceStorageConfigs"`
	// The instance type of the instance group of a SageMaker HyperPod cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#instance_type SagemakerCluster#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#on_start_deep_health_checks SagemakerCluster#on_start_deep_health_checks}
	OnStartDeepHealthChecks *[]*string `field:"optional" json:"onStartDeepHealthChecks" yaml:"onStartDeepHealthChecks"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#override_vpc_config SagemakerCluster#override_vpc_config}
	OverrideVpcConfig *SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfig `field:"optional" json:"overrideVpcConfig" yaml:"overrideVpcConfig"`
	// The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading.
	//
	// For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#threads_per_core SagemakerCluster#threads_per_core}
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// The Amazon Resource Name (ARN) of the training plan to use for this cluster restricted instance group.
	//
	// For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#training_plan_arn SagemakerCluster#training_plan_arn}
	TrainingPlanArn *string `field:"optional" json:"trainingPlanArn" yaml:"trainingPlanArn"`
}

