CREATE TABLE `exercises` (
  `id`       INT AUTO_INCREMENT NOT NULL,
  `name`     VARCHAR(255) NOT NULL UNIQUE,
  PRIMARY KEY (`id`)
);

CREATE TABLE `workouts` (
  `id` INT AUTO_INCREMENT NOT NULL,
  `exercise_id` INT NOT NULL,
  `date` DATE NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`exercise_id`) REFERENCES `exercises`(`id`)
);

CREATE TABLE `sets` (
  `id`       INT AUTO_INCREMENT NOT NULL,
  `workout_id` INT NOT NULL,
  `set_num`  INT NOT NULL,
  `weight` FLOAT NOT NULL,
  `reps` INT DEFAULT 0,
  `preps` INT DEFAULT 0,
  `rir` INT DEFAULT 0,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`workout_id`) REFERENCES `workouts`(`id`)
);
  
