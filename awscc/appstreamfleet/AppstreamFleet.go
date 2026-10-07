// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appstreamfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/appstream_fleet awscc_appstream_fleet}.
type AppstreamFleet interface {
	cdktn.TerraformResource
	Arn() *string
	AttributesToDelete() *[]*string
	SetAttributesToDelete(val *[]*string)
	AttributesToDeleteInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComputeCapacity() AppstreamFleetComputeCapacityOutputReference
	ComputeCapacityInput() interface{}
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
	DisableImdsv1() interface{}
	SetDisableImdsv1(val interface{})
	DisableImdsv1Input() interface{}
	DisconnectTimeoutInSeconds() *float64
	SetDisconnectTimeoutInSeconds(val *float64)
	DisconnectTimeoutInSecondsInput() *float64
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference
	DomainJoinInfoInput() interface{}
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
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
	IdleDisconnectTimeoutInSeconds() *float64
	SetIdleDisconnectTimeoutInSeconds(val *float64)
	IdleDisconnectTimeoutInSecondsInput() *float64
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxConcurrentSessions() *float64
	SetMaxConcurrentSessions(val *float64)
	MaxConcurrentSessionsInput() *float64
	MaxSessionsPerInstance() *float64
	SetMaxSessionsPerInstance(val *float64)
	MaxSessionsPerInstanceInput() *float64
	MaxUserDurationInSeconds() *float64
	SetMaxUserDurationInSeconds(val *float64)
	MaxUserDurationInSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
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
	RootVolumeConfig() AppstreamFleetRootVolumeConfigOutputReference
	RootVolumeConfigInput() interface{}
	SessionScriptS3Location() AppstreamFleetSessionScriptS3LocationOutputReference
	SessionScriptS3LocationInput() interface{}
	StreamView() *string
	SetStreamView(val *string)
	StreamViewInput() *string
	Tags() AppstreamFleetTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsbDeviceFilterStrings() *[]*string
	SetUsbDeviceFilterStrings(val *[]*string)
	UsbDeviceFilterStringsInput() *[]*string
	VpcConfig() AppstreamFleetVpcConfigOutputReference
	VpcConfigInput() interface{}
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
	PutComputeCapacity(value *AppstreamFleetComputeCapacity)
	PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo)
	PutRootVolumeConfig(value *AppstreamFleetRootVolumeConfig)
	PutSessionScriptS3Location(value *AppstreamFleetSessionScriptS3Location)
	PutTags(value interface{})
	PutVpcConfig(value *AppstreamFleetVpcConfig)
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
	ResetAttributesToDelete()
	ResetComputeCapacity()
	ResetDescription()
	ResetDisableImdsv1()
	ResetDisconnectTimeoutInSeconds()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetFleetType()
	ResetIamRoleArn()
	ResetIdleDisconnectTimeoutInSeconds()
	ResetImageArn()
	ResetImageName()
	ResetMaxConcurrentSessions()
	ResetMaxSessionsPerInstance()
	ResetMaxUserDurationInSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRootVolumeConfig()
	ResetSessionScriptS3Location()
	ResetStreamView()
	ResetTags()
	ResetUsbDeviceFilterStrings()
	ResetVpcConfig()
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

// The jsii proxy struct for AppstreamFleet
type jsiiProxy_AppstreamFleet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppstreamFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) AttributesToDelete() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributesToDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) AttributesToDeleteInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributesToDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacity() AppstreamFleetComputeCapacityOutputReference {
	var returns AppstreamFleetComputeCapacityOutputReference
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisableImdsv1() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableImdsv1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisableImdsv1Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableImdsv1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference {
	var returns AppstreamFleetDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxConcurrentSessions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentSessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxConcurrentSessionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentSessionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxSessionsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSessionsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxSessionsPerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSessionsPerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) RootVolumeConfig() AppstreamFleetRootVolumeConfigOutputReference {
	var returns AppstreamFleetRootVolumeConfigOutputReference
	_jsii_.Get(
		j,
		"rootVolumeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) RootVolumeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootVolumeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) SessionScriptS3Location() AppstreamFleetSessionScriptS3LocationOutputReference {
	var returns AppstreamFleetSessionScriptS3LocationOutputReference
	_jsii_.Get(
		j,
		"sessionScriptS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) SessionScriptS3LocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionScriptS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Tags() AppstreamFleetTagsList {
	var returns AppstreamFleetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) UsbDeviceFilterStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usbDeviceFilterStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) UsbDeviceFilterStringsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usbDeviceFilterStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfig() AppstreamFleetVpcConfigOutputReference {
	var returns AppstreamFleetVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/appstream_fleet awscc_appstream_fleet} Resource.
func NewAppstreamFleet(scope constructs.Construct, id *string, config *AppstreamFleetConfig) AppstreamFleet {
	_init_.Initialize()

	if err := validateNewAppstreamFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppstreamFleet{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/appstream_fleet awscc_appstream_fleet} Resource.
func NewAppstreamFleet_Override(a AppstreamFleet, scope constructs.Construct, id *string, config *AppstreamFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetAttributesToDelete(val *[]*string) {
	if err := j.validateSetAttributesToDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesToDelete",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetDisableImdsv1(val interface{}) {
	if err := j.validateSetDisableImdsv1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableImdsv1",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetDisconnectTimeoutInSeconds(val *float64) {
	if err := j.validateSetDisconnectTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetEnableDefaultInternetAccess(val interface{}) {
	if err := j.validateSetEnableDefaultInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetIdleDisconnectTimeoutInSeconds(val *float64) {
	if err := j.validateSetIdleDisconnectTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetImageArn(val *string) {
	if err := j.validateSetImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetMaxConcurrentSessions(val *float64) {
	if err := j.validateSetMaxConcurrentSessionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentSessions",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetMaxSessionsPerInstance(val *float64) {
	if err := j.validateSetMaxSessionsPerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSessionsPerInstance",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetMaxUserDurationInSeconds(val *float64) {
	if err := j.validateSetMaxUserDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUserDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetStreamView(val *string) {
	if err := j.validateSetStreamViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamView",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet)SetUsbDeviceFilterStrings(val *[]*string) {
	if err := j.validateSetUsbDeviceFilterStringsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usbDeviceFilterStrings",
		val,
	)
}

// Generates CDKTN code for importing a AppstreamFleet resource upon running "cdktn plan <stack-name>".
func AppstreamFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppstreamFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
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
func AppstreamFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.appstreamFleet.AppstreamFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppstreamFleet) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppstreamFleet) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppstreamFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppstreamFleet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppstreamFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppstreamFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppstreamFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppstreamFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppstreamFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppstreamFleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppstreamFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppstreamFleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppstreamFleet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppstreamFleet) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AppstreamFleet) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppstreamFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppstreamFleet) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppstreamFleet) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutComputeCapacity(value *AppstreamFleetComputeCapacity) {
	if err := a.validatePutComputeCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComputeCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo) {
	if err := a.validatePutDomainJoinInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutRootVolumeConfig(value *AppstreamFleetRootVolumeConfig) {
	if err := a.validatePutRootVolumeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRootVolumeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutSessionScriptS3Location(value *AppstreamFleetSessionScriptS3Location) {
	if err := a.validatePutSessionScriptS3LocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSessionScriptS3Location",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutTags(value interface{}) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutVpcConfig(value *AppstreamFleetVpcConfig) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetAttributesToDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributesToDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetComputeCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisableImdsv1() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableImdsv1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIdleDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetMaxConcurrentSessions() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentSessions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetMaxSessionsPerInstance() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxSessionsPerInstance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetMaxUserDurationInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUserDurationInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetPlatform() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatform",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetRootVolumeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRootVolumeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetSessionScriptS3Location() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionScriptS3Location",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetStreamView() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamView",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetUsbDeviceFilterStrings() {
	_jsii_.InvokeVoid(
		a,
		"resetUsbDeviceFilterStrings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleet) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

