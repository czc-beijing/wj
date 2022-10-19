import User from "../../models/user";
import { createStoreBindings } from "mobx-miniprogram-bindings";
import { timStore } from "../../store/tim";


Page({
    data: {},
    onLoad: function (options) {
        this.storeBindings = createStoreBindings(this, {
            store: timStore,
            actions: { timLogin: 'login' },
        })
    },

    onUnload() {
        this.storeBindings.destroyStoreBindings()
    },

    async handleUserInfo(event) {
        const res = await wx.getUserProfile({ desc: '完善用户信息', })

        wx.showLoading({ title: '正在授权', mask: true })

        try {
            await User.login(res)
            await this.timLogin()
            const eventChannel = this.getOpenerEventChannel()
            eventChannel.emit('login');
            wx.navigateBack()
        } catch (e) {
            wx.showModal({
                content: '登陆成功',
                showCancel: true
            })
            console.log(e)
        }
        wx.hideLoading()
        wx.switchTab({
            url: '/pages/plaza/index'
        })
    },

    handleBackHome() {
        wx.switchTab({
            url: '/pages/plaza/index'
        })
    },
});
