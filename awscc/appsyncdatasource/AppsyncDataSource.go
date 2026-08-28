// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appsyncdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source awscc_appsync_data_source}.
type AppsyncDataSource interface {
	cdktn.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
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
	DataSourceArn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamoDbConfig() AppsyncDataSourceDynamoDbConfigOutputReference
	DynamoDbConfigInput() interface{}
	ElasticsearchConfig() AppsyncDataSourceElasticsearchConfigOutputReference
	ElasticsearchConfigInput() interface{}
	EventBridgeConfig() AppsyncDataSourceEventBridgeConfigOutputReference
	EventBridgeConfigInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpConfig() AppsyncDataSourceHttpConfigOutputReference
	HttpConfigInput() interface{}
	Id() *string
	LambdaConfig() AppsyncDataSourceLambdaConfigOutputReference
	LambdaConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetricsConfig() *string
	SetMetricsConfig(val *string)
	MetricsConfigInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpenSearchServiceConfig() AppsyncDataSourceOpenSearchServiceConfigOutputReference
	OpenSearchServiceConfigInput() interface{}
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
	RelationalDatabaseConfig() AppsyncDataSourceRelationalDatabaseConfigOutputReference
	RelationalDatabaseConfigInput() interface{}
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDynamoDbConfig(value *AppsyncDataSourceDynamoDbConfig)
	PutElasticsearchConfig(value *AppsyncDataSourceElasticsearchConfig)
	PutEventBridgeConfig(value *AppsyncDataSourceEventBridgeConfig)
	PutHttpConfig(value *AppsyncDataSourceHttpConfig)
	PutLambdaConfig(value *AppsyncDataSourceLambdaConfig)
	PutOpenSearchServiceConfig(value *AppsyncDataSourceOpenSearchServiceConfig)
	PutRelationalDatabaseConfig(value *AppsyncDataSourceRelationalDatabaseConfig)
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
	ResetDescription()
	ResetDynamoDbConfig()
	ResetElasticsearchConfig()
	ResetEventBridgeConfig()
	ResetHttpConfig()
	ResetLambdaConfig()
	ResetMetricsConfig()
	ResetOpenSearchServiceConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRelationalDatabaseConfig()
	ResetServiceRoleArn()
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

// The jsii proxy struct for AppsyncDataSource
type jsiiProxy_AppsyncDataSource struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppsyncDataSource) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) DataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) DynamoDbConfig() AppsyncDataSourceDynamoDbConfigOutputReference {
	var returns AppsyncDataSourceDynamoDbConfigOutputReference
	_jsii_.Get(
		j,
		"dynamoDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) DynamoDbConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ElasticsearchConfig() AppsyncDataSourceElasticsearchConfigOutputReference {
	var returns AppsyncDataSourceElasticsearchConfigOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ElasticsearchConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) EventBridgeConfig() AppsyncDataSourceEventBridgeConfigOutputReference {
	var returns AppsyncDataSourceEventBridgeConfigOutputReference
	_jsii_.Get(
		j,
		"eventBridgeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) EventBridgeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventBridgeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) HttpConfig() AppsyncDataSourceHttpConfigOutputReference {
	var returns AppsyncDataSourceHttpConfigOutputReference
	_jsii_.Get(
		j,
		"httpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) HttpConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) LambdaConfig() AppsyncDataSourceLambdaConfigOutputReference {
	var returns AppsyncDataSourceLambdaConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) LambdaConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) MetricsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) MetricsConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) OpenSearchServiceConfig() AppsyncDataSourceOpenSearchServiceConfigOutputReference {
	var returns AppsyncDataSourceOpenSearchServiceConfigOutputReference
	_jsii_.Get(
		j,
		"openSearchServiceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) OpenSearchServiceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openSearchServiceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) RelationalDatabaseConfig() AppsyncDataSourceRelationalDatabaseConfigOutputReference {
	var returns AppsyncDataSourceRelationalDatabaseConfigOutputReference
	_jsii_.Get(
		j,
		"relationalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) RelationalDatabaseConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDataSource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source awscc_appsync_data_source} Resource.
func NewAppsyncDataSource(scope constructs.Construct, id *string, config *AppsyncDataSourceConfig) AppsyncDataSource {
	_init_.Initialize()

	if err := validateNewAppsyncDataSourceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDataSource{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source awscc_appsync_data_source} Resource.
func NewAppsyncDataSource_Override(a AppsyncDataSource, scope constructs.Construct, id *string, config *AppsyncDataSourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetMetricsConfig(val *string) {
	if err := j.validateSetMetricsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsConfig",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDataSource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a AppsyncDataSource resource upon running "cdktn plan <stack-name>".
func AppsyncDataSource_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppsyncDataSource_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
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
func AppsyncDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncDataSource_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDataSource_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsyncDataSource_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncDataSource_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncDataSource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.appsyncDataSource.AppsyncDataSource",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsyncDataSource) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsyncDataSource) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncDataSource) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncDataSource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncDataSource) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncDataSource) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncDataSource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncDataSource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncDataSource) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncDataSource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncDataSource) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsyncDataSource) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppsyncDataSource) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AppsyncDataSource) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncDataSource) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsyncDataSource) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsyncDataSource) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutDynamoDbConfig(value *AppsyncDataSourceDynamoDbConfig) {
	if err := a.validatePutDynamoDbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamoDbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutElasticsearchConfig(value *AppsyncDataSourceElasticsearchConfig) {
	if err := a.validatePutElasticsearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutEventBridgeConfig(value *AppsyncDataSourceEventBridgeConfig) {
	if err := a.validatePutEventBridgeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridgeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutHttpConfig(value *AppsyncDataSourceHttpConfig) {
	if err := a.validatePutHttpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutLambdaConfig(value *AppsyncDataSourceLambdaConfig) {
	if err := a.validatePutLambdaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutOpenSearchServiceConfig(value *AppsyncDataSourceOpenSearchServiceConfig) {
	if err := a.validatePutOpenSearchServiceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenSearchServiceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) PutRelationalDatabaseConfig(value *AppsyncDataSourceRelationalDatabaseConfig) {
	if err := a.validatePutRelationalDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelationalDatabaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDataSource) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetDynamoDbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamoDbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetElasticsearchConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetEventBridgeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridgeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetHttpConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetMetricsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetOpenSearchServiceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenSearchServiceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetRelationalDatabaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRelationalDatabaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDataSource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDataSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

