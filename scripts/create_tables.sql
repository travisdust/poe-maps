CREATE TABLE `map` (
   `id` int(11) unsigned NOT NULL,
   `name` varchar(255) NOT NULL,
   `tier` smallint(6) DEFAULT NULL,
   `unique` tinyint(4) NOT NULL DEFAULT '0',
   PRIMARY KEY (`id`)
);

CREATE TABLE `run` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
   `map_id` int(11) unsigned NOT NULL,
   PRIMARY KEY (`id`),
   KEY `fk_map_id_idx` (`map_id`),
   CONSTRAINT `fk_run_map_id` FOREIGN KEY (`map_id`) REFERENCES `map` (`id`)
);

CREATE TABLE `drop` (
    `drop_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `run_id` int(10) unsigned NOT NULL,
    `map_id` int(10) unsigned NOT NULL,
    PRIMARY KEY (`drop_id`),
    KEY `fk_run_id_idx` (`run_id`),
    KEY `fk_drop_map_id_idx` (`map_id`),
    CONSTRAINT `fk_drop_map_id` FOREIGN KEY (`map_id`) REFERENCES `map` (`id`),
    CONSTRAINT `fk_drop_run_id` FOREIGN KEY (`run_id`) REFERENCES `run` (`id`)
);
