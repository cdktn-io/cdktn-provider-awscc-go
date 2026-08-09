// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53healthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/route53healthcheck/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53HealthCheckHealthCheckConfigOutputReference interface {
	cdktn.ComplexObject
	AlarmIdentifier() Route53HealthCheckHealthCheckConfigAlarmIdentifierOutputReference
	AlarmIdentifierInput() interface{}
	ChildHealthChecks() *[]*string
	SetChildHealthChecks(val *[]*string)
	ChildHealthChecksInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableSni() interface{}
	SetEnableSni(val interface{})
	EnableSniInput() interface{}
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	FullyQualifiedDomainName() *string
	SetFullyQualifiedDomainName(val *string)
	FullyQualifiedDomainNameInput() *string
	HealthThreshold() *float64
	SetHealthThreshold(val *float64)
	HealthThresholdInput() *float64
	InsufficientDataHealthStatus() *string
	SetInsufficientDataHealthStatus(val *string)
	InsufficientDataHealthStatusInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Inverted() interface{}
	SetInverted(val interface{})
	InvertedInput() interface{}
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	MeasureLatency() interface{}
	SetMeasureLatency(val interface{})
	MeasureLatencyInput() interface{}
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RequestInterval() *float64
	SetRequestInterval(val *float64)
	RequestIntervalInput() *float64
	ResourcePath() *string
	SetResourcePath(val *string)
	ResourcePathInput() *string
	RoutingControlArn() *string
	SetRoutingControlArn(val *string)
	RoutingControlArnInput() *string
	SearchString() *string
	SetSearchString(val *string)
	SearchStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAlarmIdentifier(value *Route53HealthCheckHealthCheckConfigAlarmIdentifier)
	ResetAlarmIdentifier()
	ResetChildHealthChecks()
	ResetEnableSni()
	ResetFailureThreshold()
	ResetFullyQualifiedDomainName()
	ResetHealthThreshold()
	ResetInsufficientDataHealthStatus()
	ResetInverted()
	ResetIpAddress()
	ResetMeasureLatency()
	ResetPort()
	ResetRegions()
	ResetRequestInterval()
	ResetResourcePath()
	ResetRoutingControlArn()
	ResetSearchString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53HealthCheckHealthCheckConfigOutputReference
type jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) AlarmIdentifier() Route53HealthCheckHealthCheckConfigAlarmIdentifierOutputReference {
	var returns Route53HealthCheckHealthCheckConfigAlarmIdentifierOutputReference
	_jsii_.Get(
		j,
		"alarmIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) AlarmIdentifierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ChildHealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ChildHealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) EnableSni() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) EnableSniInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) FullyQualifiedDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) FullyQualifiedDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) HealthThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) HealthThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InsufficientDataHealthStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InsufficientDataHealthStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Inverted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inverted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InvertedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) MeasureLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) MeasureLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) RequestInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) RequestIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) RoutingControlArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) RoutingControlArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) SearchString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) SearchStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewRoute53HealthCheckHealthCheckConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Route53HealthCheckHealthCheckConfigOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53HealthCheckHealthCheckConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.route53HealthCheck.Route53HealthCheckHealthCheckConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoute53HealthCheckHealthCheckConfigOutputReference_Override(r Route53HealthCheckHealthCheckConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.route53HealthCheck.Route53HealthCheckHealthCheckConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetChildHealthChecks(val *[]*string) {
	if err := j.validateSetChildHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childHealthChecks",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetEnableSni(val interface{}) {
	if err := j.validateSetEnableSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSni",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetFullyQualifiedDomainName(val *string) {
	if err := j.validateSetFullyQualifiedDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullyQualifiedDomainName",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetHealthThreshold(val *float64) {
	if err := j.validateSetHealthThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthThreshold",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetInsufficientDataHealthStatus(val *string) {
	if err := j.validateSetInsufficientDataHealthStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataHealthStatus",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetInverted(val interface{}) {
	if err := j.validateSetInvertedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inverted",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetMeasureLatency(val interface{}) {
	if err := j.validateSetMeasureLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureLatency",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetRequestInterval(val *float64) {
	if err := j.validateSetRequestIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestInterval",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetResourcePath(val *string) {
	if err := j.validateSetResourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePath",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetRoutingControlArn(val *string) {
	if err := j.validateSetRoutingControlArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingControlArn",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetSearchString(val *string) {
	if err := j.validateSetSearchStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchString",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) PutAlarmIdentifier(value *Route53HealthCheckHealthCheckConfigAlarmIdentifier) {
	if err := r.validatePutAlarmIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAlarmIdentifier",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetAlarmIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetAlarmIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetChildHealthChecks() {
	_jsii_.InvokeVoid(
		r,
		"resetChildHealthChecks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetEnableSni() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableSni",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetFullyQualifiedDomainName() {
	_jsii_.InvokeVoid(
		r,
		"resetFullyQualifiedDomainName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetHealthThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetHealthThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetInsufficientDataHealthStatus() {
	_jsii_.InvokeVoid(
		r,
		"resetInsufficientDataHealthStatus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetInverted() {
	_jsii_.InvokeVoid(
		r,
		"resetInverted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		r,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetMeasureLatency() {
	_jsii_.InvokeVoid(
		r,
		"resetMeasureLatency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		r,
		"resetRegions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetRequestInterval() {
	_jsii_.InvokeVoid(
		r,
		"resetRequestInterval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetResourcePath() {
	_jsii_.InvokeVoid(
		r,
		"resetResourcePath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetRoutingControlArn() {
	_jsii_.InvokeVoid(
		r,
		"resetRoutingControlArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ResetSearchString() {
	_jsii_.InvokeVoid(
		r,
		"resetSearchString",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53HealthCheckHealthCheckConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

