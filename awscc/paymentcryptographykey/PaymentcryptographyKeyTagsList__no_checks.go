// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package paymentcryptographykey

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PaymentcryptographyKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PaymentcryptographyKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PaymentcryptographyKeyTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPaymentcryptographyKeyTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

