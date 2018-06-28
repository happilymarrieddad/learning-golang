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
			div
				label.uk-form-label Password
				div.uk-form-controls
					input.uk-input(type='password',v-model='password',placeholder='Password')
			div
				label.uk-form-label Confirm Password
				div.uk-form-controls
					input.uk-input(type='password',v-model='confirm_password',placeholder='Confirm Password')
			hr
			div
				button.uk-button.uk-button-primary(type='button',@click.prevent='create',:disabled='notValidForm') Create
			div(v-show='notValidForm')
				span.text-danger Passwords are required and must match

</template> 

<script>
export default {
	name:'users-create',
	data() {
		return {
			first:'',
			last:'',
			email:'',
			password:'',
			confirm_password:''
		}
	},
	methods:{
		create() {
			let vm = this

		 	let new_user = {
		 		first:vm.first,
		 		last:vm.last,
		 		email:vm.email,
		 		password:vm.password
		 	}

		 	vm.$store
		 		.dispatch('users/create',new_user)
		 		.then(newly_created_user => {
		 			vm.$router.push({ path:'/users' })
		 		})
		 		.catch(err => vm.$root.error(err))
		}
	},
	computed:{
		notValidForm() {
			return this.password.length < 1 || this.confirm_password.length < 1 || this.password != this.confirm_password
		}
	}
}
</script>