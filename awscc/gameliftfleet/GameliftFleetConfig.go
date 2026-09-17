// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftFleetConfig struct {
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
	// A descriptive label that is associated with a fleet. Fleet names do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#name GameliftFleet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for Anywhere fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#anywhere_configuration GameliftFleet#anywhere_configuration}
	AnywhereConfiguration *GameliftFleetAnywhereConfiguration `field:"optional" json:"anywhereConfiguration" yaml:"anywhereConfiguration"`
	// Determines when and how to apply fleet or location capacities.
	//
	// Allowed options are ON_UPDATE (default), ON_CREATE_AND_UPDATE and ON_CREATE_AND_UPDATE_WITH_AUTOSCALING. If you choose ON_CREATE_AND_UPDATE_WITH_AUTOSCALING, MinSize and MaxSize will still be applied on creation and on updates, but DesiredEC2Instances will only be applied once on fleet creation and will be ignored during updates to prevent conflicts with auto-scaling. During updates with ON_CREATE_AND_UPDATE_WITH_AUTOSCALING chosen, if current desired instance is lower than the new MinSize, it will be increased to the new MinSize; if current desired instance is larger than the new MaxSize, it will be decreased to the new MaxSize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#apply_capacity GameliftFleet#apply_capacity}
	ApplyCapacity *string `field:"optional" json:"applyCapacity" yaml:"applyCapacity"`
	// A unique identifier for a build to be deployed on the new fleet.
	//
	// If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#build_id GameliftFleet#build_id}
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// Indicates whether to generate a TLS/SSL certificate for the new fleet.
	//
	// TLS certificates are used for encrypting traffic between game clients and game servers running on GameLift. If this parameter is not set, certificate generation is disabled. This fleet setting cannot be changed once the fleet is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#certificate_configuration GameliftFleet#certificate_configuration}
	CertificateConfiguration *GameliftFleetCertificateConfiguration `field:"optional" json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// ComputeType to differentiate EC2 hardware managed by GameLift and Anywhere hardware managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#compute_type GameliftFleet#compute_type}
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// A human-readable description of a fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#description GameliftFleet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// [DEPRECATED] The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#desired_ec2_instances GameliftFleet#desired_ec2_instances}
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#ec2_inbound_permissions GameliftFleet#ec2_inbound_permissions}
	Ec2InboundPermissions interface{} `field:"optional" json:"ec2InboundPermissions" yaml:"ec2InboundPermissions"`
	// The name of an EC2 instance type that is supported in Amazon GameLift.
	//
	// A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#ec2_instance_type GameliftFleet#ec2_instance_type}
	Ec2InstanceType *string `field:"optional" json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Indicates whether to use On-Demand instances or Spot instances for this fleet.
	//
	// If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#fleet_type GameliftFleet#fleet_type}
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// A unique identifier for an AWS IAM role that manages access to your AWS services.
	//
	// With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#instance_role_arn GameliftFleet#instance_role_arn}
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Credentials provider implementation that loads credentials from the Amazon EC2 Instance Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#instance_role_credentials_provider GameliftFleet#instance_role_credentials_provider}
	InstanceRoleCredentialsProvider *string `field:"optional" json:"instanceRoleCredentialsProvider" yaml:"instanceRoleCredentialsProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#locations GameliftFleet#locations}.
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// This parameter is no longer used.
	//
	// When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady()
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#log_paths GameliftFleet#log_paths}
	LogPaths *[]*string `field:"optional" json:"logPaths" yaml:"logPaths"`
	// [DEPRECATED] The maximum value that is allowed for the fleet's instance count.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#max_size GameliftFleet#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The name of an Amazon CloudWatch metric group.
	//
	// A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#metric_groups GameliftFleet#metric_groups}
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// [DEPRECATED] The minimum value allowed for the fleet's instance count.
	//
	// When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#min_size GameliftFleet#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A game session protection policy to apply to all game sessions hosted on instances in this fleet.
	//
	// When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#new_game_session_protection_policy GameliftFleet#new_game_session_protection_policy}
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// A unique identifier for the AWS account with the VPC that you want to peer your Amazon GameLift fleet with.
	//
	// You can find your account ID in the AWS Management Console under account settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#peer_vpc_aws_account_id GameliftFleet#peer_vpc_aws_account_id}
	PeerVpcAwsAccountId *string `field:"optional" json:"peerVpcAwsAccountId" yaml:"peerVpcAwsAccountId"`
	// A unique identifier for a VPC with resources to be accessed by your Amazon GameLift fleet.
	//
	// The VPC must be in the same Region as your fleet. To look up a VPC ID, use the VPC Dashboard in the AWS Management Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#peer_vpc_id GameliftFleet#peer_vpc_id}
	PeerVpcId *string `field:"optional" json:"peerVpcId" yaml:"peerVpcId"`
	// Configuration for player gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#player_gateway_configuration GameliftFleet#player_gateway_configuration}
	PlayerGatewayConfiguration *GameliftFleetPlayerGatewayConfiguration `field:"optional" json:"playerGatewayConfiguration" yaml:"playerGatewayConfiguration"`
	// The player gateway mode for the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#player_gateway_mode GameliftFleet#player_gateway_mode}
	PlayerGatewayMode *string `field:"optional" json:"playerGatewayMode" yaml:"playerGatewayMode"`
	// A policy that limits the number of game sessions an individual player can create over a span of time for this fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#resource_creation_limit_policy GameliftFleet#resource_creation_limit_policy}
	ResourceCreationLimitPolicy *GameliftFleetResourceCreationLimitPolicy `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// Instructions for launching server processes on each instance in the fleet.
	//
	// Server processes run either a custom game build executable or a Realtime script. The runtime configuration defines the server executables or launch script file, launch parameters, and the number of processes to run concurrently on each instance. When creating a fleet, the runtime configuration must have at least one server process configuration; otherwise the request fails with an invalid request exception.
	//
	// This parameter is required unless the parameters ServerLaunchPath and ServerLaunchParameters are defined. Runtime configuration has replaced these parameters, but fleets that use them will continue to work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#runtime_configuration GameliftFleet#runtime_configuration}
	RuntimeConfiguration *GameliftFleetRuntimeConfiguration `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// A list of rules that control how a fleet is scaled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#scaling_policies GameliftFleet#scaling_policies}
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
	// A unique identifier for a Realtime script to be deployed on a new Realtime Servers fleet.
	//
	// The script must have been successfully uploaded to Amazon GameLift. This fleet setting cannot be changed once the fleet is created.
	//
	// Note: It is not currently possible to use the !Ref command to reference a script created with a CloudFormation template for the fleet property ScriptId. Instead, use Fn::GetAtt Script.Arn or Fn::GetAtt Script.Id to retrieve either of these properties as input for ScriptId. Alternatively, enter a ScriptId string manually.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#script_id GameliftFleet#script_id}
	ScriptId *string `field:"optional" json:"scriptId" yaml:"scriptId"`
	// This parameter is no longer used but is retained for backward compatibility.
	//
	// Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#server_launch_parameters GameliftFleet#server_launch_parameters}
	ServerLaunchParameters *string `field:"optional" json:"serverLaunchParameters" yaml:"serverLaunchParameters"`
	// This parameter is no longer used.
	//
	// Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#server_launch_path GameliftFleet#server_launch_path}
	ServerLaunchPath *string `field:"optional" json:"serverLaunchPath" yaml:"serverLaunchPath"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_fleet#tags GameliftFleet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

