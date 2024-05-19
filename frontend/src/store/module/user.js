import storageService from "@/service/storageService.js";
import userService from "@/service/userService.js";

const userModule={
    namespaced: true,
    state:{
        token:storageService.get(storageService.USER_TOKEN),
        userInfo: JSON.parse(storageService.get(storageService.USER_INFO)),
    },
    mutations:{
        //更新本地缓存服务
        SET_TOKEN(state,token){
            storageService.set(storageService.USER_TOKEN, token);
            //更新state
        state.token = token;
        },
        SET_USERINFO(state,userInfo){
            storageService.set(storageService.USER_INFO,JSON.stringify(userInfo));
            //更新state
            state.userInfo = userInfo;
        },

    },
    actions:{
        register(context, { name, telephone, password }) {
            return new Promise((resolve, reject) => {
                userService.register({ name, telephone, password }).then((res) => {
                    // 保存token
                    context.commit('SET_TOKEN', res.data.data.token);
                    return userService.info();
                }).then((res) => {
                    // 保存用户信息
                    context.commit('SET_USERINFO',res.data.data.user);
                    resolve(res);
                }).catch((err) => {
                    reject(err);
                });
            });
        },
        login(context, { telephone, password }) {
            return new Promise((resolve, reject) => {
                userService.login({  telephone, password }).then((res) => {
                    // 保存token
                    context.commit('SET_TOKEN', res.data.data.token);
                    return userService.info();
                }).then((res) => {
                    // 保存用户信息
                    context.commit('SET_USERINFO',res.data.data.user);
                    resolve(res);
                }).catch((err) => {
                    console.log('hi')
                    reject(err);
                });
            });
        },
        //登出，使得token失效
        logout(context) {
            // 清除token
            context.commit('SET_TOKEN', '');
            storageService.set(storageService.USER_TOKEN, null);
            // 清除用户信息
            context.commit('SET_USERINFO', '');
            storageService.set(storageService.USER_INFO, null);
            window.location.reload();
        },
    }
}
export default userModule