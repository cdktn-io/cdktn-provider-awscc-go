// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service awscc_ecs_service}.
type EcsService interface {
	cdktn.TerraformResource
	AvailabilityZoneRebalancing() *string
	SetAvailabilityZoneRebalancing(val *string)
	AvailabilityZoneRebalancingInput() *string
	CapacityProviderStrategy() EcsServiceCapacityProviderStrategyList
	CapacityProviderStrategyInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentConfiguration() EcsServiceDeploymentConfigurationOutputReference
	DeploymentConfigurationInput() interface{}
	DeploymentController() EcsServiceDeploymentControllerOutputReference
	DeploymentControllerInput() interface{}
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	DesiredCountInput() *float64
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableEcsManagedTagsInput() interface{}
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	ForceNewDeployment() EcsServiceForceNewDeploymentOutputReference
	ForceNewDeploymentInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriodSeconds() *float64
	SetHealthCheckGracePeriodSeconds(val *float64)
	HealthCheckGracePeriodSecondsInput() *float64
	Id() *string
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancers() EcsServiceLoadBalancersList
	LoadBalancersInput() interface{}
	Monitoring() EcsServiceMonitoringOutputReference
	MonitoringInput() interface{}
	Name() *string
	NetworkConfiguration() EcsServiceNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	PlacementConstraints() EcsServicePlacementConstraintsList
	PlacementConstraintsInput() interface{}
	PlacementStrategies() EcsServicePlacementStrategiesList
	PlacementStrategiesInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SchedulingStrategy() *string
	SetSchedulingStrategy(val *string)
	SchedulingStrategyInput() *string
	ServiceArn() *string
	ServiceConnectConfiguration() EcsServiceServiceConnectConfigurationOutputReference
	ServiceConnectConfigurationInput() interface{}
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	ServiceRegistries() EcsServiceServiceRegistriesList
	ServiceRegistriesInput() interface{}
	Tags() EcsServiceTagsList
	TagsInput() interface{}
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	TaskDefinitionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VolumeConfigurations() EcsServiceVolumeConfigurationsList
	VolumeConfigurationsInput() interface{}
	VpcLatticeConfigurations() EcsServiceVpcLatticeConfigurationsList
	VpcLatticeConfigurationsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCapacityProviderStrategy(value interface{})
	PutDeploymentConfiguration(value *EcsServiceDeploymentConfiguration)
	PutDeploymentController(value *EcsServiceDeploymentController)
	PutForceNewDeployment(value *EcsServiceForceNewDeployment)
	PutLoadBalancers(value interface{})
	PutMonitoring(value *EcsServiceMonitoring)
	PutNetworkConfiguration(value *EcsServiceNetworkConfiguration)
	PutPlacementConstraints(value interface{})
	PutPlacementStrategies(value interface{})
	PutServiceConnectConfiguration(value *EcsServiceServiceConnectConfiguration)
	PutServiceRegistries(value interface{})
	PutTags(value interface{})
	PutVolumeConfigurations(value interface{})
	PutVpcLatticeConfigurations(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAvailabilityZoneRebalancing()
	ResetCapacityProviderStrategy()
	ResetCluster()
	ResetDeploymentConfiguration()
	ResetDeploymentController()
	ResetDesiredCount()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetForceNewDeployment()
	ResetHealthCheckGracePeriodSeconds()
	ResetLaunchType()
	ResetLoadBalancers()
	ResetMonitoring()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementConstraints()
	ResetPlacementStrategies()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetRole()
	ResetSchedulingStrategy()
	ResetServiceConnectConfiguration()
	ResetServiceName()
	ResetServiceRegistries()
	ResetTags()
	ResetTaskDefinition()
	ResetVolumeConfigurations()
	ResetVpcLatticeConfigurations()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EcsService
type jsiiProxy_EcsService struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EcsService) AvailabilityZoneRebalancing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) AvailabilityZoneRebalancingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CapacityProviderStrategy() EcsServiceCapacityProviderStrategyList {
	var returns EcsServiceCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentConfiguration() EcsServiceDeploymentConfigurationOutputReference {
	var returns EcsServiceDeploymentConfigurationOutputReference
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentController() EcsServiceDeploymentControllerOutputReference {
	var returns EcsServiceDeploymentControllerOutputReference
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentControllerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DesiredCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForceNewDeployment() EcsServiceForceNewDeploymentOutputReference {
	var returns EcsServiceForceNewDeploymentOutputReference
	_jsii_.Get(
		j,
		"forceNewDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForceNewDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) HealthCheckGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LoadBalancers() EcsServiceLoadBalancersList {
	var returns EcsServiceLoadBalancersList
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Monitoring() EcsServiceMonitoringOutputReference {
	var returns EcsServiceMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) NetworkConfiguration() EcsServiceNetworkConfigurationOutputReference {
	var returns EcsServiceNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementConstraints() EcsServicePlacementConstraintsList {
	var returns EcsServicePlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementStrategies() EcsServicePlacementStrategiesList {
	var returns EcsServicePlacementStrategiesList
	_jsii_.Get(
		j,
		"placementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementStrategiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) SchedulingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceConnectConfiguration() EcsServiceServiceConnectConfigurationOutputReference {
	var returns EcsServiceServiceConnectConfigurationOutputReference
	_jsii_.Get(
		j,
		"serviceConnectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceConnectConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceConnectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceRegistries() EcsServiceServiceRegistriesList {
	var returns EcsServiceServiceRegistriesList
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceRegistriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceRegistriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Tags() EcsServiceTagsList {
	var returns EcsServiceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) VolumeConfigurations() EcsServiceVolumeConfigurationsList {
	var returns EcsServiceVolumeConfigurationsList
	_jsii_.Get(
		j,
		"volumeConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) VolumeConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) VpcLatticeConfigurations() EcsServiceVpcLatticeConfigurationsList {
	var returns EcsServiceVpcLatticeConfigurationsList
	_jsii_.Get(
		j,
		"vpcLatticeConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) VpcLatticeConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcLatticeConfigurationsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service awscc_ecs_service} Resource.
func NewEcsService(scope constructs.Construct, id *string, config *EcsServiceConfig) EcsService {
	_init_.Initialize()

	if err := validateNewEcsServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsService{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service awscc_ecs_service} Resource.
func NewEcsService_Override(e EcsService, scope constructs.Construct, id *string, config *EcsServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsService.EcsService",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcsService)SetAvailabilityZoneRebalancing(val *string) {
	if err := j.validateSetAvailabilityZoneRebalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRebalancing",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDesiredCount(val *float64) {
	if err := j.validateSetDesiredCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetHealthCheckGracePeriodSeconds(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetSchedulingStrategy(val *string) {
	if err := j.validateSetSchedulingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetTaskDefinition(val *string) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

// Generates CDKTN code for importing a EcsService resource upon running "cdktn plan <stack-name>".
func EcsService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEcsService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ecsService.EcsService",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EcsService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ecsService.EcsService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ecsService.EcsService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ecsService.EcsService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcsService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.ecsService.EcsService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcsService) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EcsService) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcsService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EcsService) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsService) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EcsService) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsService) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcsService) PutCapacityProviderStrategy(value interface{}) {
	if err := e.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutDeploymentConfiguration(value *EcsServiceDeploymentConfiguration) {
	if err := e.validatePutDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutDeploymentController(value *EcsServiceDeploymentController) {
	if err := e.validatePutDeploymentControllerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentController",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutForceNewDeployment(value *EcsServiceForceNewDeployment) {
	if err := e.validatePutForceNewDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putForceNewDeployment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutLoadBalancers(value interface{}) {
	if err := e.validatePutLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancers",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutMonitoring(value *EcsServiceMonitoring) {
	if err := e.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutNetworkConfiguration(value *EcsServiceNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutPlacementConstraints(value interface{}) {
	if err := e.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutPlacementStrategies(value interface{}) {
	if err := e.validatePutPlacementStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementStrategies",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutServiceConnectConfiguration(value *EcsServiceServiceConnectConfiguration) {
	if err := e.validatePutServiceConnectConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putServiceConnectConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutServiceRegistries(value interface{}) {
	if err := e.validatePutServiceRegistriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putServiceRegistries",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutVolumeConfigurations(value interface{}) {
	if err := e.validatePutVolumeConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVolumeConfigurations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutVpcLatticeConfigurations(value interface{}) {
	if err := e.validatePutVpcLatticeConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVpcLatticeConfigurations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EcsService) ResetAvailabilityZoneRebalancing() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZoneRebalancing",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetCluster() {
	_jsii_.InvokeVoid(
		e,
		"resetCluster",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentController() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentController",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDesiredCount() {
	_jsii_.InvokeVoid(
		e,
		"resetDesiredCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetForceNewDeployment() {
	_jsii_.InvokeVoid(
		e,
		"resetForceNewDeployment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetHealthCheckGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckGracePeriodSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetLaunchType() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetMonitoring() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPlacementStrategies() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementStrategies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		e,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetRole() {
	_jsii_.InvokeVoid(
		e,
		"resetRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetSchedulingStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetSchedulingStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetServiceConnectConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceConnectConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetServiceName() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetServiceRegistries() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceRegistries",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTaskDefinition() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskDefinition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetVolumeConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetVpcLatticeConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcLatticeConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

