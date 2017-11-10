


CREATE TABLE `recipe`.`recipe` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` TINYTEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC));


CREATE TABLE `recipe`.`recipe_recipe` (
  `parent_id` INT UNSIGNED NOT NULL,
  `child_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`parent_id`, `child_id`));
