package schema

import "fmt"

// Validate verifies a Definition
func (x *Definition) Validate() error {
	if x.Singleton && x.PrimaryKey != "" {
		return fmt.Errorf("a StateObject can not be singleton and have a primary key at the same time")
	}
	if x.Singleton && len(x.SecondaryKeys) != 0 {
		return fmt.Errorf("a StateObject can not be singleton and have secondary keys at the same time")
	}
	if !x.Singleton && x.PrimaryKey == "" {
		return fmt.Errorf("a StateObject must be a singleton or have a primary key")
	}
	return nil
}
