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
          <div v-if="isScanned" class="col-xl-3 col-md-3">
            <h2>
              <i class="ni ni-chart-bar-32"></i> <span> Scan Results</span>
            </h2>
            <!-- Via multiple directive modifiers -->
            <div class="element"
                 v-for="(result, index) in scanData.results"
            >
              <b-button v-b-toggle="'collapse-' + index + '-details'"
                        :id="'collapse-' + index"
                        class="align-left"
                        v-bind:class="[result.passed ? 'text-success' : 'passed text-danger' ]"
              >
                <i v-if="result.passed" class="ni ni-check-bold"></i>
                <i v-else class="ni ni-fat-remove"></i>
                {{result.name}}
              </b-button>
              <div v-if="result.details !== ''">
                <b-collapse :id="'collapse-' + index + '-details'" class="mt-2">
                  <b-card-text>
<!--       TODO needs a reformat; for example display the details in separate lines/points ..etc             -->
                    {{ result.details }}
                  </b-card-text>
                </b-collapse>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import LoadingBar from "~/components/LoadingBar";
import ViewContainer from "~/components/ViewContainer";
import { BIcon } from 'bootstrap-vue'

import Jsona from 'jsona';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    BIcon,
    ViewContainer,
    LoadingBar,
  },
  data() {
    return {
      container: {},
      loaded:{},
      valid:true,
      isScanned:false,
      scanData: {
        results: [
          {
            name: 'xxxxxx',
            passed: true,
            details: '',
          },
          {
            name: 'yyyyyy',
            passed: false,
            details: 'qwqwqwqwq',
          },
        ]
      }
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
          // this.scanData = response.data;
          this.isScanned = true;
          this.$nuxt.$loading.finish()
        }).catch(err=> {
            // this.isScanned = false;
            this.loaded.error="Error while requesting data please try again."});
    }
  }
};
</script>
<style scoped>
.element {
  margin-bottom: 5px;
  display: block;
}
</style>
