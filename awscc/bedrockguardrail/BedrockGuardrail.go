// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockguardrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_guardrail awscc_bedrock_guardrail}.
type BedrockGuardrail interface {
	cdktn.TerraformResource
	AutomatedReasoningPolicyConfig() BedrockGuardrailAutomatedReasoningPolicyConfigOutputReference
	AutomatedReasoningPolicyConfigInput() interface{}
	BlockedInputMessaging() *string
	SetBlockedInputMessaging(val *string)
	BlockedInputMessagingInput() *string
	BlockedOutputsMessaging() *string
	SetBlockedOutputsMessaging(val *string)
	BlockedOutputsMessagingInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentPolicyConfig() BedrockGuardrailContentPolicyConfigOutputReference
	ContentPolicyConfigInput() interface{}
	ContextualGroundingPolicyConfig() BedrockGuardrailContextualGroundingPolicyConfigOutputReference
	ContextualGroundingPolicyConfigInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	CrossRegionConfig() BedrockGuardrailCrossRegionConfigOutputReference
	CrossRegionConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FailureRecommendations() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuardrailArn() *string
	GuardrailId() *string
	Id() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
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
	SensitiveInformationPolicyConfig() BedrockGuardrailSensitiveInformationPolicyConfigOutputReference
	SensitiveInformationPolicyConfigInput() interface{}
	Status() *string
	StatusReasons() *[]*string
	Tags() BedrockGuardrailTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TopicPolicyConfig() BedrockGuardrailTopicPolicyConfigOutputReference
	TopicPolicyConfigInput() interface{}
	UpdatedAt() *string
	Version() *string
	WordPolicyConfig() BedrockGuardrailWordPolicyConfigOutputReference
	WordPolicyConfigInput() interface{}
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
	PutAutomatedReasoningPolicyConfig(value *BedrockGuardrailAutomatedReasoningPolicyConfig)
	PutContentPolicyConfig(value *BedrockGuardrailContentPolicyConfig)
	PutContextualGroundingPolicyConfig(value *BedrockGuardrailContextualGroundingPolicyConfig)
	PutCrossRegionConfig(value *BedrockGuardrailCrossRegionConfig)
	PutSensitiveInformationPolicyConfig(value *BedrockGuardrailSensitiveInformationPolicyConfig)
	PutTags(value interface{})
	PutTopicPolicyConfig(value *BedrockGuardrailTopicPolicyConfig)
	PutWordPolicyConfig(value *BedrockGuardrailWordPolicyConfig)
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
	ResetAutomatedReasoningPolicyConfig()
	ResetContentPolicyConfig()
	ResetContextualGroundingPolicyConfig()
	ResetCrossRegionConfig()
	ResetDescription()
	ResetKmsKeyArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSensitiveInformationPolicyConfig()
	ResetTags()
	ResetTopicPolicyConfig()
	ResetWordPolicyConfig()
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

// The jsii proxy struct for BedrockGuardrail
type jsiiProxy_BedrockGuardrail struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BedrockGuardrail) AutomatedReasoningPolicyConfig() BedrockGuardrailAutomatedReasoningPolicyConfigOutputReference {
	var returns BedrockGuardrailAutomatedReasoningPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"automatedReasoningPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) AutomatedReasoningPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatedReasoningPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) BlockedInputMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedInputMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) BlockedInputMessagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedInputMessagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) BlockedOutputsMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedOutputsMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) BlockedOutputsMessagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedOutputsMessagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ContentPolicyConfig() BedrockGuardrailContentPolicyConfigOutputReference {
	var returns BedrockGuardrailContentPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"contentPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ContentPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ContextualGroundingPolicyConfig() BedrockGuardrailContextualGroundingPolicyConfigOutputReference {
	var returns BedrockGuardrailContextualGroundingPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"contextualGroundingPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ContextualGroundingPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextualGroundingPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) CrossRegionConfig() BedrockGuardrailCrossRegionConfigOutputReference {
	var returns BedrockGuardrailCrossRegionConfigOutputReference
	_jsii_.Get(
		j,
		"crossRegionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) CrossRegionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) FailureRecommendations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failureRecommendations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) GuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) GuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) SensitiveInformationPolicyConfig() BedrockGuardrailSensitiveInformationPolicyConfigOutputReference {
	var returns BedrockGuardrailSensitiveInformationPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"sensitiveInformationPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) SensitiveInformationPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveInformationPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) StatusReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Tags() BedrockGuardrailTagsList {
	var returns BedrockGuardrailTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TopicPolicyConfig() BedrockGuardrailTopicPolicyConfigOutputReference {
	var returns BedrockGuardrailTopicPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"topicPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) TopicPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) WordPolicyConfig() BedrockGuardrailWordPolicyConfigOutputReference {
	var returns BedrockGuardrailWordPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"wordPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockGuardrail) WordPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wordPolicyConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_guardrail awscc_bedrock_guardrail} Resource.
func NewBedrockGuardrail(scope constructs.Construct, id *string, config *BedrockGuardrailConfig) BedrockGuardrail {
	_init_.Initialize()

	if err := validateNewBedrockGuardrailParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockGuardrail{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_guardrail awscc_bedrock_guardrail} Resource.
func NewBedrockGuardrail_Override(b BedrockGuardrail, scope constructs.Construct, id *string, config *BedrockGuardrailConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetBlockedInputMessaging(val *string) {
	if err := j.validateSetBlockedInputMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedInputMessaging",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetBlockedOutputsMessaging(val *string) {
	if err := j.validateSetBlockedOutputsMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedOutputsMessaging",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockGuardrail)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a BedrockGuardrail resource upon running "cdktn plan <stack-name>".
func BedrockGuardrail_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockGuardrail_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
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
func BedrockGuardrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockGuardrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockGuardrail_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockGuardrail_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockGuardrail_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockGuardrail_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockGuardrail_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.bedrockGuardrail.BedrockGuardrail",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockGuardrail) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockGuardrail) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockGuardrail) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := b.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockGuardrail) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockGuardrail) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockGuardrail) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutAutomatedReasoningPolicyConfig(value *BedrockGuardrailAutomatedReasoningPolicyConfig) {
	if err := b.validatePutAutomatedReasoningPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAutomatedReasoningPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutContentPolicyConfig(value *BedrockGuardrailContentPolicyConfig) {
	if err := b.validatePutContentPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContentPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutContextualGroundingPolicyConfig(value *BedrockGuardrailContextualGroundingPolicyConfig) {
	if err := b.validatePutContextualGroundingPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContextualGroundingPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutCrossRegionConfig(value *BedrockGuardrailCrossRegionConfig) {
	if err := b.validatePutCrossRegionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCrossRegionConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutSensitiveInformationPolicyConfig(value *BedrockGuardrailSensitiveInformationPolicyConfig) {
	if err := b.validatePutSensitiveInformationPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSensitiveInformationPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutTags(value interface{}) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutTopicPolicyConfig(value *BedrockGuardrailTopicPolicyConfig) {
	if err := b.validatePutTopicPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTopicPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) PutWordPolicyConfig(value *BedrockGuardrailWordPolicyConfig) {
	if err := b.validatePutWordPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putWordPolicyConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockGuardrail) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := b.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetAutomatedReasoningPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetAutomatedReasoningPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetContentPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetContentPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetContextualGroundingPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetContextualGroundingPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetCrossRegionConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetCrossRegionConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		b,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetSensitiveInformationPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetSensitiveInformationPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetTopicPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetTopicPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) ResetWordPolicyConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetWordPolicyConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockGuardrail) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockGuardrail) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}

