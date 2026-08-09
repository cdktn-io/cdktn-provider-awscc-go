// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package forecastdataset

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_ForecastDatasetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_ForecastDatasetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_ForecastDatasetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ForecastDatasetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewForecastDatasetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

