// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelpackage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelPackageDriftCheckBaselinesOutputReference interface {
	cdktn.ComplexObject
	Bias() SagemakerModelPackageDriftCheckBaselinesBiasOutputReference
	BiasInput() interface{}
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
	Explainability() SagemakerModelPackageDriftCheckBaselinesExplainabilityOutputReference
	ExplainabilityInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelDataQuality() SagemakerModelPackageDriftCheckBaselinesModelDataQualityOutputReference
	ModelDataQualityInput() interface{}
	ModelQuality() SagemakerModelPackageDriftCheckBaselinesModelQualityOutputReference
	ModelQualityInput() interface{}
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
	PutBias(value *SagemakerModelPackageDriftCheckBaselinesBias)
	PutExplainability(value *SagemakerModelPackageDriftCheckBaselinesExplainability)
	PutModelDataQuality(value *SagemakerModelPackageDriftCheckBaselinesModelDataQuality)
	PutModelQuality(value *SagemakerModelPackageDriftCheckBaselinesModelQuality)
	ResetBias()
	ResetExplainability()
	ResetModelDataQuality()
	ResetModelQuality()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageDriftCheckBaselinesOutputReference
type jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) Bias() SagemakerModelPackageDriftCheckBaselinesBiasOutputReference {
	var returns SagemakerModelPackageDriftCheckBaselinesBiasOutputReference
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) BiasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) Explainability() SagemakerModelPackageDriftCheckBaselinesExplainabilityOutputReference {
	var returns SagemakerModelPackageDriftCheckBaselinesExplainabilityOutputReference
	_jsii_.Get(
		j,
		"explainability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ExplainabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explainabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ModelDataQuality() SagemakerModelPackageDriftCheckBaselinesModelDataQualityOutputReference {
	var returns SagemakerModelPackageDriftCheckBaselinesModelDataQualityOutputReference
	_jsii_.Get(
		j,
		"modelDataQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ModelDataQualityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelDataQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ModelQuality() SagemakerModelPackageDriftCheckBaselinesModelQualityOutputReference {
	var returns SagemakerModelPackageDriftCheckBaselinesModelQualityOutputReference
	_jsii_.Get(
		j,
		"modelQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ModelQualityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageDriftCheckBaselinesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageDriftCheckBaselinesOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageDriftCheckBaselinesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageDriftCheckBaselinesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageDriftCheckBaselinesOutputReference_Override(s SagemakerModelPackageDriftCheckBaselinesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageDriftCheckBaselinesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) PutBias(value *SagemakerModelPackageDriftCheckBaselinesBias) {
	if err := s.validatePutBiasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBias",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) PutExplainability(value *SagemakerModelPackageDriftCheckBaselinesExplainability) {
	if err := s.validatePutExplainabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExplainability",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) PutModelDataQuality(value *SagemakerModelPackageDriftCheckBaselinesModelDataQuality) {
	if err := s.validatePutModelDataQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelDataQuality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) PutModelQuality(value *SagemakerModelPackageDriftCheckBaselinesModelQuality) {
	if err := s.validatePutModelQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelQuality",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		s,
		"resetBias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ResetExplainability() {
	_jsii_.InvokeVoid(
		s,
		"resetExplainability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ResetModelDataQuality() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataQuality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ResetModelQuality() {
	_jsii_.InvokeVoid(
		s,
		"resetModelQuality",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageDriftCheckBaselinesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

