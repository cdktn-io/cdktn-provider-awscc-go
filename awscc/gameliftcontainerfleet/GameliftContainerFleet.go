// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftcontainerfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet}.
type GameliftContainerFleet interface {
	cdktn.TerraformResource
	BillingType() *string
	SetBillingType(val *string)
	BillingTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentConfiguration() GameliftContainerFleetDeploymentConfigurationOutputReference
	DeploymentConfigurationInput() interface{}
	DeploymentDetails() GameliftContainerFleetDeploymentDetailsOutputReference
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FleetArn() *string
	FleetId() *string
	FleetRoleArn() *string
	SetFleetRoleArn(val *string)
	FleetRoleArnInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameServerContainerGroupDefinitionArn() *string
	GameServerContainerGroupDefinitionName() *string
	SetGameServerContainerGroupDefinitionName(val *string)
	GameServerContainerGroupDefinitionNameInput() *string
	GameServerContainerGroupsPerInstance() *float64
	SetGameServerContainerGroupsPerInstance(val *float64)
	GameServerContainerGroupsPerInstanceInput() *float64
	GameSessionCreationLimitPolicy() GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference
	GameSessionCreationLimitPolicyInput() interface{}
	Id() *string
	InstanceConnectionPortRange() GameliftContainerFleetInstanceConnectionPortRangeOutputReference
	InstanceConnectionPortRangeInput() interface{}
	InstanceInboundPermissions() GameliftContainerFleetInstanceInboundPermissionsList
	InstanceInboundPermissionsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Locations() GameliftContainerFleetLocationsList
	LocationsInput() interface{}
	LogConfiguration() GameliftContainerFleetLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	MaximumGameServerContainerGroupsPerInstance() *float64
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MetricGroupsInput() *[]*string
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	NewGameSessionProtectionPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	PerInstanceContainerGroupDefinitionArn() *string
	PerInstanceContainerGroupDefinitionName() *string
	SetPerInstanceContainerGroupDefinitionName(val *string)
	PerInstanceContainerGroupDefinitionNameInput() *string
	PlayerGatewayMode() *string
	SetPlayerGatewayMode(val *string)
	PlayerGatewayModeInput() *string
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
	ScalingPolicies() GameliftContainerFleetScalingPoliciesList
	ScalingPoliciesInput() interface{}
	Status() *string
	Tags() GameliftContainerFleetTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDeploymentConfiguration(value *GameliftContainerFleetDeploymentConfiguration)
	PutGameSessionCreationLimitPolicy(value *GameliftContainerFleetGameSessionCreationLimitPolicy)
	PutInstanceConnectionPortRange(value *GameliftContainerFleetInstanceConnectionPortRange)
	PutInstanceInboundPermissions(value interface{})
	PutLocations(value interface{})
	PutLogConfiguration(value *GameliftContainerFleetLogConfiguration)
	PutScalingPolicies(value interface{})
	PutTags(value interface{})
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
	ResetBillingType()
	ResetDeploymentConfiguration()
	ResetDescription()
	ResetGameServerContainerGroupDefinitionName()
	ResetGameServerContainerGroupsPerInstance()
	ResetGameSessionCreationLimitPolicy()
	ResetInstanceConnectionPortRange()
	ResetInstanceInboundPermissions()
	ResetInstanceType()
	ResetLocations()
	ResetLogConfiguration()
	ResetMetricGroups()
	ResetNewGameSessionProtectionPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerInstanceContainerGroupDefinitionName()
	ResetPlayerGatewayMode()
	ResetScalingPolicies()
	ResetTags()
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

// The jsii proxy struct for GameliftContainerFleet
type jsiiProxy_GameliftContainerFleet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GameliftContainerFleet) BillingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) BillingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentConfiguration() GameliftContainerFleetDeploymentConfigurationOutputReference {
	var returns GameliftContainerFleetDeploymentConfigurationOutputReference
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DeploymentDetails() GameliftContainerFleetDeploymentDetailsOutputReference {
	var returns GameliftContainerFleetDeploymentDetailsOutputReference
	_jsii_.Get(
		j,
		"deploymentDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FleetRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerContainerGroupDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameServerContainerGroupsPerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameServerContainerGroupsPerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameSessionCreationLimitPolicy() GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference {
	var returns GameliftContainerFleetGameSessionCreationLimitPolicyOutputReference
	_jsii_.Get(
		j,
		"gameSessionCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) GameSessionCreationLimitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gameSessionCreationLimitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceConnectionPortRange() GameliftContainerFleetInstanceConnectionPortRangeOutputReference {
	var returns GameliftContainerFleetInstanceConnectionPortRangeOutputReference
	_jsii_.Get(
		j,
		"instanceConnectionPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceConnectionPortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConnectionPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceInboundPermissions() GameliftContainerFleetInstanceInboundPermissionsList {
	var returns GameliftContainerFleetInstanceInboundPermissionsList
	_jsii_.Get(
		j,
		"instanceInboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceInboundPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceInboundPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Locations() GameliftContainerFleetLocationsList {
	var returns GameliftContainerFleetLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LogConfiguration() GameliftContainerFleetLogConfigurationOutputReference {
	var returns GameliftContainerFleetLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MaximumGameServerContainerGroupsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumGameServerContainerGroupsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) MetricGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) NewGameSessionProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PerInstanceContainerGroupDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perInstanceContainerGroupDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PlayerGatewayMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playerGatewayMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) PlayerGatewayModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playerGatewayModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ScalingPolicies() GameliftContainerFleetScalingPoliciesList {
	var returns GameliftContainerFleetScalingPoliciesList
	_jsii_.Get(
		j,
		"scalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) ScalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) Tags() GameliftContainerFleetTagsList {
	var returns GameliftContainerFleetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet} Resource.
func NewGameliftContainerFleet(scope constructs.Construct, id *string, config *GameliftContainerFleetConfig) GameliftContainerFleet {
	_init_.Initialize()

	if err := validateNewGameliftContainerFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerFleet{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_fleet awscc_gamelift_container_fleet} Resource.
func NewGameliftContainerFleet_Override(g GameliftContainerFleet, scope constructs.Construct, id *string, config *GameliftContainerFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetBillingType(val *string) {
	if err := j.validateSetBillingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetFleetRoleArn(val *string) {
	if err := j.validateSetFleetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetRoleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetGameServerContainerGroupDefinitionName(val *string) {
	if err := j.validateSetGameServerContainerGroupDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetGameServerContainerGroupsPerInstance(val *float64) {
	if err := j.validateSetGameServerContainerGroupsPerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameServerContainerGroupsPerInstance",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetMetricGroups(val *[]*string) {
	if err := j.validateSetMetricGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetNewGameSessionProtectionPolicy(val *string) {
	if err := j.validateSetNewGameSessionProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetPerInstanceContainerGroupDefinitionName(val *string) {
	if err := j.validateSetPerInstanceContainerGroupDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perInstanceContainerGroupDefinitionName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetPlayerGatewayMode(val *string) {
	if err := j.validateSetPlayerGatewayModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playerGatewayMode",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a GameliftContainerFleet resource upon running "cdktn plan <stack-name>".
func GameliftContainerFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
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
func GameliftContainerFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftContainerFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftContainerFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftContainerFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutDeploymentConfiguration(value *GameliftContainerFleetDeploymentConfiguration) {
	if err := g.validatePutDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutGameSessionCreationLimitPolicy(value *GameliftContainerFleetGameSessionCreationLimitPolicy) {
	if err := g.validatePutGameSessionCreationLimitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGameSessionCreationLimitPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutInstanceConnectionPortRange(value *GameliftContainerFleetInstanceConnectionPortRange) {
	if err := g.validatePutInstanceConnectionPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceConnectionPortRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutInstanceInboundPermissions(value interface{}) {
	if err := g.validatePutInstanceInboundPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceInboundPermissions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutLocations(value interface{}) {
	if err := g.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutLogConfiguration(value *GameliftContainerFleetLogConfiguration) {
	if err := g.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutScalingPolicies(value interface{}) {
	if err := g.validatePutScalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetBillingType() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameServerContainerGroupDefinitionName() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerContainerGroupDefinitionName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameServerContainerGroupsPerInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServerContainerGroupsPerInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetGameSessionCreationLimitPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetGameSessionCreationLimitPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceConnectionPortRange() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceConnectionPortRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceInboundPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceInboundPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetInstanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetMetricGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetNewGameSessionProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNewGameSessionProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetPerInstanceContainerGroupDefinitionName() {
	_jsii_.InvokeVoid(
		g,
		"resetPerInstanceContainerGroupDefinitionName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetPlayerGatewayMode() {
	_jsii_.InvokeVoid(
		g,
		"resetPlayerGatewayMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetScalingPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleet) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

