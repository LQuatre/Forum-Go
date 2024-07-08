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
	"desc"	TEXT DEFAULT NULL,
	FOREIGN KEY("category_uuid") REFERENCES "categories"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "threads" (
	"uuid"	TEXT NOT NULL,
	"topic_uuid"	INTEGER NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	"desc"	TEXT NOT NULL DEFAULT '',
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	FOREIGN KEY("topic_uuid") REFERENCES "topics"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "comments" (
	"uuid"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"thread_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("thread_uuid") REFERENCES "threads"("uuid"),
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "users" (
	"uuid"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL UNIQUE,
	"password"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL,
	"isAdmin"	INTEGER NOT NULL DEFAULT 0,
	"discord_id"	VARCHAR(50) DEFAULT NULL,
	"google_id"	VARCHAR(50) DEFAULT NULL,
	"facebook_id"	VARCHAR(50) DEFAULT NULL,
	"github_id"	VARCHAR(50) DEFAULT NULL,
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "sessions" (
	"uuid"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL,
	"isAdmin"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "tickets" (
	"uuid"	VARCHAR(50) DEFAULT NULL,
	"name"	VARCHAR(50) DEFAULT NULL,
	"desc"	VARCHAR(50) DEFAULT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATE
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);
CREATE TABLE IF NOT EXISTS "tikets-messages" (
	"uuid"	VARCHAR(50) DEFAULT NULL,
	"messages"	VARCHAR(50) DEFAULT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"ticket_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATE
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	FOREIGN KEY("ticket_uuid") REFERENCES "tickets"("uuid"),
	PRIMARY KEY("uuid")
);
INSERT INTO "users" VALUES ('39a96046-d88f-43a2-771a-ff75631455de','isma','ismailbenlaidi@gmail.com','$2a$10$Tmi8cc6s8yxPYpstM0/RRe9Cp7K3CtmV1AeGl3eb4nfToqqFE4uFK','2024-07-05 14:38:03.8988075+02:00',0,'455182152586362901','','','');
INSERT INTO "users" VALUES ('09dbb778-c4b4-4042-4b07-4d1a8601d998','AcryX','tlaucournet@gmail.com','$2a$10$z2jp5lYXYncA74H12/kLSuRTkW2C0VH8ReHPa8cchBmSIefR7e.46','2024-07-05 14:38:13.6846194+02:00',0,'235120346112458752','','','');
INSERT INTO "users" VALUES ('d19d7150-f7a5-4c7b-6114-3709fee77bff','L4.','yellow.zabuza.gaming@gmail.com','$2a$10$gdzpjgVlQnO3B01OGqsAsODRVs8itzgr1i2v4Nimzz8SGBf.yj2Xm','2024-07-05 15:17:09.0331681+02:00',0,'351702205981523988','','','');
INSERT INTO "users" VALUES ('455af70b-00fb-45ee-5099-cbc0770fcee7','saiko','delbeau.jb@gmail.com','$2a$10$wTzDknmQ4bCWtkQsrYVJ/uzzcBCCYtl3pL2ryJgB5Gqn/kdhEchZq','2024-07-05 15:17:13.461493+02:00',0,'332969890531246081','','','');
INSERT INTO "sessions" VALUES ('0dc773af-6681-4aff-4650-2103f5135d1d','unmecpasfort@gmail.com','ed8cd38b-4a81-4a62-42ca-1c4368333c25','2024-07-05 14:31:31.7404935+02:00',0);
INSERT INTO "sessions" VALUES ('f7288a56-fc79-4ba6-622b-7bbc00c2d5b8','ismailbenlaidi@gmail.com','39a96046-d88f-43a2-771a-ff75631455de','2024-07-05 14:38:03.9043425+02:00',0);
INSERT INTO "sessions" VALUES ('c48cc4f8-5467-443e-4111-7a7fff615400','tlaucournet@gmail.com','09dbb778-c4b4-4042-4b07-4d1a8601d998','2024-07-05 14:38:13.6910667+02:00',0);
INSERT INTO "sessions" VALUES ('93ea604d-ca13-4e28-75d4-8801c7b8ae93','tlaucournet@gmail.com','09dbb778-c4b4-4042-4b07-4d1a8601d998','2024-07-05 15:16:53.8652228+02:00',0);
INSERT INTO "sessions" VALUES ('20e69435-2580-48cd-6453-809f2b7c9f21','yellow.zabuza.gaming@gmail.com','d19d7150-f7a5-4c7b-6114-3709fee77bff','2024-07-05 15:17:09.1198357+02:00',0);
INSERT INTO "sessions" VALUES ('2ecc52ff-950f-4802-4d5c-39c71cc53e06','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 15:17:13.4645071+02:00',0);
INSERT INTO "sessions" VALUES ('2316c78b-94aa-4f95-4d86-41e60f949204','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 15:19:43.4315228+02:00',0);
INSERT INTO "sessions" VALUES ('f2237862-0ee8-4167-5032-d2582256de6e','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 15:19:59.6297789+02:00',0);
INSERT INTO "sessions" VALUES ('b1f7aff5-60d3-405c-5366-e551a06695ae','tlaucournet@gmail.com','09dbb778-c4b4-4042-4b07-4d1a8601d998','2024-07-05 15:24:55.2608224+02:00',0);
INSERT INTO "sessions" VALUES ('4702ae30-0818-4c68-43b6-2155afcbb514','yellow.zabuza.gaming@gmail.com','d19d7150-f7a5-4c7b-6114-3709fee77bff','2024-07-05 16:16:40.3839522+02:00',0);
INSERT INTO "sessions" VALUES ('97b730e5-188b-4457-648f-07c7c6c41db9','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 16:16:45.6314287+02:00',0);
INSERT INTO "sessions" VALUES ('33ef7d54-1ff5-4976-4b52-7ed4399f95e4','yellow.zabuza.gaming@gmail.com','d19d7150-f7a5-4c7b-6114-3709fee77bff','2024-07-05 16:17:05.5523687+02:00',0);
INSERT INTO "sessions" VALUES ('6cc08ef7-d3e4-4d70-6f01-70fff6d7f565','ismailbenlaidi@gmail.com','39a96046-d88f-43a2-771a-ff75631455de','2024-07-05 16:22:00.787794+02:00',0);
INSERT INTO "sessions" VALUES ('ee9ebdfb-ea36-42f7-7de1-98b0697d7bcf','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 16:43:48.537282+02:00',0);
INSERT INTO "sessions" VALUES ('47380d12-754b-439d-4940-6a5867a4e7c8','yellow.zabuza.gaming@gmail.com','d19d7150-f7a5-4c7b-6114-3709fee77bff','2024-07-05 16:46:42.3589439+02:00',0);
INSERT INTO "sessions" VALUES ('0ed36363-92da-40ed-7828-9bb58679a7c3','delbeau.jb@gmail.com','455af70b-00fb-45ee-5099-cbc0770fcee7','2024-07-05 16:58:00.0279677+02:00',0);
INSERT INTO "sessions" VALUES ('76b6999c-e7ed-4555-5047-f46358d5833a','yellow.zabuza.gaming@gmail.com','d19d7150-f7a5-4c7b-6114-3709fee77bff','2024-07-05 16:59:26.8357041+02:00',0);
COMMIT;
