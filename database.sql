-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 29. Jan, 2019 17:17 PM
-- Server-versjon: 10.1.28-MariaDB
-- PHP Version: 7.1.10

DROP SCHEMA

IF EXISTS cs53;
	CREATE SCHEMA cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cs53`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `assignments`
--

CREATE TABLE `assignments` (
  `id` int(11) NOT NULL,
  `courseid` int(11) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `due` date NOT NULL,
  `peer` tinyint(1) NOT NULL DEFAULT '0',
  `auto` tinyint(1) NOT NULL DEFAULT '0',
  `language` varchar(20) COLLATE utf8_danish_ci DEFAULT NULL,
  `tasktext` text COLLATE utf8_danish_ci,
  `payload` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `course`
--

CREATE TABLE `course` (
  `id` int(11) NOT NULL,
  `coursecode` varchar(10) COLLATE utf8_danish_ci NOT NULL,
  `coursename` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher` int(11) NOT NULL,
  `info` text COLLATE utf8_danish_ci,
  `link1` varchar(128) COLLATE utf8_danish_ci DEFAULT NULL,
  `link2` varchar(128) COLLATE utf8_danish_ci DEFAULT NULL,
  `link3` varchar(128) COLLATE utf8_danish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `logs`
--

CREATE TABLE `logs` (
  `userid` int(11) NOT NULL,
  `log` varchar(256) COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `peerreviews`
--

CREATE TABLE `peerreviews` (
  `assignmentid` int(11) NOT NULL,
  `submissionid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `grade` int(11) NOT NULL,
  `feedback` varchar(512) COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `reviewerpairs`
--

CREATE TABLE `reviewerpairs` (
  `assignmentid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `submissionsid1` int(11) NOT NULL,
  `submissionid2` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `submissions`
--

CREATE TABLE `submissions` (
  `id` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `assignmentid` int(11) NOT NULL,
  `repo` varchar(128) COLLATE utf8_danish_ci NOT NULL,
  `deploy` varchar(128) COLLATE utf8_danish_ci NOT NULL,
  `comment` varchar(512) COLLATE utf8_danish_ci NOT NULL,
  `grade` int(11) DEFAULT NULL,
  `test` int(11) DEFAULT NULL,
  `vet` int(11) DEFAULT NULL,
  `cycle` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `usercourse`
--

CREATE TABLE `usercourse` (
  `userid` int(11) NOT NULL,
  `courseid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `email` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher` tinyint(1) NOT NULL,
  `email2` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `password` varchar(64) COLLATE utf8_danish_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `assignments`
--
ALTER TABLE `assignments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `courseid` (`courseid`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher` (`teacher`);

--
-- Indexes for table `logs`
--
ALTER TABLE `logs`
  ADD KEY `userid` (`userid`);

--
-- Indexes for table `peerreviews`
--
ALTER TABLE `peerreviews`
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `submissionid` (`submissionid`),
  ADD KEY `userid` (`userid`);

--
-- Indexes for table `reviewerpairs`
--
ALTER TABLE `reviewerpairs`
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `userid` (`userid`),
  ADD KEY `submissionsid1` (`submissionsid1`),
  ADD KEY `submissionid2` (`submissionid2`);

--
-- Indexes for table `submissions`
--
ALTER TABLE `submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `userid` (`userid`),
  ADD KEY `assignmentid` (`assignmentid`);

--
-- Indexes for table `usercourse`
--
ALTER TABLE `usercourse`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `email2` (`email2`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `submissions`
--
ALTER TABLE `submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Begrensninger for dumpede tabeller
--

--
-- Begrensninger for tabell `assignments`
--
ALTER TABLE `assignments`
  ADD CONSTRAINT `assignments_ibfk_1` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `logs`
--
ALTER TABLE `logs`
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Begrensninger for tabell `peerreviews`
--
ALTER TABLE `peerreviews`
  ADD CONSTRAINT `peerreviews_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_2` FOREIGN KEY (`submissionid`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `peerreviews_ibfk_3` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE;

--
-- Begrensninger for tabell `reviewerpairs`
--
ALTER TABLE `reviewerpairs`
  ADD CONSTRAINT `reviewerpairs_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_3` FOREIGN KEY (`submissionsid1`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `reviewerpairs_ibfk_4` FOREIGN KEY (`submissionid2`) REFERENCES `submissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `submissions`
--
ALTER TABLE `submissions`
  ADD CONSTRAINT `submissions_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `submissions_ibfk_2` FOREIGN KEY (`assignmentid`) REFERENCES `assignments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
