<template lang='pug'>

    div#app.uk-container
        img.uk-align-center(v-show='$root.loading',src='/static/loading.gif')
        navbar(v-show='$root.authenticated')
        router-view(v-show='!$root.loading')

</template>

<script>
import Navbar from '@/components/navbar/index.vue'

export default {
  name: 'App',
  components:{ Navbar }
}
</script>

