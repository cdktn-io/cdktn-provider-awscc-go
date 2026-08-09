// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package forecastdataset

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_ForecastDatasetSchemaAttributesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_ForecastDatasetSchemaAttributesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_ForecastDatasetSchemaAttributesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetSchemaAttributesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetSchemaAttributesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetSchemaAttributesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetSchemaAttributesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewForecastDatasetSchemaAttributesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

