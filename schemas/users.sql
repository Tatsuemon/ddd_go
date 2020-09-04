CREATE TABLE `users` (
    `id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーID',
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '名前',
    `email` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'メールアドレス',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;