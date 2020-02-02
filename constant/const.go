package constant

// 配置名常量
const (
	IS_RELOAD_CONF                       = "config_auto_reload"          // 自动加载配置文件
	RELOAD_CONF_INTERVAL                 = "config_auto_reload_interval" // 自动加载配置文件间隔，单位：毫秒
	CONF_PATH                            = "config_path"                 // 配置文件地址
	IS_MANAGE_HTTP                       = "http_manage_server"          // 启动管理后台
	MANAGE_HTTP_SERVER_PORT              = "http_manage_server_port"     // 管理后台服务端口
	LOG_PATH                             = "log_path"                    // 日志写入目录
	LOG_SAVE_INTERVAL                    = "log_interval"                // 日志保存间隔，例如: 每2天对久日志压缩，日志写入新日志中
	IS_START_MC_GUI                      = "server_gui"                  // 启动gui
	WORKSPACE                            = "workspace"                   // 工作目录
	I18N                                 = "i18n"                        // 国际化
	IS_AUTO_CHANGE_MC_SERVER_REPEAT_PORT = "mc_server_port_auto_change"  // 自动更换mc服务器重复使用的端口
)

// 配置说明
const (
	IS_RELOAD_CONF_DESCREPTION          = "自动加载配置文件"
	RELOAD_CONF_INTERVAL_DESCREPTION    = "自动加载配置文件间隔，单位：毫秒"
	CONF_PATH_DESCREPTION               = "配置文件地址"
	IS_MANAGE_HTTP_DESCREPTION          = "启动管理后台"
	MANAGE_HTTP_SERVER_PORT_DESCREPTION = "管理后台服务端口"
	LOG_PATH_DESCREPTION                = "日志写入目录"
	LOG_SAVE_INTERVAL_DESCREPTION       = "日志保存间隔，例如: 每2天对久日志压缩，日志写入新日志中"
	LOG_SHOW_CODELINE_DESCREPTION       = "日志打印代码位置"
	IS_START_MC_GUI_DESCREPTION         = "启动gui"
	WORKSPACE_DESCREPTION               = "工作目录"
	I18N_DESCREPTION                    = "国际化"
)

// 配置常量
const (
	RELOAD_CONF_JOB_NAME = "reload-conf"
)

// 日志常量
const (
	DEFAULT_CHANNEL               = "default"
	EVERYDAY_JOB_NAME             = "everyday-add-log"
	LOG_SAVE_INTERVAL_EVERYDAY    = "0 0 * * *"   // cron每日表达式
	LOG_SAVE_INTERVAL_TWICEDAY    = "0 0 1/2 * ?" // cron每隔2天表达式
	LOG_SAVE_INTERVAL_EVERYMONDAY = "0 0 * * 0"   // cron每周一表达式
	LOG_SAVE_INTERVAL_EVERYMONTH  = "0 0 1 * ?"   // cron每月1日表达式
	LOG_DEBUG                     = "debug"
	LOG_INFO                      = "info"
	LOG_ERROR                     = "error"
	LOG_WARNING                   = "warning"
	LOG_FATAL                     = "fatal"
	LOG_FORMAT                    = "%s [%s]: %s\r\n"
	LOG_CODELINE_FORMAT           = "%s [%s] %s : %s\r\n"
	LOG_TIME_FORMAT               = "2006-01-02 15:04:05.000000"
	// 配置覆盖优先级
	CONF_TERMINAL_LEVEL    = 3
	CONF_ENVIRONMENT_LEVEL = 2
	CONF_FILE_LEVEL        = 1
	CONF_DEFAULT_LEVEL     = 0
)

// mc服务端常量
const (
	EULA_FILE_NAME   = "eula.txt"
	EULA             = "eula"
	TRUE_STR         = "true"
	MC_CONF_NAME     = "server.properties"
	MAX_PLAYER_PARAM = "max-players"
	PORT             = "server-port"
)
