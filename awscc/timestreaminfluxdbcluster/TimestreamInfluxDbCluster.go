// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/timestreaminfluxdbcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_influx_db_cluster awscc_timestream_influx_db_cluster}.
type TimestreamInfluxDbCluster interface {
	cdktn.TerraformResource
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	AllocatedStorageInput() *float64
	Arn() *string
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	DbInstanceType() *string
	SetDbInstanceType(val *string)
	DbInstanceTypeInput() *string
	DbParameterGroupIdentifier() *string
	SetDbParameterGroupIdentifier(val *string)
	DbParameterGroupIdentifierInput() *string
	DbStorageType() *string
	SetDbStorageType(val *string)
	DbStorageTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	Endpoint() *string
	EngineType() *string
	FailoverMode() *string
	SetFailoverMode(val *string)
	FailoverModeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InfluxAuthParametersSecretArn() *string
	InfluxDbClusterId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogDeliveryConfiguration() TimestreamInfluxDbClusterLogDeliveryConfigurationOutputReference
	LogDeliveryConfigurationInput() interface{}
	MaintenanceSchedule() TimestreamInfluxDbClusterMaintenanceScheduleOutputReference
	MaintenanceScheduleInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	NextMaintenanceTime() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReaderEndpoint() *string
	Status() *string
	Tags() TimestreamInfluxDbClusterTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutLogDeliveryConfiguration(value *TimestreamInfluxDbClusterLogDeliveryConfiguration)
	PutMaintenanceSchedule(value *TimestreamInfluxDbClusterMaintenanceSchedule)
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
	ResetAllocatedStorage()
	ResetBucket()
	ResetDbInstanceType()
	ResetDbParameterGroupIdentifier()
	ResetDbStorageType()
	ResetDeploymentType()
	ResetFailoverMode()
	ResetLogDeliveryConfiguration()
	ResetMaintenanceSchedule()
	ResetName()
	ResetNetworkType()
	ResetOrganization()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetPubliclyAccessible()
	ResetTags()
	ResetUsername()
	ResetVpcSecurityGroupIds()
	ResetVpcSubnetIds()
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

// The jsii proxy struct for TimestreamInfluxDbCluster
type jsiiProxy_TimestreamInfluxDbCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbParameterGroupIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbParameterGroupIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbStorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbStorageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DbStorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbStorageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) FailoverMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) FailoverModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) InfluxAuthParametersSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"influxAuthParametersSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) InfluxDbClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"influxDbClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) LogDeliveryConfiguration() TimestreamInfluxDbClusterLogDeliveryConfigurationOutputReference {
	var returns TimestreamInfluxDbClusterLogDeliveryConfigurationOutputReference
	_jsii_.Get(
		j,
		"logDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) LogDeliveryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) MaintenanceSchedule() TimestreamInfluxDbClusterMaintenanceScheduleOutputReference {
	var returns TimestreamInfluxDbClusterMaintenanceScheduleOutputReference
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) MaintenanceScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) NextMaintenanceTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) ReaderEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Tags() TimestreamInfluxDbClusterTagsList {
	var returns TimestreamInfluxDbClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) VpcSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamInfluxDbCluster) VpcSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSubnetIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_influx_db_cluster awscc_timestream_influx_db_cluster} Resource.
func NewTimestreamInfluxDbCluster(scope constructs.Construct, id *string, config *TimestreamInfluxDbClusterConfig) TimestreamInfluxDbCluster {
	_init_.Initialize()

	if err := validateNewTimestreamInfluxDbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamInfluxDbCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_influx_db_cluster awscc_timestream_influx_db_cluster} Resource.
func NewTimestreamInfluxDbCluster_Override(t TimestreamInfluxDbCluster, scope constructs.Construct, id *string, config *TimestreamInfluxDbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetDbInstanceType(val *string) {
	if err := j.validateSetDbInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceType",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetDbParameterGroupIdentifier(val *string) {
	if err := j.validateSetDbParameterGroupIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbParameterGroupIdentifier",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetDbStorageType(val *string) {
	if err := j.validateSetDbStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbStorageType",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetFailoverMode(val *string) {
	if err := j.validateSetFailoverModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverMode",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TimestreamInfluxDbCluster)SetVpcSubnetIds(val *[]*string) {
	if err := j.validateSetVpcSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSubnetIds",
		val,
	)
}

// Generates CDKTN code for importing a TimestreamInfluxDbCluster resource upon running "cdktn plan <stack-name>".
func TimestreamInfluxDbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTimestreamInfluxDbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
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
func TimestreamInfluxDbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamInfluxDbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamInfluxDbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamInfluxDbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamInfluxDbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamInfluxDbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TimestreamInfluxDbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.timestreamInfluxDbCluster.TimestreamInfluxDbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) PutLogDeliveryConfiguration(value *TimestreamInfluxDbClusterLogDeliveryConfiguration) {
	if err := t.validatePutLogDeliveryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogDeliveryConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) PutMaintenanceSchedule(value *TimestreamInfluxDbClusterMaintenanceSchedule) {
	if err := t.validatePutMaintenanceScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceSchedule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) PutTags(value interface{}) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetBucket() {
	_jsii_.InvokeVoid(
		t,
		"resetBucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetDbInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetDbInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetDbParameterGroupIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetDbParameterGroupIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetDbStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetDbStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetFailoverMode() {
	_jsii_.InvokeVoid(
		t,
		"resetFailoverMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetLogDeliveryConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLogDeliveryConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetMaintenanceSchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceSchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetOrganization() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		t,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ResetVpcSubnetIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSubnetIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamInfluxDbCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

