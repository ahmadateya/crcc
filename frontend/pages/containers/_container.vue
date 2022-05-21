<template>
  <div>
    <base-header class="pb-6">
      <div class="row align-items-center py-4">
        <!--  Back Button -->
        <div class="col-lg-1 col-1">
          <base-button size="sm" type="neutral" class="back-button" @click="$router.go(-1)">
            Back
          </base-button>
        </div>

      <!--   Site Title -->
        <div class="col-lg-11 col-11">
          <h1 class="h1"
              style="font-size: 2.5rem;
                  font-weight: 700; padding: 10px 5px; color: black; text-align: center">
            Container Runtime Compliance Checker
          </h1>
        </div>
      </div>
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
                <base-button size="lg"
                             @click="scanContainer"
                             class="scan-button"
                >
                  Scan
                </base-button>
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
          <!-- Analysis Part -->
          <div v-if="isScanned" class="col-xl-4 col-md-4">
            <h1>
              <i class="ni ni-chart-bar-32"></i> <span> Scan Results</span>
            </h1>
            <!-- Via multiple directive modifiers -->
            <div class="element"
                 v-for="(result, index) in scanData.results"
            >
              <b-button v-b-toggle="'collapse-' + index + '-details'"
                        :id="'collapse-' + index"
                        class="align-left w-75 p-3 mb-1 "
                        v-bind:class="[result.passed ? 'text-success' : 'passed text-danger' ]"
              >
                <i v-if="result.passed" class="ni ni-check-bold"></i>
                <i v-else class="ni ni-fat-remove"></i>
                {{result.title}}
              </b-button>
              <div v-if="result.details !== ''">
                <b-collapse :id="'collapse-' + index + '-details'" class="mt-2">
                      <!--    details is an array of objects  -->
                  <b-card-text v-if="result.details.charAt(0) === '['">
                    <ol>
                      <li v-for="detail in JSON.parse(result.details)">
                        <br>
                        <ul>
                          <li v-for="(value, key) in detail">
                            <b >{{key}} : </b> <span>{{value}}</span>
                          </li>
                        </ul>
                      </li>
                    </ol>
                  </b-card-text>
                    <!--      details is a string of data -->
                  <b-card-text v-else>
                    {{ result.details }}
                  </b-card-text>
                </b-collapse>
              </div>
            </div>
          </div>
          <!--  Chart Part  -->
          <div v-if="isScanned" class="col-xl-8 col-md-8">
            <h1>
              <i class="ni ni-chart-bar-32"></i> <span> Compliance Chart</span>
            </h1>
            <doughnut-chart
                       :height="250"
                       ref="doughnutChart"
                       :chart-data="donutChart.chartData"
            >
            </doughnut-chart>
            <pie-chart
                       :height="250"
                       ref="pieChart"
                       :chart-data="pieChart.chartData"
            >
            </pie-chart>
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
            title: '',
            passed: true,
            details: '',
          }
        ],
        compliance: '100%',
        failure: '0%',
      },
      donutChart: {
        chartData: {
          labels: [],
          datasets: [
            {
              backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
              data: [25, 20, 30, 22, 17, 100],
            },
          ],
        }
      },
      pieChart: {
        chartData: {
          labels: ['Failed', 'Passed'],
          datasets: [
            {
              backgroundColor: ['#DD1B16', '#41B883'],
              data: [75, 25],
            },
          ],
        }
      },
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
          this.isScanned = true;
          this.scanData = response.data;
          this.pieChart.chartData.datasets[0].data = [10, 90];
          // this.pieChart.chartData.datasets[0].data = [100 - this.scanData.compliance, this.scanData.compliance];
          this.$nuxt.$loading.finish()
        }).catch(err=> {
            this.isScanned = false;
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
.scan-button {
  color: #fff;
  padding: 7px 12px;
  border-color: #2496ed ;
  background-color: #2496ed;
}
.back-button {
  color: black;
  padding: 7px 12px;
  border-color: #2496ed;
}
</style>
