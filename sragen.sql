-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 30, 2023 at 05:21 PM
-- Server version: 10.1.38-MariaDB
-- PHP Version: 5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sragen`
--

-- --------------------------------------------------------

--
-- Table structure for table `auth`
--

CREATE TABLE `auth` (
  `auth_id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int(1) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `auth`
--

INSERT INTO `auth` (`auth_id`, `username`, `email`, `password`, `role_id`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 'deigo', 'deigojo@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 1, 1, '2023-07-22 08:21:35', '2023-07-21 17:00:00'),
(5, 'siahaan', 'siahaan@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 2, 0, '2023-07-22 14:03:13', '2023-07-22 11:28:49'),
(6, 'jonathan', 'jonathan@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 2, 0, '2023-07-23 01:38:44', '2023-07-22 11:30:56'),
(7, 'dri', 'a@gmail.com', '$2a$04$Bt97CDU/141HvgPhXdhRxe.dliSxRdLIFunpf0ccNJyAVnrp919ga', 2, 1, '2023-07-23 01:41:59', '2023-07-23 01:22:50'),
(8, 'sobu', 'sobu@gmail.com', '$2a$04$oYuNgFfvkUsIWp3QM5TGUOBFCnw12VhEJINn5mMdgr9868jUo2t.C', 0, 1, '2023-07-23 01:54:17', '2023-07-23 01:54:17');

-- --------------------------------------------------------

--
-- Table structure for table `master_pengajuan`
--

CREATE TABLE `master_pengajuan` (
  `id_jenis_pengajuan` int(11) NOT NULL,
  `nama_pengajuan` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_active` int(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `master_pengajuan`
--

INSERT INTO `master_pengajuan` (`id_jenis_pengajuan`, `nama_pengajuan`, `created_at`, `is_active`) VALUES
(1, 'Surat Domisili', '2023-07-30 01:57:13', 1),
(2, 'Surat Keterangan Usaha', '2023-07-30 01:57:13', 1),
(3, 'Surat Kartu Keluarga', '2023-07-30 01:57:13', 1),
(4, 'Surat Kartu Tanda Penduduk', '2023-07-30 01:57:13', 1),
(5, 'Surat Kematian', '2023-07-30 01:57:13', 1),
(6, 'Surat Kelahiran', '2023-07-30 01:57:13', 1),
(7, 'Surat kehilangan', '2023-07-30 01:57:56', 1);

-- --------------------------------------------------------

--
-- Table structure for table `pengajuan_dokumen`
--

CREATE TABLE `pengajuan_dokumen` (
  `id_pengajuan` int(11) NOT NULL,
  `auth_id` int(11) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `nik` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `jenis_kelamin` varchar(20) DEFAULT NULL,
  `agama` varchar(20) DEFAULT NULL,
  `pekerjaan` varchar(50) DEFAULT NULL,
  `tempat_lahir` varchar(50) DEFAULT NULL,
  `tanggal_lahir` timestamp NULL DEFAULT NULL,
  `alamat` varchar(100) DEFAULT NULL,
  `status_perkawinan` varchar(30) DEFAULT NULL,
  `foto` varchar(100) DEFAULT NULL,
  `ktp` varchar(100) DEFAULT NULL,
  `alasan` varchar(255) DEFAULT NULL,
  `keperluan` varchar(255) DEFAULT NULL,
  `status` int(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NULL DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `jenis_pengajuan` int(11) NOT NULL,
  `kepala_keluarga` varchar(50) DEFAULT NULL,
  `nomor_kk` varchar(30) DEFAULT NULL,
  `tanda_tangan` varchar(255) DEFAULT NULL,
  `jenis_usaha` varchar(255) DEFAULT NULL,
  `nama_anak` varchar(50) DEFAULT NULL,
  `nama_ibu` varchar(50) DEFAULT NULL,
  `nama_ayah` varchar(50) DEFAULT NULL,
  `jenis_kelamin_anak` varchar(20) DEFAULT NULL,
  `nama_pengaju` varchar(50) DEFAULT NULL,
  `hubungan_pengaju` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pengajuan_dokumen`
--

INSERT INTO `pengajuan_dokumen` (`id_pengajuan`, `auth_id`, `nama`, `nik`, `phone`, `jenis_kelamin`, `agama`, `pekerjaan`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `status_perkawinan`, `foto`, `ktp`, `alasan`, `keperluan`, `status`, `created_at`, `updated_at`, `updated_by`, `jenis_pengajuan`, `kepala_keluarga`, `nomor_kk`, `tanda_tangan`, `jenis_usaha`, `nama_anak`, `nama_ibu`, `nama_ayah`, `jenis_kelamin_anak`, `nama_pengaju`, `hubungan_pengaju`) VALUES
(1, 8, NULL, '1151352352352', '0812412414', 'laki-laki', 'kristen', 'developer', 'balige', '1996-11-13 17:00:00', 'grogol', NULL, '23243261-7a68-4889-9736-fd794efbbbb6.jpg', '7d85d9c4-8f33-40ff-b041-c8e9cd991848.jpg', 'testing', '', 0, '2023-07-24 15:55:23', '2023-07-24 15:55:23', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL),
(2, 8, NULL, '1151352352352', '0812412414', 'laki-laki', 'kristen', 'developer', 'balige', '1996-11-13 17:00:00', 'grogol', NULL, 'df686752-dc7f-47ca-8218-35b4b2d8f803.jpg', 'eaf4844d-6151-474e-9608-ee7d420d6910.jpg', 'testing', NULL, 0, '2023-07-30 03:07:49', '2023-07-30 03:07:49', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL),
(3, 0, NULL, '', '', '', 'Kristen Protestan', 'asd', 'Balige Toba Samosir', '2023-07-05 17:00:00', 'asd', NULL, '414ffb9a-bbc4-4ee1-99da-59c4abccbfc9.jpg', '711ba75c-16b4-4db1-80c8-acdfa1c23b76.jpg', '', NULL, 0, '2023-07-30 03:35:21', '2023-07-30 03:35:21', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL),
(4, 8, NULL, '', '', '', 'Kristen Protestan', 'asd', 'Balige Toba Samosir', '2023-07-05 17:00:00', 'asd', 'asd', '6993a000-b2ec-4148-9a3e-abb8af1751ce.jpg', 'b5fad330-9d41-4ad5-ba7b-3d3128a4b894.jpg', '', 'asd', 0, '2023-07-30 03:39:59', '2023-07-30 03:39:59', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL),
(5, 8, NULL, '12412412412', NULL, '', 'Kristen Protestan', NULL, NULL, NULL, 'asd', NULL, NULL, NULL, NULL, NULL, 0, '2023-07-30 14:02:17', '2023-07-30 14:02:17', NULL, 3, 'beae', '12515151551', '04b2c871-c8f9-463e-b93f-00835e97b158.jpg', '', NULL, NULL, NULL, NULL, NULL, NULL),
(6, 8, NULL, NULL, NULL, NULL, NULL, 'asd', NULL, NULL, 'sdgsdgsdg', NULL, NULL, '1b52594a-3ec4-475a-b0a9-2ba93991f761.jpg', NULL, NULL, 0, '2023-07-30 14:32:09', '2023-07-30 14:32:09', NULL, 2, NULL, NULL, NULL, 'sdgsdgsdg', NULL, NULL, NULL, NULL, NULL, NULL),
(7, 8, '', NULL, NULL, NULL, NULL, NULL, 'London', '2023-06-27 17:00:00', 'asda', NULL, NULL, '3ff13709-f2cd-434d-9daa-33e7dd4a3e0d.jpg', NULL, NULL, 0, '2023-07-30 14:57:07', '2023-07-30 14:57:07', NULL, 6, NULL, NULL, NULL, NULL, 'asd', 'dasd', 'dfjfgj', 'Laki-Laki', 'wywry', 'sdgsdg');

-- --------------------------------------------------------

--
-- Table structure for table `role`
--

CREATE TABLE `role` (
  `role_id` int(1) NOT NULL,
  `role_name` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `role`
--

INSERT INTO `role` (`role_id`, `role_name`) VALUES
(1, 'super'),
(2, 'admin'),
(3, 'user');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auth`
--
ALTER TABLE `auth`
  ADD PRIMARY KEY (`auth_id`);

--
-- Indexes for table `master_pengajuan`
--
ALTER TABLE `master_pengajuan`
  ADD PRIMARY KEY (`id_jenis_pengajuan`);

--
-- Indexes for table `pengajuan_dokumen`
--
ALTER TABLE `pengajuan_dokumen`
  ADD PRIMARY KEY (`id_pengajuan`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auth`
--
ALTER TABLE `auth`
  MODIFY `auth_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `master_pengajuan`
--
ALTER TABLE `master_pengajuan`
  MODIFY `id_jenis_pengajuan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `pengajuan_dokumen`
--
ALTER TABLE `pengajuan_dokumen`
  MODIFY `id_pengajuan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
