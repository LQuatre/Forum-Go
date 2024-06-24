BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "categories" (
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "topics" (
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"category_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("category_uuid") REFERENCES "categories"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "threads" (
	"uuid"	TEXT NOT NULL,
	"topic_uuid"	INTEGER NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("topic_uuid") REFERENCES "topics"("uuid"),
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "posts" (
	"uuid"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"thread_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	FOREIGN KEY("thread_uuid") REFERENCES "threads"("uuid"),
	PRIMARY KEY("uuid")
);

CREATE TABLE IF NOT EXISTS "users" (
	"uuid"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL UNIQUE,
	"password"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL,
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "sessions" (
	"uuid"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);
COMMIT;
