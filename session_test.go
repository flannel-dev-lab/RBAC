package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "github.com/flannel-dev-lab/RBAC/vars"
    "testing"
)


func TestCreateSession(t *testing.T) {
    setupSessionTest()
    user := database.User{Id: 1}
    session, err := sessionObject.DBService.CreateSession(user.Id, "test-session-token")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = sessionObject.DeleteSession(user.Id, session.Name)

    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownSessionTest()
}

func TestDeleteSession(t *testing.T) {
    setupSessionTest()
    user := database.User{Id: 1}
    session, err := sessionObject.DBService.CreateSession(user.Id, "test-session-token")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Delete session - this is what we are actually testing
    _, err = sessionObject.DBService.DeleteSession(user.Id, session.Name)
    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownSessionTest()
}

func TestAddActiveRole(t *testing.T) {
    user := vars.User{Id: 1}
    session := vars.Session{Id: 1, UserId: 1, Name: "123-123-123"}
    role := vars.Role{Id: 1, Name: "test-role", Description: "Reserved role for testing"}

    _, err := AddActiveRole(user, session, role.Id )

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDropActiveRole(t *testing.T) {
    user := vars.User{Id: 1}
    session := vars.Session{Id: 1, UserId: 1, Name: "123-123-123"}

    _, err := DropActiveRole(user, session, "testRole")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestCheckAccess(t *testing.T) {
    session := vars.Session{Id: 1, UserId: 1, Name: "test-session"}
    operation := Operation{Id: 1, Name: "testOperation", Description: "Reserved permission for test"}
    object := vars.Object{Id: 1, Name: "testObject", Description: "Reserved object for testing"}
    
    _, err := CheckAccess(session, operation, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}
