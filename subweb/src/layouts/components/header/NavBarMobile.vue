<template>
  <div ref="showPanel">
    <div id="navButton"><a class="toggle" @click="showNavBarMobile"></a></div>
    <div id="navPanel">
      <nav>
        <a
          v-for="i in navBarItem"
          :key="i"
          :href="i.link"
          :target="i.target"
          class="link depth-0"
          style="-webkit-tap-highlight-color: rgba(0, 0, 0, 0)"
          ><span class="indent-0"></span>{{ i.title }}</a
        >
      </nav>
    </div>
  </div>
</template>

<script>
// import { navBarItem } from './navBarItem';
export default {
  name: 'NavBarMobile',
  data() {
    return {
      navBarItem: [],
    };
  },
  created() {
    this.navBarItem = window.config.menuItem;
    document.addEventListener('click', (e) => {
      if (this.$refs.showPanel) {
        let isSelf = this.$refs.showPanel.contains(e.target);
        if (!isSelf) {
          this.emitter.emit('isShowMobileNavBar', 'isOutBodyTouch');
        }
      }
    });
  },
  methods: {
    showNavBarMobile() {
      this.emitter.emit('isShowMobileNavBar', '');
    },
  },
};
</script>

<style></style>
