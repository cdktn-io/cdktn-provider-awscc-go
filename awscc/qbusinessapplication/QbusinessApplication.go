// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinessapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_application awscc_qbusiness_application}.
type QbusinessApplication interface {
	cdktn.TerraformResource
	ApplicationArn() *string
	ApplicationId() *string
	AttachmentsConfiguration() QbusinessApplicationAttachmentsConfigurationOutputReference
	AttachmentsConfigurationInput() interface{}
	AutoSubscriptionConfiguration() QbusinessApplicationAutoSubscriptionConfigurationOutputReference
	AutoSubscriptionConfigurationInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientIdsForOidc() *[]*string
	SetClientIdsForOidc(val *[]*string)
	ClientIdsForOidcInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionConfiguration() QbusinessApplicationEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamIdentityProviderArn() *string
	SetIamIdentityProviderArn(val *string)
	IamIdentityProviderArnInput() *string
	Id() *string
	IdentityCenterApplicationArn() *string
	IdentityCenterInstanceArn() *string
	SetIdentityCenterInstanceArn(val *string)
	IdentityCenterInstanceArnInput() *string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PersonalizationConfiguration() QbusinessApplicationPersonalizationConfigurationOutputReference
	PersonalizationConfigurationInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QAppsConfiguration() QbusinessApplicationQAppsConfigurationOutputReference
	QAppsConfigurationInput() interface{}
	QuickSightConfiguration() QbusinessApplicationQuickSightConfigurationOutputReference
	QuickSightConfigurationInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Status() *string
	Tags() QbusinessApplicationTagsList
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
	PutAttachmentsConfiguration(value *QbusinessApplicationAttachmentsConfiguration)
	PutAutoSubscriptionConfiguration(value *QbusinessApplicationAutoSubscriptionConfiguration)
	PutEncryptionConfiguration(value *QbusinessApplicationEncryptionConfiguration)
	PutPersonalizationConfiguration(value *QbusinessApplicationPersonalizationConfiguration)
	PutQAppsConfiguration(value *QbusinessApplicationQAppsConfiguration)
	PutQuickSightConfiguration(value *QbusinessApplicationQuickSightConfiguration)
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
	ResetAttachmentsConfiguration()
	ResetAutoSubscriptionConfiguration()
	ResetClientIdsForOidc()
	ResetDescription()
	ResetEncryptionConfiguration()
	ResetIamIdentityProviderArn()
	ResetIdentityCenterInstanceArn()
	ResetIdentityType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersonalizationConfiguration()
	ResetQAppsConfiguration()
	ResetQuickSightConfiguration()
	ResetRoleArn()
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

// The jsii proxy struct for QbusinessApplication
type jsiiProxy_QbusinessApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_QbusinessApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) AttachmentsConfiguration() QbusinessApplicationAttachmentsConfigurationOutputReference {
	var returns QbusinessApplicationAttachmentsConfigurationOutputReference
	_jsii_.Get(
		j,
		"attachmentsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) AttachmentsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) AutoSubscriptionConfiguration() QbusinessApplicationAutoSubscriptionConfigurationOutputReference {
	var returns QbusinessApplicationAutoSubscriptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoSubscriptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) AutoSubscriptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubscriptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) ClientIdsForOidc() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientIdsForOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) ClientIdsForOidcInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientIdsForOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) EncryptionConfiguration() QbusinessApplicationEncryptionConfigurationOutputReference {
	var returns QbusinessApplicationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) EncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IamIdentityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IamIdentityProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamIdentityProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IdentityCenterApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityCenterApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IdentityCenterInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityCenterInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IdentityCenterInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityCenterInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) PersonalizationConfiguration() QbusinessApplicationPersonalizationConfigurationOutputReference {
	var returns QbusinessApplicationPersonalizationConfigurationOutputReference
	_jsii_.Get(
		j,
		"personalizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) PersonalizationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalizationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) QAppsConfiguration() QbusinessApplicationQAppsConfigurationOutputReference {
	var returns QbusinessApplicationQAppsConfigurationOutputReference
	_jsii_.Get(
		j,
		"qAppsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) QAppsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"qAppsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) QuickSightConfiguration() QbusinessApplicationQuickSightConfigurationOutputReference {
	var returns QbusinessApplicationQuickSightConfigurationOutputReference
	_jsii_.Get(
		j,
		"quickSightConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) QuickSightConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quickSightConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) Tags() QbusinessApplicationTagsList {
	var returns QbusinessApplicationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessApplication) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_application awscc_qbusiness_application} Resource.
func NewQbusinessApplication(scope constructs.Construct, id *string, config *QbusinessApplicationConfig) QbusinessApplication {
	_init_.Initialize()

	if err := validateNewQbusinessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessApplication{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_application awscc_qbusiness_application} Resource.
func NewQbusinessApplication_Override(q QbusinessApplication, scope constructs.Construct, id *string, config *QbusinessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetClientIdsForOidc(val *[]*string) {
	if err := j.validateSetClientIdsForOidcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIdsForOidc",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetIamIdentityProviderArn(val *string) {
	if err := j.validateSetIamIdentityProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamIdentityProviderArn",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetIdentityCenterInstanceArn(val *string) {
	if err := j.validateSetIdentityCenterInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityCenterInstanceArn",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_QbusinessApplication)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Generates CDKTN code for importing a QbusinessApplication resource upon running "cdktn plan <stack-name>".
func QbusinessApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateQbusinessApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
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
func QbusinessApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QbusinessApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QbusinessApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQbusinessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QbusinessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.qbusinessApplication.QbusinessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QbusinessApplication) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QbusinessApplication) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QbusinessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QbusinessApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := q.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QbusinessApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QbusinessApplication) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QbusinessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutAttachmentsConfiguration(value *QbusinessApplicationAttachmentsConfiguration) {
	if err := q.validatePutAttachmentsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAttachmentsConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutAutoSubscriptionConfiguration(value *QbusinessApplicationAutoSubscriptionConfiguration) {
	if err := q.validatePutAutoSubscriptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAutoSubscriptionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutEncryptionConfiguration(value *QbusinessApplicationEncryptionConfiguration) {
	if err := q.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutPersonalizationConfiguration(value *QbusinessApplicationPersonalizationConfiguration) {
	if err := q.validatePutPersonalizationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPersonalizationConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutQAppsConfiguration(value *QbusinessApplicationQAppsConfiguration) {
	if err := q.validatePutQAppsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putQAppsConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutQuickSightConfiguration(value *QbusinessApplicationQuickSightConfiguration) {
	if err := q.validatePutQuickSightConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putQuickSightConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := q.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetAttachmentsConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAttachmentsConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetAutoSubscriptionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAutoSubscriptionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetClientIdsForOidc() {
	_jsii_.InvokeVoid(
		q,
		"resetClientIdsForOidc",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetIamIdentityProviderArn() {
	_jsii_.InvokeVoid(
		q,
		"resetIamIdentityProviderArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetIdentityCenterInstanceArn() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityCenterInstanceArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetIdentityType() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetPersonalizationConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetPersonalizationConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetQAppsConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetQAppsConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetQuickSightConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetQuickSightConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetRoleArn() {
	_jsii_.InvokeVoid(
		q,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		q,
		"with",
		args,
		&returns,
	)

	return returns
}

