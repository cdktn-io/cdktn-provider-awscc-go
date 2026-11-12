// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnreplicationconfigurationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mgnreplicationconfigurationtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template awscc_mgn_replication_configuration_template}.
type MgnReplicationConfigurationTemplate interface {
	cdktn.TerraformResource
	Arn() *string
	AssociateDefaultSecurityGroup() interface{}
	SetAssociateDefaultSecurityGroup(val interface{})
	AssociateDefaultSecurityGroupInput() interface{}
	BandwidthThrottling() *float64
	SetBandwidthThrottling(val *float64)
	BandwidthThrottlingInput() *float64
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
	CreatePublicIp() interface{}
	SetCreatePublicIp(val interface{})
	CreatePublicIpInput() interface{}
	DataPlaneRouting() *string
	SetDataPlaneRouting(val *string)
	DataPlaneRoutingInput() *string
	DefaultLargeStagingDiskType() *string
	SetDefaultLargeStagingDiskType(val *string)
	DefaultLargeStagingDiskTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EbsEncryption() *string
	SetEbsEncryption(val *string)
	EbsEncryptionInput() *string
	EbsEncryptionKeyArn() *string
	SetEbsEncryptionKeyArn(val *string)
	EbsEncryptionKeyArnInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InternetProtocol() *string
	SetInternetProtocol(val *string)
	InternetProtocolInput() *string
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
	ReplicationConfigurationTemplateId() *string
	ReplicationServerInstanceType() *string
	SetReplicationServerInstanceType(val *string)
	ReplicationServerInstanceTypeInput() *string
	ReplicationServersSecurityGroupsIDs() *[]*string
	SetReplicationServersSecurityGroupsIDs(val *[]*string)
	ReplicationServersSecurityGroupsIDsInput() *[]*string
	StagingAreaSubnetId() *string
	SetStagingAreaSubnetId(val *string)
	StagingAreaSubnetIdInput() *string
	StagingAreaTags() *map[string]*string
	SetStagingAreaTags(val *map[string]*string)
	StagingAreaTagsInput() *map[string]*string
	StoreSnapshotOnLocalZone() interface{}
	SetStoreSnapshotOnLocalZone(val interface{})
	StoreSnapshotOnLocalZoneInput() interface{}
	Tags() MgnReplicationConfigurationTemplateTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseDedicatedReplicationServer() interface{}
	SetUseDedicatedReplicationServer(val interface{})
	UseDedicatedReplicationServerInput() interface{}
	UseFipsEndpoint() interface{}
	SetUseFipsEndpoint(val interface{})
	UseFipsEndpointInput() interface{}
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
	ResetEbsEncryptionKeyArn()
	ResetInternetProtocol()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStoreSnapshotOnLocalZone()
	ResetTags()
	ResetUseFipsEndpoint()
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

// The jsii proxy struct for MgnReplicationConfigurationTemplate
type jsiiProxy_MgnReplicationConfigurationTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) AssociateDefaultSecurityGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) AssociateDefaultSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) BandwidthThrottling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) BandwidthThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) CreatePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) CreatePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) DataPlaneRouting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) DataPlaneRoutingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) DefaultLargeStagingDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) DefaultLargeStagingDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) EbsEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) EbsEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) EbsEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) EbsEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) InternetProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) InternetProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ReplicationConfigurationTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationConfigurationTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ReplicationServerInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ReplicationServerInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIDs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIDs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIDsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIDsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StagingAreaSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StagingAreaSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StagingAreaTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StagingAreaTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StoreSnapshotOnLocalZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeSnapshotOnLocalZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) StoreSnapshotOnLocalZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeSnapshotOnLocalZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) Tags() MgnReplicationConfigurationTemplateTagsList {
	var returns MgnReplicationConfigurationTemplateTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) UseDedicatedReplicationServer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) UseDedicatedReplicationServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) UseFipsEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate) UseFipsEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template awscc_mgn_replication_configuration_template} Resource.
func NewMgnReplicationConfigurationTemplate(scope constructs.Construct, id *string, config *MgnReplicationConfigurationTemplateConfig) MgnReplicationConfigurationTemplate {
	_init_.Initialize()

	if err := validateNewMgnReplicationConfigurationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MgnReplicationConfigurationTemplate{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mgn_replication_configuration_template awscc_mgn_replication_configuration_template} Resource.
func NewMgnReplicationConfigurationTemplate_Override(m MgnReplicationConfigurationTemplate, scope constructs.Construct, id *string, config *MgnReplicationConfigurationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetAssociateDefaultSecurityGroup(val interface{}) {
	if err := j.validateSetAssociateDefaultSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetBandwidthThrottling(val *float64) {
	if err := j.validateSetBandwidthThrottlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthThrottling",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetCreatePublicIp(val interface{}) {
	if err := j.validateSetCreatePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPublicIp",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetDataPlaneRouting(val *string) {
	if err := j.validateSetDataPlaneRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPlaneRouting",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetDefaultLargeStagingDiskType(val *string) {
	if err := j.validateSetDefaultLargeStagingDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLargeStagingDiskType",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetEbsEncryption(val *string) {
	if err := j.validateSetEbsEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryption",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetEbsEncryptionKeyArn(val *string) {
	if err := j.validateSetEbsEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetInternetProtocol(val *string) {
	if err := j.validateSetInternetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetProtocol",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetReplicationServerInstanceType(val *string) {
	if err := j.validateSetReplicationServerInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServerInstanceType",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetReplicationServersSecurityGroupsIDs(val *[]*string) {
	if err := j.validateSetReplicationServersSecurityGroupsIDsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServersSecurityGroupsIDs",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetStagingAreaSubnetId(val *string) {
	if err := j.validateSetStagingAreaSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaSubnetId",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetStagingAreaTags(val *map[string]*string) {
	if err := j.validateSetStagingAreaTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaTags",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetStoreSnapshotOnLocalZone(val interface{}) {
	if err := j.validateSetStoreSnapshotOnLocalZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeSnapshotOnLocalZone",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetUseDedicatedReplicationServer(val interface{}) {
	if err := j.validateSetUseDedicatedReplicationServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDedicatedReplicationServer",
		val,
	)
}

func (j *jsiiProxy_MgnReplicationConfigurationTemplate)SetUseFipsEndpoint(val interface{}) {
	if err := j.validateSetUseFipsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFipsEndpoint",
		val,
	)
}

// Generates CDKTN code for importing a MgnReplicationConfigurationTemplate resource upon running "cdktn plan <stack-name>".
func MgnReplicationConfigurationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMgnReplicationConfigurationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
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
func MgnReplicationConfigurationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMgnReplicationConfigurationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MgnReplicationConfigurationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMgnReplicationConfigurationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MgnReplicationConfigurationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMgnReplicationConfigurationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MgnReplicationConfigurationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.mgnReplicationConfigurationTemplate.MgnReplicationConfigurationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := m.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := m.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetEbsEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		m,
		"resetEbsEncryptionKeyArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetInternetProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetInternetProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetStoreSnapshotOnLocalZone() {
	_jsii_.InvokeVoid(
		m,
		"resetStoreSnapshotOnLocalZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ResetUseFipsEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetUseFipsEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnReplicationConfigurationTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

