import Vue from 'vue'
import Router from 'vue-router'

let routes = []
Vue.use(Router)



const Dashboard = () => import('@/components/dashboard/index.vue')
routes.push({ path:'/',component:Dashboard,meta:{ auth:true } })





const Session = () => import('@/components/session/index.vue')
const SessionCreate = () => import('@/components/session/create.vue')
routes.push({ path:'/session',component:Session,children:[
	{ path:'create',component:SessionCreate,meta:{ auth:false } }
],meta:{ auth:false } })



export default new Router({
	mode: 'history',
	routes
})
