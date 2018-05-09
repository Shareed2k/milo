<template lang="pug">
  extends ../layouts/master.pug

  block content
    v-dialog(v-model='dialog', max-width='500px', persistent)
      v-btn.mb-2(color='primary', dark='', slot='activator') New Item
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
        td {{ props.item.description }}
        td.justify-center.layout.px-0
          v-btn.mx-0(icon='', @click='editItem(props.item)')
            v-icon(color='teal') edit
          v-btn.mx-0(icon='', @click='deleteItem(props.item)')
            v-icon(color='pink') delete

</template>

<script>
  import App from '../App'

  export default {
    extends: App,
    name: 'RegionPage',

    $_veeValidate: {
      validator: 'new'
    },

    data: () => ({
      dialog: false,
      valid: false,
      editedIndex: -1,
      headers: [
        { text: 'UUID', value: 'uuid' },
        { text: 'Name', value: 'name' },
        { text: 'Description', value: 'description' },
        { text: 'Actions', value: 'name', sortable: false }
      ],
      items: [],

      defaultItem: {
        name: '',
        description: ''
      },

      request: {
        name: '',
        description: ''
      }
    }),

    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
      }
    },

    methods: {
      getRegions: function () {
        this.$http.get('/api/regions')
          .then(r => { this.items = r.data.items })
      },

      close () {
        this.dialog = false
        setTimeout(() => {
          this.request = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      },

      save () {
        this.$validator.validateAll()
          .then(result => {
            if (result) {
              this.valid = result
              if (this.editedIndex > -1) {
                this.$http.put('/api/regions', this.request)
                  .then(r => Object.assign(this.items[this.editedIndex], this.request))
                  .then(() => this.close())
              } else {
                this.$http.post('/api/regions', this.request)
                  .then(r => this.items.push(r.data))
                  .then(() => this.close())
              }
            }
          })
      },

      editItem (item) {
        this.editedIndex = this.items.indexOf(item)
        this.request = Object.assign({}, item)
        this.dialog = true
      },

      deleteItem (item) {
        const index = this.items.indexOf(item)
        this.$http.delete(`/api/regions/${item.uuid}`)
          .then(r => console.log(r.data))
          .then(() => this.items.splice(index, 1))
      }
    },

    mounted () {
      this.getRegions()
    }
  }
</script>

<style scoped>

</style>
