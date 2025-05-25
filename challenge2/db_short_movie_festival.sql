-- MySQL dump 10.13  Distrib 8.0.41, for Win64 (x86_64)
--
-- Host: 172.25.56.118    Database: db_short_movie_festival
-- ------------------------------------------------------
-- Server version	8.0.42-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `artists`
--

DROP TABLE IF EXISTS `artists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `artists` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artists`
--

LOCK TABLES `artists` WRITE;
/*!40000 ALTER TABLE `artists` DISABLE KEYS */;
INSERT INTO `artists` VALUES (1,'Tom Cruise'),(2,'Tom Holland'),(3,'Johny Depp'),(4,'Dwayne Johnson'),(5,'Vin Diesel');
/*!40000 ALTER TABLE `artists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genres`
--

DROP TABLE IF EXISTS `genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `genres` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genres`
--

LOCK TABLES `genres` WRITE;
/*!40000 ALTER TABLE `genres` DISABLE KEYS */;
INSERT INTO `genres` VALUES (1,'Action'),(2,'Comedy'),(3,'Drama'),(4,'Horror'),(5,'Romance'),(6,'Sci-Fi');
/*!40000 ALTER TABLE `genres` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie_artists`
--

DROP TABLE IF EXISTS `movie_artists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_artists` (
  `movie_id` int NOT NULL,
  `artist_id` int NOT NULL,
  PRIMARY KEY (`movie_id`,`artist_id`),
  KEY `artist_id` (`artist_id`),
  CONSTRAINT `movie_artists_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `movie_artists_ibfk_2` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_artists`
--

LOCK TABLES `movie_artists` WRITE;
/*!40000 ALTER TABLE `movie_artists` DISABLE KEYS */;
INSERT INTO `movie_artists` VALUES (7,1),(11,1),(13,1),(15,1),(16,1),(17,1),(7,2),(11,2),(13,2),(18,2),(7,3),(12,3),(13,3),(14,3),(18,3),(12,5),(14,5),(15,5),(16,5),(17,5);
/*!40000 ALTER TABLE `movie_artists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie_genres`
--

DROP TABLE IF EXISTS `movie_genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_genres` (
  `movie_id` int NOT NULL,
  `genre_id` int NOT NULL,
  PRIMARY KEY (`movie_id`,`genre_id`),
  KEY `genre_id` (`genre_id`),
  CONSTRAINT `movie_genres_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `movie_genres_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_genres`
--

LOCK TABLES `movie_genres` WRITE;
/*!40000 ALTER TABLE `movie_genres` DISABLE KEYS */;
INSERT INTO `movie_genres` VALUES (12,1),(14,1),(18,1),(12,2),(14,2),(18,2),(11,3),(15,3),(16,3),(17,3),(7,4),(11,4),(13,4),(7,5),(13,5),(15,5),(16,5),(17,5);
/*!40000 ALTER TABLE `movie_genres` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movies`
--

DROP TABLE IF EXISTS `movies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_general_ci,
  `duration` time DEFAULT NULL,
  `video_file_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movies`
--

LOCK TABLES `movies` WRITE;
/*!40000 ALTER TABLE `movies` DISABLE KEYS */;
INSERT INTO `movies` VALUES (7,'Film Keren','desc film keren','02:05:01','6832df90b1fd9a9eed1cf2a2'),(11,'Fast & Furious Jakarta','action film, night race jakarta','01:50:07','6832e583b2c26e8bc4ffb777'),(12,'The Nongkrong Geng','film bapak\" main catur','00:40:00','6833153564781fa25d54cdeb'),(13,'The Queens','Kisah sekelompok gadis cantik','01:50:00','683315b964781fa25d54cdf3'),(14,'The King','Kisah sekelompok pria tampan','01:50:00','683315dd64781fa25d54ce0e'),(15,'Pendekar Rajawali','FIlm action kuno','03:07:15','68331764720f05145ca19366'),(16,'Pendekar Rajawali','FIlm action kuno tapi bagus','03:07:15','68331764720f05145ca19366'),(17,'Gamer Andalan','film siswa SMA yg jago game','01:20:09','68331965e8c0af02cc14f905'),(18,'Calon Programmer','Kisah mahasiswa yang ingin sekali menjadi programmer','00:45:00','6833153564781fa25d54cdeb');
/*!40000 ALTER TABLE `movies` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-25 21:45:06
