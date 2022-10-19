import Order from "../../models/order";

Page({
    data: {
        service: null,
        address: null,
        formData: {
            description:'',
            begin_date: '',
            end_date: ''
        },
        rules: [
            {
                name: 'begin_date',
                rules: [
                    { required: true, message: '请指定服务有效日期开始时间' },
                ],
            },
            {
                name: 'end_date',
                rules: [
                    { required: true, message: '请指定服务有效日期结束时间' },
                    {
                        validator: function (rule, value, param, models) {
                            if (moment(value).isSame(models.begin_date) || moment(value).isAfter(models.begin_date)) {
                                return null
                            }
                            return '结束时间必须大于开始时间'
                        }
                    }
                ]
            }
        ]
    },
    onLoad: function (options) {
        const service = JSON.parse(decodeURIComponent(options.service));
        this.setData({
            service
        })
    },
    onShow() {
        let userInfo = wx.getStorageSync('userInfo')

        if (userInfo.id === this.data.service.publisher.id) {
            wx.redirectTo({
                url: '/pages/service-detail/index'
            })
        }

    },
    async handleSelectAddress() {
        let address
        try {
            address = await wx.chooseAddress();
        } catch (e) {
            address = null
        }

        this.setData({
            address
        })
    },

    async handleOrder() {
        if (this.data.service.designated_place && !this.data.address) {
            await wx.showModal({
                title: '错误',
                content: '该服务必须指定服务地点',
                showCancel: false,
            })
            return
        }

        const modalRes = await wx.showModal({
            title: '注意',
            content: '是否确认预约该服务？',
            showCancel: true,
        })

        if (!modalRes.confirm) return
        const order = {
            service_id: this.data.service.id,
            address: this.data.address,
            description:this.data.formData.description,
            begin_date: this.data.formData.begin_date,
            end_date: this.data.formData.end_date
        }

        wx.showLoading({ title: '正在预约...', mask: true })
        try {
            await Order.createOrder(order)
            wx.navigateTo({
                url: '/pages/order-success/index'
            })
            wx.hideLoading()
        } catch (e) {
            wx.hideLoading()
        }
    },

    bindBeginDateChange(event) {
        this.setData({
            'formData.begin_date': event.detail.value
        })
    },

    bindEndDateChange(event) {
        this.setData({
            'formData.end_date': event.detail.value
        })
    },
    bindInputChange(event) {
        this.setData({
            'formData.description': event.detail.value
        })
    },
});
