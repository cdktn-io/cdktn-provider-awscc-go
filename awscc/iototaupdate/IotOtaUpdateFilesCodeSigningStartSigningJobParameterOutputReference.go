// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/iototaupdate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference interface {
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
	Destination() IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestinationOutputReference
	DestinationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SigningProfileName() *string
	SetSigningProfileName(val *string)
	SigningProfileNameInput() *string
	SigningProfileParameter() IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference
	SigningProfileParameterInput() interface{}
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
	PutDestination(value *IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestination)
	PutSigningProfileParameter(value *IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameter)
	ResetDestination()
	ResetSigningProfileName()
	ResetSigningProfileParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference
type jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) Destination() IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestinationOutputReference {
	var returns IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) SigningProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) SigningProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) SigningProfileParameter() IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference {
	var returns IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameterOutputReference
	_jsii_.Get(
		j,
		"signingProfileParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) SigningProfileParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signingProfileParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference {
	_init_.Initialize()

	if err := validateNewIotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference_Override(i IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotOtaUpdate.IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetSigningProfileName(val *string) {
	if err := j.validateSetSigningProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingProfileName",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) PutDestination(value *IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestination) {
	if err := i.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDestination",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) PutSigningProfileParameter(value *IotOtaUpdateFilesCodeSigningStartSigningJobParameterSigningProfileParameter) {
	if err := i.validatePutSigningProfileParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSigningProfileParameter",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		i,
		"resetDestination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ResetSigningProfileName() {
	_jsii_.InvokeVoid(
		i,
		"resetSigningProfileName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ResetSigningProfileParameter() {
	_jsii_.InvokeVoid(
		i,
		"resetSigningProfileParameter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotOtaUpdateFilesCodeSigningStartSigningJobParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

