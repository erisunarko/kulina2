-- phpMyAdmin SQL Dump
-- version 4.8.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 18, 2019 at 04:07 PM
-- Server version: 10.1.31-MariaDB
-- PHP Version: 7.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_ktp`
--

CREATE TABLE `tb_ktp` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `address` varchar(255) NOT NULL,
  `place_of_birth` text NOT NULL,
  `date_of_birth` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tb_ktp`
--

INSERT INTO `tb_ktp` (`id`, `name`, `address`, `place_of_birth`, `date_of_birth`) VALUES
(4, 'Surja Wibowo', 'Semarang', 'Kota Semarang', '1999-02-22'),
(5, 'Martis Supeno', 'Solo', 'Wonogiri', '1998-01-04');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_ktp`
--
ALTER TABLE `tb_ktp`
  ADD UNIQUE KEY `id_2` (`id`),
  ADD KEY `id` (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_ktp`
--
ALTER TABLE `tb_ktp`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
