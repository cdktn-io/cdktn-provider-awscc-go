// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSetLogicalTableMapDataTransformsOutputReference interface {
	cdktn.ComplexObject
	CastColumnTypeOperation() QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperationOutputReference
	CastColumnTypeOperationInput() interface{}
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
	CreateColumnsOperation() QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperationOutputReference
	CreateColumnsOperationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FilterOperation() QuicksightDataSetLogicalTableMapDataTransformsFilterOperationOutputReference
	FilterOperationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProjectOperation() QuicksightDataSetLogicalTableMapDataTransformsProjectOperationOutputReference
	ProjectOperationInput() interface{}
	RenameColumnOperation() QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperationOutputReference
	RenameColumnOperationInput() interface{}
	TagColumnOperation() QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperationOutputReference
	TagColumnOperationInput() interface{}
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
	PutCastColumnTypeOperation(value *QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation)
	PutCreateColumnsOperation(value *QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation)
	PutFilterOperation(value *QuicksightDataSetLogicalTableMapDataTransformsFilterOperation)
	PutProjectOperation(value *QuicksightDataSetLogicalTableMapDataTransformsProjectOperation)
	PutRenameColumnOperation(value *QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation)
	PutTagColumnOperation(value *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation)
	ResetCastColumnTypeOperation()
	ResetCreateColumnsOperation()
	ResetFilterOperation()
	ResetProjectOperation()
	ResetRenameColumnOperation()
	ResetTagColumnOperation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSetLogicalTableMapDataTransformsOutputReference
type jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) CastColumnTypeOperation() QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperationOutputReference
	_jsii_.Get(
		j,
		"castColumnTypeOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) CastColumnTypeOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"castColumnTypeOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) CreateColumnsOperation() QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperationOutputReference
	_jsii_.Get(
		j,
		"createColumnsOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) CreateColumnsOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createColumnsOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) FilterOperation() QuicksightDataSetLogicalTableMapDataTransformsFilterOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsFilterOperationOutputReference
	_jsii_.Get(
		j,
		"filterOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) FilterOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ProjectOperation() QuicksightDataSetLogicalTableMapDataTransformsProjectOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsProjectOperationOutputReference
	_jsii_.Get(
		j,
		"projectOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ProjectOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) RenameColumnOperation() QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperationOutputReference
	_jsii_.Get(
		j,
		"renameColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) RenameColumnOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) TagColumnOperation() QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperationOutputReference {
	var returns QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperationOutputReference
	_jsii_.Get(
		j,
		"tagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) TagColumnOperationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSetLogicalTableMapDataTransformsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightDataSetLogicalTableMapDataTransformsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetLogicalTableMapDataTransformsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSetLogicalTableMapDataTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightDataSetLogicalTableMapDataTransformsOutputReference_Override(q QuicksightDataSetLogicalTableMapDataTransformsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSetLogicalTableMapDataTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutCastColumnTypeOperation(value *QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation) {
	if err := q.validatePutCastColumnTypeOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCastColumnTypeOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutCreateColumnsOperation(value *QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation) {
	if err := q.validatePutCreateColumnsOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCreateColumnsOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutFilterOperation(value *QuicksightDataSetLogicalTableMapDataTransformsFilterOperation) {
	if err := q.validatePutFilterOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFilterOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutProjectOperation(value *QuicksightDataSetLogicalTableMapDataTransformsProjectOperation) {
	if err := q.validatePutProjectOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putProjectOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutRenameColumnOperation(value *QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation) {
	if err := q.validatePutRenameColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRenameColumnOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) PutTagColumnOperation(value *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation) {
	if err := q.validatePutTagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTagColumnOperation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetCastColumnTypeOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetCastColumnTypeOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetCreateColumnsOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateColumnsOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetFilterOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetProjectOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetProjectOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetRenameColumnOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetRenameColumnOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ResetTagColumnOperation() {
	_jsii_.InvokeVoid(
		q,
		"resetTagColumnOperation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapDataTransformsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

