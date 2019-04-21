INSERT INTO `rbac_operation` (name, description) VALUES 
    ('create', 'Create new records'),
    ('read', 'Read records'),
    ('update', 'Update records'),
    ('delete', 'Delete records');

INSERT INTO `rbac_object` SET `name` = "test-static-object-name", `description` = "Reserved object for testing";

INSERT INTO `rbac_role` SET `name` = "test-static-role-name", `description` = "Reserved role for testing";

INSERT INTO `rbac_permission` SET `name` = "test-static-permission-name", `description` = "Reserved permission for testing", `rbac_object_id` = 1, `rbac_operation_id` = 1;
