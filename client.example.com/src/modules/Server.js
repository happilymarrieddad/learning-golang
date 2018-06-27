import axios from 'axios'
import qs from 'qs'
import Cookies from 'js-cookie'

class Server {
	constructor(opts) {
		this.baseUrl = 'http://localhost:3000'
	}

	request(url,data,method) {
		let self = this
		method = method || 'GET'
		data = data || {}

		return new Promise((resolve,reject) => {
			if (!url || typeof url != 'string') return reject('Invalid request')

			let Uri = self.baseUrl + ( url[0] != '/' ? '/' : '' ) + url
			let cookie = Cookies.get('api.example.com')

			let packet = {
				method,
				url:Uri,
				headers: { "X-App-Token":cookie }
			}


			switch(method) {
				case "POST":
					packet.data = qs.stringify(data)
					break

				case "GET":
					packet.params = data
					break
			}


			axios(packet)
				.then(res => resolve(res.data))
				.catch(reject)
		})
	}


	login(email,password) {
		let self = this

		return new Promise((resolve,reject) => {
			if (!email || !password) return reject('Email/password is required')

			let authLoginUri = self.baseUrl + '/auth/login'

			axios({
					method:'POST',
					url:authLoginUri,
					data:qs.stringify({ email,password })
				})
				.then(res => {
					return resolve(res.data)
				})
				.catch(err => {
					return reject('Invalid credentials')
				})

		})
	}

	testAuth(cookie) {
		let self = this

		return new Promise(async (resolve,reject) => {
			if (!cookie) return resolve(false)

			let config = {
				url:self.baseUrl + '/auth/check',
				method:'GET',
				headers: { "X-App-Token":cookie }
			}

			axios(config)
				.then(data => {
					return resolve(data)
				})
				.catch(err => {
					return resolve(false)
				})

		})
	}

}





export default {
	install(Vue,opts) {
		function alertInit() {
			var options = this.$options

			if (options.server) {
				this.$server = options.server
			} else if (options.parent && options.parent.$server) {
				this.$server = options.parent.$server
			} else {
				this.$server = new Server(opts)
			}

		}

		var usesInit = Vue.config._lifecycleHooks.indexOf('init') > -1
    	Vue.mixin(usesInit ? { init: alertInit } : { beforeCreate: alertInit })
	}
}