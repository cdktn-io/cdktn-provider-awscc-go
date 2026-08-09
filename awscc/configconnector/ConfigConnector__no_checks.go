// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configconnector

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigConnector) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validatePutConnectorConfigurationParameters(value *ConfigConnectorConnectorConfiguration) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigConnector) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateConfigConnector_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateConfigConnector_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigConnector_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateConfigConnector_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConnector) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConnector) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConnector) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ConfigConnector) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewConfigConnectorParameters(scope constructs.Construct, id *string, config *ConfigConnectorConfig) error {
	return nil
}

