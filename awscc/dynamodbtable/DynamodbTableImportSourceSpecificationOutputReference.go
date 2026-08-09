// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbTableImportSourceSpecificationOutputReference interface {
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
	InputCompressionType() *string
	SetInputCompressionType(val *string)
	InputCompressionTypeInput() *string
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InputFormatOptions() DynamodbTableImportSourceSpecificationInputFormatOptionsOutputReference
	InputFormatOptionsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3BucketSource() DynamodbTableImportSourceSpecificationS3BucketSourceOutputReference
	S3BucketSourceInput() interface{}
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
	PutInputFormatOptions(value *DynamodbTableImportSourceSpecificationInputFormatOptions)
	PutS3BucketSource(value *DynamodbTableImportSourceSpecificationS3BucketSource)
	ResetInputCompressionType()
	ResetInputFormat()
	ResetInputFormatOptions()
	ResetS3BucketSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbTableImportSourceSpecificationOutputReference
type jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputFormatOptions() DynamodbTableImportSourceSpecificationInputFormatOptionsOutputReference {
	var returns DynamodbTableImportSourceSpecificationInputFormatOptionsOutputReference
	_jsii_.Get(
		j,
		"inputFormatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InputFormatOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputFormatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) S3BucketSource() DynamodbTableImportSourceSpecificationS3BucketSourceOutputReference {
	var returns DynamodbTableImportSourceSpecificationS3BucketSourceOutputReference
	_jsii_.Get(
		j,
		"s3BucketSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) S3BucketSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDynamodbTableImportSourceSpecificationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DynamodbTableImportSourceSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbTableImportSourceSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableImportSourceSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDynamodbTableImportSourceSpecificationOutputReference_Override(d DynamodbTableImportSourceSpecificationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableImportSourceSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetInputCompressionType(val *string) {
	if err := j.validateSetInputCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputCompressionType",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) PutInputFormatOptions(value *DynamodbTableImportSourceSpecificationInputFormatOptions) {
	if err := d.validatePutInputFormatOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInputFormatOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) PutS3BucketSource(value *DynamodbTableImportSourceSpecificationS3BucketSource) {
	if err := d.validatePutS3BucketSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3BucketSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ResetInputCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetInputCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ResetInputFormatOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetInputFormatOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ResetS3BucketSource() {
	_jsii_.InvokeVoid(
		d,
		"resetS3BucketSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbTableImportSourceSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

