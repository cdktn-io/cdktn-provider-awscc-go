// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccappsyncgraphqlapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccappsyncgraphqlapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/data-sources/appsync_graph_ql_api awscc_appsync_graph_ql_api}.
type DataAwsccAppsyncGraphQlApi interface {
	cdktn.TerraformDataSource
	AdditionalAuthenticationProviders() DataAwsccAppsyncGraphQlApiAdditionalAuthenticationProvidersList
	ApiId() *string
	ApiType() *string
	Arn() *string
	AuthenticationType() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	EnhancedMetricsConfig() DataAwsccAppsyncGraphQlApiEnhancedMetricsConfigOutputReference
	EnvironmentVariables() cdktn.StringMap
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GraphQlDns() *string
	GraphQlEndpointArn() *string
	GraphQlUrl() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntrospectionConfig() *string
	LambdaAuthorizerConfig() DataAwsccAppsyncGraphQlApiLambdaAuthorizerConfigOutputReference
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogConfig() DataAwsccAppsyncGraphQlApiLogConfigOutputReference
	MergedApiExecutionRoleArn() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenIdConnectConfig() DataAwsccAppsyncGraphQlApiOpenIdConnectConfigOutputReference
	OwnerContact() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	QueryDepthLimit() *float64
	// Experimental.
	RawOverrides() interface{}
	RealtimeDns() *string
	RealtimeUrl() *string
	ResolverCountLimit() *float64
	Tags() DataAwsccAppsyncGraphQlApiTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserPoolConfig() DataAwsccAppsyncGraphQlApiUserPoolConfigOutputReference
	Visibility() *string
	XrayEnabled() cdktn.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsccAppsyncGraphQlApi
type jsiiProxy_DataAwsccAppsyncGraphQlApi struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) AdditionalAuthenticationProviders() DataAwsccAppsyncGraphQlApiAdditionalAuthenticationProvidersList {
	var returns DataAwsccAppsyncGraphQlApiAdditionalAuthenticationProvidersList
	_jsii_.Get(
		j,
		"additionalAuthenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) EnhancedMetricsConfig() DataAwsccAppsyncGraphQlApiEnhancedMetricsConfigOutputReference {
	var returns DataAwsccAppsyncGraphQlApiEnhancedMetricsConfigOutputReference
	_jsii_.Get(
		j,
		"enhancedMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) EnvironmentVariables() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) GraphQlDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphQlDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) GraphQlEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphQlEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) GraphQlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphQlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) IntrospectionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) LambdaAuthorizerConfig() DataAwsccAppsyncGraphQlApiLambdaAuthorizerConfigOutputReference {
	var returns DataAwsccAppsyncGraphQlApiLambdaAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) LogConfig() DataAwsccAppsyncGraphQlApiLogConfigOutputReference {
	var returns DataAwsccAppsyncGraphQlApiLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) MergedApiExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) OpenIdConnectConfig() DataAwsccAppsyncGraphQlApiOpenIdConnectConfigOutputReference {
	var returns DataAwsccAppsyncGraphQlApiOpenIdConnectConfigOutputReference
	_jsii_.Get(
		j,
		"openIdConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) OwnerContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) QueryDepthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) RealtimeDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) RealtimeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) ResolverCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Tags() DataAwsccAppsyncGraphQlApiTagsList {
	var returns DataAwsccAppsyncGraphQlApiTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) UserPoolConfig() DataAwsccAppsyncGraphQlApiUserPoolConfigOutputReference {
	var returns DataAwsccAppsyncGraphQlApiUserPoolConfigOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi) XrayEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/data-sources/appsync_graph_ql_api awscc_appsync_graph_ql_api} Data Source.
func NewDataAwsccAppsyncGraphQlApi(scope constructs.Construct, id *string, config *DataAwsccAppsyncGraphQlApiConfig) DataAwsccAppsyncGraphQlApi {
	_init_.Initialize()

	if err := validateNewDataAwsccAppsyncGraphQlApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAppsyncGraphQlApi{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/data-sources/appsync_graph_ql_api awscc_appsync_graph_ql_api} Data Source.
func NewDataAwsccAppsyncGraphQlApi_Override(d DataAwsccAppsyncGraphQlApi, scope constructs.Construct, id *string, config *DataAwsccAppsyncGraphQlApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAppsyncGraphQlApi)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccAppsyncGraphQlApi resource upon running "cdktn plan <stack-name>".
func DataAwsccAppsyncGraphQlApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccAppsyncGraphQlApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
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
func DataAwsccAppsyncGraphQlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAppsyncGraphQlApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccAppsyncGraphQlApi_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAppsyncGraphQlApi_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccAppsyncGraphQlApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccAppsyncGraphQlApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccAppsyncGraphQlApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccAppsyncGraphQlApi.DataAwsccAppsyncGraphQlApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAppsyncGraphQlApi) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

