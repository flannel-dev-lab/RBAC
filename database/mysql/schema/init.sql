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


CREATE TABLE `rbac_session` (
  `rbac_session_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `rbac_user_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_session_id`),
  UNIQUE `UK_NAME` (`name`),
  FOREIGN KEY `FK_RBAC_USER_ID` (`rbac_user_id`)
      REFERENCES `rbac_user` (`rbac_user_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_permission` (
  `rbac_permission_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_object_name` varchar(255) NOT NULL,
  `rbac_operation_name` varchar(255) NOT NULL,
  PRIMARY KEY (`rbac_permission_id`),
  UNIQUE `UK_OBJECT_NAME_OPERATION_NAME` (`rbac_object_name`, `rbac_operation_name`),
  FOREIGN KEY `FK_RBAC_PERMISSION_RBAC_OBJECT_NAME` (`rbac_object_name`)
      REFERENCES `rbac_object` (`name`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_PERMISSION_RBAC_OPERATION_NAME` (`rbac_operation_name`)
      REFERENCES `rbac_operation` (`name`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_user_role` (
  `rbac_user_role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_user_id` int(11) unsigned NOT NULL,
  `rbac_role_name` varchar(255) NOT NULL,
  PRIMARY KEY (`rbac_user_role_id`),
  UNIQUE `UK_USER_ID_ROLE_NAME` (`rbac_user_id`, `rbac_role_name`),
  FOREIGN KEY `FK_RBAC_USER_ROLE_RBAC_ROLE_NAME` (`rbac_role_name`)
      REFERENCES `rbac_role` (`name`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_USER_ROLE_RBAC_USER_ID` (`rbac_user_id`) 
      REFERENCES `rbac_user` (`rbac_user_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;


CREATE TABLE `rbac_role_permission` (
  `rbac_role_permission_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rbac_role_name` varchar(255) NOT NULL,
  `rbac_permission_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`rbac_role_permission_id`),
  UNIQUE `UK_ROLE_NAME_PERMISSION_ID` (`rbac_role_name`, `rbac_permission_id`),
  FOREIGN KEY `FK_RBAC_ROLE_PEMISSION_RBAC_ROLE_NAME` (`rbac_role_name`)
      REFERENCES `rbac_role` (`name`)
      ON DELETE CASCADE,
  FOREIGN KEY `FK_RBAC_ROLE_PERMISSION_RBAC_PERMISSION_ID` (`rbac_permission_id`) 
      REFERENCES `rbac_permission` (`rbac_permission_id`)
      ON DELETE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE `rbac_session_role` (
    `rbac_session_role_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `rbac_role_name` varchar(255) NOT NULL,
    `rbac_user_id` int(11) unsigned NOT NULL,
    `rbac_session_name` varchar(255) NOT NULL,
    PRIMARY KEY (`rbac_session_role_id`),
    UNIQUE `UK_ROLE_NAME_USER_ID_SESSION_NAME` (`rbac_role_name`, `rbac_user_id`, `rbac_session_name`),
    FOREIGN KEY `FK_RBAC_SESSION_ROLE_RBAC_ROLE_NAME` (`rbac_role_name`)
        REFERENCES `rbac_role` (`name`)
        ON DELETE CASCADE,
    FOREIGN KEY `FK_RBAC_SESSION_ROLE_RBAC_USER_ID` (`rbac_user_id`)
        REFERENCES `rbac_user` (`rbac_user_id`)
        ON DELETE CASCADE,
    FOREIGN KEY `FK_RBAC_SESSION_ROLE_RBAC_SESSION_NAME` (`rbac_session_name`)
        REFERENCES `rbac_session` (`name`)
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
