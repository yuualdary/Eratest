-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 19, 2022 at 05:13 PM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 8.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `eratest`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_name` longtext DEFAULT NULL,
  `product_price` bigint(20) DEFAULT NULL,
  `product_description` longtext DEFAULT NULL,
  `product_quantity` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `product_price`, `product_description`, `product_quantity`) VALUES
('30e868e4-2626-4020-9fe3-a65d68e3a0da', '2022-01-19 17:57:01.001', '2022-01-19 17:57:01.001', NULL, 'Samsung FE 2021', 110000, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 100),
('3465859b-519d-4323-9ef3-8e71ba74c42e', '2022-01-19 22:28:16.587', '2022-01-19 22:28:16.587', NULL, 'Iphone SE 2020', 310000, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 13),
('3738ed51-ab36-4337-adfb-e2309e90dd3e', '2022-01-19 17:11:02.727', '2022-01-19 17:11:02.727', NULL, 'Hanphone XR', 10, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 1),
('4ed9abd6-7976-4e42-93b4-9a986acfc3aa', '2022-01-19 22:30:20.775', '2022-01-19 22:30:20.775', NULL, 'Asus ROG', 522355, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 1),
('69dfe47b-456b-47ce-9f09-7eceba8aabaa', '2022-01-19 22:29:27.124', '2022-01-19 22:29:27.124', NULL, 'Redmi Note', 12345, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 9),
('723c7c21-11b2-4518-afcc-a864b6335de6', '2022-01-19 17:54:31.417', '2022-01-19 17:54:31.417', NULL, 'Note 10', 130000, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 100),
('72dd12bc-8b6b-403b-9998-995d3dff6b66', '2022-01-19 22:29:02.578', '2022-01-19 22:29:02.578', NULL, 'Sony AAA', 12340, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 11),
('80c0c36b-cc79-471b-988f-34991c748aa6', '2022-01-19 22:29:42.692', '2022-01-19 22:29:42.692', NULL, 'Oppo F1', 22355, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 9),
('f212a1af-96e6-4f41-803e-d3aee3a46a05', '2022-01-19 22:30:10.769', '2022-01-19 22:30:10.769', NULL, 'Sony Camera', 222355, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to ', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_products_deleted_at` (`deleted_at`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
