package middleware

type Roles map[string][]string

func AuthMiddleware(role string, method string) bool {
	roles := rules()
	perms := roles[role]
	for _, perm := range perms {
		if perm == method {
			return true
		}
	}
	return false
}

func rules() Roles {
	roles := make(Roles)
	roles["Админ"] = []string{"*"}
	roles["Пользователь"] = []string{"all", "check", "myid", "whoami"}
	return roles
}
