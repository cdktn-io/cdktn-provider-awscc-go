// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisevehicle

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotfleetwisevehicle/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference interface {
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
	OnChange() *string
	SetOnChange(val *string)
	OnChangeInput() *string
	Periodic() IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodicOutputReference
	PeriodicInput() interface{}
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
	PutPeriodic(value *IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodic)
	ResetOnChange()
	ResetPeriodic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference
type jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) OnChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) OnChangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) Periodic() IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodicOutputReference {
	var returns IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodicOutputReference
	_jsii_.Get(
		j,
		"periodic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) PeriodicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"periodicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewIotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseVehicle.IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference_Override(i IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseVehicle.IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetOnChange(val *string) {
	if err := j.validateSetOnChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onChange",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) PutPeriodic(value *IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodic) {
	if err := i.validatePutPeriodicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPeriodic",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ResetOnChange() {
	_jsii_.InvokeVoid(
		i,
		"resetOnChange",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ResetPeriodic() {
	_jsii_.InvokeVoid(
		i,
		"resetPeriodic",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

