// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizedataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/personalizedataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeDatasetDatasetImportJobOutputReference interface {
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
	DatasetArn() *string
	SetDatasetArn(val *string)
	DatasetArnInput() *string
	DatasetImportJobArn() *string
	SetDatasetImportJobArn(val *string)
	DatasetImportJobArnInput() *string
	DataSource() PersonalizeDatasetDatasetImportJobDataSourceOutputReference
	DataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	PutDataSource(value *PersonalizeDatasetDatasetImportJobDataSource)
	ResetDatasetArn()
	ResetDatasetImportJobArn()
	ResetDataSource()
	ResetJobName()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersonalizeDatasetDatasetImportJobOutputReference
type jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DatasetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DatasetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DatasetImportJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetImportJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DatasetImportJobArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetImportJobArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DataSource() PersonalizeDatasetDatasetImportJobDataSourceOutputReference {
	var returns PersonalizeDatasetDatasetImportJobDataSourceOutputReference
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPersonalizeDatasetDatasetImportJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PersonalizeDatasetDatasetImportJobOutputReference {
	_init_.Initialize()

	if err := validateNewPersonalizeDatasetDatasetImportJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeDataset.PersonalizeDatasetDatasetImportJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersonalizeDatasetDatasetImportJobOutputReference_Override(p PersonalizeDatasetDatasetImportJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.personalizeDataset.PersonalizeDatasetDatasetImportJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetDatasetArn(val *string) {
	if err := j.validateSetDatasetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetArn",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetDatasetImportJobArn(val *string) {
	if err := j.validateSetDatasetImportJobArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetImportJobArn",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) PutDataSource(value *PersonalizeDatasetDatasetImportJobDataSource) {
	if err := p.validatePutDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDataSource",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ResetDatasetArn() {
	_jsii_.InvokeVoid(
		p,
		"resetDatasetArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ResetDatasetImportJobArn() {
	_jsii_.InvokeVoid(
		p,
		"resetDatasetImportJobArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		p,
		"resetDataSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ResetJobName() {
	_jsii_.InvokeVoid(
		p,
		"resetJobName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		p,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersonalizeDatasetDatasetImportJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

