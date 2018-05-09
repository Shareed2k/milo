<template lang="pug">
  extends ../layouts/master.pug

  block content
    v-dialog(v-model='dialog', max-width='500px', persistent)
      v-btn.mb-2(color='primary', dark='', slot='activator') New DataCenter
      v-card
        v-card-title
          span.headline {{ formTitle }}
        v-card-text
          form
            v-text-field(
              label='Name',
              v-model='request.name',
              :counter="10",
              :error-messages="veeErrors.collect('name')",
              v-validate="'required|max:10'",
              data-vv-name="name",
              required
            )
            v-text-field(
              label='Provider',
              v-model='request.provider',
              :counter="10",
              :error-messages="veeErrors.collect('provider')",
              v-validate="'required|max:10'",
              data-vv-name="provider",
              required
            )
            v-text-field(
              label='Description',
              v-model='request.description',
              :counter="256",
              :error-messages="veeErrors.collect('description')",
              v-validate="'max:256'",
              data-vv-name="description"
            )
        v-card-actions
          v-spacer
          v-btn(color='blue darken-1', flat='', @click.native='close') Cancel
          v-btn(color='blue darken-1', flat='', @click.native='save', :disabled='veeErrors.any()') Save

    v-data-table.elevation-1(:headers='headers', :items='items', hide-actions='')
      template(slot='items', slot-scope='props')
        td {{ props.item.uuid }}
        td {{ props.item.name }}
        td {{ props.item.provider }}
        td {{ props.item.description }}
</template>

<script>
  import App from '../App'

  export default {
    extends: App,
    name: 'DataCenterPage',

    data: () => ({
      headers: [
        { text: 'UUID', value: 'uuid' },
        { text: 'Name', value: 'name' },
        { text: 'Provider', value: 'provider' },
        { text: 'Description', value: 'description' }
      ],
      items: [],

      defaultItem: {
        name: '',
        provider: '',
        description: ''
      },

      request: {
        name: '',
        provider: '',
        description: ''
      }
    }),

    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'New DataCenter' : 'Edit DataCenter'
      }
    },

    methods: {
      getDataCenter: function () {
        this.$http.get('/api/datacenters')
          .then(r => { this.items = r.data.items })
      }
    },

    mounted () {
      this.getDataCenter()
    }
  }
</script>

<style scoped>

</style>
