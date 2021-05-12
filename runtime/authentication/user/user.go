package user

import "fmt"

// User uniquely identifies a user in the authentication context
type User interface {
	// GetName returns the unique name of the user in the system
	GetName() string
	// GetExtra can be used by custom authenticators/authorizers combos
	// to attach data to User
	GetExtra() map[string][]string
	fmt.Stringer
}

// Users defines multiple User instances
type Users interface {
	Has(names ...string) bool
	List() []User
	fmt.Stringer
}

type DefaultUser string

func (d DefaultUser) GetName() string               { return (string)(d) }
func (d DefaultUser) GetExtra() map[string][]string { return nil }
func (d DefaultUser) String() string                { return (string)(d) }

type DefaultUsers struct {
	users map[string]User
}

func (d *DefaultUsers) add(u User) bool {
	_, exists := d.users[u.GetName()]
	if exists {
		return false
	}
	d.users[u.GetName()] = u
	return true
}

func (d *DefaultUsers) Has(names ...string) bool {
	for _, name := range names {
		if !d.has(name) {
			return false
		}
	}
	return true
}

func (d *DefaultUsers) has(name string) bool {
	_, exists := d.users[name]
	return exists
}

func (d *DefaultUsers) List() []User {
	u := make([]User, 0, len(d.users))
	for _, user := range d.users {
		u = append(u, user)
	}
	return u
}

func (d *DefaultUsers) String() string {
	uStr := make([]string, 0, len(d.users))
	for _, u := range d.users {
		uStr = append(uStr, u.String())
	}
	return fmt.Sprintf("%s", uStr)
}

func NewUsersFromString(users ...string) Users {
	u := make(map[string]User, len(users))
	for _, user := range users {
		u[user] = DefaultUser(user)
	}
	return &DefaultUsers{users: u}
}

func NewUsers(users ...User) Users {
	u := make(map[string]User, len(users))
	for _, user := range users {
		u[user.GetName()] = user
	}
	return &DefaultUsers{users: u}
}

func NewUsersUnion(groups ...Users) Users {
	users := &DefaultUsers{users: map[string]User{}}
	for _, g := range groups {
		// in case of nil users we skip
		if g == nil {
			continue
		}
		for _, user := range g.List() {
			users.add(user)
		}
	}
	return users
}
