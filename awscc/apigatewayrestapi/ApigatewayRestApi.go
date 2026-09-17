// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayrestapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apigatewayrestapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api awscc_apigateway_rest_api}.
type ApigatewayRestApi interface {
	cdktn.TerraformResource
	ApiKeySourceType() *string
	SetApiKeySourceType(val *string)
	ApiKeySourceTypeInput() *string
	BinaryMediaTypes() *[]*string
	SetBinaryMediaTypes(val *[]*string)
	BinaryMediaTypesInput() *[]*string
	Body() *string
	SetBody(val *string)
	BodyInput() *string
	BodyS3Location() ApigatewayRestApiBodyS3LocationOutputReference
	BodyS3LocationInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CloneFrom() *string
	SetCloneFrom(val *string)
	CloneFromInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	DisableExecuteApiEndpointInput() interface{}
	EndpointAccessMode() *string
	SetEndpointAccessMode(val *string)
	EndpointAccessModeInput() *string
	EndpointConfiguration() ApigatewayRestApiEndpointConfigurationOutputReference
	EndpointConfigurationInput() interface{}
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
	FailOnWarningsInput() interface{}
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
	MinimumCompressionSize() *float64
	SetMinimumCompressionSize(val *float64)
	MinimumCompressionSizeInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
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
	RestApiId() *string
	RootResourceId() *string
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	Tags() ApigatewayRestApiTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutBodyS3Location(value *ApigatewayRestApiBodyS3Location)
	PutEndpointConfiguration(value *ApigatewayRestApiEndpointConfiguration)
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
	ResetApiKeySourceType()
	ResetBinaryMediaTypes()
	ResetBody()
	ResetBodyS3Location()
	ResetCloneFrom()
	ResetDescription()
	ResetDisableExecuteApiEndpoint()
	ResetEndpointAccessMode()
	ResetEndpointConfiguration()
	ResetFailOnWarnings()
	ResetMinimumCompressionSize()
	ResetMode()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPolicy()
	ResetSecurityPolicy()
	ResetTags()
	ResetVersion()
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

// The jsii proxy struct for ApigatewayRestApi
type jsiiProxy_ApigatewayRestApi struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ApigatewayRestApi) ApiKeySourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) ApiKeySourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) BinaryMediaTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) BodyS3Location() ApigatewayRestApiBodyS3LocationOutputReference {
	var returns ApigatewayRestApiBodyS3LocationOutputReference
	_jsii_.Get(
		j,
		"bodyS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) BodyS3LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) CloneFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) CloneFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) DisableExecuteApiEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) EndpointAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) EndpointAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) EndpointConfiguration() ApigatewayRestApiEndpointConfigurationOutputReference {
	var returns ApigatewayRestApiEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) EndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) FailOnWarningsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) MinimumCompressionSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) RootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Tags() ApigatewayRestApiTagsList {
	var returns ApigatewayRestApiTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayRestApi) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api awscc_apigateway_rest_api} Resource.
func NewApigatewayRestApi(scope constructs.Construct, id *string, config *ApigatewayRestApiConfig) ApigatewayRestApi {
	_init_.Initialize()

	if err := validateNewApigatewayRestApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayRestApi{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_rest_api awscc_apigateway_rest_api} Resource.
func NewApigatewayRestApi_Override(a ApigatewayRestApi, scope constructs.Construct, id *string, config *ApigatewayRestApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetApiKeySourceType(val *string) {
	if err := j.validateSetApiKeySourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeySourceType",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetBinaryMediaTypes(val *[]*string) {
	if err := j.validateSetBinaryMediaTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetCloneFrom(val *string) {
	if err := j.validateSetCloneFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneFrom",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetDisableExecuteApiEndpoint(val interface{}) {
	if err := j.validateSetDisableExecuteApiEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetEndpointAccessMode(val *string) {
	if err := j.validateSetEndpointAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointAccessMode",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetFailOnWarnings(val interface{}) {
	if err := j.validateSetFailOnWarningsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnWarnings",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetMinimumCompressionSize(val *float64) {
	if err := j.validateSetMinimumCompressionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_ApigatewayRestApi)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a ApigatewayRestApi resource upon running "cdktn plan <stack-name>".
func ApigatewayRestApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApigatewayRestApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
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
func ApigatewayRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayRestApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayRestApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayRestApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigatewayRestApi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigatewayRestApi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigatewayRestApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.apigatewayRestApi.ApigatewayRestApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigatewayRestApi) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigatewayRestApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigatewayRestApi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigatewayRestApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigatewayRestApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigatewayRestApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigatewayRestApi) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigatewayRestApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigatewayRestApi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigatewayRestApi) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_ApigatewayRestApi) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) PutBodyS3Location(value *ApigatewayRestApiBodyS3Location) {
	if err := a.validatePutBodyS3LocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBodyS3Location",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) PutEndpointConfiguration(value *ApigatewayRestApiEndpointConfiguration) {
	if err := a.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetApiKeySourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeySourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetBinaryMediaTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetBinaryMediaTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetBody() {
	_jsii_.InvokeVoid(
		a,
		"resetBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetBodyS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetBodyS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetCloneFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetCloneFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetDisableExecuteApiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableExecuteApiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetEndpointAccessMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointAccessMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetFailOnWarnings() {
	_jsii_.InvokeVoid(
		a,
		"resetFailOnWarnings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetMinimumCompressionSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumCompressionSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayRestApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayRestApi) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

