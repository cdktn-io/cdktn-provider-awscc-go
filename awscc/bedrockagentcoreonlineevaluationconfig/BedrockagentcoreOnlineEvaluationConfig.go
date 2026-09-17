// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreonlineevaluationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_online_evaluation_config awscc_bedrockagentcore_online_evaluation_config}.
type BedrockagentcoreOnlineEvaluationConfig interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClusteringConfig() BedrockagentcoreOnlineEvaluationConfigClusteringConfigOutputReference
	ClusteringConfigInput() interface{}
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
	DataSourceConfig() BedrockagentcoreOnlineEvaluationConfigDataSourceConfigOutputReference
	DataSourceConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EvaluationExecutionRoleArn() *string
	SetEvaluationExecutionRoleArn(val *string)
	EvaluationExecutionRoleArnInput() *string
	Evaluators() BedrockagentcoreOnlineEvaluationConfigEvaluatorsList
	EvaluatorsInput() interface{}
	ExecutionStatus() *string
	SetExecutionStatus(val *string)
	ExecutionStatusInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Insights() BedrockagentcoreOnlineEvaluationConfigInsightsList
	InsightsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnlineEvaluationConfigArn() *string
	OnlineEvaluationConfigId() *string
	OnlineEvaluationConfigName() *string
	SetOnlineEvaluationConfigName(val *string)
	OnlineEvaluationConfigNameInput() *string
	OutputConfig() BedrockagentcoreOnlineEvaluationConfigOutputConfigOutputReference
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
	Rule() BedrockagentcoreOnlineEvaluationConfigRuleOutputReference
	RuleInput() interface{}
	Status() *string
	Tags() BedrockagentcoreOnlineEvaluationConfigTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutClusteringConfig(value *BedrockagentcoreOnlineEvaluationConfigClusteringConfig)
	PutDataSourceConfig(value *BedrockagentcoreOnlineEvaluationConfigDataSourceConfig)
	PutEvaluators(value interface{})
	PutInsights(value interface{})
	PutRule(value *BedrockagentcoreOnlineEvaluationConfigRule)
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
	ResetClusteringConfig()
	ResetDescription()
	ResetEvaluators()
	ResetExecutionStatus()
	ResetInsights()
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

// The jsii proxy struct for BedrockagentcoreOnlineEvaluationConfig
type jsiiProxy_BedrockagentcoreOnlineEvaluationConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ClusteringConfig() BedrockagentcoreOnlineEvaluationConfigClusteringConfigOutputReference {
	var returns BedrockagentcoreOnlineEvaluationConfigClusteringConfigOutputReference
	_jsii_.Get(
		j,
		"clusteringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ClusteringConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) DataSourceConfig() BedrockagentcoreOnlineEvaluationConfigDataSourceConfigOutputReference {
	var returns BedrockagentcoreOnlineEvaluationConfigDataSourceConfigOutputReference
	_jsii_.Get(
		j,
		"dataSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) DataSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) EvaluationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) EvaluationExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Evaluators() BedrockagentcoreOnlineEvaluationConfigEvaluatorsList {
	var returns BedrockagentcoreOnlineEvaluationConfigEvaluatorsList
	_jsii_.Get(
		j,
		"evaluators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) EvaluatorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ExecutionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ExecutionStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Insights() BedrockagentcoreOnlineEvaluationConfigInsightsList {
	var returns BedrockagentcoreOnlineEvaluationConfigInsightsList
	_jsii_.Get(
		j,
		"insights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) InsightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OnlineEvaluationConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OnlineEvaluationConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OnlineEvaluationConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OnlineEvaluationConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OutputConfig() BedrockagentcoreOnlineEvaluationConfigOutputConfigOutputReference {
	var returns BedrockagentcoreOnlineEvaluationConfigOutputConfigOutputReference
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Rule() BedrockagentcoreOnlineEvaluationConfigRuleOutputReference {
	var returns BedrockagentcoreOnlineEvaluationConfigRuleOutputReference
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) Tags() BedrockagentcoreOnlineEvaluationConfigTagsList {
	var returns BedrockagentcoreOnlineEvaluationConfigTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_online_evaluation_config awscc_bedrockagentcore_online_evaluation_config} Resource.
func NewBedrockagentcoreOnlineEvaluationConfig(scope constructs.Construct, id *string, config *BedrockagentcoreOnlineEvaluationConfigConfig) BedrockagentcoreOnlineEvaluationConfig {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreOnlineEvaluationConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreOnlineEvaluationConfig{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_online_evaluation_config awscc_bedrockagentcore_online_evaluation_config} Resource.
func NewBedrockagentcoreOnlineEvaluationConfig_Override(b BedrockagentcoreOnlineEvaluationConfig, scope constructs.Construct, id *string, config *BedrockagentcoreOnlineEvaluationConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetEvaluationExecutionRoleArn(val *string) {
	if err := j.validateSetEvaluationExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetExecutionStatus(val *string) {
	if err := j.validateSetExecutionStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionStatus",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetOnlineEvaluationConfigName(val *string) {
	if err := j.validateSetOnlineEvaluationConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlineEvaluationConfigName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a BedrockagentcoreOnlineEvaluationConfig resource upon running "cdktn plan <stack-name>".
func BedrockagentcoreOnlineEvaluationConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentcoreOnlineEvaluationConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
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
func BedrockagentcoreOnlineEvaluationConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreOnlineEvaluationConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreOnlineEvaluationConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreOnlineEvaluationConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreOnlineEvaluationConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreOnlineEvaluationConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentcoreOnlineEvaluationConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.bedrockagentcoreOnlineEvaluationConfig.BedrockagentcoreOnlineEvaluationConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutClusteringConfig(value *BedrockagentcoreOnlineEvaluationConfigClusteringConfig) {
	if err := b.validatePutClusteringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putClusteringConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutDataSourceConfig(value *BedrockagentcoreOnlineEvaluationConfigDataSourceConfig) {
	if err := b.validatePutDataSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDataSourceConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutEvaluators(value interface{}) {
	if err := b.validatePutEvaluatorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEvaluators",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutInsights(value interface{}) {
	if err := b.validatePutInsightsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInsights",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutRule(value *BedrockagentcoreOnlineEvaluationConfigRule) {
	if err := b.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRule",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) PutTags(value interface{}) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := b.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetClusteringConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetClusteringConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetEvaluators() {
	_jsii_.InvokeVoid(
		b,
		"resetEvaluators",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetExecutionStatus() {
	_jsii_.InvokeVoid(
		b,
		"resetExecutionStatus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetInsights() {
	_jsii_.InvokeVoid(
		b,
		"resetInsights",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOnlineEvaluationConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

