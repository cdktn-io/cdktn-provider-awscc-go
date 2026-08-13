// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent awscc_bedrock_agent}.
type BedrockAgent interface {
	cdktn.TerraformResource
	ActionGroups() BedrockAgentActionGroupsList
	ActionGroupsInput() interface{}
	AgentArn() *string
	AgentCollaboration() *string
	SetAgentCollaboration(val *string)
	AgentCollaborationInput() *string
	AgentCollaborators() BedrockAgentAgentCollaboratorsList
	AgentCollaboratorsInput() interface{}
	AgentId() *string
	AgentName() *string
	SetAgentName(val *string)
	AgentNameInput() *string
	AgentResourceRoleArn() *string
	SetAgentResourceRoleArn(val *string)
	AgentResourceRoleArnInput() *string
	AgentStatus() *string
	AgentVersion() *string
	AutoPrepare() interface{}
	SetAutoPrepare(val interface{})
	AutoPrepareInput() interface{}
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
	CreatedAt() *string
	CustomerEncryptionKeyArn() *string
	SetCustomerEncryptionKeyArn(val *string)
	CustomerEncryptionKeyArnInput() *string
	CustomOrchestration() BedrockAgentCustomOrchestrationOutputReference
	CustomOrchestrationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FailureReasons() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	FoundationModel() *string
	SetFoundationModel(val *string)
	FoundationModelInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuardrailConfiguration() BedrockAgentGuardrailConfigurationOutputReference
	GuardrailConfigurationInput() interface{}
	Id() *string
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	IdleSessionTtlInSecondsInput() *float64
	Instruction() *string
	SetInstruction(val *string)
	InstructionInput() *string
	KnowledgeBases() BedrockAgentKnowledgeBasesList
	KnowledgeBasesInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MemoryConfiguration() BedrockAgentMemoryConfigurationOutputReference
	MemoryConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	OrchestrationType() *string
	SetOrchestrationType(val *string)
	OrchestrationTypeInput() *string
	PreparedAt() *string
	PromptOverrideConfiguration() BedrockAgentPromptOverrideConfigurationOutputReference
	PromptOverrideConfigurationInput() interface{}
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
	RecommendedActions() *[]*string
	SkipResourceInUseCheckOnDelete() interface{}
	SetSkipResourceInUseCheckOnDelete(val interface{})
	SkipResourceInUseCheckOnDeleteInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestAliasTags() *map[string]*string
	SetTestAliasTags(val *map[string]*string)
	TestAliasTagsInput() *map[string]*string
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
	PutActionGroups(value interface{})
	PutAgentCollaborators(value interface{})
	PutCustomOrchestration(value *BedrockAgentCustomOrchestration)
	PutGuardrailConfiguration(value *BedrockAgentGuardrailConfiguration)
	PutKnowledgeBases(value interface{})
	PutMemoryConfiguration(value *BedrockAgentMemoryConfiguration)
	PutPromptOverrideConfiguration(value *BedrockAgentPromptOverrideConfiguration)
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
	ResetActionGroups()
	ResetAgentCollaboration()
	ResetAgentCollaborators()
	ResetAgentResourceRoleArn()
	ResetAutoPrepare()
	ResetCustomerEncryptionKeyArn()
	ResetCustomOrchestration()
	ResetDescription()
	ResetFoundationModel()
	ResetGuardrailConfiguration()
	ResetIdleSessionTtlInSeconds()
	ResetInstruction()
	ResetKnowledgeBases()
	ResetMemoryConfiguration()
	ResetOrchestrationType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPromptOverrideConfiguration()
	ResetSkipResourceInUseCheckOnDelete()
	ResetTags()
	ResetTestAliasTags()
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

// The jsii proxy struct for BedrockAgent
type jsiiProxy_BedrockAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BedrockAgent) ActionGroups() BedrockAgentActionGroupsList {
	var returns BedrockAgentActionGroupsList
	_jsii_.Get(
		j,
		"actionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) ActionGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentCollaboration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaboration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentCollaborationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaborationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentCollaborators() BedrockAgentAgentCollaboratorsList {
	var returns BedrockAgentAgentCollaboratorsList
	_jsii_.Get(
		j,
		"agentCollaborators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentCollaboratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentCollaboratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentResourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentResourceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AutoPrepare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPrepare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AutoPrepareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPrepareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CustomerEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CustomerEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CustomOrchestration() BedrockAgentCustomOrchestrationOutputReference {
	var returns BedrockAgentCustomOrchestrationOutputReference
	_jsii_.Get(
		j,
		"customOrchestration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) CustomOrchestrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOrchestrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) FailureReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"failureReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) FoundationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) GuardrailConfiguration() BedrockAgentGuardrailConfigurationOutputReference {
	var returns BedrockAgentGuardrailConfigurationOutputReference
	_jsii_.Get(
		j,
		"guardrailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) GuardrailConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) KnowledgeBases() BedrockAgentKnowledgeBasesList {
	var returns BedrockAgentKnowledgeBasesList
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) KnowledgeBasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) MemoryConfiguration() BedrockAgentMemoryConfigurationOutputReference {
	var returns BedrockAgentMemoryConfigurationOutputReference
	_jsii_.Get(
		j,
		"memoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) MemoryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) OrchestrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) OrchestrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) PreparedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preparedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) PromptOverrideConfiguration() BedrockAgentPromptOverrideConfigurationOutputReference {
	var returns BedrockAgentPromptOverrideConfigurationOutputReference
	_jsii_.Get(
		j,
		"promptOverrideConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) PromptOverrideConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptOverrideConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) RecommendedActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recommendedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) SkipResourceInUseCheckOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) SkipResourceInUseCheckOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TestAliasTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"testAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) TestAliasTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"testAliasTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent awscc_bedrock_agent} Resource.
func NewBedrockAgent(scope constructs.Construct, id *string, config *BedrockAgentConfig) BedrockAgent {
	_init_.Initialize()

	if err := validateNewBedrockAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgent{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent awscc_bedrock_agent} Resource.
func NewBedrockAgent_Override(b BedrockAgent, scope constructs.Construct, id *string, config *BedrockAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockAgent)SetAgentCollaboration(val *string) {
	if err := j.validateSetAgentCollaborationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentCollaboration",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetAgentName(val *string) {
	if err := j.validateSetAgentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentName",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetAgentResourceRoleArn(val *string) {
	if err := j.validateSetAgentResourceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentResourceRoleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetAutoPrepare(val interface{}) {
	if err := j.validateSetAutoPrepareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPrepare",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetCustomerEncryptionKeyArn(val *string) {
	if err := j.validateSetCustomerEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetFoundationModel(val *string) {
	if err := j.validateSetFoundationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationModel",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetOrchestrationType(val *string) {
	if err := j.validateSetOrchestrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orchestrationType",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetSkipResourceInUseCheckOnDelete(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheckOnDelete",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BedrockAgent)SetTestAliasTags(val *map[string]*string) {
	if err := j.validateSetTestAliasTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testAliasTags",
		val,
	)
}

// Generates CDKTN code for importing a BedrockAgent resource upon running "cdktn plan <stack-name>".
func BedrockAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
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
func BedrockAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockAgent) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockAgent) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (b *jsiiProxy_BedrockAgent) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockAgent) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockAgent) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockAgent) PutActionGroups(value interface{}) {
	if err := b.validatePutActionGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putActionGroups",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutAgentCollaborators(value interface{}) {
	if err := b.validatePutAgentCollaboratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentCollaborators",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutCustomOrchestration(value *BedrockAgentCustomOrchestration) {
	if err := b.validatePutCustomOrchestrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomOrchestration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutGuardrailConfiguration(value *BedrockAgentGuardrailConfiguration) {
	if err := b.validatePutGuardrailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGuardrailConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutKnowledgeBases(value interface{}) {
	if err := b.validatePutKnowledgeBasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putKnowledgeBases",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutMemoryConfiguration(value *BedrockAgentMemoryConfiguration) {
	if err := b.validatePutMemoryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMemoryConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) PutPromptOverrideConfiguration(value *BedrockAgentPromptOverrideConfiguration) {
	if err := b.validatePutPromptOverrideConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPromptOverrideConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := b.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (b *jsiiProxy_BedrockAgent) ResetActionGroups() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroups",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetAgentCollaboration() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentCollaboration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetAgentCollaborators() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentCollaborators",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetAgentResourceRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentResourceRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetAutoPrepare() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoPrepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetCustomerEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomerEncryptionKeyArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetCustomOrchestration() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomOrchestration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetFoundationModel() {
	_jsii_.InvokeVoid(
		b,
		"resetFoundationModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetGuardrailConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetGuardrailConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetInstruction() {
	_jsii_.InvokeVoid(
		b,
		"resetInstruction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetKnowledgeBases() {
	_jsii_.InvokeVoid(
		b,
		"resetKnowledgeBases",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetMemoryConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetMemoryConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetOrchestrationType() {
	_jsii_.InvokeVoid(
		b,
		"resetOrchestrationType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetPromptOverrideConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptOverrideConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetSkipResourceInUseCheckOnDelete() {
	_jsii_.InvokeVoid(
		b,
		"resetSkipResourceInUseCheckOnDelete",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) ResetTestAliasTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTestAliasTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

