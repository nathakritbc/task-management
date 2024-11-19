<script lang="ts" setup>
import { itemBreadcrumbs } from "./index";
</script>

<template>
  <div>
    <v-breadcrumbs :items="itemBreadcrumbs">
      <template v-slot:divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </template>
    </v-breadcrumbs>
  </div>
</template>
