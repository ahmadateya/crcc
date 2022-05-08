<template>
  <div>
    <base-header class="pb-6">
    </base-header>
    <div class="container-fluid mt--6">
      <!--   Container Table   -->
      <div class="row">
        <h2 v-if="loaded.error">Error while fetching data please request it again.</h2>
        <h2 v-else-if="loaded.responseError">Please make sure of allowing the Rest API.</h2>
        <div class="col-xl-12">
          <div class="border-0 card-header">
            <div class="row">
              <div class="col-xl-10">
                <h1 class="mb-0 bold ">{{container.name}} </h1>
              </div>
              <div class="col-xl-2">
                <base-button size="lg" @click="scanContainer">Scan</base-button>
             </div>
            </div>
          </div>
          <view-container :container="container"/>
        </div>
      </div>

      <!--  Scan Part    -->
      <div class="col-xl-12">
        <div class="row">
          <h2 v-if="loaded.error">Error while fetching data please request it again.</h2>
          <h2 v-else-if="loaded.responseError">Please make sure of allowing the Rest API.</h2>
          <loading-bar></loading-bar>
          <h2 v-if="isScanned">Scanned</h2>
          <div>
            <!-- Via multiple directive modifiers -->
            <b-button v-b-toggle.collapse-a.collapse-b>Toggle Collapse A and B</b-button>

            <!-- Via space separated string of IDs passed to directive value -->
            <b-button v-b-toggle="'collapse-a collapse-b'">Toggle Collapse A and B</b-button>

            <!-- Via array of string IDs passed to directive value -->
            <b-button v-b-toggle="['collapse-a', 'collapse-b']">Toggle Collapse A and B</b-button>

            <!-- Elements to collapse -->
            <b-collapse id="collapse-a" class="mt-2">
              <b-card>I am collapsible content A!</b-card>
            </b-collapse>
            <b-collapse id="collapse-b" class="mt-2">
              <b-card>I am collapsible content B!</b-card>
            </b-collapse>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import LoadingBar from "~/components/LoadingBar";
import ViewContainer from "~/components/ViewContainer";
import Jsona from 'jsona';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    ViewContainer,
    LoadingBar,
  },
  data() {
    return {
      container: {},
      loaded:{},
      valid:true,
      isScanned:false,
      scanData: {}
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
  methods: {
    async scanContainer() {
      this.isScanned = true;
      // this.$nextTick(() => {
        this.$nuxt.$loading.start()
      // })

      await this.$axios.get(`${url}/containers/${this.$route.params.container}/scan`)
        .then(response => {
          if(response.status!==200){
            this.loaded.responseError=true;
            return;
          }
          this.scanData = response.data;
          this.isScanned = true;
          this.$nuxt.$loading.finish()
        }).catch(err=> {
            // this.isScanned = false;
            this.loaded.error="Error while requesting data please try again."});
    }
  }
};
</script>
