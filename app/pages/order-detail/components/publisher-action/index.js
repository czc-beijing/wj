import behavior from "../behavior";
import Order from "../../../../models/order";
import orderAction from "../../../../enum/order-action";

Component({
    behaviors: [behavior],
    properties: {},
    data: {},
    methods: {
        async handlePay() {
            const modalRes = await wx.showModal({
                title: '注意',
                content: `您即将支付该服务费用：￥${this.data.order.price}元，是否确认支付`,
                showCancel: true,
            })
            if (!modalRes.confirm) return
            // 模拟支付后订单状态改变
            await Order.updateOrderStatus(this.data.order.id, orderAction.PAY)
            // 跳转支付成功页面
            wx.navigateTo({
                url: '/pages/pay-success/index'
            })
        }
    },
   
});
