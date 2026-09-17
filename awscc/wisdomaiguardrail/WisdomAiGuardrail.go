// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiguardrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail awscc_wisdom_ai_guardrail}.
type WisdomAiGuardrail interface {
	cdktn.TerraformResource
	AiGuardrailArn() *string
	AiGuardrailId() *string
	AssistantArn() *string
	AssistantId() *string
	SetAssistantId(val *string)
	AssistantIdInput() *string
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
	ContentPolicyConfig() WisdomAiGuardrailContentPolicyConfigOutputReference
	ContentPolicyConfigInput() interface{}
	ContextualGroundingPolicyConfig() WisdomAiGuardrailContextualGroundingPolicyConfigOutputReference
	ContextualGroundingPolicyConfigInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	ModifiedTimeSeconds() *float64
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
	SensitiveInformationPolicyConfig() WisdomAiGuardrailSensitiveInformationPolicyConfigOutputReference
	SensitiveInformationPolicyConfigInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TopicPolicyConfig() WisdomAiGuardrailTopicPolicyConfigOutputReference
	TopicPolicyConfigInput() interface{}
	WordPolicyConfig() WisdomAiGuardrailWordPolicyConfigOutputReference
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
	PutContentPolicyConfig(value *WisdomAiGuardrailContentPolicyConfig)
	PutContextualGroundingPolicyConfig(value *WisdomAiGuardrailContextualGroundingPolicyConfig)
	PutSensitiveInformationPolicyConfig(value *WisdomAiGuardrailSensitiveInformationPolicyConfig)
	PutTopicPolicyConfig(value *WisdomAiGuardrailTopicPolicyConfig)
	PutWordPolicyConfig(value *WisdomAiGuardrailWordPolicyConfig)
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
	ResetContentPolicyConfig()
	ResetContextualGroundingPolicyConfig()
	ResetDescription()
	ResetName()
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

// The jsii proxy struct for WisdomAiGuardrail
type jsiiProxy_WisdomAiGuardrail struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WisdomAiGuardrail) AiGuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGuardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) AiGuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGuardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) AssistantArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assistantArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) AssistantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assistantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) AssistantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assistantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) BlockedInputMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedInputMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) BlockedInputMessagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedInputMessagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) BlockedOutputsMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedOutputsMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) BlockedOutputsMessagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedOutputsMessagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ContentPolicyConfig() WisdomAiGuardrailContentPolicyConfigOutputReference {
	var returns WisdomAiGuardrailContentPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"contentPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ContentPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ContextualGroundingPolicyConfig() WisdomAiGuardrailContextualGroundingPolicyConfigOutputReference {
	var returns WisdomAiGuardrailContextualGroundingPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"contextualGroundingPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ContextualGroundingPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextualGroundingPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) ModifiedTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modifiedTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) SensitiveInformationPolicyConfig() WisdomAiGuardrailSensitiveInformationPolicyConfigOutputReference {
	var returns WisdomAiGuardrailSensitiveInformationPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"sensitiveInformationPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) SensitiveInformationPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveInformationPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TopicPolicyConfig() WisdomAiGuardrailTopicPolicyConfigOutputReference {
	var returns WisdomAiGuardrailTopicPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"topicPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) TopicPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) WordPolicyConfig() WisdomAiGuardrailWordPolicyConfigOutputReference {
	var returns WisdomAiGuardrailWordPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"wordPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiGuardrail) WordPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wordPolicyConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail awscc_wisdom_ai_guardrail} Resource.
func NewWisdomAiGuardrail(scope constructs.Construct, id *string, config *WisdomAiGuardrailConfig) WisdomAiGuardrail {
	_init_.Initialize()

	if err := validateNewWisdomAiGuardrailParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiGuardrail{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail awscc_wisdom_ai_guardrail} Resource.
func NewWisdomAiGuardrail_Override(w WisdomAiGuardrail, scope constructs.Construct, id *string, config *WisdomAiGuardrailConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetAssistantId(val *string) {
	if err := j.validateSetAssistantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assistantId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetBlockedInputMessaging(val *string) {
	if err := j.validateSetBlockedInputMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedInputMessaging",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetBlockedOutputsMessaging(val *string) {
	if err := j.validateSetBlockedOutputsMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedOutputsMessaging",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WisdomAiGuardrail)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a WisdomAiGuardrail resource upon running "cdktn plan <stack-name>".
func WisdomAiGuardrail_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWisdomAiGuardrail_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
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
func WisdomAiGuardrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWisdomAiGuardrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WisdomAiGuardrail_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWisdomAiGuardrail_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WisdomAiGuardrail_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWisdomAiGuardrail_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WisdomAiGuardrail_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.wisdomAiGuardrail.WisdomAiGuardrail",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) PutContentPolicyConfig(value *WisdomAiGuardrailContentPolicyConfig) {
	if err := w.validatePutContentPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putContentPolicyConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) PutContextualGroundingPolicyConfig(value *WisdomAiGuardrailContextualGroundingPolicyConfig) {
	if err := w.validatePutContextualGroundingPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putContextualGroundingPolicyConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) PutSensitiveInformationPolicyConfig(value *WisdomAiGuardrailSensitiveInformationPolicyConfig) {
	if err := w.validatePutSensitiveInformationPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSensitiveInformationPolicyConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) PutTopicPolicyConfig(value *WisdomAiGuardrailTopicPolicyConfig) {
	if err := w.validatePutTopicPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTopicPolicyConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) PutWordPolicyConfig(value *WisdomAiGuardrailWordPolicyConfig) {
	if err := w.validatePutWordPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWordPolicyConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetContentPolicyConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetContentPolicyConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetContextualGroundingPolicyConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetContextualGroundingPolicyConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetSensitiveInformationPolicyConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetSensitiveInformationPolicyConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetTopicPolicyConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetTopicPolicyConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) ResetWordPolicyConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetWordPolicyConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiGuardrail) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiGuardrail) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

