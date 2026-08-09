// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermodelpackage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelPackageModelMetricsBiasOutputReference interface {
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
	PostTrainingReport() SagemakerModelPackageModelMetricsBiasPostTrainingReportOutputReference
	PostTrainingReportInput() interface{}
	PreTrainingReport() SagemakerModelPackageModelMetricsBiasPreTrainingReportOutputReference
	PreTrainingReportInput() interface{}
	Report() SagemakerModelPackageModelMetricsBiasReportOutputReference
	ReportInput() interface{}
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
	PutPostTrainingReport(value *SagemakerModelPackageModelMetricsBiasPostTrainingReport)
	PutPreTrainingReport(value *SagemakerModelPackageModelMetricsBiasPreTrainingReport)
	PutReport(value *SagemakerModelPackageModelMetricsBiasReport)
	ResetPostTrainingReport()
	ResetPreTrainingReport()
	ResetReport()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageModelMetricsBiasOutputReference
type jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PostTrainingReport() SagemakerModelPackageModelMetricsBiasPostTrainingReportOutputReference {
	var returns SagemakerModelPackageModelMetricsBiasPostTrainingReportOutputReference
	_jsii_.Get(
		j,
		"postTrainingReport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PostTrainingReportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postTrainingReportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PreTrainingReport() SagemakerModelPackageModelMetricsBiasPreTrainingReportOutputReference {
	var returns SagemakerModelPackageModelMetricsBiasPreTrainingReportOutputReference
	_jsii_.Get(
		j,
		"preTrainingReport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PreTrainingReportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preTrainingReportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) Report() SagemakerModelPackageModelMetricsBiasReportOutputReference {
	var returns SagemakerModelPackageModelMetricsBiasReportOutputReference
	_jsii_.Get(
		j,
		"report",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ReportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageModelMetricsBiasOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageModelMetricsBiasOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageModelMetricsBiasOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageModelMetricsBiasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageModelMetricsBiasOutputReference_Override(s SagemakerModelPackageModelMetricsBiasOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerModelPackage.SagemakerModelPackageModelMetricsBiasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PutPostTrainingReport(value *SagemakerModelPackageModelMetricsBiasPostTrainingReport) {
	if err := s.validatePutPostTrainingReportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPostTrainingReport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PutPreTrainingReport(value *SagemakerModelPackageModelMetricsBiasPreTrainingReport) {
	if err := s.validatePutPreTrainingReportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPreTrainingReport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) PutReport(value *SagemakerModelPackageModelMetricsBiasReport) {
	if err := s.validatePutReportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ResetPostTrainingReport() {
	_jsii_.InvokeVoid(
		s,
		"resetPostTrainingReport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ResetPreTrainingReport() {
	_jsii_.InvokeVoid(
		s,
		"resetPreTrainingReport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ResetReport() {
	_jsii_.InvokeVoid(
		s,
		"resetReport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageModelMetricsBiasOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

