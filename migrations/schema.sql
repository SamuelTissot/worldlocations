-- MySQL dump 10.13  Distrib 5.7.24, for osx10.14 (x86_64)
--
-- Host: localhost    Database: worldlocations_development
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `country_codes`
--

DROP TABLE IF EXISTS `country_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `country_codes` (
                               `alpha_2_code`       VARCHAR(2) NOT NULL,
                               `alpha_3_code`       VARCHAR(3)          DEFAULT NULL,
                               `numeric_code`       INT(11)             DEFAULT NULL,
                               `international_name` VARCHAR(255)        DEFAULT NULL,
                               `is_independant`     TINYINT(4)          DEFAULT NULL,
                               `iso_status`         VARCHAR(25)         DEFAULT NULL,
                               `created_at`         DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `updated_at`         DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`alpha_2_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `language_codes`
--

DROP TABLE IF EXISTS `language_codes`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `language_codes`
(
  `language_alpha_2_code` VARCHAR(2) NOT NULL,
  `language_alpha_3_code` VARCHAR(3)          DEFAULT NULL,
  `created_at`            DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`            DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`language_alpha_2_code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `schema_migration`
--

DROP TABLE IF EXISTS `schema_migration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migration` (
  `version` varchar(14) NOT NULL,
  UNIQUE KEY `schema_migration_version_idx` (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subdivision_codes`
--

DROP TABLE IF EXISTS `subdivision_codes`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subdivision_codes`
(
  `subdivision_code`   VARCHAR(6) NOT NULL,
  `alpha_2_code`       VARCHAR(2)          DEFAULT NULL,
  `international_name` VARCHAR(255)        DEFAULT NULL,
  `category`           VARCHAR(50)         DEFAULT NULL,
  `created_at`         DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`         DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`subdivision_code`),
  KEY `alpha_2_code` (`alpha_2_code`),
  CONSTRAINT `subdivision_codes_ibfk_1` FOREIGN KEY (`alpha_2_code`) REFERENCES `country_codes` (`alpha_2_code`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subdivision_names`
--

DROP TABLE IF EXISTS `subdivision_names`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subdivision_names`
(
  `subdivision_code`      VARCHAR(6)   NOT NULL,
  `language_alpha_2_code` VARCHAR(2)   NOT NULL,
  `name`                  VARCHAR(255) NOT NULL,
  `local_variation`       VARCHAR(255)          DEFAULT NULL,
  `created_at`            DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`            DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`subdivision_code`, `language_alpha_2_code`),
  KEY `language_alpha_2_code` (`language_alpha_2_code`),
  KEY `subdivision_names_subdivision_code_language_alpha_2_code_idx` (`subdivision_code`, `language_alpha_2_code`),
  CONSTRAINT `subdivision_names_ibfk_1` FOREIGN KEY (`subdivision_code`) REFERENCES `subdivision_codes` (`subdivision_code`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `subdivision_names_ibfk_2` FOREIGN KEY (`language_alpha_2_code`) REFERENCES `language_codes` (`language_alpha_2_code`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-02-08  8:18:09
