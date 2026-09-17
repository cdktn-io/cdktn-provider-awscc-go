// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsdbproxy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_proxy awscc_rds_db_proxy}.
type RdsDbProxy interface {
	cdktn.TerraformResource
	Auth() RdsDbProxyAuthList
	AuthInput() interface{}
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
	DbProxyArn() *string
	DbProxyName() *string
	SetDbProxyName(val *string)
	DbProxyNameInput() *string
	DebugLogging() interface{}
	SetDebugLogging(val interface{})
	DebugLoggingInput() interface{}
	DefaultAuthScheme() *string
	SetDefaultAuthScheme(val *string)
	DefaultAuthSchemeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Endpoint() *string
	EndpointNetworkType() *string
	SetEndpointNetworkType(val *string)
	EndpointNetworkTypeInput() *string
	EngineFamily() *string
	SetEngineFamily(val *string)
	EngineFamilyInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdleClientTimeout() *float64
	SetIdleClientTimeout(val *float64)
	IdleClientTimeoutInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	RequireTls() interface{}
	SetRequireTls(val interface{})
	RequireTlsInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Tags() RdsDbProxyTagsList
	TagsInput() interface{}
	TargetConnectionNetworkType() *string
	SetTargetConnectionNetworkType(val *string)
	TargetConnectionNetworkTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VpcId() *string
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
	VpcSubnetIds() *[]*string
	SetVpcSubnetIds(val *[]*string)
	VpcSubnetIdsInput() *[]*string
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
	PutAuth(value interface{})
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
	ResetAuth()
	ResetDebugLogging()
	ResetDefaultAuthScheme()
	ResetEndpointNetworkType()
	ResetIdleClientTimeout()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequireTls()
	ResetTags()
	ResetTargetConnectionNetworkType()
	ResetVpcSecurityGroupIds()
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

// The jsii proxy struct for RdsDbProxy
type jsiiProxy_RdsDbProxy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RdsDbProxy) Auth() RdsDbProxyAuthList {
	var returns RdsDbProxyAuthList
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) AuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DbProxyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DebugLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DebugLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DefaultAuthScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DefaultAuthSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAuthSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) EndpointNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) EndpointNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) EngineFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) EngineFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) IdleClientTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleClientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) IdleClientTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleClientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) RequireTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) RequireTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) Tags() RdsDbProxyTagsList {
	var returns RdsDbProxyTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TargetConnectionNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetConnectionNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TargetConnectionNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetConnectionNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) VpcSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxy) VpcSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_proxy awscc_rds_db_proxy} Resource.
func NewRdsDbProxy(scope constructs.Construct, id *string, config *RdsDbProxyConfig) RdsDbProxy {
	_init_.Initialize()

	if err := validateNewRdsDbProxyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbProxy{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_proxy awscc_rds_db_proxy} Resource.
func NewRdsDbProxy_Override(r RdsDbProxy, scope constructs.Construct, id *string, config *RdsDbProxyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetDbProxyName(val *string) {
	if err := j.validateSetDbProxyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbProxyName",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetDebugLogging(val interface{}) {
	if err := j.validateSetDebugLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugLogging",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetDefaultAuthScheme(val *string) {
	if err := j.validateSetDefaultAuthSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAuthScheme",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetEndpointNetworkType(val *string) {
	if err := j.validateSetEndpointNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointNetworkType",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetEngineFamily(val *string) {
	if err := j.validateSetEngineFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineFamily",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetIdleClientTimeout(val *float64) {
	if err := j.validateSetIdleClientTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleClientTimeout",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetRequireTls(val interface{}) {
	if err := j.validateSetRequireTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTls",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetTargetConnectionNetworkType(val *string) {
	if err := j.validateSetTargetConnectionNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetConnectionNetworkType",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxy)SetVpcSubnetIds(val *[]*string) {
	if err := j.validateSetVpcSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSubnetIds",
		val,
	)
}

// Generates CDKTN code for importing a RdsDbProxy resource upon running "cdktn plan <stack-name>".
func RdsDbProxy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRdsDbProxy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
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
func RdsDbProxy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbProxy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbProxy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbProxy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbProxy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbProxy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDbProxy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.rdsDbProxy.RdsDbProxy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDbProxy) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsDbProxy) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsDbProxy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsDbProxy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbProxy) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsDbProxy) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbProxy) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsDbProxy) PutAuth(value interface{}) {
	if err := r.validatePutAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAuth",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbProxy) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbProxy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetAuth() {
	_jsii_.InvokeVoid(
		r,
		"resetAuth",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetDebugLogging() {
	_jsii_.InvokeVoid(
		r,
		"resetDebugLogging",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetDefaultAuthScheme() {
	_jsii_.InvokeVoid(
		r,
		"resetDefaultAuthScheme",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetEndpointNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpointNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetIdleClientTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetIdleClientTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetRequireTls() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireTls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetTargetConnectionNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetConnectionNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxy) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

