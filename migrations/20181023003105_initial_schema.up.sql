CREATE TABLE `players`
(
	`id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`name` varchar(255),
	`overall_rating` decimal(6, 2) NOT NULL DEFAULT 1200.00,
	`overall_wins` int NOT NULL DEFAULT 0,
	`overall_losses` int NOT NULL DEFAULT 0,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp,
	`updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE `seasons`
(
	`id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`name` varchar(255),
	`start_date` timestamp,
	`end_date` timestamp DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp,
	`updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);


CREATE TABLE `seasonal_stats`
(
	`id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`player_id` int NOT NULL,
	`season_id` int NOT NULL,
	`rating` decimal(6, 2) NOT NULL,
	`wins` int NOT NULL DEFAULT 0,
	`losses` int NOT NULL DEFAULT 0,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp,
	`updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
	FOREIGN KEY fk_player_ss(`player_id`) REFERENCES players(`id`),
	FOREIGN KEY fk_season_ss(`season_id`) REFERENCES seasons(`id`)
);

CREATE TABLE `matches`
(
	`id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`season_id` int NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp,
	`updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
	FOREIGN KEY fk_season_m(`season_id`) REFERENCES seasons(`id`)
);

CREATE TABLE `match_competitors`
(
	`id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`match_id` int NOT NULL,
	`player_id` int NOT NULL,
	`starting_rating` decimal(6, 2) NOT NULL,
	`rating_change` decimal(6, 2) NOT NULL,
	`winner` boolean NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT current_timestamp,
	`updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
	FOREIGN KEY fk_match_mc(`match_id`) REFERENCES matches(`id`),
	FOREIGN KEY fk_player_mc(`player_id`) REFERENCES players(`id`)
);
