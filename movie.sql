-- phpMyAdmin SQL Dump
-- version 5.0.0-dev
-- https://www.phpmyadmin.net/
--
-- Host: 192.168.30.23
-- Generation Time: Jun 05, 2019 at 09:00 AM
-- Server version: 8.0.3-rc-log
-- PHP Version: 7.2.18-1+0~20190503103213.21+stretch~1.gbp101320

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rupa-movie`
--

-- --------------------------------------------------------

--
-- Table structure for table `movies`
--

CREATE TABLE `movies` (
  `id` varchar(50) NOT NULL,
  `title` varchar(50) NOT NULL,
  `price` varchar(15) NOT NULL DEFAULT 'NA',
  `description` text NOT NULL,
  `trailer_url` text,
  `audience_label` enum('NOT_RATED','GENERAL_AUDIENCE','PARENTAL_GUIDANCE','PARENTAL_GUIDANCE_13','NO_ONE_17') NOT NULL DEFAULT 'NOT_RATED',
  `ratings` float NOT NULL DEFAULT '0',
  `duration` int(11) DEFAULT 0,
  `all_votes` int(11) DEFAULT 0,
  `release_date` date DEFAULT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `category` json NOT NULL,
  `photos` json NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `movies`
--
ALTER TABLE `movies`
  ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
