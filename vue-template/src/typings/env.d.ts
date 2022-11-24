/**
 *后台服务的环境类型
 * - dev: 后台开发环境
 * - test: 后台测试环境
 * - prod: 后台生产环境
 */
type ServiceEnvType = "dev" | "test" | "prod";

/** 后台服务的环境配置 */
interface ServiceEnvConfig {
    /** 请求地址 */
    url: string;
    /** 匹配路径的正则字符串, 用于拦截地址转发代理(任意以 /开头 + 字符串, 单个/不起作用) */
    urlPattern: string;
    /** 另一个后端请求地址(有多个不同的后端服务时) */
    secondUrl: string;
    /** 匹配路径的正则字符串, 用于拦截地址转发代理(任意以 /开头 + 字符串, 单个/不起作用) */
    secondUrlPattern: string;
}
