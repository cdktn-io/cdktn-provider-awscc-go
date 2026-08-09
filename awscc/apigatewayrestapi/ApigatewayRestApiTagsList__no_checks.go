// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigatewayrestapi

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigatewayRestApiTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApigatewayRestApiTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigatewayRestApiTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRestApiTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRestApiTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRestApiTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigatewayRestApiTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigatewayRestApiTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

