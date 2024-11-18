<script lang="ts" setup>
import { Tasks } from "@/types/Tasks";
import dayjs from "dayjs";
import { defineProps, PropType, defineEmits } from "vue";

const props = defineProps({
  task: {
    type: Object as PropType<Tasks>,
    required: true,
  },
});

const emit = defineEmits([
  "dragstart",
  "handleViewItem",
  "handleUpdateItem",
  "handleDeleteItem",
]);

const onDragStart = () => {
  emit("dragstart", props.task);
};

const handleUpdateItem = (item: Tasks) => {
  emit("handleUpdateItem", item);
};

const handleViewItem = (item: Tasks) => {
  emit("handleViewItem", item);
};

const handleDeleteItem = (id: string) => {
  emit("handleDeleteItem", id);
};
</script>

<template>
  <div>
    <div
      class="mb-2 rounded-lg cursor-pointer border-2 focus:outline-none focus:ring-2 bg-surface"
      :class="{
        'border-[LightSteelBlue] focus:ring-[LightSteelBlue]':
          task.status === 'draft',
        'border-yellow-200 focus:ring-yellow-200':
          task.status === 'in_progress',
        'border-green-400 focus:ring-green-400': task.status === 'completed',
      }"
      draggable="true"
      :data-task-id="task.id"
      @dragstart="onDragStart"
    >
      <v-card-text class="bg-surface rounded-lg">
        <section class="flex align-center justify-between">
          <div>
            <h1 class="text-base text-on-surface-variant">
              {{ task.title }}
            </h1>
            <p class="text-xs text-[gray]">
              {{ dayjs(task.created_at).format("DD/MM/YYYY") }}
            </p>
          </div>

          <v-menu transition="scale-transition">
            <template v-slot:activator="{ props }">
              <v-btn
                variant="text"
                color="gray-darken-1"
                icon="mdi-dots-vertical"
                v-bind="props"
              />
            </template>
            <v-list>
              <v-list-item @click="handleViewItem(task)">
                <template v-slot:append>
                  <v-icon icon="mdi-eye"></v-icon>
                </template>
                <v-list-item-title>ดูข้อมูล</v-list-item-title>
              </v-list-item>
              <v-list-item @click="handleUpdateItem(task)">
                <template v-slot:append>
                  <v-icon icon="mdi-pencil"></v-icon>
                </template>
                <v-list-item-title>แก้ไข</v-list-item-title>
              </v-list-item>
              <v-list-item @click="handleDeleteItem(task.id)">
                <template v-slot:append>
                  <v-icon icon="mdi-delete"></v-icon>
                </template>
                <v-list-item-title>ลบ</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </section>
      </v-card-text>
    </div>
  </div>
</template>

<style scoped>
[draggable="true"]:active {
  transform: scale(1.1);
  opacity: 0.8;
  cursor: grabbing;
}
</style>
