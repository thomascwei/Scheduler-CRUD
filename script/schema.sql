CREATE TABLE `schedules`
(
    `id`               int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `time_type_id`     int             NOT NULL COMMENT 'template type',
    `interval_day`     int             NOT NULL DEFAULT -1,
    `interval_seconds` int             NOT NULL DEFAULT -1 COMMENT 'repeat in day',
    `at_time`          varchar(255)    NOT NULL DEFAULT '00:00:00',
    `start_time`       varchar(255)    NOT NULL DEFAULT '00:00:00',
    `end_time`         varchar(255)    NOT NULL DEFAULT '00:00:00',
    `command_id`       int             NOT NULL,
    `name`             varchar(255)    NOT NULL COMMENT 'template name',
    `start_date`       datetime        NOT NULL DEFAULT '0000-00-00 00:00:00',
    `end_date`         datetime        NOT NULL DEFAULT '0000-00-00 00:00:00',
    `enable`           boolean         NOT NULL DEFAULT true,
    `repeat`           boolean         NOT NULL DEFAULT false,
    `create_time`      datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `repeat_day`
(
    `id`          int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `schedule_id` int             NOT NULL,
    `day`         varchar(255)    NOT NULL,
    `create_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `repeat_month`
(
    `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `schedule_id` int          NOT NULL,
    `month`       varchar(255) NOT NULL,
    `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `repeat_weekday`
(
    `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `schedule_id` int          NOT NULL,
    `weekday`     varchar(255) NOT NULL,
    `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `commands`
(
    `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `command`     varchar(255)    NOT NULL,
    `create_time` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `schedules`
    ADD FOREIGN KEY (`command_id`) REFERENCES `commands` (`id`);

ALTER TABLE `repeat_day`
    ADD CONSTRAINT FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`);

ALTER TABLE `repeat_month`
    ADD CONSTRAINT FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`);

ALTER TABLE `repeat_weekday`
    ADD CONSTRAINT FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`);
