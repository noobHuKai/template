/** 请求服务的环境配置 */
type ServiceEnv = Record<ServiceEnvType, ServiceEnvConfig>;

/** 不同请求服务的环境配置 */
const serviceEnv: ServiceEnv = {
  dev: {
    url: 'http://172.19.210.99:8000/api/v1/system',
    urlPattern: '/api',
    secondUrl: 'http://172.28.9.179:8000',
    secondUrlPattern: '/second-url-pattern'
  },
  test: {
    url: 'http://localhost:8080',
    urlPattern: '/url-pattern',
    secondUrl: 'http://localhost:8081',
    secondUrlPattern: '/second-url-pattern'
  },
  prod: {
    url: 'http://172.28.9.179:8000',
    urlPattern: '/api/v1/system',
    secondUrl: 'http://localhost:8081',
    secondUrlPattern: '/second-url-pattern'
  }
};

/**
 * 获取当前环境模式下的请求服务的配置
 * @param env 环境
 */
export function getServiceEnvConfig(env: ImportMetaEnv) {
  const { VITE_SERVICE_ENV = 'dev' } = env;

  const config = serviceEnv[VITE_SERVICE_ENV];

  return config;
}
