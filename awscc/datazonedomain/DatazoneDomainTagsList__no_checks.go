// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datazonedomain

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneDomainTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatazoneDomainTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneDomainTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneDomainTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

