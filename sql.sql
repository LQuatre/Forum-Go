-- --------------------------------------------------------
-- Hôte:                         C:\Users\lucas\Desktop\B1\db\jilt.db
-- Version du serveur:           3.44.0
-- SE du serveur:                
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Listage de la structure de la base pour jilt
CREATE DATABASE IF NOT EXISTS "jilt";
;

-- Listage de la structure de la table jilt. categories
CREATE TABLE IF NOT EXISTS "categories" (
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL,
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.categories : 1 rows
/*!40000 ALTER TABLE "categories" DISABLE KEYS */;
INSERT INTO "categories" ("uuid", "name", "created_at") VALUES
	('8d708d9a-efea-4339-7134-3ce3bee4cce1', 'TEST', '2024-06-28 16:59:19.6267035+02:00');
/*!40000 ALTER TABLE "categories" ENABLE KEYS */;

-- Listage de la structure de la table jilt. comments
CREATE TABLE IF NOT EXISTS "comments" (
	"uuid"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"thread_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	FOREIGN KEY("thread_uuid") REFERENCES "threads"("uuid"),
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.comments : 12 rows
/*!40000 ALTER TABLE "comments" DISABLE KEYS */;
INSERT INTO "comments" ("uuid", "body", "user_uuid", "thread_uuid", "created_at") VALUES
	('30e71b08-5c99-4f4c-6542-5dfeff1100bc', 'caca
', 2ada2eaa-a6ca-4e39-606b-2f4d0ac2feb5, ba298bc0-4004-4bc8-5f06-7e4082843ac2, '2024-06-28 14:53:38.4726632+02:00'),
	('6a3fd0bc-2fed-4126-56e4-bdd796f7c75f', 'MDR
', e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, ba298bc0-4004-4bc8-5f06-7e4082843ac2, '2024-06-28 14:53:46.328377+02:00'),
	('55ae74df-dc28-4bf1-5cae-5f14a3ac7f01', 'nfjoqzrenbvoljqzenbvlùvnzerklvnzeoùnfoùgizeafezar
', 48d6154b-decf-4e63-6308-9946bddeb77e, cea3e5c2-b131-4517-7b26-8e407d27a7f0, '2024-06-28 16:09:46.3877885+02:00'),
	('ba7aabbc-a4fd-4573-568f-c9939e5e54e7', 'Mdrr', e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, cea3e5c2-b131-4517-7b26-8e407d27a7f0, '2024-06-28 16:09:54.5225706+02:00'),
	('5bdb30e6-a082-47f8-5d91-cedb15a3a83d', 'tg', 126485dc-b1c9-4972-64aa-7f67648209d5, cea3e5c2-b131-4517-7b26-8e407d27a7f0, '2024-06-28 16:25:00.463688+02:00'),
	('3bb259cd-c13f-4138-4f1c-ae531766e3eb', 'Twitch glide', 914db2db-02b3-4dc8-4a64-c61b5e4d7d3b, 9ce0bc77-918c-40af-5b93-4bc04a4c694d, '2024-06-28 16:26:29.1244883+02:00'),
	('c92a7f4e-f566-49e5-7f2b-c57906c58900', '卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍v', 52927b40-11fb-406d-4e8b-e372c22f2d1d, 1f298e78-dc8c-473d-5839-337083f74ca1, '2024-06-28 16:50:55.7101728+02:00'),
	('1f611fde-b8c4-4bcc-7263-e84337f67e1f', '卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍', 52927b40-11fb-406d-4e8b-e372c22f2d1d, 1f298e78-dc8c-473d-5839-337083f74ca1, '2024-06-28 16:50:57.6985966+02:00'),
	('cdaf3230-8704-4a72-5f42-daff34aa821d', '卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍卍v', 52927b40-11fb-406d-4e8b-e372c22f2d1d, 1f298e78-dc8c-473d-5839-337083f74ca1, '2024-06-28 16:50:59.7712905+02:00'),
	('ae5a76e3-2a59-4a7a-699d-eabe46359971', 'OK
', e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, bb993c06-e4cf-43d8-6502-24937e1d0c72, '2024-07-03 13:43:10.7033573+02:00'),
	('66473b7d-74de-449e-734c-1914cf57d3c5', 'Ok', 48d6154b-decf-4e63-6308-9946bddeb77e, bb993c06-e4cf-43d8-6502-24937e1d0c72, '2024-07-03 13:44:01.5891578+02:00'),
	('36e51b09-dc64-4ead-77c9-fc59351bb35c', 'mdr', e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, bb993c06-e4cf-43d8-6502-24937e1d0c72, '2024-07-03 13:44:32.5806875+02:00');
/*!40000 ALTER TABLE "comments" ENABLE KEYS */;

-- Listage de la structure de la table jilt. sessions
CREATE TABLE IF NOT EXISTS "sessions" (
	"uuid"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL, "isAdmin" INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.sessions : 6 rows
/*!40000 ALTER TABLE "sessions" DISABLE KEYS */;
INSERT INTO "sessions" ("uuid", "email", "user_uuid", "created_at", "isAdmin") VALUES
	('1d8d0351-4315-4c4c-6958-d4ef5c3faab2', 'tlaucournet@gmail.com', '', '2024-07-04 15:13:21.1681761+02:00', 0),
	('d77e8923-696f-4f14-651f-0c5f75da2dbe', 'tlaucournet@gmail.com', '', '2024-07-04 15:14:56.5117928+02:00', 0),
	('a0f85fa3-4654-4954-6968-b0575d89dbb8', 'tlaucournet@gmail.com', '', '2024-07-04 15:16:00.0299867+02:00', 0),
	('1410e7de-a0b9-4197-6795-2296f406b4f8', 'tlaucournet@gmail.com', '', '2024-07-04 15:16:42.8159711+02:00', 0),
	('e8c7b2dc-c5d4-49a2-5795-dd9292d742da', 'zabuz2005.zenvendark@gmail.com', '97463198-b3b8-4510-6d46-36d555cd5bfb', '2024-07-04 15:42:43.9972154+02:00', 0),
	('01150e42-7074-4e03-619c-ba1942fd3930', 'git@gmail.com', 'e6ab5dc0-9e55-4aab-5d96-fcfd87838e7a', '2024-07-04 15:44:33.7542091+02:00', 0);
/*!40000 ALTER TABLE "sessions" ENABLE KEYS */;

-- Listage de la structure de la table jilt. threads
CREATE TABLE IF NOT EXISTS "threads" (
	"uuid"	TEXT NOT NULL,
	"topic_uuid"	INTEGER NOT NULL,
	"user_uuid"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL, "desc" TEXT NOT NULL DEFAULT '',
	FOREIGN KEY("topic_uuid") REFERENCES "topics"("uuid"),
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.threads : 3 rows
/*!40000 ALTER TABLE "threads" DISABLE KEYS */;
INSERT INTO "threads" ("uuid", "topic_uuid", "user_uuid", "title", "created_at", "desc") VALUES
	('bb993c06-e4cf-43d8-6502-24937e1d0c72', cc0c2cfa-10bb-40e5-43f0-1371bd5b6c73, e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, 'DEUX', '2024-06-28 16:59:27.8053118+02:00', 'test'),
	('2f22e5fe-4ef8-4d5a-53f5-7521566ab337', cc0c2cfa-10bb-40e5-43f0-1371bd5b6c73, e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, 'qsf', '2024-06-28 17:13:58.2400811+02:00', 'qsfqsf'),
	('4491166f-a96d-4112-47e9-4fea4947e0ce', cc0c2cfa-10bb-40e5-43f0-1371bd5b6c73, e19c3993-9f60-4ea7-7873-5fd3a11e1bd4, 'qsfqqs', '2024-06-28 17:14:04.3475045+02:00', 'qsfqsfq');
/*!40000 ALTER TABLE "threads" ENABLE KEYS */;

-- Listage de la structure de la table jilt. topics
CREATE TABLE IF NOT EXISTS "topics" (
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"category_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL, "desc" TEXT NULL DEFAULT NULL,
	FOREIGN KEY("category_uuid") REFERENCES "categories"("uuid"),
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.topics : 8 rows
/*!40000 ALTER TABLE "topics" DISABLE KEYS */;
INSERT INTO "topics" ("uuid", "name", "category_uuid", "created_at", "desc") VALUES
	('7885a98c-ef27-4d96-457b-4efebb42c592', 'FiveM', 9fbe450e-7e00-47d0-6158-ab7960036a19, '2024-06-27 16:11:17.0748515+02:00', NULL),
	('1edcdde2-df7a-4d02-5acf-def31b044345', 'B1 Info A', bad1838f-550b-4f83-699d-ebfec0057bf5, '2024-06-28 15:39:18.8587343+02:00', NULL),
	('2a7b7369-6ce6-4af9-4e04-967e5aeee969', 'B1 Info B', bad1838f-550b-4f83-699d-ebfec0057bf5, '2024-06-28 15:39:24.5651713+02:00', NULL),
	('6d9eecb8-e8e3-40cb-68a5-cfb59c73d8f5', 'ADC Season 14', a60ca22b-9655-4fa2-5467-4a3439d20114, '2024-06-28 16:26:08.2368069+02:00', NULL),
	('a5bf33a6-1271-4e39-74e5-f1ae6f9cc527', 'CR7>Messi', 51a491df-6130-48e8-4215-e8b812e96373, '2024-06-28 16:26:18.4045757+02:00', NULL),
	('a3cf968c-7626-44fb-7e2d-87d58b471c5d', '卍卍卍卍卍卍卍卍卍卍', 69e301e0-d64e-4829-74b2-9778872f3f53, '2024-06-28 16:50:44.1911492+02:00', NULL),
	('cc0c2cfa-10bb-40e5-43f0-1371bd5b6c73', 'UN', 8d708d9a-efea-4339-7134-3ce3bee4cce1, '2024-06-28 16:59:24.2099134+02:00', NULL),
	('2ea80557-bac7-45a9-42df-ea59c61afda4', 'ta maman', est sympa, '2024-06-28 17:17:28.5228764+02:00', '8d708d9a-efea-4339-7134-3ce3bee4cce1');
/*!40000 ALTER TABLE "topics" ENABLE KEYS */;

-- Listage de la structure de la table jilt. users
CREATE TABLE IF NOT EXISTS "users" (
	"uuid"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL UNIQUE,
	"password"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL, "isAdmin" INTEGER NOT NULL DEFAULT 0, "discord_id" VARCHAR(50) NULL DEFAULT NULL, "google_id" VARCHAR(50) NULL DEFAULT NULL, "facebook_id" VARCHAR(50) NULL DEFAULT NULL, "github_id" VARCHAR(50) NULL DEFAULT NULL,
	PRIMARY KEY("uuid")
);

-- Listage des données de la table jilt.users : 1 rows
/*!40000 ALTER TABLE "users" DISABLE KEYS */;
INSERT INTO "users" ("uuid", "name", "email", "password", "created_at", "isAdmin", "discord_id", "google_id", "facebook_id", "github_id") VALUES
	('e6ab5dc0-9e55-4aab-5d96-fcfd87838e7a', 'git', 'git@gmail.com', '$2a$10$rJqeZRIvp7OB9l/knw1NNu1Q/zcnr9gbgG/7MAiA3bTY7eVyz1CI2', '2024-07-04 15:44:22.1932252+02:00', 0, '', '', '', '');
/*!40000 ALTER TABLE "users" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
