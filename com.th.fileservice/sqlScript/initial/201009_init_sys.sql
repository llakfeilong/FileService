DROP TABLE IF EXISTS `file_info`;
CREATE TABLE `file_info` (
  `file_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件编号',
  `file_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '获取文件key',
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `file_original_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '原文件名',
  `id` bigint NOT NULL AUTO_INCREMENT,
  `file_suffix_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件后缀名',
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '内部存储路径',
  `file_size` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件大小',
  `file_type_code` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件类型代码',
  `file_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `update_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编辑人',
  `update_date` datetime DEFAULT NULL COMMENT '编辑时间',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  `create_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `deleted` varchar(1) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否已删除Y：已删除，N：未删除',
  `delete_date` datetime DEFAULT NULL COMMENT '删除时间',
  `delete_name` varchar(40) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `file_type_manager`;
CREATE TABLE `file_type_manager` (
  `file_type_id` bigint NOT NULL AUTO_INCREMENT COMMENT '文件类型主键',
  `file_type_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件类型名称',
  `file_type_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件类型代码唯一键\r\n作为存储当前类型文件的文件夹名称',
  `file_format` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件格式',
  `file_max_size` int DEFAULT NULL COMMENT '最大文件限制（KB）',
  `file_expiry_date` int DEFAULT NULL COMMENT '文件有效期（自然月）',
  `enable` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否启用该文件类型 0 启用   1不启用',
  `update_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编辑人',
  `update_date` datetime DEFAULT NULL COMMENT '编辑时间',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  `create_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `deleted` varchar(1) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '是否已删除Y：已删除，N：未删除',
  `delete_date` datetime DEFAULT NULL COMMENT '删除时间',
  `delete_name` varchar(40) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`file_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `tenfine_fileservice`.`file_type_manager`(`file_type_id`, `file_type_name`, `file_type_code`, `file_format`, `file_max_size`, `file_expiry_date`, `enable`, `update_name`, `update_date`, `create_date`, `create_name`) VALUES (1, '默认照片类型', 'tenfineImages', '.jpg.jpeg.png.bmp', 30720, NULL, '0', NULL, NULL, NULL, NULL);
INSERT INTO `tenfine_fileservice`.`file_type_manager`(`file_type_id`, `file_type_name`, `file_type_code`, `file_format`, `file_max_size`, `file_expiry_date`, `enable`, `update_name`, `update_date`, `create_date`, `create_name`) VALUES (2, '默认文件类型', 'tenfineFiles', '.word.pdf.rar.zip.doc.docx', 30720, NULL, '0', NULL, NULL, NULL, NULL);




