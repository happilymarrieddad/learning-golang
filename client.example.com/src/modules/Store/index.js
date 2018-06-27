import users from '@/modules/Store/Users.js'

export default {
	modules:{
		users
	},
	state:{
		$root:{}
	},
	actions:{
		setRoot(context,data) {
			context.state.$root = data
		}
	}
}