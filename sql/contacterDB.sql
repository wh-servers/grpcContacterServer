CREATE database contacter_db;

USE contacter_db;

/*fieldName type(size) (unsigned) Can_be_null auto_increament/default+value*/
CREATE TABLE contacter_changelog_tab (
id bigint(11) unsigned NOT NULL AUTO_INCREMENT,
contacterid bigint(11) unsigned NOT NULL,
time int(11) unsigned NOT NULL DEFAULT '0',
old blob,
new blob,
operator varchar(128) NOT NULL DEFAULT '',
extinfo blob,
PRIMARY KEY (id),
KEY `k_id_time` (`id`,`time`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;

CREATE TABLE contacter_tab (
id bigint(11) unsigned NOT NULL AUTO_INCREMENT,
name varchar(512) NOT NULL DEFAULT '',
country char(4) NOT NULL DEFAULT '',
status int(11) unsigned NOT NULL DEFAULT '0',
image varchar(256) NOT NULL DEFAULT '', 
headphone int(11) unsigned NOT NULL DEFAULT '0',
homephone int(11) unsigned NOT NULL DEFAULT '0',
email varchar(256) NOT NULL DEFAULT '', 
ctime int(11) unsigned NOT NULL DEFAULT '0',
mtime int(11) unsigned NOT NULL DEFAULT '0',
extinfo blob,
PRIMARY KEY (id),
KEY `k_name_headphone` (`name`,`headphone`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4;
