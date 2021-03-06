package state

var createTables = `
CREATE TABLE components (
	id VARCHAR(36) PRIMARY KEY NOT NULL,
	component_type VARCHAR(32) NOT NULL,
	component_path TEXT NOT NULL,
	specification_path TEXT NOT NULL,
	created_at INTEGER NOT NULL
);

CREATE TABLE flows (
	id VARCHAR(36) PRIMARY KEY NOT NULL,
	specification_path TEXT NOT NULL,
	created_at INTEGER NOT NULL
);

CREATE TABLE builds (
	id VARCHAR(36) PRIMARY KEY NOT NULL,
	component_id VARCHAR(36) NOT NULL,
	created_at INTEGER NOT NULL
);

CREATE TABLE executions (
	id VARCHAR(36) PRIMARY KEY NOT NULL,
	build_id VARCHAR(36) NOT NULL,
	component_id VARCHAR(36) NOT NULL,
	created_at INTEGER NOT NULL,
	flow_id VARCHAR(36)
);
`
