// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftredshiftidcapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/redshiftredshiftidcapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application awscc_redshift_redshift_idc_application}.
type RedshiftRedshiftIdcApplication interface {
	cdktn.TerraformResource
	ApplicationType() *string
	SetApplicationType(val *string)
	ApplicationTypeInput() *string
	AuthorizedTokenIssuerList() RedshiftRedshiftIdcApplicationAuthorizedTokenIssuerListStructList
	AuthorizedTokenIssuerListInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	IdcDisplayName() *string
	SetIdcDisplayName(val *string)
	IdcDisplayNameInput() *string
	IdcInstanceArn() *string
	SetIdcInstanceArn(val *string)
	IdcInstanceArnInput() *string
	IdcManagedApplicationArn() *string
	IdcOnboardStatus() *string
	IdentityNamespace() *string
	SetIdentityNamespace(val *string)
	IdentityNamespaceInput() *string
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
	RedshiftIdcApplicationArn() *string
	RedshiftIdcApplicationName() *string
	SetRedshiftIdcApplicationName(val *string)
	RedshiftIdcApplicationNameInput() *string
	ServiceIntegrations() RedshiftRedshiftIdcApplicationServiceIntegrationsList
	ServiceIntegrationsInput() interface{}
	SsoTagKeys() *[]*string
	SetSsoTagKeys(val *[]*string)
	SsoTagKeysInput() *[]*string
	Tags() RedshiftRedshiftIdcApplicationTagsList
	TagsInput() interface{}
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
	PutAuthorizedTokenIssuerList(value interface{})
	PutServiceIntegrations(value interface{})
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
	ResetApplicationType()
	ResetAuthorizedTokenIssuerList()
	ResetIdentityNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceIntegrations()
	ResetSsoTagKeys()
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

// The jsii proxy struct for RedshiftRedshiftIdcApplication
type jsiiProxy_RedshiftRedshiftIdcApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ApplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ApplicationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) AuthorizedTokenIssuerList() RedshiftRedshiftIdcApplicationAuthorizedTokenIssuerListStructList {
	var returns RedshiftRedshiftIdcApplicationAuthorizedTokenIssuerListStructList
	_jsii_.Get(
		j,
		"authorizedTokenIssuerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) AuthorizedTokenIssuerListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizedTokenIssuerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcManagedApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcManagedApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdcOnboardStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcOnboardStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdentityNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) IdentityNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) RedshiftIdcApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) RedshiftIdcApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) RedshiftIdcApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ServiceIntegrations() RedshiftRedshiftIdcApplicationServiceIntegrationsList {
	var returns RedshiftRedshiftIdcApplicationServiceIntegrationsList
	_jsii_.Get(
		j,
		"serviceIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) ServiceIntegrationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceIntegrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) SsoTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ssoTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) SsoTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ssoTagKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) Tags() RedshiftRedshiftIdcApplicationTagsList {
	var returns RedshiftRedshiftIdcApplicationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application awscc_redshift_redshift_idc_application} Resource.
func NewRedshiftRedshiftIdcApplication(scope constructs.Construct, id *string, config *RedshiftRedshiftIdcApplicationConfig) RedshiftRedshiftIdcApplication {
	_init_.Initialize()

	if err := validateNewRedshiftRedshiftIdcApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftRedshiftIdcApplication{}

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/redshift_redshift_idc_application awscc_redshift_redshift_idc_application} Resource.
func NewRedshiftRedshiftIdcApplication_Override(r RedshiftRedshiftIdcApplication, scope constructs.Construct, id *string, config *RedshiftRedshiftIdcApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetApplicationType(val *string) {
	if err := j.validateSetApplicationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationType",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetIdcDisplayName(val *string) {
	if err := j.validateSetIdcDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idcDisplayName",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetIdcInstanceArn(val *string) {
	if err := j.validateSetIdcInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idcInstanceArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetIdentityNamespace(val *string) {
	if err := j.validateSetIdentityNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityNamespace",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetRedshiftIdcApplicationName(val *string) {
	if err := j.validateSetRedshiftIdcApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftIdcApplicationName",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplication)SetSsoTagKeys(val *[]*string) {
	if err := j.validateSetSsoTagKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoTagKeys",
		val,
	)
}

// Generates CDKTN code for importing a RedshiftRedshiftIdcApplication resource upon running "cdktn plan <stack-name>".
func RedshiftRedshiftIdcApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRedshiftRedshiftIdcApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
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
func RedshiftRedshiftIdcApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftRedshiftIdcApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftRedshiftIdcApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftRedshiftIdcApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftRedshiftIdcApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftRedshiftIdcApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedshiftRedshiftIdcApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) PutAuthorizedTokenIssuerList(value interface{}) {
	if err := r.validatePutAuthorizedTokenIssuerListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAuthorizedTokenIssuerList",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) PutServiceIntegrations(value interface{}) {
	if err := r.validatePutServiceIntegrationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServiceIntegrations",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetApplicationType() {
	_jsii_.InvokeVoid(
		r,
		"resetApplicationType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetAuthorizedTokenIssuerList() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthorizedTokenIssuerList",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetIdentityNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentityNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetServiceIntegrations() {
	_jsii_.InvokeVoid(
		r,
		"resetServiceIntegrations",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetSsoTagKeys() {
	_jsii_.InvokeVoid(
		r,
		"resetSsoTagKeys",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

