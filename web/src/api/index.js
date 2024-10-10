import request from '@/utils/request'

export function healthz() {
    return request({
        url: '/v1/healthz',
        method: 'get'
    })
}

export function search(keyword) {
    return request({
        url: '/v1/command/search',
        method: 'get',
        params: {
            keyword
        }
    })
}

export function show(keyword) {
    return request({
        url: '/v1/command/show',
        method: 'get',
        params: {
            keyword
        }
    })
}