<script setup lang="ts">
definePage({
  meta: {
    layout: "dashboardLayout",
  },
});
import { useDisplay } from "vuetify";
import { useTaskStore } from "@/stores/task";

const tab = ref(null);
const taskStore = useTaskStore();
const { xs } = useDisplay();
const addItemDialog = reactive({
  visible: false,
  title: "เพิ่มรายการงาน (Task List)",
});

const handleAddItem = () => {
  addItemDialog.visible = true;
};

onMounted(async () => {
  taskStore.search.search = "";
  taskStore.search.status = null;

  await taskStore.findAllTasks();
});

watch(
  () => tab.value,
  (newValue: any) => {
    console.log("tab.value", newValue);
    taskStore.search.status = null;
  }
);
</script>

<template>
  <div class="p-4 w-full h-full">
    <!-- tabbar -->
    <section class="w-full h-full">
      <v-card>
        <v-tabs v-model="tab">
          <v-tab value="table">
            <h1 class="text-md font-semibold">ตาราง</h1>
          </v-tab>
          <v-tab v-tab value="board">
            <h1 class="text-md font-semibold">บอร์ด</h1>
          </v-tab>
        </v-tabs>

        <v-card-text>
          <!-- display contetnt -->
          <v-tabs-window v-model="tab">
            <v-tabs-window-item value="table">
              <div class="border-none h-full">
                <v-card-title>
                  <section
                    class="w-full flex flex-col align-center justify-between md:flex-row"
                  >
                    <h1
                      class="text-2xl text-on-surface-varian mb-5 font-semibold"
                    >
                      รายการงาน
                    </h1>
                    <section
                      class="flex align-center gap-2 mb-2.5"
                      :class="{ 'w-full': xs }"
                    >
                      <!-- <v-btn prepend-icon="mdi-magnify" variant="outlined"
                        >ค้นหา
                        </v-btn
                      > -->
                      <v-btn
                        :block="xs"
                        prepend-icon="mdi-plus"
                        size="large"
                        @click="handleAddItem"
                        >เพิ่มรายการ
                      </v-btn>
                    </section>
                  </section>
                </v-card-title>
                <v-card-text>
                  <TableTask />
                </v-card-text>
              </div>
            </v-tabs-window-item>
            <v-tabs-window-item value="board">
              <BoardTask />
            </v-tabs-window-item>
          </v-tabs-window>
          <!-- display contetnt -->
        </v-card-text>
      </v-card>
    </section>
    <!-- tabbar -->
  </div>

  <AddTaskDialog
    v-model:isDialogVisible="addItemDialog.visible"
    v-model:title="addItemDialog.title"
  />
</template>
