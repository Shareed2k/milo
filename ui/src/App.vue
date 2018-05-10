<template lang="pug">
  router-view
</template>

<script>
  import {mapActions, mapGetters} from 'vuex'
  import cookie from 'js-cookie'
  import _ from 'lodash'

  export default {
    data: () => ({
      clipped: false,
      drawer: true,
      fixed: false,
      menu: [
        {
          icon: 'restaurant',
          title: 'Menu',
          active: true,
          items: [
            {
              title: 'Home',
              link: '/',
              icon: 'plus'
            },
            {
              title: 'Region',
              link: '/regions',
              icon: 'plus'
            },
            {
              title: 'Provider',
              link: '/providers',
              icon: 'plus'
            },
            {
              title: 'DataCenter',
              link: '/datacenters',
              icon: 'plus'
            }
          ]
        }
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'Milo'
    }),
    name: 'App',

    computed: {
      ...mapGetters({
        isAuthorized: 'isAuthorized'
      })
    },

    methods: {
      ...mapActions([
        'setUser',
        'setToken'
      ])
    },

    beforeMount () {
      if (!this.isAuthorized) {
        let token = cookie.get('milo_token')
        if (!_.isUndefined(token)) {
          this.setToken(token)
            .then(t => { this.$http.defaults.headers.common['Authorization'] = `Bearer ${t}` })
            .then(() => this.$http.get('/api/bootdata')
              .then(r => this.setUser(r.data.user)))
              .then(() => this.$router.push('/'))
        }
      }
    }
}
</script>
