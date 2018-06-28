<template lang='pug'>

	table.uk-table
		caption
			router-link(to='/users/create') Create New User
		thead
			tr
				th First Name
				th Last Name
				th Email
				th
		tbody
			tr(v-for='user in users')
				td {{ user.first }}
				td {{ user.last }}
				td {{ user.email }}
				td
					button.uk-button.uk-button-default Edit
					button.uk-button.uk-button-primary Delete


</template>

<script>
import { mapState } from 'vuex'

export default {
	name:'users-index',
	computed:{
		...mapState({
			users:state => state.users.items
		})
	},
	mounted() {
		let vm = this

		vm.$store.dispatch('users/index')
	}
}
</script>
