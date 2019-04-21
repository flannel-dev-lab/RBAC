CREATE TABLE `rbac_role` (
  `rbac_role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(3000) NOT NULL,
  PRIMARY KEY (`rbac_role_id`),
  UNIQUE `UK_NAME` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_user` (
  `rbac_user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`rbac_user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_object` (
  `rbac_object_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(3000) NOT NULL,
  PRIMARY KEY (`rbac_object_id`),
  UNIQUE `UK_NAME` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_operation` (
  `rbac_operation_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(3000) NOT NULL,
  PRIMARY KEY (`rbac_operation_id`),
  UNIQUE `UK_NAME` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_permission` (
  `rbac_permission_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_object_id` int(11) unsigned NOT NULL,
  `rbac_operation_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_permission_id`),
  UNIQUE `UK_OBJECT_ID_OPERATION_ID` (`rbac_object_id`, `rbac_operation_id`),
  FOREIGN KEY `FK_RBAC_PERMISSION_RBAC_OBJECT_ID` (`rbac_object_id`)
      REFERENCES `rbac_object` (`rbac_object_id`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_PERMISSION_RBAC_OPERATION_ID` (`rbac_operation_id`)
      REFERENCES `rbac_operation` (`rbac_operation_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_user_role` (
  `rbac_user_role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_user_id` int(11) unsigned NOT NULL,
  `rbac_role_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_user_role_id`),
  UNIQUE `UK_USER_ID_ROLE_ID` (`rbac_user_id`, `rbac_role_id`),
  FOREIGN KEY `FK_RBAC_USER_ROLE_RBAC_ROLE_ID` (`rbac_role_id`) 
      REFERENCES `rbac_role` (`rbac_role_id`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_USER_ROLE_RBAC_USER_ID` (`rbac_user_id`) 
      REFERENCES `rbac_user` (`rbac_user_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_role_permission` (
  `rbac_role_permission_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_role_id` int(11) unsigned NOT NULL,
  `rbac_permission_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_role_permission_id`),
  UNIQUE `UK_ROLE_ID_PERMISSION_ID` (`rbac_role_id`, `rbac_permission_id`),
  FOREIGN KEY `FK_RBAC_ROLE_PEMISSION_RBAC_ROLE_ID` (`rbac_role_id`) 
      REFERENCES `rbac_role` (`rbac_role_id`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_ROLE_PERMISSION_RBAC_PERMISSION_ID` (`rbac_permission_id`) 
      REFERENCES `rbac_permission` (`rbac_permission_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_role_role` (
  `rbac_role_role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `left_role_id` int(11) unsigned NOT NULL,
  `right_role_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_role_role_id`),
  UNIQUE `UK_ROLE_ID_ROLE_ID` (`left_role_id`, `right_role_id`),
  FOREIGN KEY `FK_RBAC_ROLE_ROLE_LEFT_ROLE_ID` (`left_role_id`) 
      REFERENCES `rbac_role` (`rbac_role_id`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_ROLE_ROLE_RIGHT_ROLE_ID` (`right_role_id`) 
      REFERENCES `rbac_role` (`rbac_role_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;



