DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255),
  `detail` VARCHAR(255),
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);