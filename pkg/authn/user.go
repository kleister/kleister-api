package authn

// User defines a model filles by the authentication provider.
type User struct {
	Ident string
	Login string
	Name  string
	Email string
	Roles []string
	Raw   map[string]interface{}
}
