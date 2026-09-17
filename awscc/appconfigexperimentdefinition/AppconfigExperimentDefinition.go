// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appconfigexperimentdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_experiment_definition awscc_appconfig_experiment_definition}.
type AppconfigExperimentDefinition interface {
	cdktn.TerraformResource
	ApplicationId() *string
	ApplicationIdentifier() *string
	SetApplicationIdentifier(val *string)
	ApplicationIdentifierInput() *string
	AudienceDescription() *string
	SetAudienceDescription(val *string)
	AudienceDescriptionInput() *string
	AudienceRule() *string
	SetAudienceRule(val *string)
	AudienceRuleInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ConfigurationProfileIdentifier() *string
	SetConfigurationProfileIdentifier(val *string)
	ConfigurationProfileIdentifierInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Control() AppconfigExperimentDefinitionControlOutputReference
	ControlInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvironmentIdentifier() *string
	SetEnvironmentIdentifier(val *string)
	EnvironmentIdentifierInput() *string
	ExperimentDefinitionId() *string
	FlagKey() *string
	SetFlagKey(val *string)
	FlagKeyInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hypothesis() *string
	SetHypothesis(val *string)
	HypothesisInput() *string
	Id() *string
	LaunchCriteria() *string
	SetLaunchCriteria(val *string)
	LaunchCriteriaInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Status() *string
	Tags() AppconfigExperimentDefinitionTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Treatments() AppconfigExperimentDefinitionTreatmentsList
	TreatmentsInput() interface{}
	UpdatedAt() *string
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
	PutControl(value *AppconfigExperimentDefinitionControl)
	PutTags(value interface{})
	PutTreatments(value interface{})
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
	ResetAudienceDescription()
	ResetHypothesis()
	ResetLaunchCriteria()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for AppconfigExperimentDefinition
type jsiiProxy_AppconfigExperimentDefinition struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ApplicationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ApplicationIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) AudienceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) AudienceDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) AudienceRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) AudienceRuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ConfigurationProfileIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ConfigurationProfileIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Control() AppconfigExperimentDefinitionControlOutputReference {
	var returns AppconfigExperimentDefinitionControlOutputReference
	_jsii_.Get(
		j,
		"control",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) EnvironmentIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) EnvironmentIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ExperimentDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) FlagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) FlagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Hypothesis() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypothesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) HypothesisInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypothesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) LaunchCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) LaunchCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Tags() AppconfigExperimentDefinitionTagsList {
	var returns AppconfigExperimentDefinitionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) Treatments() AppconfigExperimentDefinitionTreatmentsList {
	var returns AppconfigExperimentDefinitionTreatmentsList
	_jsii_.Get(
		j,
		"treatments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) TreatmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigExperimentDefinition) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_experiment_definition awscc_appconfig_experiment_definition} Resource.
func NewAppconfigExperimentDefinition(scope constructs.Construct, id *string, config *AppconfigExperimentDefinitionConfig) AppconfigExperimentDefinition {
	_init_.Initialize()

	if err := validateNewAppconfigExperimentDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppconfigExperimentDefinition{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appconfig_experiment_definition awscc_appconfig_experiment_definition} Resource.
func NewAppconfigExperimentDefinition_Override(a AppconfigExperimentDefinition, scope constructs.Construct, id *string, config *AppconfigExperimentDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetApplicationIdentifier(val *string) {
	if err := j.validateSetApplicationIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetAudienceDescription(val *string) {
	if err := j.validateSetAudienceDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audienceDescription",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetAudienceRule(val *string) {
	if err := j.validateSetAudienceRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audienceRule",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetConfigurationProfileIdentifier(val *string) {
	if err := j.validateSetConfigurationProfileIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationProfileIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetEnvironmentIdentifier(val *string) {
	if err := j.validateSetEnvironmentIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetFlagKey(val *string) {
	if err := j.validateSetFlagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flagKey",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetHypothesis(val *string) {
	if err := j.validateSetHypothesisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hypothesis",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetLaunchCriteria(val *string) {
	if err := j.validateSetLaunchCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchCriteria",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigExperimentDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a AppconfigExperimentDefinition resource upon running "cdktn plan <stack-name>".
func AppconfigExperimentDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppconfigExperimentDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
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
func AppconfigExperimentDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppconfigExperimentDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppconfigExperimentDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppconfigExperimentDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppconfigExperimentDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppconfigExperimentDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigExperimentDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.appconfigExperimentDefinition.AppconfigExperimentDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) PutControl(value *AppconfigExperimentDefinitionControl) {
	if err := a.validatePutControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putControl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) PutTreatments(value interface{}) {
	if err := a.validatePutTreatmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTreatments",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ResetAudienceDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetAudienceDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ResetHypothesis() {
	_jsii_.InvokeVoid(
		a,
		"resetHypothesis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ResetLaunchCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigExperimentDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppconfigExperimentDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

