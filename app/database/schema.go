package database

import "github.com/rs/zerolog/log"

func createUsersTable() error {

	query := `CREATE TABLE IF NOT EXISTS users (
		id smallint(2) unsigned NOT NULL auto_increment,
		userid varchar(64) NOT NULL default '',
		uid smallint(6) unsigned default NULL,
		gid smallint(6) unsigned default NULL,
		passwd varchar(265) NOT NULL default '',
		homedir varchar(255) NOT NULL default '',
		comment varchar(255) NOT NULL default '',
		disabled smallint(2) unsigned NOT NULL default '0',
		shell varchar(32) NOT NULL default '/bin/false',
		sshpubkey varchar(8192) NOT NULL default '',
		email varchar(255) NOT NULL default '',
		name varchar(255) NOT NULL default '',
		title varchar(5) NOT NULL default '',
		company varchar(255) NOT NULL default '',
		quota_type enum('user','group','class','all') NOT NULL default 'user',
		per_session enum('false','true') NOT NULL default 'false',
		limit_type enum('soft','hard') NOT NULL default 'soft',
		bytes_in_used bigint(20) unsigned NOT NULL default '0',
		bytes_out_used bigint(20) unsigned NOT NULL default '0',
		bytes_xfer_avail int(10) unsigned NOT NULL default '0',
		files_in_used bigint(20) unsigned NOT NULL default '0',
		files_out_used bigint(20) unsigned NOT NULL default '0',
		files_xfer_avail int(10) unsigned NOT NULL default '0',
		login_count int(11) unsigned NOT NULL default '0',
		last_login datetime NOT NULL default '0000-00-00 00:00:00',
		last_modified datetime NOT NULL default '0000-00-00 00:00:00',
		expiration datetime NOT NULL default '0000-00-00 00:00:00',
		PRIMARY KEY  (id),
		UNIQUE KEY userid (userid)
	  ) ENGINE=MyISAM COMMENT='ProFTPd user table';`

	return DB().Exec(query).Error
}

func createGroupsTable() error {

	query := `CREATE TABLE IF NOT EXISTS groups (
		groupname varchar(32) NOT NULL default '',
		gid smallint(6) unsigned NOT NULL auto_increment,
		members varchar(255) NOT NULL default '',
		PRIMARY KEY  (gid),
		UNIQUE KEY groupname (groupname)
	  ) ENGINE=MyISAM COMMENT='ProFTPd group table';`

	return DB().Exec(query).Error

}

func createQuotaTable() error {
	query := `CREATE TABLE IF NOT EXISTS quotatallies (
		name varchar(30) NOT NULL default '',
		quota_type enum('user','group','class','all') NOT NULL default 'user',
		bytes_in_used int(10) unsigned NOT NULL default '0',
		bytes_out_used int(10) unsigned NOT NULL default '0',
		bytes_xfer_used int(10) unsigned NOT NULL default '0',
		files_in_used int(10) unsigned NOT NULL default '0',
		files_out_used int(10) unsigned NOT NULL default '0',
		files_xfer_used int(10) unsigned NOT NULL default '0'
	  ) ENGINE=MyISAM COMMENT='ProFTPd Quota table';`

	return DB().Exec(query).Error
}

func initTables() {
	if err := createUsersTable(); err != nil {
		log.Fatal().Err(err).Msg("Failed to create users table")
	}

	if err := createGroupsTable(); err != nil {
		log.Fatal().Err(err).Msg("Failed to create groups table")
	}

	if err := createQuotaTable(); err != nil {
		log.Fatal().Err(err).Msg("Failed to create quota table")
	}
}
