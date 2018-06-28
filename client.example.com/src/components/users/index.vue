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
				td
					router-link(:to="'/users/' + user.id + '/edit'") {{ user.first }}
				td {{ user.last }}
				td {{ user.email }}
				td
					button.uk-button.uk-button-primary(@click.prevent='destroy(user.id)') Delete


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
	},
	methods:{
		destroy(id) {
			let vm = this
			
			vm.$store.dispatch('users/destroy',id)
				.then(res => {
					vm.$store.dispatch('users/index')
				})
		}
	}
}
</script>
