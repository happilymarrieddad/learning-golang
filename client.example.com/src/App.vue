<template lang='pug'>
    div#app.uk-container
        img.uk-align-center(v-show='$root.loading',src='/static/loading.gif')
        router-view(v-show='!$root.loading')
</template>

<script>
export default {
  name: 'App'
}
</script>

