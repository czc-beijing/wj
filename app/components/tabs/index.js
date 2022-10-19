import { throttle } from "../../utils/utils";

Component({
    options: {
        multipleSlots: true
    },
    properties: {
        tabs: Array,
        active: {
            type: Number,
            value: 0
        },
    },
    data: {
        currentTabIndex: 0,
    },
    observers: {
        active: function (value) {
            this.setData({
                currentTabIndex: this.data.active
            })
        }
    },
    methods: {
        handleSwitchTab: throttle(async function (event) {
            const index = event.currentTarget.dataset.index
            if (this.data.currentTabIndex === index) {
                return
            }
            this.setData({
                currentTabIndex: index,
            })
            this.triggerEvent('change', { index })
        }),
        handleTouchmove: function (event) {
            const direction = event.direction

            const currentTabIndex = this.data.currentTabIndex
            const targetTabIndex = currentTabIndex + direction

            if (targetTabIndex < 0 || targetTabIndex > this.data.tabs.length - 1) {
                return
            }

            const customEvent = {
                currentTarget: {
                    dataset: {
                        index: targetTabIndex
                    }
                }
            }
            this.handleSwitchTab(customEvent)
        }
    }
});
