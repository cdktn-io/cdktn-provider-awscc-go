// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/codedeploydeploymentgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group awscc_codedeploy_deployment_group}.
type CodedeployDeploymentGroup interface {
	cdktn.TerraformResource
	AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference
	AlarmConfigurationInput() interface{}
	ApplicationName() *string
	SetApplicationName(val *string)
	ApplicationNameInput() *string
	AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	AutoRollbackConfigurationInput() interface{}
	AutoScalingGroups() *[]*string
	SetAutoScalingGroups(val *[]*string)
	AutoScalingGroupsInput() *[]*string
	BlueGreenDeploymentConfiguration() CodedeployDeploymentGroupBlueGreenDeploymentConfigurationOutputReference
	BlueGreenDeploymentConfigurationInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() CodedeployDeploymentGroupDeploymentOutputReference
	DeploymentConfigName() *string
	SetDeploymentConfigName(val *string)
	DeploymentConfigNameInput() *string
	DeploymentGroupName() *string
	SetDeploymentGroupName(val *string)
	DeploymentGroupNameInput() *string
	DeploymentInput() interface{}
	DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference
	DeploymentStyleInput() interface{}
	Ec2TagFilters() CodedeployDeploymentGroupEc2TagFiltersList
	Ec2TagFiltersInput() interface{}
	Ec2TagSet() CodedeployDeploymentGroupEc2TagSetOutputReference
	Ec2TagSetInput() interface{}
	EcsServices() CodedeployDeploymentGroupEcsServicesList
	EcsServicesInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	LoadBalancerInfoInput() interface{}
	// The tree node.
	Node() constructs.Node
	OnPremisesInstanceTagFilters() CodedeployDeploymentGroupOnPremisesInstanceTagFiltersList
	OnPremisesInstanceTagFiltersInput() interface{}
	OnPremisesTagSet() CodedeployDeploymentGroupOnPremisesTagSetOutputReference
	OnPremisesTagSetInput() interface{}
	OutdatedInstancesStrategy() *string
	SetOutdatedInstancesStrategy(val *string)
	OutdatedInstancesStrategyInput() *string
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
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	Tags() CodedeployDeploymentGroupTagsList
	TagsInput() interface{}
	TerminationHookEnabled() interface{}
	SetTerminationHookEnabled(val interface{})
	TerminationHookEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TriggerConfigurations() CodedeployDeploymentGroupTriggerConfigurationsList
	TriggerConfigurationsInput() interface{}
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
	PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration)
	PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration)
	PutBlueGreenDeploymentConfiguration(value *CodedeployDeploymentGroupBlueGreenDeploymentConfiguration)
	PutDeployment(value *CodedeployDeploymentGroupDeployment)
	PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle)
	PutEc2TagFilters(value interface{})
	PutEc2TagSet(value *CodedeployDeploymentGroupEc2TagSet)
	PutEcsServices(value interface{})
	PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo)
	PutOnPremisesInstanceTagFilters(value interface{})
	PutOnPremisesTagSet(value *CodedeployDeploymentGroupOnPremisesTagSet)
	PutTags(value interface{})
	PutTriggerConfigurations(value interface{})
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
	ResetAlarmConfiguration()
	ResetAutoRollbackConfiguration()
	ResetAutoScalingGroups()
	ResetBlueGreenDeploymentConfiguration()
	ResetDeployment()
	ResetDeploymentConfigName()
	ResetDeploymentGroupName()
	ResetDeploymentStyle()
	ResetEc2TagFilters()
	ResetEc2TagSet()
	ResetEcsServices()
	ResetLoadBalancerInfo()
	ResetOnPremisesInstanceTagFilters()
	ResetOnPremisesTagSet()
	ResetOutdatedInstancesStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTerminationHookEnabled()
	ResetTriggerConfigurations()
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

// The jsii proxy struct for CodedeployDeploymentGroup
type jsiiProxy_CodedeployDeploymentGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfiguration() CodedeployDeploymentGroupAlarmConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAlarmConfigurationOutputReference
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AlarmConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfiguration() CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference {
	var returns CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoRollbackConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoScalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoScalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) AutoScalingGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoScalingGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfiguration() CodedeployDeploymentGroupBlueGreenDeploymentConfigurationOutputReference {
	var returns CodedeployDeploymentGroupBlueGreenDeploymentConfigurationOutputReference
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) BlueGreenDeploymentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Deployment() CodedeployDeploymentGroupDeploymentOutputReference {
	var returns CodedeployDeploymentGroupDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyle() CodedeployDeploymentGroupDeploymentStyleOutputReference {
	var returns CodedeployDeploymentGroupDeploymentStyleOutputReference
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) DeploymentStyleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFilters() CodedeployDeploymentGroupEc2TagFiltersList {
	var returns CodedeployDeploymentGroupEc2TagFiltersList
	_jsii_.Get(
		j,
		"ec2TagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSet() CodedeployDeploymentGroupEc2TagSetOutputReference {
	var returns CodedeployDeploymentGroupEc2TagSetOutputReference
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Ec2TagSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsServices() CodedeployDeploymentGroupEcsServicesList {
	var returns CodedeployDeploymentGroupEcsServicesList
	_jsii_.Get(
		j,
		"ecsServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) EcsServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfo() CodedeployDeploymentGroupLoadBalancerInfoOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoOutputReference
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) LoadBalancerInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFilters() CodedeployDeploymentGroupOnPremisesInstanceTagFiltersList {
	var returns CodedeployDeploymentGroupOnPremisesInstanceTagFiltersList
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesInstanceTagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesTagSet() CodedeployDeploymentGroupOnPremisesTagSetOutputReference {
	var returns CodedeployDeploymentGroupOnPremisesTagSetOutputReference
	_jsii_.Get(
		j,
		"onPremisesTagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OnPremisesTagSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesTagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OutdatedInstancesStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) OutdatedInstancesStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) Tags() CodedeployDeploymentGroupTagsList {
	var returns CodedeployDeploymentGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerminationHookEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerminationHookEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfigurations() CodedeployDeploymentGroupTriggerConfigurationsList {
	var returns CodedeployDeploymentGroupTriggerConfigurationsList
	_jsii_.Get(
		j,
		"triggerConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroup) TriggerConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfigurationsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group awscc_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup(scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) CodedeployDeploymentGroup {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentGroup{}

	_jsii_.Create(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group awscc_codedeploy_deployment_group} Resource.
func NewCodedeployDeploymentGroup_Override(c CodedeployDeploymentGroup, scope constructs.Construct, id *string, config *CodedeployDeploymentGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetApplicationName(val *string) {
	if err := j.validateSetApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetAutoScalingGroups(val *[]*string) {
	if err := j.validateSetAutoScalingGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroups",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetDeploymentConfigName(val *string) {
	if err := j.validateSetDeploymentConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetDeploymentGroupName(val *string) {
	if err := j.validateSetDeploymentGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetOutdatedInstancesStrategy(val *string) {
	if err := j.validateSetOutdatedInstancesStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outdatedInstancesStrategy",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroup)SetTerminationHookEnabled(val interface{}) {
	if err := j.validateSetTerminationHookEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationHookEnabled",
		val,
	)
}

// Generates CDKTN code for importing a CodedeployDeploymentGroup resource upon running "cdktn plan <stack-name>".
func CodedeployDeploymentGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
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
func CodedeployDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CodedeployDeploymentGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodedeployDeploymentGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodedeployDeploymentGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.codedeployDeploymentGroup.CodedeployDeploymentGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAlarmConfiguration(value *CodedeployDeploymentGroupAlarmConfiguration) {
	if err := c.validatePutAlarmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlarmConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutAutoRollbackConfiguration(value *CodedeployDeploymentGroupAutoRollbackConfiguration) {
	if err := c.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutBlueGreenDeploymentConfiguration(value *CodedeployDeploymentGroupBlueGreenDeploymentConfiguration) {
	if err := c.validatePutBlueGreenDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBlueGreenDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutDeployment(value *CodedeployDeploymentGroupDeployment) {
	if err := c.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeployment",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutDeploymentStyle(value *CodedeployDeploymentGroupDeploymentStyle) {
	if err := c.validatePutDeploymentStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeploymentStyle",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEc2TagFilters(value interface{}) {
	if err := c.validatePutEc2TagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2TagFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEc2TagSet(value *CodedeployDeploymentGroupEc2TagSet) {
	if err := c.validatePutEc2TagSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2TagSet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutEcsServices(value interface{}) {
	if err := c.validatePutEcsServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEcsServices",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutLoadBalancerInfo(value *CodedeployDeploymentGroupLoadBalancerInfo) {
	if err := c.validatePutLoadBalancerInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancerInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutOnPremisesInstanceTagFilters(value interface{}) {
	if err := c.validatePutOnPremisesInstanceTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnPremisesInstanceTagFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutOnPremisesTagSet(value *CodedeployDeploymentGroupOnPremisesTagSet) {
	if err := c.validatePutOnPremisesTagSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnPremisesTagSet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) PutTriggerConfigurations(value interface{}) {
	if err := c.validatePutTriggerConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTriggerConfigurations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAlarmConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetAutoScalingGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetBlueGreenDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetBlueGreenDeploymentConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeployment() {
	_jsii_.InvokeVoid(
		c,
		"resetDeployment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetDeploymentStyle() {
	_jsii_.InvokeVoid(
		c,
		"resetDeploymentStyle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEc2TagSet() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2TagSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetEcsServices() {
	_jsii_.InvokeVoid(
		c,
		"resetEcsServices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetLoadBalancerInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancerInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOnPremisesInstanceTagFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetOnPremisesInstanceTagFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOnPremisesTagSet() {
	_jsii_.InvokeVoid(
		c,
		"resetOnPremisesTagSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOutdatedInstancesStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetOutdatedInstancesStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTerminationHookEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationHookEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ResetTriggerConfigurations() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerConfigurations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

