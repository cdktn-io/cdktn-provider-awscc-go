// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsecurityprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotSecurityProfileBehaviorsCriteriaOutputReference interface {
	cdktn.ComplexObject
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	ComparisonOperatorInput() *string
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
	ConsecutiveDatapointsToAlarm() *float64
	SetConsecutiveDatapointsToAlarm(val *float64)
	ConsecutiveDatapointsToAlarmInput() *float64
	ConsecutiveDatapointsToClear() *float64
	SetConsecutiveDatapointsToClear(val *float64)
	ConsecutiveDatapointsToClearInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DurationSeconds() *float64
	SetDurationSeconds(val *float64)
	DurationSecondsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MlDetectionConfig() IotSecurityProfileBehaviorsCriteriaMlDetectionConfigOutputReference
	MlDetectionConfigInput() interface{}
	StatisticalThreshold() IotSecurityProfileBehaviorsCriteriaStatisticalThresholdOutputReference
	StatisticalThresholdInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() IotSecurityProfileBehaviorsCriteriaValueOutputReference
	ValueInput() interface{}
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
	PutMlDetectionConfig(value *IotSecurityProfileBehaviorsCriteriaMlDetectionConfig)
	PutStatisticalThreshold(value *IotSecurityProfileBehaviorsCriteriaStatisticalThreshold)
	PutValue(value *IotSecurityProfileBehaviorsCriteriaValue)
	ResetComparisonOperator()
	ResetConsecutiveDatapointsToAlarm()
	ResetConsecutiveDatapointsToClear()
	ResetDurationSeconds()
	ResetMlDetectionConfig()
	ResetStatisticalThreshold()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecurityProfileBehaviorsCriteriaOutputReference
type jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ConsecutiveDatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveDatapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ConsecutiveDatapointsToAlarmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveDatapointsToAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ConsecutiveDatapointsToClear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveDatapointsToClear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ConsecutiveDatapointsToClearInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveDatapointsToClearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) DurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) DurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) MlDetectionConfig() IotSecurityProfileBehaviorsCriteriaMlDetectionConfigOutputReference {
	var returns IotSecurityProfileBehaviorsCriteriaMlDetectionConfigOutputReference
	_jsii_.Get(
		j,
		"mlDetectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) MlDetectionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlDetectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) StatisticalThreshold() IotSecurityProfileBehaviorsCriteriaStatisticalThresholdOutputReference {
	var returns IotSecurityProfileBehaviorsCriteriaStatisticalThresholdOutputReference
	_jsii_.Get(
		j,
		"statisticalThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) StatisticalThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statisticalThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) Value() IotSecurityProfileBehaviorsCriteriaValueOutputReference {
	var returns IotSecurityProfileBehaviorsCriteriaValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewIotSecurityProfileBehaviorsCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotSecurityProfileBehaviorsCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecurityProfileBehaviorsCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotSecurityProfile.IotSecurityProfileBehaviorsCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecurityProfileBehaviorsCriteriaOutputReference_Override(i IotSecurityProfileBehaviorsCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotSecurityProfile.IotSecurityProfileBehaviorsCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetConsecutiveDatapointsToAlarm(val *float64) {
	if err := j.validateSetConsecutiveDatapointsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveDatapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetConsecutiveDatapointsToClear(val *float64) {
	if err := j.validateSetConsecutiveDatapointsToClearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveDatapointsToClear",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetDurationSeconds(val *float64) {
	if err := j.validateSetDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationSeconds",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) PutMlDetectionConfig(value *IotSecurityProfileBehaviorsCriteriaMlDetectionConfig) {
	if err := i.validatePutMlDetectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMlDetectionConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) PutStatisticalThreshold(value *IotSecurityProfileBehaviorsCriteriaStatisticalThreshold) {
	if err := i.validatePutStatisticalThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putStatisticalThreshold",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) PutValue(value *IotSecurityProfileBehaviorsCriteriaValue) {
	if err := i.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetComparisonOperator() {
	_jsii_.InvokeVoid(
		i,
		"resetComparisonOperator",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetConsecutiveDatapointsToAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetConsecutiveDatapointsToAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetConsecutiveDatapointsToClear() {
	_jsii_.InvokeVoid(
		i,
		"resetConsecutiveDatapointsToClear",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetDurationSeconds() {
	_jsii_.InvokeVoid(
		i,
		"resetDurationSeconds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetMlDetectionConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetMlDetectionConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetStatisticalThreshold() {
	_jsii_.InvokeVoid(
		i,
		"resetStatisticalThreshold",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		i,
		"resetValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

