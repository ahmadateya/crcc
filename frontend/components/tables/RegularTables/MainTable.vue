<template>
  <div class="card">
    <div class="border-0 card-header">
      <h1 class="mb-0 bold ">Running Containers</h1>
    </div>

    <el-table class="table-responsive table-flush"
              header-row-class-name="thead-light"
              :data="rows">
      <el-table-column label="Container Name" prop="id" width="200">
        min-width="310px"
        prop="name"
        sortable>
        <template v-slot="{row}">
          <div class="media align-items-center">
            <a href="#">
              <span class="font-weight-600 name mb-0 text-sm">{{row.names[0]}}</span>
            </a>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Container ID"
                       prop="budget"
                       min-width="140px"
                       sortable>
                       <template v-slot="{row}">
          {{ row.id }}
          
        </template>
      </el-table-column>

      <el-table-column label="Status"
                       min-width="170px"
                       prop="status"
                       sortable>
        <template v-slot="{row}">
          <badge class="badge-dot mr-4" type="">
            <i :class="`bg-${row.ports[0].type}`"></i>
            <span class="status">{{row.status}}</span>
          </badge>
        </template>
      </el-table-column>

      <el-table-column label="Image" min-width="190px">
       <template v-slot="{row}">
          {{ row.image }}
        </template>
      </el-table-column>

      
      <el-table-column min-width="180px">
        <template v-slot="{row}">
          <el-dropdown trigger="click" class="dropdown">
                    <span class="btn btn-sm btn-icon-only text-light">
                      <i class="fas fa-ellipsis-v mt-2"></i>
                    </span>
            <el-dropdown-menu class="dropdown-menu dropdown-menu-arrow show" slot="dropdown">
              <a class="dropdown-item" href="#">Action</a>
              <a class="dropdown-item" href="#">Another action</a>
              <a class="dropdown-item" href="#">Something else here</a>
            </el-dropdown-menu>
          </el-dropdown>
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
