// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/rdsdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment awscc_rds_deployment}.
type RdsDeployment interface {
	cdktn.TerraformResource
	Arn() *string
	BlueGreenDeploymentIdentifier() *string
	BlueGreenDeploymentName() *string
	SetBlueGreenDeploymentName(val *string)
	BlueGreenDeploymentNameInput() *string
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// The tree node.
	Node() constructs.Node
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
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Status() *string
	SwitchoverDetails() RdsDeploymentSwitchoverDetailsList
	Tags() RdsDeploymentTagsList
	TagsInput() interface{}
	Target() *string
	TargetAllocatedStorage() *float64
	SetTargetAllocatedStorage(val *float64)
	TargetAllocatedStorageInput() *float64
	TargetDbClusterParameterGroupName() *string
	SetTargetDbClusterParameterGroupName(val *string)
	TargetDbClusterParameterGroupNameInput() *string
	TargetDbInstanceClass() *string
	SetTargetDbInstanceClass(val *string)
	TargetDbInstanceClassInput() *string
	TargetDbParameterGroupName() *string
	SetTargetDbParameterGroupName(val *string)
	TargetDbParameterGroupNameInput() *string
	TargetEngineVersion() *string
	SetTargetEngineVersion(val *string)
	TargetEngineVersionInput() *string
	TargetIops() *float64
	SetTargetIops(val *float64)
	TargetIopsInput() *float64
	TargetStorageThroughput() *float64
	SetTargetStorageThroughput(val *float64)
	TargetStorageThroughputInput() *float64
	TargetStorageType() *string
	SetTargetStorageType(val *string)
	TargetStorageTypeInput() *string
	Tasks() RdsDeploymentTasksList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpgradeTargetStorageConfig() interface{}
	SetUpgradeTargetStorageConfig(val interface{})
	UpgradeTargetStorageConfigInput() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTargetAllocatedStorage()
	ResetTargetDbClusterParameterGroupName()
	ResetTargetDbInstanceClass()
	ResetTargetDbParameterGroupName()
	ResetTargetEngineVersion()
	ResetTargetIops()
	ResetTargetStorageThroughput()
	ResetTargetStorageType()
	ResetUpgradeTargetStorageConfig()
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

// The jsii proxy struct for RdsDeployment
type jsiiProxy_RdsDeployment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RdsDeployment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) BlueGreenDeploymentIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueGreenDeploymentIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) BlueGreenDeploymentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueGreenDeploymentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) BlueGreenDeploymentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueGreenDeploymentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) SwitchoverDetails() RdsDeploymentSwitchoverDetailsList {
	var returns RdsDeploymentSwitchoverDetailsList
	_jsii_.Get(
		j,
		"switchoverDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Tags() RdsDeploymentTagsList {
	var returns RdsDeploymentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetDbParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetEngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEngineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetStorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStorageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TargetStorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStorageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) Tasks() RdsDeploymentTasksList {
	var returns RdsDeploymentTasksList
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) UpgradeTargetStorageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeTargetStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDeployment) UpgradeTargetStorageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeTargetStorageConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment awscc_rds_deployment} Resource.
func NewRdsDeployment(scope constructs.Construct, id *string, config *RdsDeploymentConfig) RdsDeployment {
	_init_.Initialize()

	if err := validateNewRdsDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDeployment{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/rds_deployment awscc_rds_deployment} Resource.
func NewRdsDeployment_Override(r RdsDeployment, scope constructs.Construct, id *string, config *RdsDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsDeployment)SetBlueGreenDeploymentName(val *string) {
	if err := j.validateSetBlueGreenDeploymentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blueGreenDeploymentName",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetAllocatedStorage(val *float64) {
	if err := j.validateSetTargetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetTargetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetDbInstanceClass(val *string) {
	if err := j.validateSetTargetDbInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetDbParameterGroupName(val *string) {
	if err := j.validateSetTargetDbParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetEngineVersion(val *string) {
	if err := j.validateSetTargetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetEngineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetIops(val *float64) {
	if err := j.validateSetTargetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIops",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetStorageThroughput(val *float64) {
	if err := j.validateSetTargetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetTargetStorageType(val *string) {
	if err := j.validateSetTargetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStorageType",
		val,
	)
}

func (j *jsiiProxy_RdsDeployment)SetUpgradeTargetStorageConfig(val interface{}) {
	if err := j.validateSetUpgradeTargetStorageConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeTargetStorageConfig",
		val,
	)
}

// Generates CDKTN code for importing a RdsDeployment resource upon running "cdktn plan <stack-name>".
func RdsDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRdsDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
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
func RdsDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.rdsDeployment.RdsDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDeployment) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsDeployment) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsDeployment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsDeployment) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsDeployment) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDeployment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RdsDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetDbInstanceClass() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetDbInstanceClass",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetDbParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetDbParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetEngineVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetEngineVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetIops() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetIops",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetStorageThroughput() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetStorageThroughput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetTargetStorageType() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetStorageType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) ResetUpgradeTargetStorageConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetUpgradeTargetStorageConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDeployment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

