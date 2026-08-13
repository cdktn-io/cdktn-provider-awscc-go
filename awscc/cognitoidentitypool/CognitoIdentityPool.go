// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitoidentitypool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool}.
type CognitoIdentityPool interface {
	cdktn.TerraformResource
	AllowClassicFlow() interface{}
	SetAllowClassicFlow(val interface{})
	AllowClassicFlowInput() interface{}
	AllowUnauthenticatedIdentities() interface{}
	SetAllowUnauthenticatedIdentities(val interface{})
	AllowUnauthenticatedIdentitiesInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CognitoEvents() *string
	SetCognitoEvents(val *string)
	CognitoEventsInput() *string
	CognitoIdentityProviders() CognitoIdentityPoolCognitoIdentityProvidersList
	CognitoIdentityProvidersInput() interface{}
	CognitoStreams() CognitoIdentityPoolCognitoStreamsOutputReference
	CognitoStreamsInput() interface{}
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
	DeveloperProviderName() *string
	SetDeveloperProviderName(val *string)
	DeveloperProviderNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdentityPoolId() *string
	IdentityPoolName() *string
	SetIdentityPoolName(val *string)
	IdentityPoolNameInput() *string
	IdentityPoolTags() CognitoIdentityPoolIdentityPoolTagsList
	IdentityPoolTagsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenIdConnectProviderArNs() *[]*string
	SetOpenIdConnectProviderArNs(val *[]*string)
	OpenIdConnectProviderArNsInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushSync() CognitoIdentityPoolPushSyncOutputReference
	PushSyncInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SamlProviderArNs() *[]*string
	SetSamlProviderArNs(val *[]*string)
	SamlProviderArNsInput() *[]*string
	SupportedLoginProviders() *string
	SetSupportedLoginProviders(val *string)
	SupportedLoginProvidersInput() *string
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
	PutCognitoIdentityProviders(value interface{})
	PutCognitoStreams(value *CognitoIdentityPoolCognitoStreams)
	PutIdentityPoolTags(value interface{})
	PutPushSync(value *CognitoIdentityPoolPushSync)
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
	ResetAllowClassicFlow()
	ResetCognitoEvents()
	ResetCognitoIdentityProviders()
	ResetCognitoStreams()
	ResetDeveloperProviderName()
	ResetIdentityPoolName()
	ResetIdentityPoolTags()
	ResetOpenIdConnectProviderArNs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushSync()
	ResetSamlProviderArNs()
	ResetSupportedLoginProviders()
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

// The jsii proxy struct for CognitoIdentityPool
type jsiiProxy_CognitoIdentityPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowClassicFlowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClassicFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) AllowUnauthenticatedIdentitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUnauthenticatedIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitoEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProviders() CognitoIdentityPoolCognitoIdentityProvidersList {
	var returns CognitoIdentityPoolCognitoIdentityProvidersList
	_jsii_.Get(
		j,
		"cognitoIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoIdentityProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoIdentityProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoStreams() CognitoIdentityPoolCognitoStreamsOutputReference {
	var returns CognitoIdentityPoolCognitoStreamsOutputReference
	_jsii_.Get(
		j,
		"cognitoStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) CognitoStreamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cognitoStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) DeveloperProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolTags() CognitoIdentityPoolIdentityPoolTagsList {
	var returns CognitoIdentityPoolIdentityPoolTagsList
	_jsii_.Get(
		j,
		"identityPoolTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) IdentityPoolTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityPoolTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenIdConnectProviderArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openIdConnectProviderArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) OpenIdConnectProviderArNsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"openIdConnectProviderArNsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) PushSync() CognitoIdentityPoolPushSyncOutputReference {
	var returns CognitoIdentityPoolPushSyncOutputReference
	_jsii_.Get(
		j,
		"pushSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) PushSyncInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArNs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArNs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SamlProviderArNsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"samlProviderArNsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProviders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedLoginProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) SupportedLoginProvidersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedLoginProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoIdentityPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool} Resource.
func NewCognitoIdentityPool(scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) CognitoIdentityPool {
	_init_.Initialize()

	if err := validateNewCognitoIdentityPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoIdentityPool{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cognito_identity_pool awscc_cognito_identity_pool} Resource.
func NewCognitoIdentityPool_Override(c CognitoIdentityPool, scope constructs.Construct, id *string, config *CognitoIdentityPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetAllowClassicFlow(val interface{}) {
	if err := j.validateSetAllowClassicFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicFlow",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetAllowUnauthenticatedIdentities(val interface{}) {
	if err := j.validateSetAllowUnauthenticatedIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUnauthenticatedIdentities",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetCognitoEvents(val *string) {
	if err := j.validateSetCognitoEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cognitoEvents",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetDeveloperProviderName(val *string) {
	if err := j.validateSetDeveloperProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerProviderName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetIdentityPoolName(val *string) {
	if err := j.validateSetIdentityPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolName",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetOpenIdConnectProviderArNs(val *[]*string) {
	if err := j.validateSetOpenIdConnectProviderArNsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openIdConnectProviderArNs",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetSamlProviderArNs(val *[]*string) {
	if err := j.validateSetSamlProviderArNsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlProviderArNs",
		val,
	)
}

func (j *jsiiProxy_CognitoIdentityPool)SetSupportedLoginProviders(val *string) {
	if err := j.validateSetSupportedLoginProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedLoginProviders",
		val,
	)
}

// Generates CDKTN code for importing a CognitoIdentityPool resource upon running "cdktn plan <stack-name>".
func CognitoIdentityPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
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
func CognitoIdentityPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoIdentityPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoIdentityPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoIdentityPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoIdentityPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.cognitoIdentityPool.CognitoIdentityPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutCognitoIdentityProviders(value interface{}) {
	if err := c.validatePutCognitoIdentityProvidersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCognitoIdentityProviders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutCognitoStreams(value *CognitoIdentityPoolCognitoStreams) {
	if err := c.validatePutCognitoStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCognitoStreams",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutIdentityPoolTags(value interface{}) {
	if err := c.validatePutIdentityPoolTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentityPoolTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) PutPushSync(value *CognitoIdentityPoolPushSync) {
	if err := c.validatePutPushSyncParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPushSync",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetAllowClassicFlow() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowClassicFlow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoEvents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoIdentityProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoIdentityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetCognitoStreams() {
	_jsii_.InvokeVoid(
		c,
		"resetCognitoStreams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetDeveloperProviderName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeveloperProviderName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetIdentityPoolName() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityPoolName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetIdentityPoolTags() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityPoolTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetOpenIdConnectProviderArNs() {
	_jsii_.InvokeVoid(
		c,
		"resetOpenIdConnectProviderArNs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetPushSync() {
	_jsii_.InvokeVoid(
		c,
		"resetPushSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSamlProviderArNs() {
	_jsii_.InvokeVoid(
		c,
		"resetSamlProviderArNs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) ResetSupportedLoginProviders() {
	_jsii_.InvokeVoid(
		c,
		"resetSupportedLoginProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoIdentityPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoIdentityPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

