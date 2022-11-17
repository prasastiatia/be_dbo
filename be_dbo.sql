-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 17, 2022 at 12:35 PM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `be_dbo`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `tanggal_lahir` longtext DEFAULT NULL,
  `mobile` longtext DEFAULT NULL,
  `alamat` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `tanggal_lahir`, `mobile`, `alamat`) VALUES
(1, '2022-11-15 19:12:14.413', '2022-11-15 19:12:14.413', NULL, 'Prasastia Aryani Saliha', '15-08-1997', '081249263092', 'Depok'),
(2, '2022-11-15 19:34:49.472', '2022-11-15 19:34:49.472', '2022-11-16 03:29:48.287', 'Safina A. F', '20-02-2003', '088612342345', 'Depok'),
(4, '2022-11-16 03:31:59.809', '2022-11-16 03:31:59.809', NULL, 'Safina A. Febriani', '20-02-2003', '0886123423451', 'Depok'),
(5, '2022-11-16 03:33:00.711', '2022-11-16 03:33:00.711', NULL, 'Fina', '15-07-1967', '082134516723', 'Depok 2'),
(6, '2022-11-16 03:33:22.186', '2022-11-16 03:33:22.186', NULL, 'Testing', '16-07-1967', '082134516744', 'Depok 2'),
(7, '2022-11-16 03:33:39.169', '2022-11-16 03:33:39.169', NULL, 'Testingc 2', '20-07-1967', '082134516776', 'Depok 1');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `order_id` bigint(20) UNSIGNED NOT NULL,
  `customer_name` longtext DEFAULT NULL,
  `ordered_at` datetime(3) DEFAULT NULL,
  `items` longtext DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `total_bayar` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `created_at`, `updated_at`, `deleted_at`, `order_id`, `customer_name`, `ordered_at`, `items`, `quantity`, `total_bayar`) VALUES
(1, '2022-11-16 15:38:22.988', '2022-11-16 16:02:46.357', NULL, 0, 'Prasastia Aryani Saliha', '0000-00-00 00:00:00.000', 'Celana Legging', 4, 160000),
(2, '2022-11-16 16:03:39.210', '2022-11-16 16:03:39.210', '2022-11-16 16:05:02.969', 0, 'Prasastia Aryani Saliha', '0000-00-00 00:00:00.000', 'Shirt', 3, 500000),
(3, '2022-11-16 16:04:16.870', '2022-11-16 16:04:16.870', NULL, 0, 'Safina A. Febriani', '0000-00-00 00:00:00.000', 'Kemeja', 1, 250000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `username`, `email`, `password`) VALUES
(1, '2022-11-15 18:27:02.121', '2022-11-15 18:27:02.121', NULL, 'Prasastia Aryani Saliha', 'prasastiatia', 'prasastia@mail.com', '$2a$14$cXHU8.k1pIzSNYHAy77d1ejKBcCrXxScaR0ziVlTJkcT5JN85K4Pe');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `mobile` (`mobile`) USING HASH,
  ADD KEY `idx_customers_deleted_at` (`deleted_at`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`,`order_id`),
  ADD KEY `idx_orders_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
