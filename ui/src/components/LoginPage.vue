<template lang="pug">
  v-app#inspire
    v-content
      v-container(fluid='', fill-height='')
        v-layout(align-center='', justify-center='')
          v-flex(xs12='', sm8='', md4='')
            v-card.elevation-12
              v-toolbar(dark='', color='primary')
                v-toolbar-title Login form
                v-spacer
              v-card-text
                v-form(lazy-validation)
                  v-text-field(
                    prepend-icon='person',
                    v-model='request.email',
                    name='Email',
                    label='Email',
                    type='email',
                    v-validate="'required|email'",
                    :error-messages="veeErrors.collect('email')",
                    data-vv-name="email",
                    required
                  )
                  v-text-field#password(
                    prepend-icon='lock',
                    v-model='request.password',
                    name='password',
                    label='Password',
                    type='password',
                    v-validate="'required|max:20'",
                    :error-messages="veeErrors.collect('password')",
                    data-vv-name="password",
                    :counter="20",
                    required
                  )
              v-card-actions
                v-spacer
                v-btn(color='primary', @click.prevent='login', :disabled='veeErrors.any()') Login
</template>

<script>
  import {mapActions} from 'vuex'

  export default {
    name: 'login-page',

    $_veeValidate: {
      validator: 'new'
    },

    data: () => ({
      request: {
        email: '',
        password: ''
      }
    }),

    methods: {
      ...mapActions([
        'setUser'
      ]),

      login: function () {
        this.$validator.validateAll()
          .then(result => {
            if (result) {
              this.$http.post('/login', this.request)
                .then(r =>
                  this.setUser(r.data)
                    .then(u => {
                      this.$http.defaults.headers.common['Authorization'] = `Bearer ${u.token}`
                    })
                    .then(() => this.$router.push('/'))
                )
            }
          })
      }
    }
  }
</script>

<style scoped>

</style>
