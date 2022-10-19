import behavior from "../behavior";


Component({
    behaviors: [behavior],
    properties: {},
    data: {},
    methods: {
        handlePay: function (event) {
            this.triggerEvent('pay')
        },

        handleRefund() {
            this.triggerEvent('refund')
        },

        handleRating() {
            this.triggerEvent('rating')
        },

        handleUpdateOrder() {
          const newData = JSON.stringify(this.data.order);
          wx.navigateTo({
              url: `/pages/order-management/index?order=` + encodeURIComponent(newData)
          })
      },
    }
});
