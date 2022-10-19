Component({
    properties: {
        userInfo: Object,
    },
    data: {},
    methods: {
        handleChat() {
            wx.makePhoneCall({
                phoneNumber: '13429208394'
              })
        }
    }
});
