<template>
  <v-app>
    <v-content>
      <auto-deleter autoDelete="true"></auto-deleter>
    </v-content>
  </v-app>
</template>

<script>
import AutoDeleter from './components/EngineAutoDelete.vue';

export default {
  name: 'App',

  components: {
    AutoDeleter,
  },

  data: () => ({
    //
  }),
};
</script>
