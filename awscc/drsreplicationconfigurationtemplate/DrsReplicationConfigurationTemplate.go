// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drsreplicationconfigurationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/drsreplicationconfigurationtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template awscc_drs_replication_configuration_template}.
type DrsReplicationConfigurationTemplate interface {
	cdktn.TerraformResource
	Arn() *string
	AssociateDefaultSecurityGroup() interface{}
	SetAssociateDefaultSecurityGroup(val interface{})
	AssociateDefaultSecurityGroupInput() interface{}
	AutoReplicateNewDisks() interface{}
	SetAutoReplicateNewDisks(val interface{})
	AutoReplicateNewDisksInput() interface{}
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
	PitPolicy() DrsReplicationConfigurationTemplatePitPolicyList
	PitPolicyInput() interface{}
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
	Tags() DrsReplicationConfigurationTemplateTagsList
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
	PutPitPolicy(value interface{})
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
	ResetAssociateDefaultSecurityGroup()
	ResetAutoReplicateNewDisks()
	ResetCreatePublicIp()
	ResetDataPlaneRouting()
	ResetDefaultLargeStagingDiskType()
	ResetEbsEncryptionKeyArn()
	ResetInternetProtocol()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplicationServerInstanceType()
	ResetTags()
	ResetUseDedicatedReplicationServer()
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

// The jsii proxy struct for DrsReplicationConfigurationTemplate
type jsiiProxy_DrsReplicationConfigurationTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) AssociateDefaultSecurityGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) AssociateDefaultSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) AutoReplicateNewDisks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) AutoReplicateNewDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) BandwidthThrottling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) BandwidthThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) CreatePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) CreatePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) DataPlaneRouting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) DataPlaneRoutingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) DefaultLargeStagingDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) DefaultLargeStagingDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) EbsEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) EbsEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) EbsEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) EbsEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) InternetProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) InternetProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internetProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) PitPolicy() DrsReplicationConfigurationTemplatePitPolicyList {
	var returns DrsReplicationConfigurationTemplatePitPolicyList
	_jsii_.Get(
		j,
		"pitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) PitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ReplicationConfigurationTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationConfigurationTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ReplicationServerInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ReplicationServerInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIDs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIDs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIDsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIDsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) StagingAreaSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) StagingAreaSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) StagingAreaTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) StagingAreaTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) Tags() DrsReplicationConfigurationTemplateTagsList {
	var returns DrsReplicationConfigurationTemplateTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) UseDedicatedReplicationServer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate) UseDedicatedReplicationServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template awscc_drs_replication_configuration_template} Resource.
func NewDrsReplicationConfigurationTemplate(scope constructs.Construct, id *string, config *DrsReplicationConfigurationTemplateConfig) DrsReplicationConfigurationTemplate {
	_init_.Initialize()

	if err := validateNewDrsReplicationConfigurationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DrsReplicationConfigurationTemplate{}

	_jsii_.Create(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template awscc_drs_replication_configuration_template} Resource.
func NewDrsReplicationConfigurationTemplate_Override(d DrsReplicationConfigurationTemplate, scope constructs.Construct, id *string, config *DrsReplicationConfigurationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetAssociateDefaultSecurityGroup(val interface{}) {
	if err := j.validateSetAssociateDefaultSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetAutoReplicateNewDisks(val interface{}) {
	if err := j.validateSetAutoReplicateNewDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoReplicateNewDisks",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetBandwidthThrottling(val *float64) {
	if err := j.validateSetBandwidthThrottlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthThrottling",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetCreatePublicIp(val interface{}) {
	if err := j.validateSetCreatePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPublicIp",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetDataPlaneRouting(val *string) {
	if err := j.validateSetDataPlaneRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPlaneRouting",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetDefaultLargeStagingDiskType(val *string) {
	if err := j.validateSetDefaultLargeStagingDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLargeStagingDiskType",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetEbsEncryption(val *string) {
	if err := j.validateSetEbsEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryption",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetEbsEncryptionKeyArn(val *string) {
	if err := j.validateSetEbsEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetInternetProtocol(val *string) {
	if err := j.validateSetInternetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetProtocol",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetReplicationServerInstanceType(val *string) {
	if err := j.validateSetReplicationServerInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServerInstanceType",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetReplicationServersSecurityGroupsIDs(val *[]*string) {
	if err := j.validateSetReplicationServersSecurityGroupsIDsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServersSecurityGroupsIDs",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetStagingAreaSubnetId(val *string) {
	if err := j.validateSetStagingAreaSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaSubnetId",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetStagingAreaTags(val *map[string]*string) {
	if err := j.validateSetStagingAreaTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaTags",
		val,
	)
}

func (j *jsiiProxy_DrsReplicationConfigurationTemplate)SetUseDedicatedReplicationServer(val interface{}) {
	if err := j.validateSetUseDedicatedReplicationServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDedicatedReplicationServer",
		val,
	)
}

// Generates CDKTN code for importing a DrsReplicationConfigurationTemplate resource upon running "cdktn plan <stack-name>".
func DrsReplicationConfigurationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDrsReplicationConfigurationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
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
func DrsReplicationConfigurationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsReplicationConfigurationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DrsReplicationConfigurationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsReplicationConfigurationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DrsReplicationConfigurationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDrsReplicationConfigurationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DrsReplicationConfigurationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.drsReplicationConfigurationTemplate.DrsReplicationConfigurationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) PutPitPolicy(value interface{}) {
	if err := d.validatePutPitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPitPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetAssociateDefaultSecurityGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetAssociateDefaultSecurityGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetAutoReplicateNewDisks() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoReplicateNewDisks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetCreatePublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatePublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetDataPlaneRouting() {
	_jsii_.InvokeVoid(
		d,
		"resetDataPlaneRouting",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetDefaultLargeStagingDiskType() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultLargeStagingDiskType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetEbsEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetEbsEncryptionKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetInternetProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetInternetProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetReplicationServerInstanceType() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationServerInstanceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ResetUseDedicatedReplicationServer() {
	_jsii_.InvokeVoid(
		d,
		"resetUseDedicatedReplicationServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DrsReplicationConfigurationTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

