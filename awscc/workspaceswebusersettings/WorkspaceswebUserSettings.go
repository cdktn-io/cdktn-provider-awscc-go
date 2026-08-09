// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspaceswebusersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/workspacesweb_user_settings awscc_workspacesweb_user_settings}.
type WorkspaceswebUserSettings interface {
	cdktn.TerraformResource
	AdditionalEncryptionContext() *map[string]*string
	SetAdditionalEncryptionContext(val *map[string]*string)
	AdditionalEncryptionContextInput() *map[string]*string
	AssociatedPortalArns() *[]*string
	BrandingConfiguration() WorkspaceswebUserSettingsBrandingConfigurationOutputReference
	BrandingConfigurationInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CookieSynchronizationConfiguration() WorkspaceswebUserSettingsCookieSynchronizationConfigurationOutputReference
	CookieSynchronizationConfigurationInput() interface{}
	CopyAllowed() *string
	SetCopyAllowed(val *string)
	CopyAllowedInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerManagedKey() *string
	SetCustomerManagedKey(val *string)
	CustomerManagedKeyInput() *string
	DeepLinkAllowed() *string
	SetDeepLinkAllowed(val *string)
	DeepLinkAllowedInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisconnectTimeoutInMinutes() *float64
	SetDisconnectTimeoutInMinutes(val *float64)
	DisconnectTimeoutInMinutesInput() *float64
	DownloadAllowed() *string
	SetDownloadAllowed(val *string)
	DownloadAllowedInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdleDisconnectTimeoutInMinutes() *float64
	SetIdleDisconnectTimeoutInMinutes(val *float64)
	IdleDisconnectTimeoutInMinutesInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasteAllowed() *string
	SetPasteAllowed(val *string)
	PasteAllowedInput() *string
	PrintAllowed() *string
	SetPrintAllowed(val *string)
	PrintAllowedInput() *string
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
	Tags() WorkspaceswebUserSettingsTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ToolbarConfiguration() WorkspaceswebUserSettingsToolbarConfigurationOutputReference
	ToolbarConfigurationInput() interface{}
	UploadAllowed() *string
	SetUploadAllowed(val *string)
	UploadAllowedInput() *string
	UserSettingsArn() *string
	WebAuthnAllowed() *string
	SetWebAuthnAllowed(val *string)
	WebAuthnAllowedInput() *string
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
	PutBrandingConfiguration(value *WorkspaceswebUserSettingsBrandingConfiguration)
	PutCookieSynchronizationConfiguration(value *WorkspaceswebUserSettingsCookieSynchronizationConfiguration)
	PutTags(value interface{})
	PutToolbarConfiguration(value *WorkspaceswebUserSettingsToolbarConfiguration)
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
	ResetAdditionalEncryptionContext()
	ResetBrandingConfiguration()
	ResetCookieSynchronizationConfiguration()
	ResetCustomerManagedKey()
	ResetDeepLinkAllowed()
	ResetDisconnectTimeoutInMinutes()
	ResetIdleDisconnectTimeoutInMinutes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetToolbarConfiguration()
	ResetWebAuthnAllowed()
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

// The jsii proxy struct for WorkspaceswebUserSettings
type jsiiProxy_WorkspaceswebUserSettings struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AdditionalEncryptionContext() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AdditionalEncryptionContextInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) AssociatedPortalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedPortalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) BrandingConfiguration() WorkspaceswebUserSettingsBrandingConfigurationOutputReference {
	var returns WorkspaceswebUserSettingsBrandingConfigurationOutputReference
	_jsii_.Get(
		j,
		"brandingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) BrandingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"brandingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CookieSynchronizationConfiguration() WorkspaceswebUserSettingsCookieSynchronizationConfigurationOutputReference {
	var returns WorkspaceswebUserSettingsCookieSynchronizationConfigurationOutputReference
	_jsii_.Get(
		j,
		"cookieSynchronizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CookieSynchronizationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieSynchronizationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CopyAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CopyAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CustomerManagedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) CustomerManagedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DeepLinkAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DeepLinkAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DownloadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) DownloadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) IdleDisconnectTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PasteAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PasteAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PrintAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) PrintAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) Tags() WorkspaceswebUserSettingsTagsList {
	var returns WorkspaceswebUserSettingsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ToolbarConfiguration() WorkspaceswebUserSettingsToolbarConfigurationOutputReference {
	var returns WorkspaceswebUserSettingsToolbarConfigurationOutputReference
	_jsii_.Get(
		j,
		"toolbarConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) ToolbarConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolbarConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UploadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UploadAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) UserSettingsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSettingsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) WebAuthnAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAuthnAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettings) WebAuthnAllowedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAuthnAllowedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/workspacesweb_user_settings awscc_workspacesweb_user_settings} Resource.
func NewWorkspaceswebUserSettings(scope constructs.Construct, id *string, config *WorkspaceswebUserSettingsConfig) WorkspaceswebUserSettings {
	_init_.Initialize()

	if err := validateNewWorkspaceswebUserSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebUserSettings{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/workspacesweb_user_settings awscc_workspacesweb_user_settings} Resource.
func NewWorkspaceswebUserSettings_Override(w WorkspaceswebUserSettings, scope constructs.Construct, id *string, config *WorkspaceswebUserSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetAdditionalEncryptionContext(val *map[string]*string) {
	if err := j.validateSetAdditionalEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCopyAllowed(val *string) {
	if err := j.validateSetCopyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetCustomerManagedKey(val *string) {
	if err := j.validateSetCustomerManagedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKey",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDeepLinkAllowed(val *string) {
	if err := j.validateSetDeepLinkAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deepLinkAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetDownloadAllowed(val *string) {
	if err := j.validateSetDownloadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downloadAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetIdleDisconnectTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleDisconnectTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetPasteAllowed(val *string) {
	if err := j.validateSetPasteAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pasteAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetPrintAllowed(val *string) {
	if err := j.validateSetPrintAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetUploadAllowed(val *string) {
	if err := j.validateSetUploadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadAllowed",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettings)SetWebAuthnAllowed(val *string) {
	if err := j.validateSetWebAuthnAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAuthnAllowed",
		val,
	)
}

// Generates CDKTN code for importing a WorkspaceswebUserSettings resource upon running "cdktn plan <stack-name>".
func WorkspaceswebUserSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
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
func WorkspaceswebUserSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceswebUserSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspaceswebUserSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspaceswebUserSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspaceswebUserSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutBrandingConfiguration(value *WorkspaceswebUserSettingsBrandingConfiguration) {
	if err := w.validatePutBrandingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBrandingConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutCookieSynchronizationConfiguration(value *WorkspaceswebUserSettingsCookieSynchronizationConfiguration) {
	if err := w.validatePutCookieSynchronizationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookieSynchronizationConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutTags(value interface{}) {
	if err := w.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTags",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) PutToolbarConfiguration(value *WorkspaceswebUserSettingsToolbarConfiguration) {
	if err := w.validatePutToolbarConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putToolbarConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetAdditionalEncryptionContext() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalEncryptionContext",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetBrandingConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetBrandingConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetCookieSynchronizationConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieSynchronizationConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetCustomerManagedKey() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomerManagedKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetDeepLinkAllowed() {
	_jsii_.InvokeVoid(
		w,
		"resetDeepLinkAllowed",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetIdleDisconnectTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetIdleDisconnectTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetToolbarConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetToolbarConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ResetWebAuthnAllowed() {
	_jsii_.InvokeVoid(
		w,
		"resetWebAuthnAllowed",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettings) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

