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
              <div class="col-xl-9">
                <h1 class="mb-0 bold ">{{ container.name }} </h1>
              </div>
              <div class="col-xl-3">
                <div class="input-group mb-3">
  <input id="cnninput" type="text" class="form-control" placeholder="Enter A Path" aria-label="CNN Scan" aria-describedby="basic-addon2">
  <div class="input-group-append">
    <base-button size="lg"
                            @click="cnnScan"
                             
                             class="scan-button"
                >
                  CNN Scan
                </base-button>
  </div>
</div>
                
                <base-button size="lg"
                             @click="scanContainer"
                             class="scan-button"
                >
                  Scan
                </base-button>

                

                  <a type="button" class="btn history-button"
                     :href="'/containers' + container.name + '/history'">
                    History
                  </a>
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
          <div v-if="isScanned" class="col-xl-6 col-md-6">
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
                {{ result.title }}
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
                            <b>{{ key }} : </b> <span>{{ value }}</span>
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
          <div v-if="isScanned" class="col-xl-6 col-md-6">
            <h1>
              <i class="ni ni-chart-pie-35"></i>
              <span> Compliance Percentage</span>
              <span
                  v-bind:class="[scanData.compliance >= 50 ? 'text-success' : 'passed text-danger' ]"
              >{{ scanData.compliance }} %</span>
            </h1>
            <pie-chart
                v-if="isScanned"
                :height="250"
                ref="pieChart"
                :chart-data="chartData"
            >
            </pie-chart>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import LoadingBar from "@/components/LoadingBar";
import ViewContainer from "@/components/ViewContainer";
import {BIcon} from 'bootstrap-vue'

import Jsona from 'jsona';
import BaseInput from "../../../components/argon-core/Inputs/BaseInput.vue";
import BaseAlert from "../../../components/argon-core/BaseAlert.vue";

const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    BIcon,
    ViewContainer,
    LoadingBar,
    BaseInput,
    BaseAlert
},
  data() {
    return {
      container: {},
      loaded: {},
      valid: true,
      isScanned: false,
      scanData: {
        results: [
          {
            title: '',
            passed: true,
            details: '',
          }
        ],
        compliance: '',
      },
      chartData: {},
      options: {},
    }
  },
  async fetch() {
    await this.$axios.get(`${url}/containers/${this.$route.params.container}`)
        .then(response => {
          if (response.status !== 200) {
            this.loaded.responseError = true;
            return;
          }
          this.container = response.data;
        }).catch(err => {
          this.loaded.error = "Error while requesting data please try again."
        });
  },
  methods: {
    async scanContainer() {
      this.scanData=[];
      this.isScanned = true;
      this.$nuxt.$loading.start()

      await this.$axios.get(`${url}/containers/${this.$route.params.container}/scan`)
          .then(response => {
            if (response.status !== 200) {
              this.loaded.responseError = true;
              return;
            }
            this.scanData = response.data;
            this.chartData = {
              labels: ['Failed', 'Passed'],
              datasets: [
                {
                  backgroundColor: ['#DD1B16', '#41B883'],
                  data: [100 - this.scanData.compliance, this.scanData.compliance],
                },
              ],
            };
            this.isScanned = true;
            this.$nuxt.$loading.finish()
          }).catch((err) => {
            this.$nuxt.$loading.finish()
            this.isScanned = false;
            this.loaded.error = "Error while requesting data please try again.";
            console.log(err, "error");
          });
    },
    async cnnScan(){
      
     
      
      if(document.getElementById("cnninput").value==="" || document.getElementById("cnninput").value===undefined){
        return
      }
      this.scanData=[];
      this.isScanned=true;
      this.$nuxt.$loading.start();

      await this.$axios.get("http://172.21.0.4:5001/container/cnnscan?container=6c3e8d39dbe5"+
      "&path="+document.getElementById("cnninput").value).then(response=>{
        if (response.status !== 200) {
              this.loaded.responseError = true;
              return;
            }
        this.scanData = response.data;
            this.chartData = {
              labels: ['Failed', 'Passed'],
              datasets: [
                {
                  backgroundColor: ['#DD1B16', '#41B883'],
                  data: [100-this.scanData.compliance, this.scanData.compliance],
                },
              ],
            };
            
            this.isScanned = true;
            this.$nuxt.$loading.finish()
          }).catch((err) => {
            this.$nuxt.$loading.finish()
            this.isScanned = false;
            this.loaded.error = "Error while requesting data please try again.";
            console.log(err, "error");
          });

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
  border-color: #2496ed;
  background-color: #2496ed;
}

.back-button {
  color: black;
  padding: 7px 12px;
  border-color: #2496ed;
}

.history-button {
  color: #fff;
  padding: 7px 12px;
  border-color: #485967;
  background-color: #485967;
}
</style>
