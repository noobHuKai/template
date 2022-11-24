<template>
  <n-card :bordered="false" class="rounded-16px shadow-sm">
    <n-card :bordered="false" class="rounded-16px shadow-sm">
      <div ref="pieRef1" class="h-400px"></div>
    </n-card>
  </n-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Ref } from 'vue';
import { type ECOption, useEcharts } from '@/composables';
defineOptions({ name: 'DashboardMonitorhHeader' });

const pieOptions1 = ref<ECOption>({
  tooltip: {
    formatter: '{a} <br/>{b} : {c}%'
  },
  series: [
    {
      name: 'Pressure',
      type: 'gauge',
      progress: {
        show: true
      },
      detail: {
        valueAnimation: true,
        formatter: '{value}%'
      },
      data: [
        {
          value: 50,
          name: '内存'
        }
      ]
    }
  ]
}) as Ref<ECOption>;
const { domRef: pieRef1 } = useEcharts(pieOptions1);
</script>

<style scoped></style>
