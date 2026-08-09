// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneprojectprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneprojectprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneProjectProfileEnvironmentConfigurationsOutputReference interface {
	cdktn.ComplexObject
	AwsAccount() DatazoneProjectProfileEnvironmentConfigurationsAwsAccountOutputReference
	AwsAccountInput() interface{}
	AwsRegion() DatazoneProjectProfileEnvironmentConfigurationsAwsRegionOutputReference
	AwsRegionInput() interface{}
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
	ConfigurationParameters() DatazoneProjectProfileEnvironmentConfigurationsConfigurationParametersOutputReference
	ConfigurationParametersInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentMode() *string
	SetDeploymentMode(val *string)
	DeploymentModeInput() *string
	DeploymentOrder() *float64
	SetDeploymentOrder(val *float64)
	DeploymentOrderInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentBlueprintId() *string
	SetEnvironmentBlueprintId(val *string)
	EnvironmentBlueprintIdInput() *string
	EnvironmentConfigurationId() *string
	SetEnvironmentConfigurationId(val *string)
	EnvironmentConfigurationIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutAwsAccount(value *DatazoneProjectProfileEnvironmentConfigurationsAwsAccount)
	PutAwsRegion(value *DatazoneProjectProfileEnvironmentConfigurationsAwsRegion)
	PutConfigurationParameters(value *DatazoneProjectProfileEnvironmentConfigurationsConfigurationParameters)
	ResetAwsAccount()
	ResetAwsRegion()
	ResetConfigurationParameters()
	ResetDeploymentMode()
	ResetDeploymentOrder()
	ResetDescription()
	ResetEnvironmentBlueprintId()
	ResetEnvironmentConfigurationId()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneProjectProfileEnvironmentConfigurationsOutputReference
type jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) AwsAccount() DatazoneProjectProfileEnvironmentConfigurationsAwsAccountOutputReference {
	var returns DatazoneProjectProfileEnvironmentConfigurationsAwsAccountOutputReference
	_jsii_.Get(
		j,
		"awsAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) AwsAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) AwsRegion() DatazoneProjectProfileEnvironmentConfigurationsAwsRegionOutputReference {
	var returns DatazoneProjectProfileEnvironmentConfigurationsAwsRegionOutputReference
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) AwsRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ConfigurationParameters() DatazoneProjectProfileEnvironmentConfigurationsConfigurationParametersOutputReference {
	var returns DatazoneProjectProfileEnvironmentConfigurationsConfigurationParametersOutputReference
	_jsii_.Get(
		j,
		"configurationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ConfigurationParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) DeploymentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) DeploymentOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) DeploymentOrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) EnvironmentBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) EnvironmentBlueprintIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentBlueprintIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) EnvironmentConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) EnvironmentConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazoneProjectProfileEnvironmentConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatazoneProjectProfileEnvironmentConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneProjectProfileEnvironmentConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneProjectProfile.DatazoneProjectProfileEnvironmentConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatazoneProjectProfileEnvironmentConfigurationsOutputReference_Override(d DatazoneProjectProfileEnvironmentConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneProjectProfile.DatazoneProjectProfileEnvironmentConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetDeploymentMode(val *string) {
	if err := j.validateSetDeploymentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetDeploymentOrder(val *float64) {
	if err := j.validateSetDeploymentOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentOrder",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetEnvironmentBlueprintId(val *string) {
	if err := j.validateSetEnvironmentBlueprintIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentBlueprintId",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetEnvironmentConfigurationId(val *string) {
	if err := j.validateSetEnvironmentConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentConfigurationId",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) PutAwsAccount(value *DatazoneProjectProfileEnvironmentConfigurationsAwsAccount) {
	if err := d.validatePutAwsAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAccount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) PutAwsRegion(value *DatazoneProjectProfileEnvironmentConfigurationsAwsRegion) {
	if err := d.validatePutAwsRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsRegion",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) PutConfigurationParameters(value *DatazoneProjectProfileEnvironmentConfigurationsConfigurationParameters) {
	if err := d.validatePutConfigurationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigurationParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetAwsAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetConfigurationParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigurationParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetDeploymentMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDeploymentMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetDeploymentOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetDeploymentOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetEnvironmentBlueprintId() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentBlueprintId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetEnvironmentConfigurationId() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironmentConfigurationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneProjectProfileEnvironmentConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

