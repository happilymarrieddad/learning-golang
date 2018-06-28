// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import Alertifyjs from 'vue2-alertifyjs'
import Cookies from 'js-cookie'
 
import App from './App'
import router from './router'
import Server from '@/modules/Server.js'
import Store from '@/modules/Store'



Vue.config.productionTip = false




import 'uikit/dist/css/uikit.min.css'
import 'uikit/dist/js/uikit.min.js'
import 'uikit/dist/js/uikit-icons.min.js'
import 'alertifyjs/build/alertify.min.js'
import 'alertifyjs/build/css/alertify.min.css'
import 'alertifyjs/build/css/themes/default.min.css'
Vue.use(Alertifyjs,{
  notifier:{
      delay:5,
      position:'top-right',
      closeButton: false
  }
})



let app_started = false
router.beforeEach((to,from,next) => {
	if (!app_started) return

	if (to.meta.auth && !router.app.authenticated) return next('/session/create')
	if (to.path == '/session/create' && router.app && router.app.authenticated) return next('/')


	return next()
})




router.afterEach((to,from) => {
	setTimeout(() => router.app.loading = false,100)
})




Vue.use(Server,{
	test:'1'
})

Vue.use(Vuex)

/* eslint-disable no-new */
new Vue({
	el: '#app',
	data() {
		return {
			authenticated:false,
			loading:true,
			user:{}
		}
	},
	store:new Vuex.Store(Store),
	router,
	components: { App },
	template: '<App/>',
	methods:{
		auth(data) {
			let vm = this
			let { user,token } = data

			Cookies.set('api.example.com',token)
			vm.user = user || {}
			
			vm.authenticated = true
			vm.$router.push({ path:'/' })
		},
		info(msg) {
			alert(msg)
		},
		error(msg) {
			alert(msg)
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch('setRoot',vm)

		let cookie = Cookies.get('api.example.com')
		this.$server.testAuth(cookie)
			.then(data => {

				let path = '/session/create'
				if (data) {
					let { user,token } = data
					
					vm.authenticated = true
					path = '/'
				}

				app_started = true
				vm.$router.push({ path })
			})
	}
})
