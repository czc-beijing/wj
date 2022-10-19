import { wxToPromise } from "../utils/wx";
import APIConfig from "../config/api";
import authUserinfo from "../enum/auth-userinfo";
import Http from "../utils/http";

class Token {

    static tokenUrl = 'v1/token'

    
    static async getToken(code) {
        const sendcode = code
        console.log(sendcode)
        const res = await wxToPromise('request', {
            url: `${APIConfig.baseUrl}/${this.tokenUrl}`,
            data: {
                code: sendcode
            },
            method: 'GET'
        })
        wx.setStorageSync('token', res.data.data.token)
        return res.data.data.token
    }

    static async verifyToken() {
        const token = wx.getStorageSync('token');
        return await Http.request({
            url: `v1/token/verify`,
            data: { token },
            method: 'POST'
        })
    }

    static async getAuthUserInfoStatus() {
        const setting = await wx.getSetting({})
        const userInfoSetting = await setting.authSetting['scope.userInfo']
        if (userInfoSetting === undefined) {
            return authUserinfo.NOT_AUTH
        }
        if (userInfoSetting === false) {
            return authUserinfo.DENY
        }
        if (userInfoSetting === true) {
            return authUserinfo.AUTHORIZED
        }
    }
}

export default Token
