package vars

// User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
	Id   int    `json:"rbac_user_id"` // should come from the underlying system
	Name string `json:"name"`         // this might need to be removed for target system
}

// Role is a job function within the context of an organization
type Role struct {
	Id          int    `json:"rbac_role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Session represents a user as owner and an active role set
type Session struct {
	Id     int    `json:"rbac_session_id"`
	Name   string `json:"name"`
	UserId int    `json:"rbac_user_id"`
}

// Object can be any system resource subject to access control
type Object struct {
	Id          int    `json:"rbac_object_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Operation is a method, which upon invocation exexutes some function for the user
type Operation struct {
	Id          int    `json:"rbac_operation_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Permission is a a combination of object & operation that can be enforced
type Permission struct {
	Id            int    `json:"rbac_permission_id"`
	ObjectName    string `json:"rbac_object_name"`
	OperationName string `json:"rbac_operation_name"`
}
