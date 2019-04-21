INSERT INTO `rbac_operation` (name, description) VALUES 
    ('create', 'Create new records'),
    ('read', 'Read records'),
    ('update', 'Update records'),
    ('delete', 'Delete records');

INSERT INTO `rbac_object` SET `name` = "testObject", `description` = "Reserved object for testing";

INSERT INTO `rbac_role` SET `name` = "testRole", `description` = "Reserved role for testing";

INSERT INTO `rbac_permission` SET `name` = "testPermission", `description` = "Reserved permission for testing", `rbac_object_id` = 1, `rbac_operation_id` = 1;
