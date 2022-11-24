<template>
  <n-grid cols="s:1 m:2 l:4" responsive="screen" :x-gap="16" :y-gap="16">
    <n-grid-item v-for="item in cardData" :key="item.id">
      <gradient-bg class="h-100px" :start-color="item.colors[0]" :end-color="item.colors[1]">
        <h3 class="text-16px">{{ item.title }}</h3>
        <div class="flex justify-between pt-12px">
          <div class="text-30px text-white dark:text-dark">{{ item.value }}</div>
        </div>
      </gradient-bg>
    </n-grid-item>
  </n-grid>
</template>

<script setup lang="ts">
import { GradientBg } from './components';

defineOptions({ name: 'DashboardAnalysisDataCard' });

interface CardData {
  id: string;
  title: string;
  value: number;
  colors: [string, string];
}

const cardData: CardData[] = [
  {
    id: 'cpu',
    title: 'CPU',
    value: 1000000,
    colors: ['#ec4786', '#b955a4']
  },
  {
    id: 'memory',
    title: '内存',
    value: 234567.89,
    colors: ['#865ec0', '#5144b4']
  },
  {
    id: 'disk',
    title: '磁盘',
    value: 666666,
    colors: ['#56cdf3', '#719de3']
  }
];
</script>

<style scoped></style>
