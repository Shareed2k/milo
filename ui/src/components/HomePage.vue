<template lang="pug">
  extends ../layouts/master.pug

  block content
    v-data-table.elevation-1(:headers='headers', :items='items', hide-actions='')
      template(slot='items', slot-scope='props')
        td {{ props.item.name }}
        td.text-xs-right {{ props.item.calories }}
        td.text-xs-right {{ props.item.fat }}
        td.text-xs-right {{ props.item.carbs }}
        td.text-xs-right {{ props.item.protein }}
        td.text-xs-right {{ props.item.iron }}
</template>

<script>
  import App from '../App'

  export default {
    extends: App,
    name: 'home-page',

    data: () => ({
      headers: [
        {
          text: 'Dessert (100g serving)',
          align: 'left',
          sortable: false,
          value: 'name'
        },
        { text: 'Calories', value: 'calories' },
        { text: 'Fat (g)', value: 'fat' },
        { text: 'Carbs (g)', value: 'carbs' },
        { text: 'Protein (g)', value: 'protein' },
        { text: 'Iron (%)', value: 'iron' }
      ],
      items: []
    })
  }
</script>

<style scoped>

</style>
