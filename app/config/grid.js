/**
 * @Created by WebStorm
 * @Author: 沁塵
 * @Desc:
 */
import roleType from "../enum/role-type";
import serviceType from "../enum/service-type";

const appointWithMeGrid = {
    rowNum: 5,
    title: '订单',
    extend: '查看全部',
    extendCell: {
        role: roleType.PUBLISHER,
        status: -1,
    },
    gridItems: [
        {
            badge: 'unapproved',
            cell: { role: roleType.PUBLISHER, status: 0 },
            icon: 'operation',
            text: '待处理'
        },
        {
            badge: 'unpaid',
            cell: { role: roleType.PUBLISHER, status: 1 },
            icon: 'file-open',
            text: '工作中'
        },
        {
            badge: 'unconfirmed',
            cell: { role: roleType.PUBLISHER, status: 2 },
            icon: 'complete',
            text: '工作完成'
        },
        {
            badge: 'unrated',
            cell: { role: roleType.PUBLISHER, status: 3 },
            icon: 'comment',
            text: '待评价'
        },
        {
            badge: '',
            cell: { role: roleType.PUBLISHER , status:4},
            icon: 'tuipiao',
            text: '已完成'
        }
    ]
}

const myAppointGrid = {
    rowNum: 5,
    title: '我的预约',
    extend: '查看全部',
    extendCell: {
        role: roleType.CONSUMER,
        status: -1,
    },
    gridItems: [
        {
            badge: 'unapproved',
            cell: { role: roleType.CONSUMER, status: 0 },
            icon: 'operation',
            text: '待同意'
        },
        {
            badge: 'unpaid',
            cell: { role: roleType.CONSUMER, status: 1 },
            icon: 'file-open',
            text: '工作中'
        },
        {
            badge: 'unconfirmed',
            cell: { role: roleType.CONSUMER, status: 2 },
            icon: 'complete',
            text: '工作完成'
        },
        {
            badge: 'unrated',
            cell: { role: roleType.CONSUMER, status: 3 },
            icon: 'comment',
            text: '待评价'
        },
        {
            badge: '',
            cell: { role: roleType.CONSUMER, status:4 },
            icon: 'tuipiao',
            text: '已完成'
        }
    ]
}

const myProvideGird = {
    rowNum: 4,
    title: '我的挖机',
    extend: '',
    extendCell: null,
    gridItems: [
        {
            badge: 'pending',
            cell: { type: serviceType.PROVIDE, status: 0 },
            icon: 'time',
            text: '待审核'
        },
        {
            badge: 'unpublished',
            cell: { type: serviceType.PROVIDE, status: 1 },
            icon: 'upload',
            text: '待发布'
        },
        {
            badge: '',
            cell: { type: serviceType.PROVIDE, status: 2 },
            icon: 'file-common',
            text: '已发布'
        },
        {
            badge: '',
            cell: { type: serviceType.PROVIDE, status: -1 },
            icon: 'comment',
            text: '查看全部'
        },
    ]
}

export { myAppointGrid, myProvideGird , appointWithMeGrid}
