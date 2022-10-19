import Order from "../../models/order";
import orderStatus from "../../enum/order-status";
import roleType from "../../enum/role-type";

const orderModel = new Order()

Page({
    data: {
        loading: {
            hideTabsLoading: false,
            hideOrderLoading: false,
        },
        tabs: ['全部订单', '待同意', '工作中', '待确认', '待评价', '已完成'],
        currentTabIndex: 0,
        orderList: [],
        orderStatus,
        roleType,
        role: null,
    },
    onLoad: async function (options) {
        // status: -1:全部  0：待同意、1:待收款、2：待确认、3：待评价 4已完成
        // tabs:    0：全部  1：待同意、 2、待收款、 3 待确认 4 待评价 5 已完
        const status = parseInt(options.status)
        const role = parseInt(options.role)
        wx.setNavigationBarTitle({
            title: role === roleType.PUBLISHER ? '订单' : '我的预约'
        })
        this.setData({
            currentTabIndex: status + 1,
            role
        })
        this.data.status = status < 0 ? '' : status

    },

    onShow() {
        this._getOrderList()
    },

    handleTabChange: async function (event) {
        const index = event.detail.index
        this.data.status = index < 1 ? '' : index - 1

        this.setData({
            currentTabIndex: index,
        })

        await this._getOrderList()
    },

    async _getOrderList() {

        this.setData({
            ['loading.hideOrderLoading']: false,
        })

        const orderList = await orderModel.reset()
            .getMyOrderList(this.data.role, this.data.status)

        this.setData({
            ['loading.hideOrderLoading']: true,
            orderList,
        })
    },

    handleNavDetail(event) {
        var nextDatas = JSON.stringify(event.detail.order)
        wx.navigateTo({
            url: `/pages/order-detail/index?role=${this.data.role}&order=` + encodeURIComponent(nextDatas)
        })
    },

    handleRefund(event) {
        const order = event.detail.order
        wx.navigateTo({
            url: `/pages/refund/index?order=${JSON.stringify(order)}`
        })
    },

    handleChat(event) {
        const { order } = event.detail
        const targetUserId = order[this.data.role === roleType.PUBLISHER ? 'consumer' : 'publisher'].id

        wx.navigateTo({ url: `/pages/conversation/index?targetUserId=${targetUserId}&service=${JSON.stringify(order.service_snap)}` })
    },

    async onPullDownRefresh() {
        await this._getOrderList()
        wx.stopPullDownRefresh()
    },

    async onReachBottom() {
        if (!orderModel.hasMoreData) {
            return
        }
        const orderList = await orderModel.getMyOrderList(this.data.role, this.data.status);
        this.setData({
            orderList
        })
    },

    searchInput: function (e) {
        let value = e.detail.value //搜索框输入的信息
        this._searchOrderList(value)
      },

      async _searchOrderList(serach_value) {
        this.setData({
            ['loading.hideOrderLoading']: false,
        })

        const orderList = await orderModel.reset()
            .searchMyOrderList(this.data.role, this.data.status, serach_value)

        this.setData({
            ['loading.hideOrderLoading']: true,
            orderList,
        })
    },

});
