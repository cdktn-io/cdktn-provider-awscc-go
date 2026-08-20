// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncfunctionconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appsyncfunctionconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration awscc_appsync_function_configuration}.
type AppsyncFunctionConfiguration interface {
	cdktn.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Code() *string
	SetCode(val *string)
	CodeInput() *string
	CodeS3Location() *string
	SetCodeS3Location(val *string)
	CodeS3LocationInput() *string
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
	DataSourceName() *string
	SetDataSourceName(val *string)
	DataSourceNameInput() *string
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
	FunctionArn() *string
	FunctionId() *string
	FunctionVersion() *string
	SetFunctionVersion(val *string)
	FunctionVersionInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	MaxBatchSizeInput() *float64
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
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	RequestMappingTemplateInput() *string
	RequestMappingTemplateS3Location() *string
	SetRequestMappingTemplateS3Location(val *string)
	RequestMappingTemplateS3LocationInput() *string
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	ResponseMappingTemplateInput() *string
	ResponseMappingTemplateS3Location() *string
	SetResponseMappingTemplateS3Location(val *string)
	ResponseMappingTemplateS3LocationInput() *string
	Runtime() AppsyncFunctionConfigurationRuntimeOutputReference
	RuntimeInput() interface{}
	SyncConfig() AppsyncFunctionConfigurationSyncConfigOutputReference
	SyncConfigInput() interface{}
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
	PutRuntime(value *AppsyncFunctionConfigurationRuntime)
	PutSyncConfig(value *AppsyncFunctionConfigurationSyncConfig)
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
	ResetCode()
	ResetCodeS3Location()
	ResetDescription()
	ResetFunctionVersion()
	ResetMaxBatchSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestMappingTemplate()
	ResetRequestMappingTemplateS3Location()
	ResetResponseMappingTemplate()
	ResetResponseMappingTemplateS3Location()
	ResetRuntime()
	ResetSyncConfig()
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

// The jsii proxy struct for AppsyncFunctionConfiguration
type jsiiProxy_AppsyncFunctionConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) CodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) CodeS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) CodeS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) DataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) DataSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) FunctionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) MaxBatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RequestMappingTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RequestMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RequestMappingTemplateS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ResponseMappingTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ResponseMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) ResponseMappingTemplateS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) Runtime() AppsyncFunctionConfigurationRuntimeOutputReference {
	var returns AppsyncFunctionConfigurationRuntimeOutputReference
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) RuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) SyncConfig() AppsyncFunctionConfigurationSyncConfigOutputReference {
	var returns AppsyncFunctionConfigurationSyncConfigOutputReference
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) SyncConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration awscc_appsync_function_configuration} Resource.
func NewAppsyncFunctionConfiguration(scope constructs.Construct, id *string, config *AppsyncFunctionConfigurationConfig) AppsyncFunctionConfiguration {
	_init_.Initialize()

	if err := validateNewAppsyncFunctionConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncFunctionConfiguration{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_function_configuration awscc_appsync_function_configuration} Resource.
func NewAppsyncFunctionConfiguration_Override(a AppsyncFunctionConfiguration, scope constructs.Construct, id *string, config *AppsyncFunctionConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetCode(val *string) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetCodeS3Location(val *string) {
	if err := j.validateSetCodeS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeS3Location",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetDataSourceName(val *string) {
	if err := j.validateSetDataSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceName",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetFunctionVersion(val *string) {
	if err := j.validateSetFunctionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetMaxBatchSize(val *float64) {
	if err := j.validateSetMaxBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetRequestMappingTemplate(val *string) {
	if err := j.validateSetRequestMappingTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetRequestMappingTemplateS3Location(val *string) {
	if err := j.validateSetRequestMappingTemplateS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetResponseMappingTemplate(val *string) {
	if err := j.validateSetResponseMappingTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionConfiguration)SetResponseMappingTemplateS3Location(val *string) {
	if err := j.validateSetResponseMappingTemplateS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseMappingTemplateS3Location",
		val,
	)
}

// Generates CDKTN code for importing a AppsyncFunctionConfiguration resource upon running "cdktn plan <stack-name>".
func AppsyncFunctionConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncFunctionConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
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
func AppsyncFunctionConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncFunctionConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncFunctionConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncFunctionConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncFunctionConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncFunctionConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncFunctionConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.appsyncFunctionConfiguration.AppsyncFunctionConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AppsyncFunctionConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) PutRuntime(value *AppsyncFunctionConfigurationRuntime) {
	if err := a.validatePutRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) PutSyncConfig(value *AppsyncFunctionConfigurationSyncConfig) {
	if err := a.validatePutSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetCode() {
	_jsii_.InvokeVoid(
		a,
		"resetCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetCodeS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetFunctionVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetMaxBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetRequestMappingTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestMappingTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetRequestMappingTemplateS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestMappingTemplateS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetResponseMappingTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseMappingTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetResponseMappingTemplateS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseMappingTemplateS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ResetSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

