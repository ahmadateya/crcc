<template>
  <div>
    <base-header class="pb-6">
    </base-header>
    <div class="container-fluid mt--6">
      <div class="row">
        <div class="col">
          <main-table :rows="containers"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import users from '~/components/tables/users'
import LightTable from "~/components/tables/RegularTables/LightTable";
import MainTable from "~/components/tables/RegularTables/MainTable";

import Jsona from 'jsona';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    MainTable,
    // [Dropdown.name]: Dropdown,
    // [DropdownItem.name]: DropdownItem,
    // [DropdownMenu.name]: DropdownMenu,
    // [Table.name]: Table,
    // [TableColumn.name]: TableColumn
  },
  data() {
    return {
      containers: [],
    }
  },
  // created() {
  //   this.getProfile();
  // },
  async fetch() {
    this.containers = await this.$http.$get(`${url}/containers`);
    // this.containers =  await this.$axios.get(`${url}/containers`)
    //     .then(response => {
    //       console.log(response.data);
    //       return {
    //         data: jsona.deserialize(response.data),
    //       };
    //     });
  },
  // methods: {
  //   async getProfile() {
  //     awiat this.$axios.get(`${url}/containers`)
  //       .then(response => {
  //         console.log(response.data);
  //         return {
  //         data: jsona.deserialize(response.data),
  //       };
  //   });
  //   }
  // }
};
</script>
