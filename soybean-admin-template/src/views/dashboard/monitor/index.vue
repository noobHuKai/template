<template>
  <n-space :vertical="true" :size="16">
    <data-card />
    <monitor-header />
  </n-space>
</template>

<script lang="ts" setup>
import { MonitorHeader, DataCard } from './components';
</script>

<style scoped></style>
