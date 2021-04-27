package v1alpha1

import meta "github.com/fdymylja/tmos/module/meta/v1alpha1"

func (m *CurrentAccountNumber) GetObjectMeta() *meta.ObjectMeta {
	return &meta.ObjectMeta{Id: "account_number"}
}
