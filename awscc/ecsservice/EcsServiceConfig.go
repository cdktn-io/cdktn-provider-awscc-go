// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsServiceConfig struct {
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
	// Indicates whether to use Availability Zone rebalancing for the service.
	//
	// For more information, see [Balancing an Amazon ECS service across Availability Zones](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-rebalancing.html) in the *Amazon Elastic Container Service Developer Guide*.
	//  The default behavior of ``AvailabilityZoneRebalancing`` differs between create and update requests:
	//   +  For create service requests, when no value is specified for ``AvailabilityZoneRebalancing``, Amazon ECS defaults the value to ``ENABLED``.
	//   +  For update service requests, when no value is specified for ``AvailabilityZoneRebalancing``, Amazon ECS defaults to the existing service’s ``AvailabilityZoneRebalancing`` value. If the service never had an ``AvailabilityZoneRebalancing`` value set, Amazon ECS treats this as ``DISABLED``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#availability_zone_rebalancing EcsService#availability_zone_rebalancing}
	AvailabilityZoneRebalancing *string `field:"optional" json:"availabilityZoneRebalancing" yaml:"availabilityZoneRebalancing"`
	// The capacity provider strategy to use for the service.
	//
	// If a ``capacityProviderStrategy`` is specified, the ``launchType`` parameter must be omitted. If no ``capacityProviderStrategy`` or ``launchType`` is specified, the ``defaultCapacityProviderStrategy`` for the cluster is used.
	//  A capacity provider strategy can contain a maximum of 20 capacity providers.
	//   To remove this property from your service resource, specify an empty ``CapacityProviderStrategyItem`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#capacity_provider_strategy EcsService#capacity_provider_strategy}
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.
	//
	// If you do not specify a cluster, the default cluster is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#cluster EcsService#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#deployment_configuration EcsService#deployment_configuration}
	DeploymentConfiguration *EcsServiceDeploymentConfiguration `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// The deployment controller to use for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#deployment_controller EcsService#deployment_controller}
	DeploymentController *EcsServiceDeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The number of instantiations of the specified task definition to place and keep running in your service.
	//
	// For new services, if a desired count is not specified, a default value of ``1`` is used. When using the ``DAEMON`` scheduling strategy, the desired count is not required.
	//  For existing services, if a desired count is not specified, it is omitted from the operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#desired_count EcsService#desired_count}
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide*.
	//  When you use Amazon ECS managed tags, you must set the ``propagateTags`` request parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#enable_ecs_managed_tags EcsService#enable_ecs_managed_tags}
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Determines whether the execute command functionality is turned on for the service.
	//
	// If ``true``, the execute command functionality is turned on for all containers in tasks as part of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#enable_execute_command EcsService#enable_execute_command}
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Determines whether to force a new deployment of the service.
	//
	// By default, deployments aren't forced. You can use this option to start a new deployment with no service definition changes. For example, you can update a service's tasks to use a newer Docker image with the same image/tag combination (``my_image:latest``) or to roll Fargate tasks onto a newer platform version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#force_new_deployment EcsService#force_new_deployment}
	ForceNewDeployment *EcsServiceForceNewDeployment `field:"optional" json:"forceNewDeployment" yaml:"forceNewDeployment"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task has first started.
	//
	// If you do not specify a health check grace period value, the default value of 0 is used. If you do not use any of the health checks, then ``healthCheckGracePeriodSeconds`` is unused.
	//  If your service has more running tasks than desired, unhealthy tasks in the grace period might be stopped to reach the desired count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#health_check_grace_period_seconds EcsService#health_check_grace_period_seconds}
	HealthCheckGracePeriodSeconds *float64 `field:"optional" json:"healthCheckGracePeriodSeconds" yaml:"healthCheckGracePeriodSeconds"`
	// The launch type on which to run your service.
	//
	// For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide*.
	//   If you want to use Managed Instances, you must use the ``capacityProviderStrategy`` request parameter
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#launch_type EcsService#launch_type}
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// A list of load balancer objects to associate with the service.
	//
	// If you specify the ``Role`` property, ``LoadBalancers`` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide*.
	//   To remove this property from your service resource, specify an empty ``LoadBalancer`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#load_balancers EcsService#load_balancers}
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The optional monitoring configuration for the service, which defines the resolution for the service-level ``CPUUtilization`` and ``MemoryUtilization`` Amazon CloudWatch metrics.
	//
	// When not specified, Amazon ECS uses the default resolution of ``60`` seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#monitoring EcsService#monitoring}
	Monitoring *EcsServiceMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The network configuration for the service.
	//
	// This parameter is required for task definitions that use the ``awsvpc`` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#network_configuration EcsService#network_configuration}
	NetworkConfiguration *EcsServiceNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for tasks in your service.
	//
	// You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
	//   To remove this property from your service resource, specify an empty ``PlacementConstraint`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#placement_constraints EcsService#placement_constraints}
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for tasks in your service.
	//
	// You can specify a maximum of 5 strategy rules for each service.
	//   To remove this property from your service resource, specify an empty ``PlacementStrategy`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#placement_strategies EcsService#placement_strategies}
	PlacementStrategies interface{} `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// The platform version that your tasks in the service are running on.
	//
	// A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the ``LATEST`` platform version is used. For more information, see [platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#platform_version EcsService#platform_version}
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags aren't propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
	//  You must set this to a value other than ``NONE`` when you use Cost Explorer. For more information, see [Amazon ECS usage reports](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/usage-reports.html) in the *Amazon Elastic Container Service Developer Guide*.
	//  The default is ``NONE``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#propagate_tags EcsService#propagate_tags}
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.
	//
	// This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the ``awsvpc`` network mode. If you specify the ``role`` parameter, you must also specify a load balancer object with the ``loadBalancers`` parameter.
	//   If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the ``awsvpc`` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide*.
	//   If your specified role has a path other than ``/``, then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name ``bar`` has a path of ``/foo/`` then you would specify ``/foo/bar`` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#role EcsService#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The scheduling strategy to use for the service.
	//
	// For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).
	//  There are two service scheduler strategies available:
	//   +  ``REPLICA``-The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the ``CODE_DEPLOY`` or ``EXTERNAL`` deployment controller types.
	//   +  ``DAEMON``-The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.
	//   Tasks using the Fargate launch type or the ``CODE_DEPLOY`` or ``EXTERNAL`` deployment controller types don't support the ``DAEMON`` scheduling strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#scheduling_strategy EcsService#scheduling_strategy}
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.
	//
	// Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#service_connect_configuration EcsService#service_connect_configuration}
	ServiceConnectConfiguration *EcsServiceServiceConnectConfiguration `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// The name of your service.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.
	//   The stack update fails if you change any properties that require replacement and the ``ServiceName`` is configured. This is because AWS CloudFormation creates the replacement service first, but each ``ServiceName`` must be unique in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#service_name EcsService#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The details of the service discovery registry to associate with this service.
	//
	// For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).
	//   Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
	//    To remove this property from your service resource, specify an empty ``ServiceRegistry`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#service_registries EcsService#service_registries}
	ServiceRegistries interface{} `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
	// The metadata that you apply to the service to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well.
	//  The following basic restrictions apply to tags:
	//   +  Maximum number of tags per resource - 50
	//   +  For each resource, each tag key must be unique, and each tag key can have only one value.
	//   +  Maximum key length - 128 Unicode characters in UTF-8
	//   +  Maximum value length - 256 Unicode characters in UTF-8
	//   +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ``family`` and ``revision`` (``family:revision``) or full ARN of the task definition to run in your service.
	//
	// If a ``revision`` isn't specified, the latest ``ACTIVE`` revision is used.
	//  A task definition must be specified if the service uses either the ``ECS`` or ``CODE_DEPLOY`` deployment controllers.
	//  For more information about deployment types, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#task_definition EcsService#task_definition}
	TaskDefinition *string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// The configuration for a volume specified in the task definition as a volume that is configured at launch time.
	//
	// Currently, the only supported volume type is an Amazon EBS volume.
	//   To remove this property from your service resource, specify an empty ``ServiceVolumeConfiguration`` array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#volume_configurations EcsService#volume_configurations}
	VolumeConfigurations interface{} `field:"optional" json:"volumeConfigurations" yaml:"volumeConfigurations"`
	// The VPC Lattice configuration for the service being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_service#vpc_lattice_configurations EcsService#vpc_lattice_configurations}
	VpcLatticeConfigurations interface{} `field:"optional" json:"vpcLatticeConfigurations" yaml:"vpcLatticeConfigurations"`
}

