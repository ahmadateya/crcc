<template>
  <div>
      <ul v-for="row in rows" :key="row.id">
          <li>heap{{row.id}}</li>
      </ul>
  </div>
</template>

<script>
export default {
props:["rows"]
}
</script>

<style>

</style>