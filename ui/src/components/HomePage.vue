<template lang="pug">
  extends ../layouts/master.pug

  block content
    v-data-table.elevation-1(:headers='headers', :items='items', hide-actions='')
      template(slot='items', slot-scope='props')
        td {{ props.item.ID }}
        td {{ props.item.uuid }}
        td {{ props.item.public_ip }}
        td {{ props.item.private_ip }}
</template>

<script>
  import App from '../App'

  export default {
    extends: App,
    name: 'home-page',

    data: () => ({
      headers: [
        {
          text: 'ID',
          align: 'left',
          sortable: false,
          value: 'ID'
        },
        { text: 'UUID', value: 'uuid' },
        { text: 'Public IP', value: 'public_ip' },
        { text: 'Private IP', value: 'private_ip' }
      ],
      items: []
    }),

    mounted () {
      this.$http.get('/api/servers')
        .then(r => { this.items = r.data.items })
    }
  }
</script>

<style scoped>

</style>
