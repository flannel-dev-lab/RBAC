package rbac

/*import (
    "github.com/flannel-dev-lab/RBAC/database"
    "github.com/flannel-dev-lab/RBAC/vars"
    "testing"
)


var (
    TestRoleStatic = vars.Role{Id: 1, Name: "test-fluid-role-name", Description: "Reserved role for fluid testing"}
)




func TestAddRole(t *testing.T) {
    setupRoleTest()
    role, err := roleObject.AddRole("test-role", "Test role description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = roleObject.DeleteRole(role.Id)
    tearDownRoleTest()
}

func TestDeleteRole(t *testing.T) {
    setupRoleTest()
    role, err := roleObject.AddRole("test-role", "Test role description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = roleObject.DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownRoleTest()
}

func TestAssignUser(t *testing.T) {
    setupRoleTest()
    setupUserTest()
    // Create a new user
    user, userErr := userObject.AddUser("test-user")
    if userErr != nil {
        t.Errorf("%v", userErr)
    }

    // Create a new role
    role, roleErr := roleObject.AddRole("test-role", "Test role description")
    if roleErr != nil {
        t.Errorf("%v", roleErr)
    }

    // Assign the user - what we are actually testing
    _, err := roleObject.AssignUser(user.Id, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = roleObject.DeassignUser(user.Id, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = userObject.DeleteUser(user.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = roleObject.DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownRoleTest()
    tearDownUserTest()
}

func TestDeassignUser(t *testing.T) {
    setupRoleTest()
    setupUserTest()
    // Create a new user
    user, userErr := userObject.AddUser("test-user")
    if userErr != nil {
        t.Errorf("%v", userErr)
    }

    // Create a new role
    role, roleErr := roleObject.AddRole("test-role", "Test role description")
    if roleErr != nil {
        t.Errorf("%v", roleErr)
    }

    // Assign the user
    _, err := roleObject.AssignUser(user.Id, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Deassign the user - what we are actually testing
    _, err = roleObject.DeassignUser(user.Id, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = userObject.DeleteUser(user.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = roleObject.DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownUserTest()
    tearDownRoleTest()
}

func TestAssignedUsers(t *testing.T) {
    setupRoleTest()
    role := database.Role{Id: 1, Name: "roleName", Description:"Reserved role for testing"}
    _, err := roleObject.AssignedUsers(role.Id)

    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownRoleTest()
}

func TestSessionRoles(t *testing.T) {
    session := vars.Session{Id: 1, UserId: 1, Name: "123-123-123"}
    _, err := SessionRoles(session);

    if err != nil {
        t.Errorf("%v", err);
    }
}
*/
