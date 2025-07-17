package config

func RolePermission(userRole string, action string) bool {

	if userRole == "admin" {
		return true
	} else if userRole == "user" {
		return action == "read:book" || action == "borrow:book" || action == "return:book"
	} else {
		return false
	}

}
