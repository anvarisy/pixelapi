-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 28, 2021 at 02:46 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pixelapi`
--

-- --------------------------------------------------------

--
-- Table structure for table `casbin_rule`
--

CREATE TABLE `casbin_rule` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `casbin_rule`
--

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
(6, 'g', 'jerapahku', 'Buyer', '', '', '', ''),
(5, 'g', 'koala', 'Admin', '', '', '', ''),
(1, 'p', 'Admin', 'Admin', 'read', '', '', ''),
(2, 'p', 'Admin', 'Admin', 'write', '', '', ''),
(4, 'p', 'Buyer', 'Buyer', 'read', '', '', ''),
(3, 'p', 'Buyer', 'Buyer', 'write', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `stuffs`
--

CREATE TABLE `stuffs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `stuff_name` longtext DEFAULT NULL,
  `stuff_price` bigint(20) DEFAULT NULL,
  `stuff_stock` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stuffs`
--

INSERT INTO `stuffs` (`id`, `stuff_name`, `stuff_price`, `stuff_stock`, `created_at`, `updated_at`, `deleted`) VALUES
(1, 'Indomie', 2700, 20, '2021-10-27 22:09:54.326', '2021-10-27 22:09:54.326', '2021-10-27 22:13:33.301'),
(2, 'Sarimie', 2500, 30, '2021-10-27 22:10:17.842', '2021-10-27 22:10:17.842', '2021-10-27 22:19:28.616'),
(3, 'Sedapmie', 2600, 40, '2021-10-27 22:10:31.431', '2021-10-27 22:10:31.431', '2021-10-27 22:13:33.416'),
(4, 'Supermie', 2400, 30, '2021-10-27 22:10:57.438', '2021-10-27 22:20:20.114', NULL),
(5, 'Pop Mie Soto', 5000, 5, '2021-10-28 06:17:53.819', '2021-10-28 06:17:53.819', NULL),
(6, 'Pop Mie Ayam Bawang', 5200, 3, '2021-10-28 06:18:11.162', '2021-10-28 06:18:11.162', NULL),
(7, 'Pop Mie Bakso', 5300, 10, '2021-10-28 06:18:21.973', '2021-10-28 06:18:21.973', NULL),
(8, 'Pop Mie Kari', 5400, 7, '2021-10-28 06:18:42.998', '2021-10-28 06:18:42.998', NULL),
(9, 'Pop Mie Goreng', 5600, 12, '2021-10-28 06:19:00.772', '2021-10-28 06:19:00.772', NULL),
(10, 'Pop Mie Original', 5600, 2, '2021-10-28 06:19:15.950', '2021-10-28 06:19:15.950', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` varchar(16) DEFAULT NULL,
  `amount` bigint(20) DEFAULT NULL,
  `transaction_date` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_paid` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `amount`, `transaction_date`, `created_at`, `updated_at`, `deleted_at`, `is_paid`) VALUES
(1, 'jerapahku', 5600, '2021-10-28 06:21:39.000', '2021-10-28 06:21:39.000', '2021-10-28 06:21:39.000', NULL, 0),
(2, 'jerapahku', 16000, '2021-10-28 06:40:22.000', '2021-10-28 06:40:22.000', '2021-10-28 06:40:22.000', NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `transaction_id` bigint(20) UNSIGNED DEFAULT NULL,
  `stuff_id` bigint(20) UNSIGNED DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `transaction_id`, `stuff_id`, `count`) VALUES
(1, 1, 10, 1),
(2, 2, 10, 1),
(3, 2, 6, 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `username` varchar(16) NOT NULL,
  `user_password` longtext NOT NULL,
  `user_fullname` longtext DEFAULT NULL,
  `user_mobile` longtext DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`username`, `user_password`, `user_fullname`, `user_mobile`, `is_admin`, `created_at`, `updated_at`, `deleted`) VALUES
('jerapahku', '$2a$10$IiOjKYPt3/ewCfQ5hwmV5O6DGoUxays4AUi/9lLQQrQ5ube55OmxS', 'Jerapah ku', '6285811521000', 0, '2021-10-27 20:26:16.583', '2021-10-27 20:26:16.583', NULL),
('koala', '$2a$10$yvEbVGSa8HMfN9XjjYYOleNHGy.tfEUwFByo9JS8nfWHGNVHKi4Ee', 'Koala Panda', '6285219529352', 1, '2021-10-27 20:18:48.652', '2021-10-27 20:18:48.652', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `casbin_rule`
--
ALTER TABLE `casbin_rule`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`);

--
-- Indexes for table `stuffs`
--
ALTER TABLE `stuffs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_stuffs_deleted` (`deleted`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transactions_deleted_at` (`deleted_at`),
  ADD KEY `fk_transactions_u` (`user_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_transaction_details_s` (`stuff_id`),
  ADD KEY `fk_transactions_detail` (`transaction_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`username`),
  ADD KEY `idx_users_deleted` (`deleted`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `casbin_rule`
--
ALTER TABLE `casbin_rule`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `stuffs`
--
ALTER TABLE `stuffs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_transactions_u` FOREIGN KEY (`user_id`) REFERENCES `users` (`username`) ON DELETE SET NULL ON UPDATE CASCADE;

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `fk_transaction_details_s` FOREIGN KEY (`stuff_id`) REFERENCES `stuffs` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_transactions_detail` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
