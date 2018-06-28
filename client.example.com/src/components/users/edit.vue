<template lang='pug'>

	div
		router-link(to='/users') Back To Users
		form.uk-form-stacked
			div
				label.uk-form-label First Name
				div.uk-form-controls
					input.uk-input(v-model='first',placeholder='First Name')
			div
				label.uk-form-label Last Name
				div.uk-form-controls
					input.uk-input(v-model='last',placeholder='Last Name')
			div
				label.uk-form-label Email
				div.uk-form-controls
					input.uk-input(v-model='email',placeholder='Email')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='save') Save

</template> 

<script>
export default {
	name:'users-edit',
	data() {
		return {
			first:'',
			last:'',
			email:''
		}
	},
	computed: {
		id() { return this.$route.params.id }
	},
	methods:{
		save() {
			let vm = this

		 	let updated_user = {
		 		id:vm.id,
		 		first:vm.first,
		 		last:vm.last,
		 		email:vm.email
		 	}

		 	vm.$store.dispatch('users/update',updated_user)
		 		.then(newly_edited_user => {
		 			vm.$router.push({ path:'/users' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	mounted() {
		let vm = this

		vm.$store.dispatch(`users/edit`,vm.id)
			.then(user => {
				vm.first = user.first
				vm.last = user.last
				vm.email = user.email
			})
		 	.catch(err => vm.$root.error(err))
	}
}
</script>