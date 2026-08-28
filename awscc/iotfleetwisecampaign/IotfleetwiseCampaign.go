// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotfleetwisecampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotfleetwise_campaign awscc_iotfleetwise_campaign}.
type IotfleetwiseCampaign interface {
	cdktn.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CollectionScheme() IotfleetwiseCampaignCollectionSchemeOutputReference
	CollectionSchemeInput() interface{}
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
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
	DataDestinationConfigs() IotfleetwiseCampaignDataDestinationConfigsList
	DataDestinationConfigsInput() interface{}
	DataExtraDimensions() *[]*string
	SetDataExtraDimensions(val *[]*string)
	DataExtraDimensionsInput() *[]*string
	DataPartitions() IotfleetwiseCampaignDataPartitionsList
	DataPartitionsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiagnosticsMode() *string
	SetDiagnosticsMode(val *string)
	DiagnosticsModeInput() *string
	ExpiryTime() *string
	SetExpiryTime(val *string)
	ExpiryTimeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastModificationTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PostTriggerCollectionDuration() *float64
	SetPostTriggerCollectionDuration(val *float64)
	PostTriggerCollectionDurationInput() *float64
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	SignalCatalogArn() *string
	SetSignalCatalogArn(val *string)
	SignalCatalogArnInput() *string
	SignalsToCollect() IotfleetwiseCampaignSignalsToCollectList
	SignalsToCollectInput() interface{}
	SignalsToFetch() IotfleetwiseCampaignSignalsToFetchList
	SignalsToFetchInput() interface{}
	SpoolingMode() *string
	SetSpoolingMode(val *string)
	SpoolingModeInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	Status() *string
	Tags() IotfleetwiseCampaignTagsList
	TagsInput() interface{}
	TargetArn() *string
	SetTargetArn(val *string)
	TargetArnInput() *string
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
	PutCollectionScheme(value *IotfleetwiseCampaignCollectionScheme)
	PutDataDestinationConfigs(value interface{})
	PutDataPartitions(value interface{})
	PutSignalsToCollect(value interface{})
	PutSignalsToFetch(value interface{})
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
	ResetAction()
	ResetCompression()
	ResetDataDestinationConfigs()
	ResetDataExtraDimensions()
	ResetDataPartitions()
	ResetDescription()
	ResetDiagnosticsMode()
	ResetExpiryTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostTriggerCollectionDuration()
	ResetPriority()
	ResetSignalsToCollect()
	ResetSignalsToFetch()
	ResetSpoolingMode()
	ResetStartTime()
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

// The jsii proxy struct for IotfleetwiseCampaign
type jsiiProxy_IotfleetwiseCampaign struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IotfleetwiseCampaign) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) CollectionScheme() IotfleetwiseCampaignCollectionSchemeOutputReference {
	var returns IotfleetwiseCampaignCollectionSchemeOutputReference
	_jsii_.Get(
		j,
		"collectionScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) CollectionSchemeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataDestinationConfigs() IotfleetwiseCampaignDataDestinationConfigsList {
	var returns IotfleetwiseCampaignDataDestinationConfigsList
	_jsii_.Get(
		j,
		"dataDestinationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataDestinationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDestinationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataExtraDimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataExtraDimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataExtraDimensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataExtraDimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataPartitions() IotfleetwiseCampaignDataPartitionsList {
	var returns IotfleetwiseCampaignDataPartitionsList
	_jsii_.Get(
		j,
		"dataPartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DataPartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataPartitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DiagnosticsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diagnosticsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) DiagnosticsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diagnosticsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) ExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) ExpiryTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) LastModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) PostTriggerCollectionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"postTriggerCollectionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) PostTriggerCollectionDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"postTriggerCollectionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalCatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalCatalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalCatalogArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalCatalogArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalsToCollect() IotfleetwiseCampaignSignalsToCollectList {
	var returns IotfleetwiseCampaignSignalsToCollectList
	_jsii_.Get(
		j,
		"signalsToCollect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalsToCollectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalsToCollectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalsToFetch() IotfleetwiseCampaignSignalsToFetchList {
	var returns IotfleetwiseCampaignSignalsToFetchList
	_jsii_.Get(
		j,
		"signalsToFetch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SignalsToFetchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalsToFetchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SpoolingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spoolingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) SpoolingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spoolingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) Tags() IotfleetwiseCampaignTagsList {
	var returns IotfleetwiseCampaignTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaign) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotfleetwise_campaign awscc_iotfleetwise_campaign} Resource.
func NewIotfleetwiseCampaign(scope constructs.Construct, id *string, config *IotfleetwiseCampaignConfig) IotfleetwiseCampaign {
	_init_.Initialize()

	if err := validateNewIotfleetwiseCampaignParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotfleetwiseCampaign{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotfleetwise_campaign awscc_iotfleetwise_campaign} Resource.
func NewIotfleetwiseCampaign_Override(i IotfleetwiseCampaign, scope constructs.Construct, id *string, config *IotfleetwiseCampaignConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetDataExtraDimensions(val *[]*string) {
	if err := j.validateSetDataExtraDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataExtraDimensions",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetDiagnosticsMode(val *string) {
	if err := j.validateSetDiagnosticsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diagnosticsMode",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetExpiryTime(val *string) {
	if err := j.validateSetExpiryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiryTime",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetPostTriggerCollectionDuration(val *float64) {
	if err := j.validateSetPostTriggerCollectionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postTriggerCollectionDuration",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetSignalCatalogArn(val *string) {
	if err := j.validateSetSignalCatalogArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signalCatalogArn",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetSpoolingMode(val *string) {
	if err := j.validateSetSpoolingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spoolingMode",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaign)SetTargetArn(val *string) {
	if err := j.validateSetTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

// Generates CDKTN code for importing a IotfleetwiseCampaign resource upon running "cdktn plan <stack-name>".
func IotfleetwiseCampaign_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIotfleetwiseCampaign_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
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
func IotfleetwiseCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotfleetwiseCampaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotfleetwiseCampaign_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotfleetwiseCampaign_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotfleetwiseCampaign_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotfleetwiseCampaign_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotfleetwiseCampaign_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaign",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := i.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutCollectionScheme(value *IotfleetwiseCampaignCollectionScheme) {
	if err := i.validatePutCollectionSchemeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCollectionScheme",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutDataDestinationConfigs(value interface{}) {
	if err := i.validatePutDataDestinationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDataDestinationConfigs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutDataPartitions(value interface{}) {
	if err := i.validatePutDataPartitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDataPartitions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutSignalsToCollect(value interface{}) {
	if err := i.validatePutSignalsToCollectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSignalsToCollect",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutSignalsToFetch(value interface{}) {
	if err := i.validatePutSignalsToFetchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSignalsToFetch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetCompression() {
	_jsii_.InvokeVoid(
		i,
		"resetCompression",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetDataDestinationConfigs() {
	_jsii_.InvokeVoid(
		i,
		"resetDataDestinationConfigs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetDataExtraDimensions() {
	_jsii_.InvokeVoid(
		i,
		"resetDataExtraDimensions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetDataPartitions() {
	_jsii_.InvokeVoid(
		i,
		"resetDataPartitions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetDiagnosticsMode() {
	_jsii_.InvokeVoid(
		i,
		"resetDiagnosticsMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetExpiryTime() {
	_jsii_.InvokeVoid(
		i,
		"resetExpiryTime",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetPostTriggerCollectionDuration() {
	_jsii_.InvokeVoid(
		i,
		"resetPostTriggerCollectionDuration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetPriority() {
	_jsii_.InvokeVoid(
		i,
		"resetPriority",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetSignalsToCollect() {
	_jsii_.InvokeVoid(
		i,
		"resetSignalsToCollect",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetSignalsToFetch() {
	_jsii_.InvokeVoid(
		i,
		"resetSignalsToFetch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetSpoolingMode() {
	_jsii_.InvokeVoid(
		i,
		"resetSpoolingMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetStartTime() {
	_jsii_.InvokeVoid(
		i,
		"resetStartTime",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaign) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaign) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

