import axios from 'axios'
import qs from 'qs'


class Server {
	constructor(opts) {
		this.baseUrl = 'http://localhost:3000'
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