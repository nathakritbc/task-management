<script setup lang="ts">
import { useTaskStore } from "@/stores/task";
import { useDisplay } from "vuetify";
import { Tasks } from "@/types/Tasks";
import _ from "lodash";

const headers: any = [
  { title: "หัวข้องาน", align: "start", key: "title", width: "23%" },
  { title: "รายละเอียยด", align: "start", key: "description", width: "40%" },
  { title: "สถานะ", align: "center", key: "status", width: "20%" },
  { title: "Acctions", align: "center", key: "actions" },
];
const { xs, sm } = useDisplay();
const itemsPerPage = ref(5);
const taskStore = useTaskStore();
const items = computed<Tasks[]>(() => taskStore.tasks.result);

const viewItemDialog = reactive({
  visible: false,
  title: "รายละเอียดข้อมูล",
  description: "",
  status: "",
});

const updateItemDialog = reactive({
  visible: false,
  title: "เเก้ไขรายการ",
  description: "",
  status: "",
});

const deleteItemDialog = reactive({
  id: "",
  visible: false,
  title: "ยืนยันการทำรายการ?",
  text: "ถ้าหากยืนยันการทำรายการนี้ จะไม่สามารถกู้คืนได้",
});

const handleViewItem = (item: Tasks) => {
  if (item.id) {
    taskStore.task = item;
  }
  viewItemDialog.visible = true;
};

const handleUpdateItem = (item: Tasks) => {
  if (item.id) {
    taskStore.task = item;
  }
  updateItemDialog.visible = true;
};

const handleDeleteItem = (id: string) => {
  deleteItemDialog.title = `ยืนยันการลบรายการ?`;
  deleteItemDialog.visible = true;
  deleteItemDialog.id = id;
};

const deleteItem = () => {
  taskStore.deleteTaskById(deleteItemDialog.id);
};

const renderColorStatus = (status: string): string => {
  switch (status) {
    case "draft":
      return "blue-grey-lighten-1";

    case "in_progress":
      return "warning";

    case "completed":
      return "success";

    default:
      return "blue-grey-darken-1";
  }
};

const renderTextStatus = (status: string): string => {
  switch (status) {
    case "draft":
      return "ฉบับร่าง";

    case "in_progress":
      return "กําลังดําเนินการ";

    case "completed":
      return "ดำเนินการเสร็จสิ้น";

    default:
      return "ฉบับร่าง";
  }
};

const handleSearch = async () => {
  await taskStore.findAllTasks();
};

// ใช้ Lodash debouncing
const debouncedSearch = _.debounce(handleSearch, 400);

watch(taskStore.search, () => {
  debouncedSearch();
});

onMounted(() => {
  if (taskStore.tasks.result.length === 0) {
    taskStore.findAllTasks();
  }
});
</script>

<template>
  <div>
    <section class="w-full flex align-center mb-4 gap-3">
      <v-text-field
        v-model="taskStore.search.search"
        class="w-full none"
        variant="outlined"
        single-line
        hide-details
        label="ค้นหา"
        clearable
      />
      <v-select
        v-show="!xs"
        v-model="taskStore.search.status"
        class="grow w-[30%]"
        :class="{
          'w-full': xs,
          'w-[50%]': sm,
        }"
        variant="outlined"
        single-line
        hide-details
        clearable
        :item-title="(item) => item.label"
        :item-value="(item) => item.value"
        label="สถานะ"
        :items="taskStore.statusList"
      ></v-select>
    </section>
    <v-data-table
      :loading="taskStore.tasks.loading"
      class="w-full h-[365px] table grow"
      :headers="headers"
      :items="items"
      :items-per-page="itemsPerPage"
    >
      <!-- หัวข้องาน -->
      <template #item.title="{ item }">
        <div class="line-clamp-2 my-2 flex align-start">
          <p>{{ item.title }}</p>
        </div>
      </template>
      <!-- หัวข้องาน -->

      <!-- รายละเอียยด -->
      <template #item.description="{ item }">
        <p class="line-clamp-2 my-2">
          {{ item.description }}
        </p>
      </template>
      <!-- รายละเอียยด -->

      <!-- สถานะ -->
      <template #item.status="{ item }">
        <v-chip
          class="currsor-pointer rounded-md"
          @click="handleUpdateItem(item)"
          label
          size="small"
          :color="renderColorStatus(item.status)"
        >
          {{ renderTextStatus(item.status) }}
        </v-chip>
      </template>
      <!-- สถานะ -->

      <!-- actions -->
      <template #item.actions="{ item }">
        <section class="flex align-center justify-center">
          <v-menu transition="scale-transition">
            <template v-slot:activator="{ props }">
              <v-icon icon="mdi-dots-vertical" v-bind="props"></v-icon>
            </template>
            <v-list>
              <v-list-item @click="handleViewItem(item)">
                <template v-slot:append>
                  <v-icon icon="mdi-eye"></v-icon>
                </template>
                <v-list-item-title>ดูข้อมูล</v-list-item-title>
              </v-list-item>
              <v-list-item @click="handleUpdateItem(item)">
                <template v-slot:append>
                  <v-icon icon="mdi-pencil"></v-icon>
                </template>
                <v-list-item-title>แก้ไข</v-list-item-title>
              </v-list-item>
              <v-list-item @click="handleDeleteItem(item.id)">
                <template v-slot:append>
                  <v-icon icon="mdi-delete"></v-icon>
                </template>
                <v-list-item-title>ลบ</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </section>
      </template>
      <!-- actions -->
    </v-data-table>
  </div>

  <!-- 👉 Confirm Dialog -->
  <ConfirmDeleteDialog
    v-model:isDialogVisible="deleteItemDialog.visible"
    v-model:title="deleteItemDialog.title"
    v-model:text="deleteItemDialog.text"
    @confirm="deleteItem"
  />

  <ViewTaskDialog
    v-model:isDialogVisible="viewItemDialog.visible"
    v-model:title="viewItemDialog.title"
  />

  <UpdateTaskDialog
    v-model:isDialogVisible="updateItemDialog.visible"
    v-model:title="updateItemDialog.title"
  />
</template>

<style lang="scss" scoped>
// :deep(.v-data-table) thead tr {
//   color: rgb(var(--v-theme-primary)) !important;
//   font-weight: 700;
// }

:deep(.v-data-table-header__content) span {
  color: rgb(var(--v-theme-on-surface-variant)) !important;
  white-space: nowrap;
}

// :deep(.v-data-table) tbody {
//   color: rgb(var(--v-theme-secondary)) !important;
//   font-weight: 600;
// }

// :deep(.v-data-table) tbody tr td i {
//   color: #616161;
// }
</style>
