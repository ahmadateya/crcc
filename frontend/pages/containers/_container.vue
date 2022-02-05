<template>
  <div>
    <base-header class="pb-6">
    </base-header>
    <div class="container-fluid mt--6">
      <div class="row">
        <div class="col">
          <view-container :container="container"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import MainTable from "~/components/tables/RegularTables/MainTable";
import ViewContainer from "~/components/ViewContainer";
import Jsona from 'jsona';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    ViewContainer,
  },
  data() {
    return {
      container: {},
      loaded:{},
      valid:true
    }
  },
  async fetch() {
    await this.$axios.get(`${url}/containers/${this.$route.params.container}`)
        .then(response => {
          if(response.status!==200){
            this.loaded.responseError=true;
            return;
          }
          this.container = response.data;
        }).catch(err=> {
          this.loaded.error="Error while requesting data please try again."
        });
  },
};
</script>
