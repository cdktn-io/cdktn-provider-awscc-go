// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lightsaildomain

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDomainTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LightsailDomainTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDomainTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDomainTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LightsailDomainTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDomainTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDomainTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDomainTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

