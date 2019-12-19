/*
* 创建用户数据表
 */
CREATE TABLE  IF NOT EXISTS `user_info`(
  `id`  int(11) unsigned not null auto_increment comment '主键，自增',
  `user_id` varchar(40) not null comment '用户登录号',
  `passwd` varchar(40) not null comment '用户登录密码',
  `nick` varchar(30) not null default "kity" comment '用户昵称',
  `nick` varchar(200) comment '备注',
  `ip` varchar(30) not null default "127.0.0.1" commit '用户当前ip',
  `status` varchar(10) not null default "active" comment '该用户的状态 active、unactive、delete',
  `role` varchar(100) not null default "nothing" comment '管理者分配的角色',
  `role_name` varchar(200) not null default "普通操作员" comment '操作员分配的角色名称',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp on update current_timestamp comment '最后一次修改时间',
  primary key (`id`),
  unique key `u_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';
insert into user_info (user_id, passwd, nick, status, role) values ("10086", "FFB23F80E5F0DA11ED14BA13FCF528DD", "admin", "active", "nothing");

/*
* 创建一级菜单表
 */
CREATE TABLE IF NOT EXISTS `menu_info`(
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `menu_order` int(5) unsigned not null default 0 comment '一级菜单的排名顺序',
  `menu_uid` varchar(40) not null comment '一级菜单的唯一标识',
  `first_menu` varchar(50) not null comment '一级菜单名称，字符不能超过50',
  `second_menu` text comment '二级菜单名称，每个之间用|隔开',
  `creater` varchar(20) not null comment '创建者的id',
  `status` varchar(10) not null default "active" comment '菜单的状态情况，默认是active',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp comment '最近更新时间',
  primary key (`id`),
  unique key `u_first_menu` (`first_menu`),
  unique key `u_menu_uid` (`menu_uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放左侧栏的菜单';

/*
* 创建一级菜单表
 */
CREATE TABLE IF NOT EXISTS `second_menu_info`(
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `first_menu_order` int(5) unsigned not null default 0 comment '一级菜单对应的顺序',
  `menu_order` int(5) unsigned not null default 0 comment '二级菜单的排名顺序',
  `first_menu_uid` varchar(40) not null comment '二级菜单的唯一标识',
  `first_menu` varchar(50) not null comment '一级菜单名称，字符不能超过50',
  `second_menu_uid` varchar(40) not null comment '二级菜单唯一标识',
  `second_menu` varchar(225) not null comment '二级菜单名称',
  `second_router` varchar(200) not null comment '二级菜单路由',
  `creater` varchar(20) not null comment '创建者的id',
  `status` varchar(10) not null default "active" comment '菜单的状态情况，默认是active',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp comment '最近更新时间',
  primary key (`id`),
  unique key `u_second_menu` (`second_menu`),
  unique key `u_second_menu_uid` (`second_menu_uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放左侧栏的二级菜单';

/*
* 创建权限项表
 */
CREATE TABLE IF NOT EXISTS `power_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `first_menu_uid` varchar(40) not null comment '一级菜单的唯一标识',
  `second_menu_uid` varchar(40) not null comment '二级菜单的唯一标识',
  `second_menu` varchar(50) not null comment '二级菜单的名称',
  `power_item` varchar(50) not null comment '权限项的名称',
  `power_id` varchar(200) not null comment '权限的ID',
  `creater` varchar(20) not null comment '创建者的id',
  `status` varchar(10) not null default "active" comment '菜单的状态情况，默认是active',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp comment '最近更新时间',
  primary key (`id`),
  unique key `u_power_id` (`power_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='存放控制页面的一些功能操作';

/*
* 创建角色
 */
CREATE TABLE IF NOT EXISTS `role_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `role_name` varchar(100) not null comment '角色名称',
  `role_uid` varchar(200) not null comment '角色唯一标识号',
  `show_first_menu` text not null comment '可以展示的一级菜单名',
  `show_first_uid` text not null comment '可以展示的一级菜单uid',
  `show_second_menu` text not null comment '可以展示的二级菜单名',
  `show_second_uid` text not null comment '可以展示的二级菜单uid',
  `show_power` text not null comment '可以展示的权限项名称',
  `show_power_uid` text not null comment '可以展示的权限项uid',
  `remark` text not null comment '角色描述',
  `creater` varchar(20) not null comment '创建者的id',
  `status` varchar(10) not null default "active" comment '菜单的状态情况，默认是active',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp comment '最近更新时间',
  primary key (`id`),
  unique key `u_power_name` (`role_name`),
  unique key `u_role_uid` (`role_uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

/*
* 银行卡管理表
 */
CREATE TABLE IF NOT EXISTS `bank_card_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `uid` varchar(100) not  null comment '唯一标识',
  `user_name` varchar(100) not null comment '用户名称',
  `bank_name` varchar(100) not null comment '银行名称',
  `bank_code` varchar(30) not null comment '银行编码',
  `bank_account_type` varchar(20) not null comment '银行账号类型',
  `account_name` varchar(50) not null comment '银行账户名称',
  `bank_no` varchar(50) not null comment '银行账号',
  `identify_card` varchar(100) not null comment '证件类型',
  `certificate_no` varchar(100) not null comment '证件号码',
  `phone_no` varchar(50) not null comment '手机号码',
  `bank_address` varchar(200) not null comment '银行地址',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp comment '最近更新时间',
  primary key (`id`),
  unique `uid` (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='银行卡表';

/*
* 通道数据表
 */
CREATE TABLE IF NOT EXISTS `road_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" '通道状态',
  `road_name` varchar(100) not null comment '通道名称',
  `road_uid` varchar(100) not null comment '通道唯一id',
  `remark` varchar(100) comment '备注',
  `product_name` varchar(100) not null comment '上游产品名称',
  `product_uid` varchar(100) not null  comment '上游产品编号',
  `pay_type` varchar(50) not null  comment '支付类型',
  `basic_fee` double not null comment '基本汇率/成本汇率',
  `settle_fee` double not null comment '代付手续费',
  `total_limit` double not null comment '通道总额度',
  `today_limit` double not null comment '每日最多额度',
  `single_min_limit` double not null comment '单笔最小金额',
  `single_max_limit` double not null comment '单笔最大金额',
  `star_hour` int not null comment '通道开始时间',
  `end_hour` int not null comment '通道结束时间',
  `params` text comment '参数json格式',
  `today_income` decimal(20, 3) not null default 0 comment '当天的收入',
  `total_income` decimal(20, 3) not null default 0 comment '通道总收入',
  `today_profit` decimal(20, 3) not null default 0 comment '当天的收益',
  `total_profit` decimal(20 ,3) not null default 0 comment '通道总收益',
  `balance` decimal(20, 3) not null default 0 comment '通道的余额',
  `request_all` int default 0 comment '请求总次数',
  `request_success` int  default 0 comment '请求成功次数',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `road_name` (`road_name`),
  unique `road_uid` (`road_uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通道数据表';

/*
* 通道池数据表
 */
CREATE TABLE IF NOT EXISTS `road_pool_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" comment '通道池状态',
  `road_pool_name` varchar(100) not null comment '通道池名称',
  `road_pool_code` varchar(100) not null comment '通道池编号',
  `road_uid_pool` text comment '通道池里面的通道uid',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `road_pool_name` (`road_pool_name`),
  unique `road_pool_code` (`road_pool_code`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通道池';

/*
* 商户配置表
 */
CREATE TABLE IF NOT EXISTS `merchant_info`(
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" comment '商户状态状态',
  `belong_agent_uid` vachar(100) comment '所属代理uid',
  `belong_agent_name` vachar(100) comment '所属代理名称',
  `merchant_name` varchar(100) not null default '客户' comment '商户名称',
  `merchant_uid` varchar(100) not null comment '商户uid',
  `merchant_key` varchar(100) not null comment '商户key',
  `merchant_secret` varchar(100) not null comment '商户密钥',
  `login_account` varchar(100) not null comment '登录账号',
  `login_password` varchar(100) not null comment '登录密码',
  `auto_settle` varchar(10) not null default "YES" comment '是否自动结算',
  `auto_pay_for` varchar(10) not null default "YES" comment '是否自动代付',
  `white_ips` text comment '配置ip白名单',
  `remark` text comment '备注',
  `single_pay_for_road_uid` varchar(100) comment '单代付代付通道uid',
  `single_pay_for_road_name` varchar(200) comment '单代付通道名称',
  `roll_pay_for_road_code` varchar(100) comment '轮询代付通道编码',
  `roll_pay_for_road_name` varchar(200) comment '轮询代付通道名称',
  `payfor_fee` double comment '代付手续费',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `merchant_uid` (`merchant_uid`),
  unique `merchant_name` (`merchant_name`),
  unique `merchant_key` (`merchant_key`),
  unique `merchant_secret` (`merchant_secret`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商户支付配置表';

/*
* 商戶对应通道表
 */
CREATE TABLE IF NOT EXISTS `merchant_deploy_info`(
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" comment '商户状态状态',
  `merchant_uid` varchar(100) not null comment '商户uid',
  `pay_type` varchar(50) comment '支付配置',
  `single_road_uid` varchar(100) comment '单通道uid',
  `single_road_name` varchar(200) comment '单通道名称',
  `single_road_platform_rate` decimal(20,3) not null default 0.000 comment '单通到平台净利率',
  `single_road_agent_rate` decimal(20,3) not null default 0.000 comment '单通到代理净利率',
  `roll_road_code` varchar(100) comment '轮询通道编码',
  `roll_road_name` varchar(200) comment '轮询通道名称',
  `roll_road_platform_rate` decimal(20,3) not null default 0.000 comment '轮询通道平台净利率',
  `roll_road_agent_rate` decimal(20,3) not null default 0.000 comment '轮询通道代理净利率',
  `is_loan` varchar(10) not null default "NO" comment '是否押款',
  `loan_rate` decimal(20,3) not null default 0.000 comment '押款比例，默认是0',
  `loan_day` int not null default 0 comment '押款的天数，默认0天',
  `unfreeze_hour` int not null default 0 comment '每天解款的时间点，默认是凌晨',
  `wait_unfreeze_amount` decimal(20,3) comment '等待解款的金额',
  `loan_amount` decimal(20,3) comment '押款中的金额',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商户通道配置；\r\n单通道给商户的汇率=single_road_platform_rate+single_road_agent_rate+basic_fee；\r\n轮询通道汇率=roll_road_platform_rate+roll_road_agent_rate+basic_fee；';

/*
* 商户对应的每条通道的押款
 */
 CREATE TABLE IF NOT EXISTS `merchant_load_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default 'no' comment 'no-没有结算，yes-结算',
  `merchant_uid` varchar(100) not null comment '商户uid',
  `road_uid` varchar(50) not null comment '通道uid',
  `load_date` varchar(50) not null comment '押款日期，格式2019-10-10',
  `load_amount` decimal(20,3) not null default 0 comment '押款金额',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商户对应的每条通道的押款信息';

/*
* 账户表，记录商户和代理的资金情况
 */
CREATE TABLE IF NOT EXISTS `account_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" comment '状态',
  `account_uid` varchar(100) not null comment '账户uid，对应为merchant_uid或者agent_uid',
  `account_name` varchar(100) not null comment '账户名称，对应的是merchant_name或者agent_name',
  `balance` decimal(20, 3) not null default 0 comment '账户余额',
  `settle_amount` decimal(20,3) not null default 0 comment '已经结算了的金额',
  `loan_amount` decimal(20,3) not null default 0 comment '押款金额',
  `wait_amount` decimal(20,3) not null default 0 comment '待结算资金',
  `freeze_amount` decimal(20,3) not null default 0 comment '账户冻结金额',
  `payfor_amount` decimal(20,3) not null default 0 comment '账户代付中金额',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `account_uid` (`account_uid`),
  unique `account_name` (`account_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='账户记录表';

/*
* 账户动账表,可以追溯每笔资金的动向
 */
CREATE TABLE IF NOT EXISTS `account_history_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `account_uid` varchar(100) not null comment '账号uid',
  `account_name` varchar(100) not null comment '账户名称',
  `type` varchar(20) not null default "" comment '减款，加款',
  `amount` decimal(20,3) not null default 0 comment '操作对应金额对应的金额',
  `balance` decimal(20,3) not null default 0 comment '操作后的当前余额',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='账户账户资金动向表';


/*
* 代理表
 */
CREATE TABLE IF NOT EXISTS `agent_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `status` varchar(20) not null default "active" comment '代理状态状态',
  `agent_name` varchar(100) not null comment '代理名称',
  `agent_password` varchar(50) not null comment '代理登录密码',
  `pay_password` varchar(50) not null comment '支付密码',
  `agent_uid` varchar(100) not null comment '代理编号',
  `agent_phone` varchar(15) not null comment '代理手机号',
  `agent_remark` text comment '备注',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `agent_name` (`agent_name`),
  unique `agent_uid` (`agent_uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='代理';

/*
* 订单表
 */
CREATE TABLE IF NOT EXISTS `order_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `merchant_order_id` varchar(50) not null comment '下游商户提交过来的订单id',
  `shop_name` varchar(100) not null comment '商品名称',
  `order_period` varchar(3) not null default "0" comment '订单有效时间，小时制',
  `bank_order_id` varchar(50) not null comment '平台自身的订单id',
  `bank_trans_id` varchar(50) not null comment '上游返回的订单id',
  `order_amount` decimal(20,3) not null default 0 comment '订单提交金额',
  `show_amount` decimal(20,3) not null default 0 comment '展示在用户面前待支付的金额',
  `fact_amount` decimal(20,3) not null default 0 comment '实际支付金额',
  `roll_pool_code` varchar(50) comment '轮询产品编码',
  `roll_pool_name` varchar(100) comment '轮询产品名称',
  `road_uid` varchar(100) not null comment '通道uid',
  `road_name` varchar(200) not null comment '通道名称',
  `pay_product_code` varchar(100) not null comment '支付产品编码',
  `pay_product_name` varchar(200) not null comment '支付产品名称',
  `pay_type_code` varchar(50) not null comment '支付类型编码',
  `pay_type_name` varchar(100) not null comment '支付类型名称',
  `os_type` varchar(5) not null comment '平台类型，苹果app-0， 安卓app-1，苹果H5-3，安卓H5-4，pc-5',
  `status` varchar(20) not null default "wait" comment '等待支付-wait,支付成功-success, 支付失败-failed',
  `refund` varchar(5) not null default "no" comment '退款-yes， 没有退款-no',
  `refund_time` varchar(100) comment '退款时间',
  `freeze` varchar(5) not null default "no" comment '冻结-yes， 没有-no',
  `freeze_time` varchar(100) comment '冻结时间',
  `unfreeze` varchar(5) not null default "no" comment '解冻-yes，没有-no',
  `unfreeze_time` varchar(100) comment '解冻时间',
  `return_url` text comment '订单支付后，跳转的地址',
  `notify_url` text comment '订单回调给下游的地址',
  `merchant_uid` varchar(100) not null comment '商户uid，表示订单是哪个商户的',
  `merchant_name` varchar(200) not null comment '商户名称',
  `agent_uid` varchar(100) comment '代理uid，表示该商户是谁的代理',
  `agent_name` varchar(200) comment '代理名称',
  `response` text comment '上游返回的结果',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `merchant_order_id` (`merchant_order_id`),
  unique `bank_order_id` (`bank_order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';

/*
* 订单利润表
 */
 CREATE TABLE IF NOT EXISTS `order_profit_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `merchant_name` varchar(100) not null comment '商户名称',
  `merchant_uid` varchar(50) not null comment '商户uid',
  `agent_uid` varchar(100) comment '代理uid，表示该商户是谁的代理',
  `agent_name` varchar(200) comment '代理名称',
  `pay_product_code` varchar(100) not null comment '支付产品编码',
  `pay_product_name` varchar(200) not null comment '支付产品名称',
  `pay_type_code` varchar(50) not null comment '支付类型编码',
  `pay_type_name` varchar(100) not null comment '支付类型名称',
  `status` varchar(20) not null default "wait" comment '等待支付-wait,支付成功-success, 支付失败-failed',
  `merchant_order_id` varchar(50) not null comment '下游商户提交过来的订单id',
  `bank_order_id` varchar(50) not null comment '平台自身的订单id',
  `bank_trans_id` varchar(50) not null comment '上游返回的订单id',
  `order_amount` decimal(20,3) not null default 0 comment '订单提交金额',
  `show_amount` decimal(20,3) not null default 0 comment '展示在用户面前待支付的金额',
  `fact_amount` decimal(20,3) not null default 0 comment '实际支付金额',
  `user_in_amount` decimal(20,3) not null default 0 comment '商户入账金额',
  `all_profit` decimal(20,3) not null default 0 comment '总的利润，包括上游，平台，代理',
  `supplier_rate` decimal(20,3) not null default 0 comment '上游的汇率',
  `platform_rate` decimal(20,3) not null default 0 comment '平台自己的手续费率',
  `agent_rate` decimal(20,3) not null default 0 comment '代理的手续费率',
  `supplier_profit` decimal(20,3) not null default 0 comment '上游的利润',
  `platform_profit` decimal(20,3) not null default 0 comment '平台利润',
  `agent_profit` decimal(20, 3) not null default 0 comment '代理利润',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `merchant_order_id` (`merchant_order_id`),
  unique `bank_order_id` (`bank_order_id`)
 )ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单利润表';

/*
* 订单结算表
 */
CREATE TABLE IF NOT EXISTS `order_settle_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `pay_product_code` varchar(100) not null comment '支付产品编码',
  `pay_product_name` varchar(200) not null comment '支付产品名称',
  `pay_type_code` varchar(50) not null comment '支付类型编码',
  `pay_type_name` varchar(100) not null comment '支付类型名称',
  `merchant_uid` varchar(100) not null comment '商户uid，表示订单是哪个商户的',
  `road_uid` varchar(50) not null comment '通道uid',
  `merchant_name` varchar(200) not null comment '商户名称',
  `merchant_order_id` varchar(50) not null comment '下游商户提交过来的订单id',
  `bank_order_id` varchar(50) not null comment '平台自身的订单id',
  `settle_amount` decimal(20,3) not null default 0 comment '结算金额',
  `is_allow_settle` varchar(10) not null default 'yes' comment '是否允许结算，允许-yes，不允许-no',
  `is_complete_settle` varchar(10) not null default 'no' comment '该笔订单是否结算完毕，没有结算-no，结算完毕-yes',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `merchant_order_id` (`merchant_order_id`),
  unique `bank_order_id` (`bank_order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单结算表';

/*
* 回调记录表
 */
CREATE TABLE IF NOT EXISTS `notify_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `type` varchar(10) not null comment '支付订单-order， 代付订单-payfor',
  `bank_order_id` varchar(50) not null comment '系统订单id',
  `merchant_order_id` varchar(50) not null comment '下游商户订单id',
  `status` varchar(20) not null default "wait" comment '状态字段',
  `times` int not null default 0 comment '回调次数',
  `url` text comment '回调的url',
  `response` text comment '回调返回的结果',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `merchant_order_id` (`merchant_order_id`),
  unique `bank_order_id` (`bank_order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='支付回调';

/*
* 代付表
 */
CREATE TABLE IF NOT EXISTS `payfor_info` (
  `id` int(11) unsigned not null auto_increment comment '主键，自增',
  `payfor_uid` varchar(100) not null comment '代付唯一uid',
  `merchant_uid` varchar(100) not null comment '发起代付的商户uid',
  `merchant_name` varchar(200) not null comment '发起代付的商户名称',
  `merchant_order_id` varchar(50) comment '下游代付订单id',
  `bank_order_id` varchar(50) not null comment '系统代付订单id',
  `bank_trans_id` varchar(50) not null comment '上游返回的代付订单id',
  `road_uid` varchar(100) not null comment '所用的代付通道uid',
  `road_name` varchar(200) not null comment '所有通道的名称',
  `roll_pool_code` varchar(100) comment '所用轮询池编码',
  `roll_pool_name` varchar(200) comment '所用轮询池的名称',
  `payfor_fee` decimal(20,3) not null default 0 comment '代付手续费',
  `payfor_amount` decimal(20,3) not null default 0 comment '代付到账金额',
  `payfor_total_amount` decimal(20,3) not null default 0 comment '代付总金额',
  `bank_code` varchar(20) not null comment '银行编码',
  `bank_name` varchar(100) not null comment '银行名称',
  `bank_account_name` varchar(100) not null comment '银行开户名称',
  `bank_account_no` varchar(50) not null comment '银行开户账号',
  `bank_account_type` int not null default 0 comment '银行卡类型，对私-0，对公-1',
  `country` varchar(50) not null default "中国" comment '开户所属国家',
  `province` varchar(50) not null default "" comment '银行卡开户所属省',
  `city` varchar(50) not null default "" comment '银行卡开户所属城市',
  `ares` varchar(50) comment '所属地区',
  `bank_account_address` text comment '银行开户具体街道',
  `phone_no` varchar(20) not null comment '开户所用手机号',
  `give_type` varchar(50) not null default "payfor_road" comment '下发类型，payfor_road-通道打款，payfor_hand-手动打款，payfor_refuse-拒绝打款',
  `type` varchar(10) not null default "auto" comment '代付类型，self_api-系统发下， 管理员手动下发给商户-self_merchant，管理自己提现-self_help',
  `notify_url` text comment '代付结果回调给下游的地址',
  `status` varchar(20) not null default "wait" comment '审核-payfor_confirm,系统处理中-payfor_solving，银行处理中-payfor_banking，代付成功-success, 代付失败-failed',
  `is_send` varchar(10) not null default "no" comment '未发送-no，已经发送-yes',
  `request_time` timestamp not null comment '发起请求时间',
  `response_time` timestamp not null comment '上游做出响应的时间',
  `response_content` text comment '代付的最终结果',
  `remark` text comment '代付备注',
  `update_time` timestamp not null comment '更新时间',
  `create_time` timestamp not null comment '创建时间',
  primary key (`id`),
  unique `payfor_uid` (`payfor_uid`),
  unique `merchant_order_id` (`merchant_order_id`),
  unique `bank_order_id` (`bank_order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='代付表';