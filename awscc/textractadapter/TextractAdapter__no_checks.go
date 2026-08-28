// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package textractadapter

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TextractAdapter) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (t *jsiiProxy_TextractAdapter) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateTextractAdapter_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateTextractAdapter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTextractAdapter_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTextractAdapter_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetAdapterNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetAutoUpdateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetFeatureTypesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_TextractAdapter) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewTextractAdapterParameters(scope constructs.Construct, id *string, config *TextractAdapterConfig) error {
	return nil
}

