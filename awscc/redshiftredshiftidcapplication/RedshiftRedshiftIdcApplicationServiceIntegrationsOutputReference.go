// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftredshiftidcapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/redshiftredshiftidcapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference interface {
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
	LakeFormation() RedshiftRedshiftIdcApplicationServiceIntegrationsLakeFormationList
	LakeFormationInput() interface{}
	Redshift() RedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList
	RedshiftInput() interface{}
	S3AccessGrants() RedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList
	S3AccessGrantsInput() interface{}
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
	PutLakeFormation(value interface{})
	PutRedshift(value interface{})
	PutS3AccessGrants(value interface{})
	ResetLakeFormation()
	ResetRedshift()
	ResetS3AccessGrants()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference
type jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) LakeFormation() RedshiftRedshiftIdcApplicationServiceIntegrationsLakeFormationList {
	var returns RedshiftRedshiftIdcApplicationServiceIntegrationsLakeFormationList
	_jsii_.Get(
		j,
		"lakeFormation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) LakeFormationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakeFormationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) Redshift() RedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList {
	var returns RedshiftRedshiftIdcApplicationServiceIntegrationsRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) RedshiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) S3AccessGrants() RedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList {
	var returns RedshiftRedshiftIdcApplicationServiceIntegrationsS3AccessGrantsList
	_jsii_.Get(
		j,
		"s3AccessGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) S3AccessGrantsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3AccessGrantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference {
	_init_.Initialize()

	if err := validateNewRedshiftRedshiftIdcApplicationServiceIntegrationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference_Override(r RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftRedshiftIdcApplication.RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) PutLakeFormation(value interface{}) {
	if err := r.validatePutLakeFormationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLakeFormation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) PutRedshift(value interface{}) {
	if err := r.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRedshift",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) PutS3AccessGrants(value interface{}) {
	if err := r.validatePutS3AccessGrantsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putS3AccessGrants",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ResetLakeFormation() {
	_jsii_.InvokeVoid(
		r,
		"resetLakeFormation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		r,
		"resetRedshift",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ResetS3AccessGrants() {
	_jsii_.InvokeVoid(
		r,
		"resetS3AccessGrants",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_RedshiftRedshiftIdcApplicationServiceIntegrationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

