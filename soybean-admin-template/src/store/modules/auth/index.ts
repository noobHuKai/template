// import { unref } from 'vue';
import { defineStore } from 'pinia';
// import { router } from '@/router';
import { fetchLogin, fetchUserInfo } from '@/service';
import { useRouterPush } from '@/composables';
import { localStg } from '@/utils';
import { MD5Encrypt } from '@/utils/crypto';
import { useTabStore } from '../tab';
import { useRouteStore } from '../route';
import { getToken, getUserInfo, clearAuthStorage } from './helpers';
import { router } from '~/src/router';

interface AuthState {
  /** 用户信息 */
  userInfo: Auth.UserInfo;
  /** 用户token */
  token: string;
  /** 登录的加载状态 */
  loginLoading: boolean;
}

export const useAuthStore = defineStore('auth-store', {
  state: (): AuthState => ({
    userInfo: getUserInfo(),
    token: getToken(),
    loginLoading: false
  }),
  getters: {
    /** 是否登录 */
    isLogin(state) {
      return Boolean(state.token);
    }
  },
  actions: {
    /** 重置auth状态 */
    /** 不是双 token ，是 redis+token */
    resetAuthStore(isRefresh :boolean=false) {
      const { toLogin } = useRouterPush(false);
      const { resetTabStore } = useTabStore();
      const { resetRouteStore } = useRouteStore();
      // const route = unref(router.currentRoute);

      clearAuthStorage();
      this.$reset();

      resetTabStore();
      resetRouteStore();
      if (isRefresh){
        // 刷新页面，不然不会加载
        router.go(0);
      }
      // if (route.meta.requiresAuth) {
      toLogin();
      // }
    },
    /**
     * 处理登录后成功或失败的逻辑
     * @param backendToken - 返回的token
     */
    async handleActionAfterLogin(backendToken: ApiAuth.Token) {
      const { toLoginRedirect } = useRouterPush(false);

      const loginSuccess = await this.loginByToken(backendToken);

      if (loginSuccess) {
        // 跳转登录后的地址
        toLoginRedirect();

        // 登录成功弹出欢迎提示
        window.$notification?.success({
          title: '登录成功!',
          content: `欢迎回来，${this.userInfo.userName}!`,
          duration: 3000
        });

        return;
      }

      // 不成功则重置状态
      this.resetAuthStore();
    },
    /**
     * 根据token进行登录
     * @param backendToken - 返回的token
     */
    async loginByToken(backendToken: ApiAuth.Token) {
      let successFlag = false;

      // 先把token存储到缓存中(后面接口的请求头需要token)
      const { token } = backendToken;
      localStg.set('token', token);
      // localStg.set('refreshToken', refreshToken);

      // 获取用户信息
      const { data } = await fetchUserInfo();
      if (data) {
        // 成功后把用户信息存储到缓存中
        localStg.set('userInfo', data);

        // 更新状态
        this.userInfo = data;
        this.token = token;

        successFlag = true;
      }

      return successFlag;
    },
    /**
     * 登录
     * @param userName - 用户名
     * @param password - 密码
     */
    async login(userName: string, password: string) {
      this.loginLoading = true;
      const { data } = await fetchLogin(userName, MD5Encrypt(password));
      if (data) {
        await this.handleActionAfterLogin(data);
      }
      this.loginLoading = false;
    },
    /**
     * 更换用户权限(切换账号)
     * @param userRole
     */
    async updateUserRole(userRole: Auth.RoleType) {
      const { resetRouteStore, initAuthRoute } = useRouteStore();

      const accounts: Record<Auth.RoleType, { userName: string; password: string }> = {
        super: {
          userName: 'Super',
          password: 'super123'
        },
        admin: {
          userName: 'Admin',
          password: 'admin123'
        },
        user: {
          userName: 'User01',
          password: 'user01123'
        }
      };
      const { userName, password } = accounts[userRole];
      const { data } = await fetchLogin(userName, password);
      if (data) {
        await this.loginByToken(data);
        resetRouteStore();
        initAuthRoute();
      }
    }
  }
});
