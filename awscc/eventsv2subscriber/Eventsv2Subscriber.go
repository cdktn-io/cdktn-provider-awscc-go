// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsv2subscriber/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber awscc_eventsv2_subscriber}.
type Eventsv2Subscriber interface {
	cdktn.TerraformResource
	BatchConfiguration() Eventsv2SubscriberBatchConfigurationOutputReference
	BatchConfigurationInput() interface{}
	BusName() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EventBusArn() *string
	SetEventBusArn(val *string)
	EventBusArnInput() *string
	FilterConfiguration() Eventsv2SubscriberFilterConfigurationOutputReference
	FilterConfigurationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InvokeConfiguration() Eventsv2SubscriberInvokeConfigurationOutputReference
	InvokeConfigurationInput() interface{}
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogConfiguration() Eventsv2SubscriberLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnFailureConfiguration() Eventsv2SubscriberOnFailureConfigurationOutputReference
	OnFailureConfigurationInput() interface{}
	PointInTimeConfiguration() Eventsv2SubscriberPointInTimeConfigurationOutputReference
	PointInTimeConfigurationInput() interface{}
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
	ResumePosition() *string
	SetResumePosition(val *string)
	ResumePositionInput() *string
	RetryPolicy() Eventsv2SubscriberRetryPolicyOutputReference
	RetryPolicyInput() interface{}
	StartingPosition() *string
	SetStartingPosition(val *string)
	StartingPositionInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	SubscriberArn() *string
	Tags() Eventsv2SubscriberTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Transformer() Eventsv2SubscriberTransformerOutputReference
	TransformerInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutBatchConfiguration(value *Eventsv2SubscriberBatchConfiguration)
	PutFilterConfiguration(value *Eventsv2SubscriberFilterConfiguration)
	PutInvokeConfiguration(value *Eventsv2SubscriberInvokeConfiguration)
	PutLogConfiguration(value *Eventsv2SubscriberLogConfiguration)
	PutOnFailureConfiguration(value *Eventsv2SubscriberOnFailureConfiguration)
	PutPointInTimeConfiguration(value *Eventsv2SubscriberPointInTimeConfiguration)
	PutRetryPolicy(value *Eventsv2SubscriberRetryPolicy)
	PutTags(value interface{})
	PutTransformer(value *Eventsv2SubscriberTransformer)
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
	ResetBatchConfiguration()
	ResetDescription()
	ResetFilterConfiguration()
	ResetLogConfiguration()
	ResetOnFailureConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeConfiguration()
	ResetResumePosition()
	ResetRetryPolicy()
	ResetStartingPosition()
	ResetState()
	ResetTags()
	ResetTransformer()
	ResetType()
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

// The jsii proxy struct for Eventsv2Subscriber
type jsiiProxy_Eventsv2Subscriber struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Eventsv2Subscriber) BatchConfiguration() Eventsv2SubscriberBatchConfigurationOutputReference {
	var returns Eventsv2SubscriberBatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"batchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) BatchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) BusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) EventBusArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) EventBusArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) FilterConfiguration() Eventsv2SubscriberFilterConfigurationOutputReference {
	var returns Eventsv2SubscriberFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) InvokeConfiguration() Eventsv2SubscriberInvokeConfigurationOutputReference {
	var returns Eventsv2SubscriberInvokeConfigurationOutputReference
	_jsii_.Get(
		j,
		"invokeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) InvokeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) LogConfiguration() Eventsv2SubscriberLogConfigurationOutputReference {
	var returns Eventsv2SubscriberLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) OnFailureConfiguration() Eventsv2SubscriberOnFailureConfigurationOutputReference {
	var returns Eventsv2SubscriberOnFailureConfigurationOutputReference
	_jsii_.Get(
		j,
		"onFailureConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) OnFailureConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onFailureConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) PointInTimeConfiguration() Eventsv2SubscriberPointInTimeConfigurationOutputReference {
	var returns Eventsv2SubscriberPointInTimeConfigurationOutputReference
	_jsii_.Get(
		j,
		"pointInTimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) PointInTimeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) ResumePosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resumePosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) ResumePositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resumePositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) RetryPolicy() Eventsv2SubscriberRetryPolicyOutputReference {
	var returns Eventsv2SubscriberRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) RetryPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) SubscriberArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriberArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Tags() Eventsv2SubscriberTagsList {
	var returns Eventsv2SubscriberTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Transformer() Eventsv2SubscriberTransformerOutputReference {
	var returns Eventsv2SubscriberTransformerOutputReference
	_jsii_.Get(
		j,
		"transformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TransformerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eventsv2Subscriber) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber awscc_eventsv2_subscriber} Resource.
func NewEventsv2Subscriber(scope constructs.Construct, id *string, config *Eventsv2SubscriberConfig) Eventsv2Subscriber {
	_init_.Initialize()

	if err := validateNewEventsv2SubscriberParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Eventsv2Subscriber{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber awscc_eventsv2_subscriber} Resource.
func NewEventsv2Subscriber_Override(e Eventsv2Subscriber, scope constructs.Construct, id *string, config *Eventsv2SubscriberConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetEventBusArn(val *string) {
	if err := j.validateSetEventBusArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBusArn",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetResumePosition(val *string) {
	if err := j.validateSetResumePositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resumePosition",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Eventsv2Subscriber)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a Eventsv2Subscriber resource upon running "cdktn plan <stack-name>".
func Eventsv2Subscriber_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEventsv2Subscriber_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
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
func Eventsv2Subscriber_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventsv2Subscriber_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Eventsv2Subscriber_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventsv2Subscriber_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Eventsv2Subscriber_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventsv2Subscriber_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Eventsv2Subscriber_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.eventsv2Subscriber.Eventsv2Subscriber",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Eventsv2Subscriber) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Eventsv2Subscriber) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Eventsv2Subscriber) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_Eventsv2Subscriber) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutBatchConfiguration(value *Eventsv2SubscriberBatchConfiguration) {
	if err := e.validatePutBatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBatchConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutFilterConfiguration(value *Eventsv2SubscriberFilterConfiguration) {
	if err := e.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutInvokeConfiguration(value *Eventsv2SubscriberInvokeConfiguration) {
	if err := e.validatePutInvokeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInvokeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutLogConfiguration(value *Eventsv2SubscriberLogConfiguration) {
	if err := e.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutOnFailureConfiguration(value *Eventsv2SubscriberOnFailureConfiguration) {
	if err := e.validatePutOnFailureConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnFailureConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutPointInTimeConfiguration(value *Eventsv2SubscriberPointInTimeConfiguration) {
	if err := e.validatePutPointInTimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPointInTimeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutRetryPolicy(value *Eventsv2SubscriberRetryPolicy) {
	if err := e.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) PutTransformer(value *Eventsv2SubscriberTransformer) {
	if err := e.validatePutTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTransformer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetBatchConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetOnFailureConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetOnFailureConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetPointInTimeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetPointInTimeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetResumePosition() {
	_jsii_.InvokeVoid(
		e,
		"resetResumePosition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		e,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetState() {
	_jsii_.InvokeVoid(
		e,
		"resetState",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetTransformer() {
	_jsii_.InvokeVoid(
		e,
		"resetTransformer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eventsv2Subscriber) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eventsv2Subscriber) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

