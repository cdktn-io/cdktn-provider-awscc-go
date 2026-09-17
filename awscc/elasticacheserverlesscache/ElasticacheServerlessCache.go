// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticacheserverlesscache/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_serverless_cache awscc_elasticache_serverless_cache}.
type ElasticacheServerlessCache interface {
	cdktn.TerraformResource
	Arn() *string
	CacheUsageLimits() ElasticacheServerlessCacheCacheUsageLimitsOutputReference
	CacheUsageLimitsInput() interface{}
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
	CreateTime() *string
	DailySnapshotTime() *string
	SetDailySnapshotTime(val *string)
	DailySnapshotTimeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Endpoint() ElasticacheServerlessCacheEndpointOutputReference
	EndpointInput() interface{}
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	FinalSnapshotName() *string
	SetFinalSnapshotName(val *string)
	FinalSnapshotNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullEngineVersion() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MajorEngineVersion() *string
	SetMajorEngineVersion(val *string)
	MajorEngineVersionInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
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
	ReaderEndpoint() ElasticacheServerlessCacheReaderEndpointOutputReference
	ReaderEndpointInput() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	ServerlessCacheName() *string
	SetServerlessCacheName(val *string)
	ServerlessCacheNameInput() *string
	SnapshotArnsToRestore() *[]*string
	SetSnapshotArnsToRestore(val *[]*string)
	SnapshotArnsToRestoreInput() *[]*string
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	SnapshotRetentionLimitInput() *float64
	Status() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() ElasticacheServerlessCacheTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserGroupId() *string
	SetUserGroupId(val *string)
	UserGroupIdInput() *string
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
	PutCacheUsageLimits(value *ElasticacheServerlessCacheCacheUsageLimits)
	PutEndpoint(value *ElasticacheServerlessCacheEndpoint)
	PutReaderEndpoint(value *ElasticacheServerlessCacheReaderEndpoint)
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
	ResetCacheUsageLimits()
	ResetDailySnapshotTime()
	ResetDescription()
	ResetEndpoint()
	ResetFinalSnapshotName()
	ResetKmsKeyId()
	ResetMajorEngineVersion()
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReaderEndpoint()
	ResetSecurityGroupIds()
	ResetSnapshotArnsToRestore()
	ResetSnapshotRetentionLimit()
	ResetSubnetIds()
	ResetTags()
	ResetUserGroupId()
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

// The jsii proxy struct for ElasticacheServerlessCache
type jsiiProxy_ElasticacheServerlessCache struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ElasticacheServerlessCache) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) CacheUsageLimits() ElasticacheServerlessCacheCacheUsageLimitsOutputReference {
	var returns ElasticacheServerlessCacheCacheUsageLimitsOutputReference
	_jsii_.Get(
		j,
		"cacheUsageLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) CacheUsageLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUsageLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) DailySnapshotTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailySnapshotTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) DailySnapshotTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailySnapshotTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Endpoint() ElasticacheServerlessCacheEndpointOutputReference {
	var returns ElasticacheServerlessCacheEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) FinalSnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) FinalSnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) FullEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) MajorEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) MajorEngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorEngineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ReaderEndpoint() ElasticacheServerlessCacheReaderEndpointOutputReference {
	var returns ElasticacheServerlessCacheReaderEndpointOutputReference
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ReaderEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ServerlessCacheName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessCacheName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) ServerlessCacheNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessCacheNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SnapshotArnsToRestore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsToRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SnapshotArnsToRestoreInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsToRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) Tags() ElasticacheServerlessCacheTagsList {
	var returns ElasticacheServerlessCacheTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) UserGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheServerlessCache) UserGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGroupIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_serverless_cache awscc_elasticache_serverless_cache} Resource.
func NewElasticacheServerlessCache(scope constructs.Construct, id *string, config *ElasticacheServerlessCacheConfig) ElasticacheServerlessCache {
	_init_.Initialize()

	if err := validateNewElasticacheServerlessCacheParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheServerlessCache{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_serverless_cache awscc_elasticache_serverless_cache} Resource.
func NewElasticacheServerlessCache_Override(e ElasticacheServerlessCache, scope constructs.Construct, id *string, config *ElasticacheServerlessCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetDailySnapshotTime(val *string) {
	if err := j.validateSetDailySnapshotTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailySnapshotTime",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetFinalSnapshotName(val *string) {
	if err := j.validateSetFinalSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetMajorEngineVersion(val *string) {
	if err := j.validateSetMajorEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"majorEngineVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetServerlessCacheName(val *string) {
	if err := j.validateSetServerlessCacheNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessCacheName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetSnapshotArnsToRestore(val *[]*string) {
	if err := j.validateSetSnapshotArnsToRestoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArnsToRestore",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ElasticacheServerlessCache)SetUserGroupId(val *string) {
	if err := j.validateSetUserGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userGroupId",
		val,
	)
}

// Generates CDKTN code for importing a ElasticacheServerlessCache resource upon running "cdktn plan <stack-name>".
func ElasticacheServerlessCache_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateElasticacheServerlessCache_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
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
func ElasticacheServerlessCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheServerlessCache_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheServerlessCache_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheServerlessCache_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheServerlessCache_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheServerlessCache_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticacheServerlessCache_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.elasticacheServerlessCache.ElasticacheServerlessCache",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticacheServerlessCache) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticacheServerlessCache) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticacheServerlessCache) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_ElasticacheServerlessCache) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) PutCacheUsageLimits(value *ElasticacheServerlessCacheCacheUsageLimits) {
	if err := e.validatePutCacheUsageLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCacheUsageLimits",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) PutEndpoint(value *ElasticacheServerlessCacheEndpoint) {
	if err := e.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) PutReaderEndpoint(value *ElasticacheServerlessCacheReaderEndpoint) {
	if err := e.validatePutReaderEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putReaderEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetCacheUsageLimits() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheUsageLimits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetDailySnapshotTime() {
	_jsii_.InvokeVoid(
		e,
		"resetDailySnapshotTime",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetFinalSnapshotName() {
	_jsii_.InvokeVoid(
		e,
		"resetFinalSnapshotName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetMajorEngineVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetMajorEngineVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetNetworkType() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetReaderEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetReaderEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetSnapshotArnsToRestore() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotArnsToRestore",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) ResetUserGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetUserGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheServerlessCache) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheServerlessCache) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

