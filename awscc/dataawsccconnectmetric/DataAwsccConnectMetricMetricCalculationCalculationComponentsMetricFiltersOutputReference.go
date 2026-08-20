// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccconnectmetric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccconnectmetric/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference interface {
	cdktn.ComplexObject
	BooleanCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFilters
	SetInternalValue(val *DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFilters)
	MetricFilterKey() *string
	Negate() cdktn.IResolvable
	NumberCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference
	StringCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference
type jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) BooleanCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference {
	var returns DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersBooleanConditionOutputReference
	_jsii_.Get(
		j,
		"booleanCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InternalValue() *DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFilters {
	var returns *DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) MetricFilterKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricFilterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Negate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) NumberCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference {
	var returns DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersNumberConditionOutputReference
	_jsii_.Get(
		j,
		"numberCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) StringCondition() DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference {
	var returns DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringConditionOutputReference
	_jsii_.Get(
		j,
		"stringCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccConnectMetric.DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference_Override(d DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccConnectMetric.DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetInternalValue(val *DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectMetricMetricCalculationCalculationComponentsMetricFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

