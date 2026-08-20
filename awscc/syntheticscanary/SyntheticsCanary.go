// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/syntheticscanary/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary awscc_synthetics_canary}.
type SyntheticsCanary interface {
	cdktn.TerraformResource
	ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference
	ArtifactConfigInput() interface{}
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	ArtifactS3LocationInput() *string
	BrowserConfigs() SyntheticsCanaryBrowserConfigsList
	BrowserConfigsInput() interface{}
	CanaryId() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Code() SyntheticsCanaryCodeOutputReference
	CodeInput() interface{}
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
	DeleteLambdaResourcesOnCanaryDeletion() interface{}
	SetDeleteLambdaResourcesOnCanaryDeletion(val interface{})
	DeleteLambdaResourcesOnCanaryDeletionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DryRunAndUpdate() interface{}
	SetDryRunAndUpdate(val interface{})
	DryRunAndUpdateInput() interface{}
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	FailureRetentionPeriodInput() *float64
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionedResourceCleanup() *string
	SetProvisionedResourceCleanup(val *string)
	ProvisionedResourceCleanupInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourcesToReplicateTags() *[]*string
	SetResourcesToReplicateTags(val *[]*string)
	ResourcesToReplicateTagsInput() *[]*string
	RunConfig() SyntheticsCanaryRunConfigOutputReference
	RunConfigInput() interface{}
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	Schedule() SyntheticsCanaryScheduleOutputReference
	ScheduleInput() interface{}
	StartCanaryAfterCreation() interface{}
	SetStartCanaryAfterCreation(val interface{})
	StartCanaryAfterCreationInput() interface{}
	State() *string
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	SuccessRetentionPeriodInput() *float64
	Tags() SyntheticsCanaryTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VisualReference() SyntheticsCanaryVisualReferenceOutputReference
	VisualReferenceInput() interface{}
	VisualReferences() SyntheticsCanaryVisualReferencesList
	VisualReferencesInput() interface{}
	VpcConfig() SyntheticsCanaryVpcConfigOutputReference
	VpcConfigInput() interface{}
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
	PutArtifactConfig(value *SyntheticsCanaryArtifactConfig)
	PutBrowserConfigs(value interface{})
	PutCode(value *SyntheticsCanaryCode)
	PutRunConfig(value *SyntheticsCanaryRunConfig)
	PutSchedule(value *SyntheticsCanarySchedule)
	PutTags(value interface{})
	PutVisualReference(value *SyntheticsCanaryVisualReference)
	PutVisualReferences(value interface{})
	PutVpcConfig(value *SyntheticsCanaryVpcConfig)
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
	ResetArtifactConfig()
	ResetBrowserConfigs()
	ResetDeleteLambdaResourcesOnCanaryDeletion()
	ResetDryRunAndUpdate()
	ResetFailureRetentionPeriod()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProvisionedResourceCleanup()
	ResetResourcesToReplicateTags()
	ResetRunConfig()
	ResetStartCanaryAfterCreation()
	ResetSuccessRetentionPeriod()
	ResetTags()
	ResetVisualReference()
	ResetVisualReferences()
	ResetVpcConfig()
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

// The jsii proxy struct for SyntheticsCanary
type jsiiProxy_SyntheticsCanary struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfig() SyntheticsCanaryArtifactConfigOutputReference {
	var returns SyntheticsCanaryArtifactConfigOutputReference
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ArtifactS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) BrowserConfigs() SyntheticsCanaryBrowserConfigsList {
	var returns SyntheticsCanaryBrowserConfigsList
	_jsii_.Get(
		j,
		"browserConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) BrowserConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) CanaryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Code() SyntheticsCanaryCodeOutputReference {
	var returns SyntheticsCanaryCodeOutputReference
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) CodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DeleteLambdaResourcesOnCanaryDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambdaResourcesOnCanaryDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DeleteLambdaResourcesOnCanaryDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambdaResourcesOnCanaryDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DryRunAndUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dryRunAndUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) DryRunAndUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dryRunAndUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FailureRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ProvisionedResourceCleanup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourceCleanup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ProvisionedResourceCleanupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourceCleanupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ResourcesToReplicateTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToReplicateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ResourcesToReplicateTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesToReplicateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfig() SyntheticsCanaryRunConfigOutputReference {
	var returns SyntheticsCanaryRunConfigOutputReference
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RunConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Schedule() SyntheticsCanaryScheduleOutputReference {
	var returns SyntheticsCanaryScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanaryAfterCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryAfterCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) StartCanaryAfterCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryAfterCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) SuccessRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) Tags() SyntheticsCanaryTagsList {
	var returns SyntheticsCanaryTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VisualReference() SyntheticsCanaryVisualReferenceOutputReference {
	var returns SyntheticsCanaryVisualReferenceOutputReference
	_jsii_.Get(
		j,
		"visualReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VisualReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VisualReferences() SyntheticsCanaryVisualReferencesList {
	var returns SyntheticsCanaryVisualReferencesList
	_jsii_.Get(
		j,
		"visualReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VisualReferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualReferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfig() SyntheticsCanaryVpcConfigOutputReference {
	var returns SyntheticsCanaryVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanary) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary awscc_synthetics_canary} Resource.
func NewSyntheticsCanary(scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) SyntheticsCanary {
	_init_.Initialize()

	if err := validateNewSyntheticsCanaryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsCanary{}

	_jsii_.Create(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/synthetics_canary awscc_synthetics_canary} Resource.
func NewSyntheticsCanary_Override(s SyntheticsCanary, scope constructs.Construct, id *string, config *SyntheticsCanaryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetArtifactS3Location(val *string) {
	if err := j.validateSetArtifactS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetDeleteLambdaResourcesOnCanaryDeletion(val interface{}) {
	if err := j.validateSetDeleteLambdaResourcesOnCanaryDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteLambdaResourcesOnCanaryDeletion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetDryRunAndUpdate(val interface{}) {
	if err := j.validateSetDryRunAndUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dryRunAndUpdate",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetFailureRetentionPeriod(val *float64) {
	if err := j.validateSetFailureRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetProvisionedResourceCleanup(val *string) {
	if err := j.validateSetProvisionedResourceCleanupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedResourceCleanup",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetResourcesToReplicateTags(val *[]*string) {
	if err := j.validateSetResourcesToReplicateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcesToReplicateTags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetStartCanaryAfterCreation(val interface{}) {
	if err := j.validateSetStartCanaryAfterCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startCanaryAfterCreation",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCanary)SetSuccessRetentionPeriod(val *float64) {
	if err := j.validateSetSuccessRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

// Generates CDKTN code for importing a SyntheticsCanary resource upon running "cdktn plan <stack-name>".
func SyntheticsCanary_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsCanary_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
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
func SyntheticsCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCanary_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCanary_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCanary_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsCanary_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.syntheticsCanary.SyntheticsCanary",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsCanary) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsCanary) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsCanary) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := s.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCanary) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsCanary) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCanary) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutArtifactConfig(value *SyntheticsCanaryArtifactConfig) {
	if err := s.validatePutArtifactConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putArtifactConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutBrowserConfigs(value interface{}) {
	if err := s.validatePutBrowserConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBrowserConfigs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutCode(value *SyntheticsCanaryCode) {
	if err := s.validatePutCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCode",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutRunConfig(value *SyntheticsCanaryRunConfig) {
	if err := s.validatePutRunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRunConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutSchedule(value *SyntheticsCanarySchedule) {
	if err := s.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSchedule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutVisualReference(value *SyntheticsCanaryVisualReference) {
	if err := s.validatePutVisualReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVisualReference",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutVisualReferences(value interface{}) {
	if err := s.validatePutVisualReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVisualReferences",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) PutVpcConfig(value *SyntheticsCanaryVpcConfig) {
	if err := s.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCanary) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := s.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetArtifactConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetArtifactConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetBrowserConfigs() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowserConfigs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetDeleteLambdaResourcesOnCanaryDeletion() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteLambdaResourcesOnCanaryDeletion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetDryRunAndUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetDryRunAndUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetFailureRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureRetentionPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetProvisionedResourceCleanup() {
	_jsii_.InvokeVoid(
		s,
		"resetProvisionedResourceCleanup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetResourcesToReplicateTags() {
	_jsii_.InvokeVoid(
		s,
		"resetResourcesToReplicateTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetRunConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRunConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetStartCanaryAfterCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetStartCanaryAfterCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetSuccessRetentionPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessRetentionPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetVisualReference() {
	_jsii_.InvokeVoid(
		s,
		"resetVisualReference",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetVisualReferences() {
	_jsii_.InvokeVoid(
		s,
		"resetVisualReferences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCanary) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanary) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

