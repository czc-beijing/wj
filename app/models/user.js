import Http from "../utils/http";
import Token from "./token";


class User {

    static async login(userRes) {
        wx.login({
            async success (res) {
                 if (res.code) {
                const token = await Token.getToken(res.code)
                wx.setStorageSync('token', token)
                wx.setStorageSync('isLogin', true)
              }
              await User.updateUserInfo(userRes.userInfo)
            }
          })
    }

    static newLogin() {
        wx.login({
            async success (res) {
                 if (res.code) {
                const token = await Token.getToken(res.code)
                wx.setStorageSync('token', token)
                wx.setStorageSync('isLogin', true)
              }
            }
          })
    }


    static getUserInfoByLocal() {
        return wx.getStorageSync('userInfo')
    }

    static async getUserInfo() {
        const userInfo = await Http.request({ url: 'v1/user' })
        if (userInfo) {
            return userInfo
        } else {
            return null
        }
    }

    static async getUserSign() {
        return await Http.request({
            url: 'v1/user/sign'
        })

    }

    static async updateUserInfo(data) {
        console.log(66666, data)
        const res = await Http.request({
            url: 'v1/user',
            data: {
                nickname: data.nickName,
                avatar: data.avatarUrl,
            },
            method: 'POST'
        });
        console.log(7777777, res)
        wx.setStorageSync('userInfo', res)
    }

}

export default User
