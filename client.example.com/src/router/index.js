import Vue from 'vue'
import Router from 'vue-router'

let routes = []
Vue.use(Router)



const Dashboard = () => import('@/components/dashboard/index.vue')
routes.push({ path:'/',component:Dashboard,meta:{ auth:true } })






const Users = () => import('@/components/users/index.vue')
const UsersCreate = () => import('@/components/users/create.vue')
const UsersEdit = () => import('@/components/users/edit.vue')
routes.push({ path:'/users',component:Users,meta:{ auth:true } })
routes.push({ path:'/users/create',component:UsersCreate,meta:{ auth:true } })
routes.push({ path:'/users/:id/edit',component:UsersEdit,meta:{ auth:true } })




const Session = () => import('@/components/session/index.vue')
const SessionCreate = () => import('@/components/session/create.vue')
const SessionDestroy = () => import('@/components/session/destroy.vue')
routes.push({ path:'/session',component:Session,children:[
	{ path:'create',component:SessionCreate,meta:{ auth:false } },
	{ path:'destroy',component:SessionDestroy,meta:{ auth:true } }
],meta:{ auth:false } })



export default new Router({
	mode: 'history',
	routes
})
