<template lang="pug">
  router-view
</template>

<script>
export default {
  data: () => ({
    clipped: false,
    drawer: true,
    fixed: false,
    items: [{
      icon: 'bubble_chart',
      title: 'Inspire'
    }],
    miniVariant: false,
    right: true,
    rightDrawer: false,
    title: 'Vuetify.js'
  }),
  name: 'App'
}
</script>
