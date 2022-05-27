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
      <!--  Title Part    -->
      <div class="row">
        <h2 v-if="loaded.error">Error while fetching data please request it again.</h2>
        <h2 v-else-if="loaded.responseError">Please make sure of allowing the Rest API.</h2>
        <div class="col-xl-12">
          <div class="border card-header ">
            <div class="row">
              <div class="col-xl-10">
                <h1 class="mb-0 bold ">{{ containerName }} </h1>
              </div>
            </div>
          </div>
        </div>
      </div>
      <br><br>
      <!--  History Part    -->
      <div v-for="(scan, scanIndex) in history" class="col-xl-12">
        <div class="border card-body ">
        <div class="row">
          <!-- Analysis Part -->
          <div class="col-xl-6 col-md-6">
            <h1>
              <i class="ni ni-chart-bar-32"></i> <span> {{ scan.CreatedAt }}</span>
            </h1>
            <!-- Via multiple directive modifiers -->
            {{ }}
            <div class="element"
                 v-for="(result, index) in JSON.parse(scan.scan).results"
            >
              <b-button v-b-toggle="'collapse-' + scanIndex + '-' + index + '-details'"
                        :id="'collapse-' + scanIndex + '-' + index "
                        class="align-left w-75 p-3 mb-1 "
                        v-bind:class="[result.passed ? 'text-success' : 'passed text-danger' ]"
              >
                <i v-if="result.passed" class="ni ni-check-bold"></i>
                <i v-else class="ni ni-fat-remove"></i>
                {{ result.title }}
              </b-button>
              <div v-if="result.details !== ''">
                <b-collapse :id="'collapse-' + scanIndex + '-' + index + '-details'" class="mt-2">
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
          <div class="col-xl-6 col-md-6">
            <h1>
              <i class="ni ni-chart-pie-35"></i>
              <span> Compliance Percentage</span>
              <span
                  v-bind:class="[scan.result >= 50 ? 'text-success' : 'passed text-danger' ]"
              >{{scan.result}} %</span>
            </h1>
            <!--            <pie-chart-->
            <!--                        v-if="isScanned"-->
            <!--                       :height="250"-->
            <!--                       ref="pieChart"-->
            <!--                       :chart-data="chartData"-->
            <!--            >-->
            <!--            </pie-chart>-->
          </div>
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
      history: [],
      loaded: {},
      containerName: '',
      valid: true,
      chartData: {},
      options: {},
    }
  },
  async fetch() {
    await this.$axios.get(`${url}/containers/${this.$route.params.container}/history`)
        .then(response => {
          if (response.status !== 200) {
            this.loaded.responseError = true;
            return;
          }
          this.containerName = this.$route.params.container;
          this.history = response.data;
        }).catch(err => {
          this.loaded.error = "Error while requesting data please try again."
        });
  },
  methods: {
    async scanContainer() {
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
    }
  }
};
</script>
<style scoped>
.element {
  margin-bottom: 5px;
  display: block;
}

.back-button {
  color: black;
  padding: 7px 12px;
  border-color: #2496ed;
}

.card-header {
}
</style>
