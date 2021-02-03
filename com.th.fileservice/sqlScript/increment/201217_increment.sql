ALTER TABLE `th_file_down_upload_service`.`file_info`
MODIFY COLUMN `id` bigint(0) NOT NULL AUTO_INCREMENT AFTER `file_original_name`;


ALTER TABLE `th_file_down_upload_service`.`file_type_manager`
    MODIFY COLUMN `file_type_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文件类型主键' FIRST;

ALTER TABLE `th_file_down_upload_service`.`file_info`
ADD COLUMN `deleted` varchar(1) NULL COMMENT '是否已删除Y：已删除，N：未删除' AFTER `create_name`,
ADD COLUMN `delete_date` datetime NULL COMMENT '删除时间' AFTER `deleted`,
ADD COLUMN `delete_name` varchar(40) NULL COMMENT '删除人' AFTER `delete_date`;

ALTER TABLE `th_file_down_upload_service`.`file_type_manager`
ADD COLUMN `deleted` varchar(1) NULL COMMENT '是否已删除Y：已删除，N：未删除' AFTER `create_name`,
ADD COLUMN `delete_date` datetime NULL COMMENT '删除时间' AFTER `deleted`,
ADD COLUMN `delete_name` varchar(40) NULL COMMENT '删除人' AFTER `delete_date`;