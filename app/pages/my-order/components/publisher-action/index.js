const myBehavior = require('../behavior')

Component({
    behaviors: [myBehavior],
    properties: {},
    data: {},
    methods: {
        handleChat(event) {
            const newphone = event.currentTarget.dataset['tel']
            wx.makePhoneCall({
                phoneNumber: newphone
              })
        }
    }
});


