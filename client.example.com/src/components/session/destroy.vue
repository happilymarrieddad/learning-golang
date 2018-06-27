<template lang='pug'>

	h3 Logging out!

</template>

<script>
import Cookies from 'js-cookie'

export default {
	name:'session-destroy',
	beforeRouteEnter(to,from,next) {
		Cookies.remove('api.example.com')

		location.reload(true)
	}
}
</script>