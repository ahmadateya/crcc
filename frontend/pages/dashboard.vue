<template>
  <div>
    <base-header class="pb-6">
    </base-header>
    <div class="container-fluid mt--6">
      <div class="row">
        <div class="col">
          <main-table :rows="projects"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import projects from '~/components/tables/projects'
import users from '~/components/tables/users'
import LightTable from "~/components/tables/RegularTables/LightTable";
import MainTable from "~/components/tables/RegularTables/MainTable";

export default {
  layout: 'DashboardLayout',
  components: {
    MainTable,
    // [Dropdown.name]: Dropdown,
    // [DropdownItem.name]: DropdownItem,
    // [DropdownMenu.name]: DropdownMenu,
    // [Table.name]: Table,
    // [TableColumn.name]: TableColumn
  },
  data() {
    return {
      projects,
      // users
    };
  }
};
</script>
