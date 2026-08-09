// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccinspectorv2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccinspectorv2filter/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference interface {
	cdktn.ComplexObject
	Architecture() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference
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
	Epoch() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference
	FilePath() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesFilePathOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackages
	SetInternalValue(val *DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackages)
	Name() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference
	Release() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference
	SourceLambdaLayerArn() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference
	SourceLayerHash() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Version() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference
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

// The jsii proxy struct for DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference
type jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Architecture() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesArchitectureOutputReference
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Epoch() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesEpochOutputReference
	_jsii_.Get(
		j,
		"epoch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) FilePath() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesFilePathOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesFilePathOutputReference
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InternalValue() *DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackages {
	var returns *DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackages
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Name() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Release() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesReleaseOutputReference
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLambdaLayerArn() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnOutputReference
	_jsii_.Get(
		j,
		"sourceLambdaLayerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLayerHash() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHashOutputReference
	_jsii_.Get(
		j,
		"sourceLayerHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Version() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesVersionOutputReference
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewDataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInspectorv2Filter.DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference_Override(d DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInspectorv2Filter.DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetInternalValue(val *DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackages) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

