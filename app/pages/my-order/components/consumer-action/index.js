const myBehavior = require('../behavior')

Component({
    behaviors: [myBehavior],
    properties: {},
    data: {},
    methods: {
        handleChat(event) {
            const newphone = event.currentTarget.dataset['tel']
            console.log(newphone)
            wx.makePhoneCall({
                phoneNumber: newphone
              })
        }
    },
   
});
