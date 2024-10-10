import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '',
            redirect: "/search"
        },
        {
            path: '/search',
            name: 'Search',
            component: () => import('@/components/Search')
        }
        ,
        {
            path: '/show/:command',
            name: 'Show',
            component: () => import('@/components/Show')
        }
    ],

});

export default router;
