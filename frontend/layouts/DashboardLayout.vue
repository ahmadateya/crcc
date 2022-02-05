<template>
  <div class="wrapper">
<!--    <notifications></notifications>-->
    <side-bar>
      <template slot-scope="props" slot="links">
        <sidebar-item
          :link="{
            name: 'Dashboard',
            icon: 'ni ni-shop text-primary',
            path: '/dashboard',
          }"
        >
        </sidebar-item>

        <sidebar-item
          opened
          :link="{
            name: 'Examples (API)',
            icon: 'fab fa-vuejs',
          }"
        >
          <sidebar-item
            opened
            :link="{ name: 'User Profile', path: '/examples/user-profile' }"
          />
          <sidebar-item
            opened
            :link="{
              name: 'User Management',
              path: '/examples/user-management',
            }"
          />
        </sidebar-item>
        <sidebar-item
          :link="{
            name: 'Icons',
            path: '/components/icons',
            icon: 'ni ni-planet',
          }"
        />
        <sidebar-item
          :link="{
            name: 'Tables',
            icon: 'ni ni-align-left-2 text-default',
            path: '/tables/regular',
          }"
        >
        </sidebar-item>

        <sidebar-item
          :link="{
            name: 'Google',
            icon: 'ni ni-pin-3',
            path: '/maps/google',
          }"
        >
        </sidebar-item>

        <hr
          class="my-3"
          style="
            border: 0;
            border-top: 1px solid rgba(0, 0, 0, 0.1);
            min-width: 80%;
            overflow: visible;
            box-sizing: content-box;
            height: 0;
          "
        />

      </template>
    </side-bar>
    <div class="main-content">
      <dashboard-navbar
        :type="$route.name === 'alternative' ? 'light' : 'default'"
      ></dashboard-navbar>

      <div @click="$sidebar.displaySidebar(false)">
        <nuxt></nuxt>
      </div>
      <content-footer v-if="!$route.meta.hideFooter"></content-footer>
    </div>
  </div>
</template>
<script>
/* eslint-disable no-new */
import PerfectScrollbar from "perfect-scrollbar";
import "perfect-scrollbar/css/perfect-scrollbar.css";

function hasElement(className) {
  return document.getElementsByClassName(className).length > 0;
}

function initScrollbar(className) {
  if (hasElement(className)) {
    new PerfectScrollbar(`.${className}`);
  } else {
    // try to init it later in case this component is loaded async
    setTimeout(() => {
      initScrollbar(className);
    }, 100);
  }
}

import DashboardNavbar from "~/components/layouts/argon/DashboardNavbar.vue";
import ContentFooter from "~/components/layouts/argon/ContentFooter.vue";
import DashboardContent from "~/components/layouts/argon/Content.vue";
import Vuex from "vuex";

export default {
  components: {
    DashboardNavbar,
    ContentFooter,
    DashboardContent,
  },
  methods: {
    initScrollbar() {
      let isWindows = navigator.platform.startsWith("Win");
      if (isWindows) {
        initScrollbar("scrollbar-inner");
      }
    },
  },
  mounted() {
    this.initScrollbar(), this.$store.dispatch("profile/me");
  },
};
</script>
<style lang="scss">
</style>
