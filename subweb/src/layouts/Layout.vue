<template>
  <body :class="bodyClass">
    <nav-bar-mobile />
    <div id="page-wrapper">
      <header-view />
      <content-view />
      <footer-view />
    </div>
  </body>
</template>

<script>
import NavBarMobile from './components/header/NavBarMobile';
import HeaderView from './HeaderView';
import ContentView from './ContentView';
import FooterView from './FooterView';

export default {
  name: 'LayoutView',
  components: {
    NavBarMobile,
    HeaderView,
    ContentView,
    FooterView,
  },
  data() {
    return {
      screenWidth: document.body.clientWidth,
      bodyClass: ['', ''],
    };
  },
  created() {
    this.emitter.on('isLanding', (isLanding) => {
      this.bodyClass[0] = isLanding;
    });
    // 监听事件开关 NavBarMobile
    this.emitter.on('isShowMobileNavBar', (isOutBodyTouch) => {
      if (isOutBodyTouch == 'isOutBodyTouch') {
        this.bodyClass[1] = '';
      } else if (this.bodyClass[1] == '') {
        this.bodyClass[1] = 'navPanel-visible';
      } else {
        this.bodyClass[1] = '';
      }
    });
  },
};
</script>
