<template>
  <div>
    <!-- Base Header -->
    <base-header class="pb-6">
      <h1 class="h1"
          style="font-size: 2.5rem;
                  font-weight: 700; padding: 10px; color: black; text-align: center">
        Container Runtime Compliance Checker
      </h1>

      <!-- Card stats -->
      <div class="row">
        <div class="col-xl-4 col-md-6">
          <stats-card
              title="Total Containers"
              type="gradient-blue"
              :sub-title="containers.length"
              icon="ni ni-chart-pie-35"
          >
          </stats-card>
           
        </div>
        
        <div class="col-xl-4 col-md-6">
          <stats-card
              title="Total Networks"
              type="gradient-blue"
              :sub-title="networks.size()"
              icon="ni ni-chart-pie-35"
          >
          </stats-card>
           
        </div>

        <div class="col-xl-4 col-md-6">
          <stats-card
              title="Total Imagess"
              type="gradient-blue"
              :sub-title="images.size()"
              icon="ni ni-chart-pie-35"
          >
          </stats-card>
           
        </div>
      </div>


      
    </base-header>
    <!-- End Base Header -->

    <!--  Main Section in the Page  -->
    
    <div class="container-fluid">
      <div class="mt--6">
        <!-- Containers Table-->
        <div class="row">
          <div class="col-xl-12">
            <main-table v-if="containers.length!==0" :rows="containers"/>
            <h2 v-else-if="loaded.error">Error while fetching data please request it again.</h2>
            <h2 v-else-if="loaded.reponseError">Please make sure of allowing the Rest API.</h2>
            <h2 v-else-if="loaded.length===0">No Running Containers</h2>
          </div>

          <!--  Pie chart   -->
          
          <!--  end Pie chart   -->


        </div>
        <!-- End Containers Table-->
        <div class="row">
          <div class="col-xl-4">
            <card header-classes="bg-transparent">
              <div slot="header" class="row align-items-center">
                <div class="col">
<!--                  <h6 class="text-uppercase text-muted ls-1 mb-1">Operating Systems</h6>-->
                  <h5 class="h3 mb-0">Containers Images</h5>
                </div>
              </div>
              <pie-chart v-if="containers.length!=0"
                  :height="350"
                  ref="pieChart"
                  :chart-data="pieChart.chartData"
                 
              >
              </pie-chart>
            </card>
          </div>

           <div class="col-xl-4">
            <card header-classes="bg-transparent">
              <div slot="header" class="row align-items-center">
                <div class="col">
<!--                  <h6 class="text-uppercase text-muted ls-1 mb-1">Operating Systems</h6>-->
                  <h5 class="h3 mb-0">Containers Networks</h5>
                </div>
              </div>
              <pie-chart v-if="containers.length!=0"
                  :height="350"
                  ref="pieChart"
                  :chart-data="networksPieChart.chartData"
              >
              </pie-chart>
            </card>
          </div>

           <div class="col-xl-4">
            <card header-classes="bg-transparent">
              <div slot="header" class="row align-items-center">
                <div class="col">
<!--                  <h6 class="text-uppercase text-muted ls-1 mb-1">Operating Systems</h6>-->
                  <h5 class="h3 mb-0">Containers Status</h5>
                </div>
              </div>
              <pie-chart v-if="containers.length!=0"
                  :height="350"
                  ref="pieChart"
                  :chart-data="statusPieChart.chartData"
                  
              >
              </pie-chart>
            </card>
          </div>
        </div>
        <!--Charts-->
<!--        <div class="row">-->
<!--          <div class="col-xl-8">-->
<!--            <card type="default" header-classes="bg-transparent">-->
<!--              <div slot="header" class="row align-items-center">-->
<!--                <div class="col">-->
<!--                  <h6 class="text-light text-uppercase ls-1 mb-1">Overview</h6>-->
<!--                  <h5 class="h3 text-white mb-0">Sales value</h5>-->
<!--                </div>-->
<!--                <div class="col">-->
<!--                  <ul class="nav nav-pills justify-content-end">-->
<!--                    <li class="nav-item mr-2 mr-md-0">-->
<!--                      <a-->
<!--                          class="nav-link py-2 px-3"-->
<!--                          href="#"-->
<!--                          :class="{ active: bigLineChart.activeIndex === 0 }"-->
<!--                          @click.prevent="initBigChart(0)"-->
<!--                      >-->
<!--                        <span class="d-none d-md-block">Month</span>-->
<!--                        <span class="d-md-none">M</span>-->
<!--                      </a>-->
<!--                    </li>-->
<!--                    <li class="nav-item">-->
<!--                      <a-->
<!--                          class="nav-link py-2 px-3"-->
<!--                          href="#"-->
<!--                          :class="{ active: bigLineChart.activeIndex === 1 }"-->
<!--                          @click.prevent="initBigChart(1)"-->
<!--                      >-->
<!--                        <span class="d-none d-md-block">Week</span>-->
<!--                        <span class="d-md-none">W</span>-->
<!--                      </a>-->
<!--                    </li>-->
<!--                  </ul>-->
<!--                </div>-->
<!--              </div>-->
<!--              <line-chart-->
<!--                  :height="350"-->
<!--                  ref="bigChart"-->
<!--                  :chart-data="bigLineChart.chartData"-->
<!--                  :extra-options="bigLineChart.extraOptions"-->
<!--              >-->
<!--              </line-chart>-->
<!--            </card>-->
<!--          </div>-->
<!--        </div>-->
        <!-- End charts-->
      </div>
    </div>
    <!--  End Main Section in the Page  -->
  </div>
</template>
<script>
// Charts
import * as chartConfigs from "@/components/argon-core/Charts/config";
import LineChart from "@/components/argon-core/Charts/LineChart";
import BarChart from "@/components/argon-core/Charts/BarChart";
import PieChart from "@/components/argon-core/Charts/PieChart";
import StatsCard from "@/components/argon-core/Cards/StatsCard";

// tables
import MainTable from "~/components/tables/RegularTables/MainTable";
import Jsona from 'jsona';
import { map } from 'd3';
const url = process.env.apiUrl;
const jsona = new Jsona();

export default {
  layout: 'DashboardLayout',

  components: {
    MainTable,
    LineChart,
    BarChart,
    StatsCard,
    PieChart,
  },
  data() {
    return {
      // containers related data
      containers: [],
      networks: new map(),
      images: new map(), 
      status: new map(),  
      loaded:{},
      valid:true,
      // charts related data
      bigLineChart: {
        allData: [
          [0, 20, 10, 30, 15, 40, 20, 60, 60],
          [0, 20, 5, 25, 10, 30, 15, 40, 40],
        ],
        activeIndex: 0,
        chartData: {
          datasets: [
            {
              label: "Performance",
              data: [0, 20, 10, 30, 15, 40, 20, 60, 60],
            },
          ],
          labels: ["May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
        },
        extraOptions: chartConfigs.blueChartOptions,
      },
      redBarChart: {
        chartData: {
          labels: ["Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
          datasets: [
            {
              label: "Sales",
              data: [25, 20, 30, 22, 17, 100],
            },
          ],
        },
      },
      pieChart: {
        chartData: {
          labels: [],
          datasets: [
            {
              //backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
              data: [25, 20, 30, 22, 17, 100],
            },
          ],
        }
      },
      networksPieChart: {
        chartData: {
          labels: [],
          datasets: [
            {
              //backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
              data: [25, 20, 30, 22, 17, 100],
            },
          ],
        }
      },
      statusPieChart: {
        chartData: {
          labels: [],
          datasets: [
            {
              //backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
              data: [25, 20, 30, 22, 17, 100],
            },
          ],
        }
      },
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
           
           this.networkandImageView();
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
  methods: {
    initBigChart(index) {
      let chartData = {
        datasets: [
          {
            label: "Performance",
            data: this.bigLineChart.allData[index],
          },
        ],
        labels: [],
      };
      this.bigLineChart.chartData = chartData;
      this.bigLineChart.activeIndex = index;
    },
    networkandImageView(){
      for(var i=0;i<this.containers.length;i++){
             if(this.containers[i].networksettings.Networks != undefined){
               for(var name in this.containers[i].networksettings.Networks){
                 if(this.containers[i].networksettings.Networks.hasOwnProperty(name)){
                   if(this.networks.get(name) == undefined){
                     this.networks.set(name,1);
                   }else{
                     this.networks[name]+=1
                   }
                 }
               }
             }
             
           }
           this.networksPieChart.chartData.labels=this.networks.keys()
           this.networksPieChart.chartData.datasets[0].data=this.networks.values()
    this.imagesView()
    this.statusView()
           
           
    },
    imagesView(){
      for(var i=0;i<this.containers.length;i++){
             if(this.images.get(this.containers[i].image) == undefined){
               this.images.set(this.containers[i].image,1)
             }else{
               this.images[this.containers[i].image]+=1
             }
             
           }
           this.pieChart.chartData.labels=this.images.keys()
           
          this.pieChart.chartData.datasets[0].data=this.images.values()
    },
    statusView(){
      for(var i=0;i<this.containers.length;i++){
             if(this.status.get(this.containers[i].names[0]) == undefined){
               
               this.status.set(this.containers[i].names[0],this.containers[i].status.match(/(\d+)/)[0])
             }
             
           }
           this.statusPieChart.chartData.labels=this.status.keys()
           
          this.statusPieChart.chartData.datasets[0].data=this.status.values()
    }
  },
  mounted() {
    this.initBigChart(0);
  },
};
</script>
