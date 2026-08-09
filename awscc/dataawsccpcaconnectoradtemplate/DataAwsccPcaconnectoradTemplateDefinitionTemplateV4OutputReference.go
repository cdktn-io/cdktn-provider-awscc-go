// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccpcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccpcaconnectoradtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference interface {
	cdktn.ComplexObject
	CertificateValidity() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference
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
	EnrollmentFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference
	Extensions() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference
	// Experimental.
	Fqn() *string
	GeneralFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference
	HashAlgorithm() *string
	InternalValue() *DataAwsccPcaconnectoradTemplateDefinitionTemplateV4
	SetInternalValue(val *DataAwsccPcaconnectoradTemplateDefinitionTemplateV4)
	PrivateKeyAttributes() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference
	PrivateKeyFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference
	SubjectNameFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference
	SupersededTemplates() *[]*string
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

// The jsii proxy struct for DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference
type jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) CertificateValidity() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4CertificateValidityOutputReference
	_jsii_.Get(
		j,
		"certificateValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) EnrollmentFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4EnrollmentFlagsOutputReference
	_jsii_.Get(
		j,
		"enrollmentFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) Extensions() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4ExtensionsOutputReference
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GeneralFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4GeneralFlagsOutputReference
	_jsii_.Get(
		j,
		"generalFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) InternalValue() *DataAwsccPcaconnectoradTemplateDefinitionTemplateV4 {
	var returns *DataAwsccPcaconnectoradTemplateDefinitionTemplateV4
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyAttributes() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesOutputReference
	_jsii_.Get(
		j,
		"privateKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) PrivateKeyFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4PrivateKeyFlagsOutputReference
	_jsii_.Get(
		j,
		"privateKeyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) SubjectNameFlags() DataAwsccPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference {
	var returns DataAwsccPcaconnectoradTemplateDefinitionTemplateV4SubjectNameFlagsOutputReference
	_jsii_.Get(
		j,
		"subjectNameFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) SupersededTemplates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccPcaconnectoradTemplate.DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference_Override(d DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccPcaconnectoradTemplate.DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetInternalValue(val *DataAwsccPcaconnectoradTemplateDefinitionTemplateV4) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccPcaconnectoradTemplateDefinitionTemplateV4OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

