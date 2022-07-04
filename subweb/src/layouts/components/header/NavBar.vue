<template>
  <header id="header" class="reveal">
    <h1><a href="/">Subconverter Web</a></h1>
    <nav id="nav" ref="navBar">
      <ul>
        <li v-for="i in navBarItem" :key="i" :class="i.style">
          <a :href="i.link" :class="i.class" :target="i.target">{{
            i.title
          }}</a>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script>
import { navBarItem } from './navBarItem';
export default {
  name: 'NavBar',
  data() {
    return {
      navBarItem: [],
    };
  },
  created() {
    this.navBarItem = navBarItem;
  },
};
</script>
