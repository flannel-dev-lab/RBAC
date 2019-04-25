INSERT INTO `rbac_operation` (name, description) VALUES 
    ('create', 'Create new records'),
    ('read', 'Read records'),
    ('update', 'Update records'),
    ('delete', 'Delete records');

INSERT INTO `rbac_object` SET `rbac_object_id` = 1,`name` = "test-static-object-name", `description` = "Reserved object for testing";

INSERT INTO `rbac_role` SET `rbac_role_id` = 1, `name` = "test-static-role-name", `description` = "Reserved role for testing";

INSERT INTO `rbac_user` SET `rbac_user_id` = 1, `name` = "test-static-user-name";

INSERT INTO `rbac_permission` SET `rbac_permission_id` = 1, `rbac_object_id` = 1, `rbac_operation_id` = 1;
