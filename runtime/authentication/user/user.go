package user

// User uniquely identifies a user in the authentication context
type User interface {
	// GetName returns the unique name of the user in the system
	GetName() string
	// GetExtra can be used by custom authenticators/authorizers combos
	// to attach data to User
	GetExtra() map[string][]string
}

// Users defines multiple User instances
type Users interface {
	Has(names ...string) bool
	List() []User
}

type DefaultUser string

func (d DefaultUser) GetName() string               { return (string)(d) }
func (d DefaultUser) GetExtra() map[string][]string { return nil }

type DefaultUsers struct {
	users map[string]User
}

func (d DefaultUsers) add(u User) bool {
	_, exists := d.users[u.GetName()]
	if exists {
		return false
	}
	d.users[u.GetName()] = u
	return true
}

func (d DefaultUsers) Has(names ...string) bool {
	for _, name := range names {
		if !d.has(name) {
			return false
		}
	}
	return true
}

func (d DefaultUsers) has(name string) bool {
	_, exists := d.users[name]
	return exists
}

func (d DefaultUsers) List() []User {
	panic("implement me")
}

func NewUsersFromString(users ...string) Users {
	u := make(map[string]User, len(users))
	for _, user := range users {
		u[user] = DefaultUser(user)
	}
	return DefaultUsers{users: u}
}

func NewUsers(users ...User) Users {
	u := make(map[string]User, len(users))
	for _, user := range users {
		u[user.GetName()] = user
	}
	return DefaultUsers{users: u}
}

func NewUsersUnion(groups ...Users) Users {
	users := DefaultUsers{users: map[string]User{}}
	for _, g := range groups {
		for _, user := range g.List() {
			users.add(user)
		}
	}
	return users
}
