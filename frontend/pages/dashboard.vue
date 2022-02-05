<template>
  <div>
    <base-header class="pb-6">
    </base-header>
    <div class="container-fluid mt--6">
      <div class="row">
        <div class="col">
          <main-table v-if="containers.length!==0" :rows="containers"/>
          <h2 v-else-if="loaded.error">Error while fetching data please request it again.</h2>
          <h2 v-else-if="loaded.reponseError">Please make sure of allowing the Rest API.</h2>
          <h2 v-else-if="loaded.length===0">No Running Containers</h2>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import MainTable from "~/components/tables/RegularTables/MainTable";
import Jsona from 'jsona';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    MainTable,
  },
  data() {
    return {
      containers: [],
      loaded:{},
      valid:true
    }
  },
  async fetch() {
   // this.containers = await this.$axios.$get(`${url}/containers`);
     await this.$axios.get(`${url}/containers`)
         .then(response => {
           if(response.status!==200){
             this.loaded.responseError=true;
             return;
           }
           this.containers=response.data;
           if(this.containers.length===0){
             this.loaded.length=0;
           }
        }).catch(err=> {
          this.loaded.error="Error while requesting data please try again."
        });
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
