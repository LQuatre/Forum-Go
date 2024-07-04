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

-- Les données exportées n'étaient pas sélectionnées.

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

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de la table jilt. sessions
CREATE TABLE IF NOT EXISTS "sessions" (
	"uuid"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL,
	"user_uuid"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL, "isAdmin" INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("user_uuid") REFERENCES "users"("uuid"),
	PRIMARY KEY("uuid")
);

-- Les données exportées n'étaient pas sélectionnées.

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

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de la table jilt. topics
CREATE TABLE IF NOT EXISTS "topics" (
	"uuid"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"category_uuid"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL, "desc" TEXT NULL DEFAULT NULL,
	FOREIGN KEY("category_uuid") REFERENCES "categories"("uuid"),
	PRIMARY KEY("uuid")
);

-- Les données exportées n'étaient pas sélectionnées.

-- Listage de la structure de la table jilt. users
CREATE TABLE IF NOT EXISTS "users" (
	"uuid"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"email"	VARCHAR(255) NOT NULL UNIQUE,
	"password"	VARCHAR(255) NOT NULL,
	"created_at"	DATETIME NOT NULL, "isAdmin" INTEGER NOT NULL DEFAULT 0, "discord_id" VARCHAR(50) NULL DEFAULT NULL, "google_id" VARCHAR(50) NULL DEFAULT NULL, "facebook_id" VARCHAR(50) NULL DEFAULT NULL, "github_id" VARCHAR(50) NULL DEFAULT NULL,
	PRIMARY KEY("uuid")
);

-- Les données exportées n'étaient pas sélectionnées.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
