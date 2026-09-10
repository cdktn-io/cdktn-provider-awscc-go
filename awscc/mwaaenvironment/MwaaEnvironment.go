// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mwaaenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mwaa_environment awscc_mwaa_environment}.
type MwaaEnvironment interface {
	cdktn.TerraformResource
	AirflowConfigurationOptions() *string
	SetAirflowConfigurationOptions(val *string)
	AirflowConfigurationOptionsInput() *string
	AirflowVersion() *string
	SetAirflowVersion(val *string)
	AirflowVersionInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CeleryExecutorQueue() *string
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
	DagS3Path() *string
	SetDagS3Path(val *string)
	DagS3PathInput() *string
	DatabaseVpcEndpointService() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointManagement() *string
	SetEndpointManagement(val *string)
	EndpointManagementInput() *string
	EnvironmentClass() *string
	SetEnvironmentClass(val *string)
	EnvironmentClassInput() *string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoggingConfiguration() MwaaEnvironmentLoggingConfigurationOutputReference
	LoggingConfigurationInput() interface{}
	MaxWebservers() *float64
	SetMaxWebservers(val *float64)
	MaxWebserversInput() *float64
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MaxWorkersInput() *float64
	MinWebservers() *float64
	SetMinWebservers(val *float64)
	MinWebserversInput() *float64
	MinWorkers() *float64
	SetMinWorkers(val *float64)
	MinWorkersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfiguration() MwaaEnvironmentNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	PluginsS3ObjectVersion() *string
	SetPluginsS3ObjectVersion(val *string)
	PluginsS3ObjectVersionInput() *string
	PluginsS3Path() *string
	SetPluginsS3Path(val *string)
	PluginsS3PathInput() *string
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
	RequirementsS3ObjectVersion() *string
	SetRequirementsS3ObjectVersion(val *string)
	RequirementsS3ObjectVersionInput() *string
	RequirementsS3Path() *string
	SetRequirementsS3Path(val *string)
	RequirementsS3PathInput() *string
	Schedulers() *float64
	SetSchedulers(val *float64)
	SchedulersInput() *float64
	SourceBucketArn() *string
	SetSourceBucketArn(val *string)
	SourceBucketArnInput() *string
	StartupScriptS3ObjectVersion() *string
	SetStartupScriptS3ObjectVersion(val *string)
	StartupScriptS3ObjectVersionInput() *string
	StartupScriptS3Path() *string
	SetStartupScriptS3Path(val *string)
	StartupScriptS3PathInput() *string
	Tags() *string
	SetTags(val *string)
	TagsInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebserverAccessMode() *string
	SetWebserverAccessMode(val *string)
	WebserverAccessModeInput() *string
	WebserverUrl() *string
	WebserverVpcEndpointService() *string
	WeeklyMaintenanceWindowStart() *string
	SetWeeklyMaintenanceWindowStart(val *string)
	WeeklyMaintenanceWindowStartInput() *string
	WorkerReplacementStrategy() *string
	SetWorkerReplacementStrategy(val *string)
	WorkerReplacementStrategyInput() *string
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
	PutLoggingConfiguration(value *MwaaEnvironmentLoggingConfiguration)
	PutNetworkConfiguration(value *MwaaEnvironmentNetworkConfiguration)
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
	ResetAirflowConfigurationOptions()
	ResetAirflowVersion()
	ResetDagS3Path()
	ResetEndpointManagement()
	ResetEnvironmentClass()
	ResetExecutionRoleArn()
	ResetKmsKey()
	ResetLoggingConfiguration()
	ResetMaxWebservers()
	ResetMaxWorkers()
	ResetMinWebservers()
	ResetMinWorkers()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginsS3ObjectVersion()
	ResetPluginsS3Path()
	ResetRequirementsS3ObjectVersion()
	ResetRequirementsS3Path()
	ResetSchedulers()
	ResetSourceBucketArn()
	ResetStartupScriptS3ObjectVersion()
	ResetStartupScriptS3Path()
	ResetTags()
	ResetWebserverAccessMode()
	ResetWeeklyMaintenanceWindowStart()
	ResetWorkerReplacementStrategy()
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

// The jsii proxy struct for MwaaEnvironment
type jsiiProxy_MwaaEnvironment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_MwaaEnvironment) AirflowConfigurationOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowConfigurationOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowConfigurationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) AirflowVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) CeleryExecutorQueue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"celeryExecutorQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DagS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EndpointManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EndpointManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EnvironmentClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) EnvironmentClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) LoggingConfiguration() MwaaEnvironmentLoggingConfigurationOutputReference {
	var returns MwaaEnvironmentLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) LoggingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWebserversInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWebserversInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) MinWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NetworkConfiguration() MwaaEnvironmentNetworkConfigurationOutputReference {
	var returns MwaaEnvironmentNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) PluginsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) RequirementsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) SchedulersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) SourceBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) SourceBucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) StartupScriptS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) StartupScriptS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) Tags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WeeklyMaintenanceWindowStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WorkerReplacementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerReplacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwaaEnvironment) WorkerReplacementStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerReplacementStrategyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mwaa_environment awscc_mwaa_environment} Resource.
func NewMwaaEnvironment(scope constructs.Construct, id *string, config *MwaaEnvironmentConfig) MwaaEnvironment {
	_init_.Initialize()

	if err := validateNewMwaaEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwaaEnvironment{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mwaa_environment awscc_mwaa_environment} Resource.
func NewMwaaEnvironment_Override(m MwaaEnvironment, scope constructs.Construct, id *string, config *MwaaEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetAirflowConfigurationOptions(val *string) {
	if err := j.validateSetAirflowConfigurationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowConfigurationOptions",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetAirflowVersion(val *string) {
	if err := j.validateSetAirflowVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetDagS3Path(val *string) {
	if err := j.validateSetDagS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dagS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetEndpointManagement(val *string) {
	if err := j.validateSetEndpointManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointManagement",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetEnvironmentClass(val *string) {
	if err := j.validateSetEnvironmentClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentClass",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetMaxWebservers(val *float64) {
	if err := j.validateSetMaxWebserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWebservers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetMaxWorkers(val *float64) {
	if err := j.validateSetMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetMinWebservers(val *float64) {
	if err := j.validateSetMinWebserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWebservers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetMinWorkers(val *float64) {
	if err := j.validateSetMinWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWorkers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetPluginsS3ObjectVersion(val *string) {
	if err := j.validateSetPluginsS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetPluginsS3Path(val *string) {
	if err := j.validateSetPluginsS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetRequirementsS3ObjectVersion(val *string) {
	if err := j.validateSetRequirementsS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirementsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetRequirementsS3Path(val *string) {
	if err := j.validateSetRequirementsS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirementsS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetSchedulers(val *float64) {
	if err := j.validateSetSchedulersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulers",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetSourceBucketArn(val *string) {
	if err := j.validateSetSourceBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBucketArn",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetStartupScriptS3ObjectVersion(val *string) {
	if err := j.validateSetStartupScriptS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScriptS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetStartupScriptS3Path(val *string) {
	if err := j.validateSetStartupScriptS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScriptS3Path",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetTags(val *string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetWebserverAccessMode(val *string) {
	if err := j.validateSetWebserverAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webserverAccessMode",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetWeeklyMaintenanceWindowStart(val *string) {
	if err := j.validateSetWeeklyMaintenanceWindowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceWindowStart",
		val,
	)
}

func (j *jsiiProxy_MwaaEnvironment)SetWorkerReplacementStrategy(val *string) {
	if err := j.validateSetWorkerReplacementStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerReplacementStrategy",
		val,
	)
}

// Generates CDKTN code for importing a MwaaEnvironment resource upon running "cdktn plan <stack-name>".
func MwaaEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMwaaEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
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
func MwaaEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwaaEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwaaEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwaaEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MwaaEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMwaaEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MwaaEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.mwaaEnvironment.MwaaEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MwaaEnvironment) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MwaaEnvironment) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MwaaEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MwaaEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := m.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwaaEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MwaaEnvironment) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MwaaEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MwaaEnvironment) PutLoggingConfiguration(value *MwaaEnvironmentLoggingConfiguration) {
	if err := m.validatePutLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLoggingConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironment) PutNetworkConfiguration(value *MwaaEnvironmentNetworkConfiguration) {
	if err := m.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MwaaEnvironment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := m.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetAirflowConfigurationOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetAirflowConfigurationOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetAirflowVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetAirflowVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetDagS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetDagS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetEndpointManagement() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpointManagement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetEnvironmentClass() {
	_jsii_.InvokeVoid(
		m,
		"resetEnvironmentClass",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		m,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetKmsKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetLoggingConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetLoggingConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMaxWebservers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxWebservers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMinWebservers() {
	_jsii_.InvokeVoid(
		m,
		"resetMinWebservers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetMinWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMinWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetPluginsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginsS3ObjectVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetPluginsS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginsS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetRequirementsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetRequirementsS3ObjectVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetRequirementsS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetRequirementsS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetSchedulers() {
	_jsii_.InvokeVoid(
		m,
		"resetSchedulers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetSourceBucketArn() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceBucketArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetStartupScriptS3ObjectVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetStartupScriptS3ObjectVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetStartupScriptS3Path() {
	_jsii_.InvokeVoid(
		m,
		"resetStartupScriptS3Path",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetWebserverAccessMode() {
	_jsii_.InvokeVoid(
		m,
		"resetWebserverAccessMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetWeeklyMaintenanceWindowStart() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklyMaintenanceWindowStart",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) ResetWorkerReplacementStrategy() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkerReplacementStrategy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwaaEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwaaEnvironment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

