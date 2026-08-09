// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/applicationsignalsservicelevelobjective/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricSourceAttributes() *map[string]*string
	SetMetricSourceAttributes(val *map[string]*string)
	MetricSourceAttributesInput() *map[string]*string
	MetricSourceKeyAttributes() *map[string]*string
	SetMetricSourceKeyAttributes(val *map[string]*string)
	MetricSourceKeyAttributesInput() *map[string]*string
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
	ResetMetricSourceAttributes()
	ResetMetricSourceKeyAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference
type jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) MetricSourceAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metricSourceAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) MetricSourceAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metricSourceAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) MetricSourceKeyAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metricSourceKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) MetricSourceKeyAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metricSourceKeyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference_Override(a ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetMetricSourceAttributes(val *map[string]*string) {
	if err := j.validateSetMetricSourceAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricSourceAttributes",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetMetricSourceKeyAttributes(val *map[string]*string) {
	if err := j.validateSetMetricSourceKeyAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricSourceKeyAttributes",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ResetMetricSourceAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricSourceAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ResetMetricSourceKeyAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricSourceKeyAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

