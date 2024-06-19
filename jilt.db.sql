BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "categories" (
	"id"	INTEGER,
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT=1)
);
CREATE TABLE IF NOT EXISTS "topics" (
	"id"	INTEGER,
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"category_id"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("category_id") REFERENCES "categories"("id"),
	PRIMARY KEY("id" AUTOINCREMENT=1)
);
CREATE TABLE IF NOT EXISTS "threads" (
	"id"	INTEGER,
	"uuid"	TEXT NOT NULL,
	"topic_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("topic_id") REFERENCES "topics"("id"),
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id" AUTOINCREMENT=1)
);
CREATE TABLE IF NOT EXISTS "posts" (
	"id"	INTEGER,
	"uuid"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"thread_id"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	FOREIGN KEY("thread_id") REFERENCES "threads"("id"),
	PRIMARY KEY("id" AUTOINCREMENT=1)
);

CREATE TABLE IF NOT EXISTS "users" (
	"id"	INT AUTO_INCREMENT=1,
	"uuid"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL UNIQUE,
	"password"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "sessions" (
	"id"	INT AUTO_INCREMENT=1,
	"uuid"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL,
	"user_id"	INT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id"),
	PRIMARY KEY("id")
);
COMMIT;
