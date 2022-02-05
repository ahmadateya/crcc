<template>
  <div class="card">
    <div class="border-0 card-header">
      <h1 class="mb-0 bold ">Running Containers</h1>
    </div>

    <el-table class="table-responsive table-flush"
              header-row-class-name="thead-light"
              :data="rows">

      <el-table-column label="Container Names"
                       width="200"
                        min-width="310px"
                        prop="name"
                        sortable>
        <template v-slot="{row}">
          <div class="media align-items-center">
              <li v-for="name in row.names" :key="name"  class="list-unstyled">
                  <router-link :to="`containers/${name}`">
                      <span class="font-weight-600 name mb-0 text-sm">
                                        {{ name }}
                      </span>
                  </router-link>
              </li>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="Status"
                       min-width="200px"
                       prop="status"
                       width="250"
                       sortable>
        <template v-slot="{row}">
          <badge class="badge-dot mr-4" type="">
            <!--            <i :class="`bg-${row.ports[0].type}`"></i>-->
            <span class="status">{{row.status}}</span>
          </badge>
        </template>
      </el-table-column>

      <el-table-column label="Image"
        sortable>
        <template v-slot="{row}">
          {{ row.image }}
        </template>
      </el-table-column>


      <el-table-column label="Container ID"
                       prop="id"
                       min-width="140px"
                       sortable>
                       <template v-slot="{row}">
          {{ row.id }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import projects from './../projects'
import { Table, TableColumn, DropdownMenu, DropdownItem, Dropdown} from 'element-ui'
import { json } from 'd3';
export default {
  name: 'main-table',
  props: ["rows"],
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn,
    [Dropdown.name]: Dropdown,
    [DropdownItem.name]: DropdownItem,
    [DropdownMenu.name]: DropdownMenu,
  },
  
}
</script>
