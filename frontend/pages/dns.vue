<template>
  <div>
    <!-- Base Header -->
    <base-header class="pb-6">
      <h1 class="h1"
          style="font-size: 2.5rem;
                  font-weight: 700; padding: 10px; color: black; text-align: center">
        Container Runtime Compliance Checker
      </h1>

     <div  class="col-xl-15" style="text-align: center;">
            
            <!-- Via multiple directive modifiers -->
            <div v-on:click="checkToggle()" class="element" style="text-align: center;"
            >
              <b-button v-b-toggle="'collapse-' + '-details'"
                        id="collapse"
                        class="align-center w-75 p-3 mb-1 passed btn-default"
                        
              >
                
                Add DNS
              </b-button>
              <div >
                <b-collapse :id="'collapse-'  + '-details'" class="mt-2">
                  <!--    details is an array of objects  -->
                  
                  <b-card-text >
                    <form @submit.prevent="edit? editRule(): addDns()">
                    <div class="input-group">
  
  
  <input id="dns" type="text" class="form-control" placeholder="Dns:">
  <input id="desc" type="text" class="form-control" placeholder="Description:">
  <input id="impact" type="text" class="form-control" placeholder="Impact:"
  pattern="(high|medium|low)" 
                  title="Impact have to be one of these : high, medium, low">
  
 <input type="checkbox" class="btn-check" id="btn-check-2-outlined" checked autocomplete="off">
<label size="md" style="background-color: white; color: black;" class="btn btn-outline-secondary" for="btn-check-2-outlined">White</label><br>
  
  <div class="input-group-prepend">
     <base-button size="md"
                            @click="edit? editRule(): addDns()"
                             
                             class="scan-button"
                >
                  {{ edit? 'Edit Rule' : 'Add Rule' }}
                </base-button>
  </div>
</div>
</form>
                  </b-card-text>
                </b-collapse>
              </div>
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
            <main-table @edit="updateToggle" v-if="dns.length!==0" :rows="dns" title="DNS Whitelist Rules"/>
            <h2 v-else-if="loaded.error">Error while fetching data please request it again.</h2>
            <h2 v-else-if="loaded.reponseError">Please make sure of allowing the Rest API.</h2>
            <h2 v-else-if="loaded.length===0">No Running Containers</h2>
          </div>

          <!--  Pie chart   -->
          
          <!--  end Pie chart   -->


        </div>
        <!-- End Containers Table-->
    
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
<!--                          href="#"--></form>
<!--                   document.getElementById('your_input_id').validity.valid       :class="{ active: bigLineChart.activeIndex === 1 }"-->
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
import Swal from "sweetalert2";

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
      dns: [], 
      loaded:{},
      valid:true,
      edit: false,
      editIndex: -1,
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
   this.$nextTick(() => {
      this.$nuxt.$loading.start()
    })
     await this.$axios.get(`${url}/dnsrules`)
         .then(response => {
           if(response.status!==200){
             this.loaded.responseError=true;
             return;
           }
           this.dns=response.data;
           if(this.dns.length===0){
             this.loaded.length=0;
           }
           
           
        }).catch(err=> {
          this.loaded.error="Error while requesting data please try again."
        });
        this.$nuxt.$loading.finish()
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
  async addDns(){
    if(document.getElementById('dns').value!=""){
  var data={dns: document.getElementById("dns").value,white: document.getElementById("btn-check-2-outlined").checked,
      description: document.getElementById("desc").value , impact:document.getElementById("impact").value};

      await this.$axios.post(`${url}/adddns`,data)
         .then(response => {
           if(response.status!==200){
             //this.loaded.responseError=true;
             Swal.fire({icon:'error',title:"Error Creating The Rule"})
             return;
           }
           Swal.fire({icon:'success',title:"The Rule Created"})
           this.dns=response.data;
      
           
        }).catch(err=> {
          Swal.fire({icon:'error',title:"Error Creating The Rule"})
          //this.loaded.error="Error while requesting data please try again."
        });
      
    }
  },checkToggle(){
    if (document.getElementById("collapse").classList.contains("collapsed")){
      this.edit=false
    document.getElementById("dns").value="";
    document.getElementById("btn-check-2-outlined").checked=true;
    document.getElementById("desc").value="";
    document.getElementById("impact").value="";
    }
  },
  updateToggle(value){
    
    if (document.getElementById("collapse").classList.contains("collapsed")){
      document.getElementById("collapse").click();
    }
    document.getElementById("dns").value=value.row.dns;
    document.getElementById("btn-check-2-outlined").checked=value.row.white;
    document.getElementById("desc").value=value.row.description;
    document.getElementById("impact").value=value.row.impact;
    this.edit=true;
    this.editIndex=value.index;
    window.scrollTo(0, 0);
      },
      async editRule(){
      if(document.getElementById("dns").value != ""  ){
  var data={dns: document.getElementById("dns").value,white: document.getElementById("btn-check-2-outlined").checked,
      description: document.getElementById("desc").value , impact:document.getElementById("impact").value};

      await this.$axios.put(`${url}/editdns/`+this.editIndex,data)
         .then(response => {
           if(response.status!==200){
             //this.loaded.responseError=true;
             Swal.fire({icon:'error',title:"Error Editing The Rule"})
             return;
           }
           Swal.fire({icon:'success',title:"The Rule Edited"})
           this.dns=response.data
           
        }).catch(err=> {
          Swal.fire({icon:'error',title:"Error Editing The Rule"})
          //this.loaded.error="Error while requesting data please try again."
        });
      
    }
  }
  }
  
};
</script>
