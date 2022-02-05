<template>
  <div class="card">
    <div class="border-0 card-header">
      <h1 class="mb-0 bold ">Container {{container.name}} </h1>
    </div>
    <table class="table table-bordered">
      <tbody>
      <tr>
        <th scope="row"><h3>ID</h3></th>
        <td>{{container.id}}</td>
      </tr>
      <tr>
        <th scope="row"><h3>Name</h3></th>
        <td>{{container.name}}</td>
      </tr>
      <tr>
        <th scope="row"><h3>Image</h3></th>
        <td>{{container.image}}</td>
      </tr>
      <tr>
        <th scope="row"><h3>Created</h3></th>
        <td>{{container.created}}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
export default {
  name: 'view-container',
  props: ["container"],
}
</script>
