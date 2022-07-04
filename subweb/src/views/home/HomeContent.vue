<template>
  <section id="main" class="container medium">
    <home-header />
    <home-form />
  </section>
</template>

<script>
import HomeHeader from './HomeHeader';
import HomeForm from './HomeForm';
export default {
  name: 'HomeContent',
  components: {
    HomeHeader,
    HomeForm,
  },
};
</script>
