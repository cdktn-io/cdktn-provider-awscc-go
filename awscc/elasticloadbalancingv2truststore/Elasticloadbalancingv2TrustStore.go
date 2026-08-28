// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2truststore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2truststore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store}.
type Elasticloadbalancingv2TrustStore interface {
	cdktn.TerraformResource
	CaCertificatesBundleS3Bucket() *string
	SetCaCertificatesBundleS3Bucket(val *string)
	CaCertificatesBundleS3BucketInput() *string
	CaCertificatesBundleS3Key() *string
	SetCaCertificatesBundleS3Key(val *string)
	CaCertificatesBundleS3KeyInput() *string
	CaCertificatesBundleS3ObjectVersion() *string
	SetCaCertificatesBundleS3ObjectVersion(val *string)
	CaCertificatesBundleS3ObjectVersionInput() *string
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
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumberOfCaCertificates() *float64
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
	Status() *string
	Tags() Elasticloadbalancingv2TrustStoreTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustStoreArn() *string
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
	ResetCaCertificatesBundleS3Bucket()
	ResetCaCertificatesBundleS3Key()
	ResetCaCertificatesBundleS3ObjectVersion()
	ResetName()
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

// The jsii proxy struct for Elasticloadbalancingv2TrustStore
type jsiiProxy_Elasticloadbalancingv2TrustStore struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CaCertificatesBundleS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificatesBundleS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) NumberOfCaCertificates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCaCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) Tags() Elasticloadbalancingv2TrustStoreTagsList {
	var returns Elasticloadbalancingv2TrustStoreTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store} Resource.
func NewElasticloadbalancingv2TrustStore(scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreConfig) Elasticloadbalancingv2TrustStore {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2TrustStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2TrustStore{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticloadbalancingv2_trust_store awscc_elasticloadbalancingv2_trust_store} Resource.
func NewElasticloadbalancingv2TrustStore_Override(e Elasticloadbalancingv2TrustStore, scope constructs.Construct, id *string, config *Elasticloadbalancingv2TrustStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3Bucket(val *string) {
	if err := j.validateSetCaCertificatesBundleS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3Bucket",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3Key(val *string) {
	if err := j.validateSetCaCertificatesBundleS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3Key",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCaCertificatesBundleS3ObjectVersion(val *string) {
	if err := j.validateSetCaCertificatesBundleS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificatesBundleS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStore)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a Elasticloadbalancingv2TrustStore resource upon running "cdktn plan <stack-name>".
func Elasticloadbalancingv2TrustStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
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
func Elasticloadbalancingv2TrustStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStore_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TrustStore_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TrustStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Elasticloadbalancingv2TrustStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.elasticloadbalancingv2TrustStore.Elasticloadbalancingv2TrustStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3Bucket() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3Bucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3Key() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3Key",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetCaCertificatesBundleS3ObjectVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetCaCertificatesBundleS3ObjectVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStore) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

