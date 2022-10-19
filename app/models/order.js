import Http from "../utils/http";
import Base from "./base";

class Order extends Base {

    static async getOrderStatus(role) {
        return Http.request({
            url: `v1/order/count?role=${role}`,
        })
    }

    static async createOrder(order) {
        return Http.request({
            url: 'v1/order',
            data: order,
            method: 'POST'
        })
    }

    static async updateOrderStatus(orderId, action) {
        return Http.request({
            url: `v1/order/${orderId}`,
            data: {
                action
            },
            method: 'POST'
        })
    }

    static async getOrderById(orderId) {
        return await Http.request({
            url: `v1/order/${orderId}`,
        });
    }

    async getMyOrderList(role, status) {
        if (!this.hasMoreData) {
            return this.data
        }
        const orderList = await Http.request({
            url: 'v1/order/my',
            data: {
                page: this.page,
                count: this.count,
                role,
                status
            }
        });

        this.data = this.data.concat(orderList.data)
        this.hasMoreData = this.page !== orderList.last_page
        this.page++
        return this.data
    }

    async searchMyOrderList(role, status, serach_value) {
        if (!this.hasMoreData) {
            return this.data
        }
        const orderList = await Http.request({
            url: 'v1/order/my',
            data: {
                page: this.page,
                count: this.count,
                role:role,
                status:status,
                address_user_name:serach_value
            }
        });

        this.data = this.data.concat(orderList.data)
        this.hasMoreData = this.page !== orderList.last_page
        this.page++
        return this.data
    }
}

export default Order
