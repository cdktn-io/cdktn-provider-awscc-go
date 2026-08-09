// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datasynctask/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncTaskTaskReportConfigOverridesOutputReference interface {
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
	Deleted() DatasyncTaskTaskReportConfigOverridesDeletedOutputReference
	DeletedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Skipped() DatasyncTaskTaskReportConfigOverridesSkippedOutputReference
	SkippedInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Transferred() DatasyncTaskTaskReportConfigOverridesTransferredOutputReference
	TransferredInput() interface{}
	Verified() DatasyncTaskTaskReportConfigOverridesVerifiedOutputReference
	VerifiedInput() interface{}
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
	PutDeleted(value *DatasyncTaskTaskReportConfigOverridesDeleted)
	PutSkipped(value *DatasyncTaskTaskReportConfigOverridesSkipped)
	PutTransferred(value *DatasyncTaskTaskReportConfigOverridesTransferred)
	PutVerified(value *DatasyncTaskTaskReportConfigOverridesVerified)
	ResetDeleted()
	ResetSkipped()
	ResetTransferred()
	ResetVerified()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatasyncTaskTaskReportConfigOverridesOutputReference
type jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Deleted() DatasyncTaskTaskReportConfigOverridesDeletedOutputReference {
	var returns DatasyncTaskTaskReportConfigOverridesDeletedOutputReference
	_jsii_.Get(
		j,
		"deleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) DeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Skipped() DatasyncTaskTaskReportConfigOverridesSkippedOutputReference {
	var returns DatasyncTaskTaskReportConfigOverridesSkippedOutputReference
	_jsii_.Get(
		j,
		"skipped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) SkippedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skippedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Transferred() DatasyncTaskTaskReportConfigOverridesTransferredOutputReference {
	var returns DatasyncTaskTaskReportConfigOverridesTransferredOutputReference
	_jsii_.Get(
		j,
		"transferred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) TransferredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Verified() DatasyncTaskTaskReportConfigOverridesVerifiedOutputReference {
	var returns DatasyncTaskTaskReportConfigOverridesVerifiedOutputReference
	_jsii_.Get(
		j,
		"verified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) VerifiedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedInput",
		&returns,
	)
	return returns
}


func NewDatasyncTaskTaskReportConfigOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatasyncTaskTaskReportConfigOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDatasyncTaskTaskReportConfigOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datasyncTask.DatasyncTaskTaskReportConfigOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatasyncTaskTaskReportConfigOverridesOutputReference_Override(d DatasyncTaskTaskReportConfigOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datasyncTask.DatasyncTaskTaskReportConfigOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) PutDeleted(value *DatasyncTaskTaskReportConfigOverridesDeleted) {
	if err := d.validatePutDeletedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeleted",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) PutSkipped(value *DatasyncTaskTaskReportConfigOverridesSkipped) {
	if err := d.validatePutSkippedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSkipped",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) PutTransferred(value *DatasyncTaskTaskReportConfigOverridesTransferred) {
	if err := d.validatePutTransferredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTransferred",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) PutVerified(value *DatasyncTaskTaskReportConfigOverridesVerified) {
	if err := d.validatePutVerifiedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVerified",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ResetDeleted() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ResetSkipped() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipped",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ResetTransferred() {
	_jsii_.InvokeVoid(
		d,
		"resetTransferred",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ResetVerified() {
	_jsii_.InvokeVoid(
		d,
		"resetVerified",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatasyncTaskTaskReportConfigOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

