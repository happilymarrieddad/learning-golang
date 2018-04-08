<template lang='pug'>
	
	div.uk-grid

		div.uk-width-1-3
		div.uk-width-1-3
			form.uk-form-stacked(style='margin-top:100px')

				div
					label.uk-form-label Email
					div.uk-form-controls
						input.uk-input(type='text',v-model='email')

				br
				div
					label.uk-form-label Password
					div.uk-form-controls
						input.uk-input(type='password',v-model='password')

				br
				div
					button.uk-button.uk-button-success(type='button',@click.prevent='login') Login



		div.uk-width-1-3




</template>

<script>
export default {
	name:'session-create',
	data() {
		return {
			email:'',
			password:''
		}
	},
	methods:{
		login() {
			let vm = this

			vm.$server.login(vm.email,vm.password)
				.then(data => {
					vm.$root.auth(data)
				})
				.catch(err => {
					vm.$notify(err,'success')
				})
		}
	}
}
</script>