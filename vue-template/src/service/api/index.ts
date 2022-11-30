import { request } from "../request";

/**
 * 登录
 * @param userName - 用户名
 * @param password - 密码
 */
export function fetchLogin(userName: string, password: string) {
    return request.post<ApiAuth.Token>("/login", { userName, password });
}
