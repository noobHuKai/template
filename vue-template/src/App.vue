<template>
  <n-config-provider :theme="theme.naiveTheme" class="h-full">
      <naive-provider>
          <router-view />
      </naive-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { useThemeStore } from "@/store";

const theme = useThemeStore();
</script>
<style scoped></style>
