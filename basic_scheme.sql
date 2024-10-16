-- stage.users definition

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `age` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- stage.days definition

CREATE TABLE stage.days (
	id INT auto_increment NOT NULL,
	current_timestamp TIMESTAMP NULL,
	CONSTRAINT days_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

-- stage.event definition

CREATE TABLE stage.event (
	id INT auto_increment NOT NULL,
	label varchar(100) DEFAULT '""' NULL,
	day_id varchar(100) NULL,
	CONSTRAINT event_pk PRIMARY KEY (id),
	CONSTRAINT event_days_FK FOREIGN KEY (id) REFERENCES stage.days(id) ON UPDATE CASCADE
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS stage.event;