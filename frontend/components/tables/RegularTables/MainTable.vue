<template>
  <div class="card">
    <div class="border-0 card-header">
      <h1 class="mb-0 bold ">{{ title }}</h1>
    </div>

    <el-table class="table-responsive table-flush"
              header-row-class-name="thead-light"
              :data="rows">

      <el-table-column v-if="title == 'Running Containers'" label="Container Names"
                       width="200"
                       min-width="310px"
                       prop="name"
                       sortable>
        <template v-slot="{row}">
          <div class="media align-items-center">
            <li v-for="name in row.names" :key="name" class="list-unstyled">
              <a :href="'/containers' + name" class="text-sm font-weight-bold">
                  <span class="font-weight-600 name mb-0 text-sm">
                          {{ name }}
                  </span>
              </a>
            </li>
          </div>
        </template>
      </el-table-column>

      <el-table-column :label="title == 'Running Containers'? 'Status' : title == 'Processes Rules'? 'Cmd' :
      title == 'Files Rules' ? 'File': title == 'Network Rules' ? 'Port': title == 'DNS Whitelist Rules'?'DNS':''
      "
                       min-width="200px"
                       prop="status"
                       width="200"
                       sortable>
        <template v-slot="{row}">

          <!--            <i :class="`bg-${row.ports[0].type}`"></i>-->
          {{
            title == "Running Containers" ? row.status :
                title == 'Processes Rules' ? row.cmd :
                    title == 'Files Rules' ? row.file : title == "Network Rules" ? row.port : title == 'DNS Whitelist Rules' ? row.dns : ''
          }}

        </template>
      </el-table-column>

      <el-table-column :label="title == 'Running Containers'? 'Image' : title == 'Processes Rules'? 'User' :
      title == 'Files Rules' || title == 'Network Rules' ? 'Description': title == 'DNS Whitelist Rules'?'Description':'' "
                       min-width="200px"
                       width="200"
                       sortable>
        <template v-slot="{row}">
          {{
            title == "Running Containers" ? row.image : title == 'Processes Rules' ? row.user
                :
                title == 'Files Rules' || title == 'Network Rules' || title == 'DNS Whitelist Rules' ? row.description : ''
          }}
        </template>
      </el-table-column>

      <el-table-column :label="title == 'Running Containers'? 'Id'
      : title == 'Processes Rules'? 'Description' :
      title == 'Files Rules' || title == 'Network Rules' || title == 'DNS Whitelist Rules' ? 'Impact': ''"
                       sortable>
        <template v-slot="{row}">
          {{
            title == "Running Containers" ? row.id : title == 'Processes Rules' ? row.description :
                title == 'Files Rules' || title == 'Network Rules' || title == 'DNS Whitelist Rules' ? row.impact : ''
          }}
        </template>
      </el-table-column>

      <el-table-column v-if="title == 'Processes Rules' || title == 'DNS Whitelist Rules'" :label="title == 'Processes Rules'?
      'Impact' : ''"
                       min-width="200px"
                       width="150"
                       sortable>
        <template v-slot="{row}">
          {{ title == 'Processes Rules' ? row.impact : title == 'DNS Whitelist Rules' ? row.white : '' }}
        </template>
      </el-table-column>

      <el-table-column v-if="title !== 'Running Containers'" label="Options"
                       sortable>
        <template scope="scope">
          <base-button size="md"

                       @click="title == 'Processes Rules' ? deleteProcess(scope.$index) :
                             title == 'Files Rules' ? deleteFile(scope.$index): title == 
                             'Network Rules' ? deletePort(scope.$index) : deleteDns(scope.$index)"
                       class="btn-default"
          >
            Delete
          </base-button>
          <base-button size="md"

                       @click="$emit('edit',{index: scope.$index,row: scope.row})"
                       class="scan-button"
          >
            Edit
          </base-button>
        </template>
      </el-table-column>


      <!--      <el-table-column label="Container ID"-->
      <!--                       prop="id"-->
      <!--                       min-width="140px"-->
      <!--                       sortable>-->
      <!--                       <template v-slot="{row}">-->
      <!--          {{ row.id }}-->
      <!--        </template>-->
      <!--      </el-table-column>-->
    </el-table>
  </div>
</template>
<script>
import {Table, TableColumn, DropdownMenu, DropdownItem, Dropdown} from 'element-ui'

const url = process.env.apiUrl;
export default {
  name: 'main-table',
  props: ["rows", "title"],
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn,
    [Dropdown.name]: Dropdown,
    [DropdownItem.name]: DropdownItem,
    [DropdownMenu.name]: DropdownMenu,
  },
  methods: {
    async deleteProcess(index) {
      await this.$axios.delete(`${url}/deleteprocess/` + index)
          .then(response => {
            if (response.status !== 200) {
              //this.loaded.responseError=true;
              return;
            }
            this.rows = response.data;
            if (this.rows.length === 0) {
              // this.loaded.length=0;
            }


          }).catch(err => {
            //this.loaded.error="Error while requesting data please try again."
          });
    },
    async deleteFile(index) {
      await this.$axios.delete(`${url}/deletefile/` + index)
          .then(response => {
            if (response.status !== 200) {
              //this.loaded.responseError=true;
              return;
            }
            this.rows = response.data;
            if (this.rows.length === 0) {
              // this.loaded.length=0;
            }


          }).catch(err => {
            //this.loaded.error="Error while requesting data please try again."
          });
    },
    async deletePort(index) {
      await this.$axios.delete(`${url}/deleteport/` + index)
          .then(response => {
            if (response.status !== 200) {
              //this.loaded.responseError=true;
              return;
            }
            this.rows = response.data;
            if (this.rows.length === 0) {
              // this.loaded.length=0;
            }


          }).catch(err => {
            //this.loaded.error="Error while requesting data please try again."
          });
    },
    async deleteDns(index) {
      await this.$axios.delete(`${url}/deleteDns/` + index)
          .then(response => {
            if (response.status !== 200) {
              //this.loaded.responseError=true;
              return;
            }
            this.rows = response.data;
            if (this.rows.length === 0) {
              // this.loaded.length=0;
            }


          }).catch(err => {
            //this.loaded.error="Error while requesting data please try again."
          });
    }
  }

}
</script>
