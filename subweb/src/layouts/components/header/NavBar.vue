<template>
  <header id="header" class="reveal">
    <h1>
      <a href="/">{{ siteName }}</a>
    </h1>
    <nav id="nav" ref="navBar">
      <ul>
        <li v-for="i in navBarItem" :key="i">
          <a :href="i.link" :target="i.target">{{ i.title }}</a>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script>
// import { navBarItem } from './navBarItem';
export default {
  name: 'NavBar',
  data() {
    return {
      navBarItem: [],
      siteName: '',
    };
  },
  created() {
    this.navBarItem = window.config.menuItem;
    this.siteName = window.config.siteName;
  },
};
</script>
