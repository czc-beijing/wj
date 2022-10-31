import Service from "../../models/service";
import Category from "../../models/category";
import { throttle } from "../../utils/utils";
import { setTabBarBadge } from "../../utils/wx";

const serviceModel = new Service()

Page({
    data: {
        loading: true,
        tabs: ['军明挖机家族'],
        serviceList: [],
        currentCategoryId: 0,
        multiple: 0,
        categoryList: [],
        currentTabIndex: 0,
        showStatus: false,
    },

    onLoad: async function (options) {
        const categoryList = await Category.getCategoryListWithAll();
        this.setData({
            categoryList,
            multiple: 2
        })
        await this._getInitServiceList(this.data.currentTabIndex)
        this.setData({
            loading: false
        })

    },

    onShow: function () {
        const unreadCount = wx.getStorageSync('unread-count')
        setTabBarBadge(unreadCount)
    },

    handleTabChange: function (event) {
        const index = event.detail.index
        this._getInitServiceList(index, this.data.currentCategoryId)
    },

    handleChangeCategory: throttle(function (event) {
        const categoryId = event.currentTarget.dataset.id
        if (categoryId === this.data.currentCategoryId) {
            return
        }

        this._getInitServiceList(this.data.currentTabIndex, categoryId)
    }),

    handleSelect: function (event) {
        const service = event.detail.service;
        wx.navigateTo({
            url: `/pages/service-detail/index?id=${service.id}`
        })
    },

    async _getInitServiceList(currentTabIndex = 0, categoryId = 0) {
        this.setData({
            currentTabIndex: currentTabIndex,
            currentCategoryId: categoryId,
            showStatus: false,
        })
        const serviceList = await serviceModel.reset().getServiceList(currentTabIndex, categoryId);

        this.setData({
            showStatus: !serviceList.length,
            serviceList,
        })
    },

    handleNavRemark: function () {
        wx.navigateTo({
            url: "/pages/remark/index"
        })
    },

    onPullDownRefresh: function () {
        this._getInitServiceList(this.data.currentTabIndex, this.data.currentCategoryId)
        wx.stopPullDownRefresh()
    },

    onReachBottom: async function () {
        if (!serviceModel.hasMoreData) {
            return
        }
        const serviceList = await serviceModel.getServiceList(this.data.currentTabIndex, this.data.currentCategoryId);
        this.setData({
            serviceList
        })
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
});
