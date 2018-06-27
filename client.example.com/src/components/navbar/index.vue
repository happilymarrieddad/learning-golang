<template lang='pug'>

	nav.uk-navbar-container(uk-navbar)

		div.uk-navbar-left
			ul.uk-navbar-nav
				li
					router-link(to='/') Dashboard
				li
					router-link(to='/users') Users

		div.uk-navbar-right
			ul.uk-navbar-nav
				li
					router-link(to='/session/destroy') Logout

</template> 

<script>
export default {
	name:'navbar'
}
</script>